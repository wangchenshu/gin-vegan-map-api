package v1

import (
	"gin-vegan-map-api/server/common"
	e "gin-vegan-map-api/server/enum"
	"gin-vegan-map-api/server/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// GetUsers - 取得用戶列表
// @Summary 取得用戶列表
// @Description 取得用戶列表
// @Tags 用戶
// @Accept  json
// @Produce  json
// @Success 200 {object} models.JSONResult{data=string} "desc"
// @Failure 500 {object} models.JSONResult{data=string} "desc"
// @Router /api/v1/users [get]
func GetUsers(c *gin.Context) {
	result, err := models.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			ErrBody(
				http.StatusInternalServerError,
				e.StatusEnum.String(e.GetFailed),
			))
		return
	}

	c.JSON(http.StatusOK, OkBody(result))
}

// GetUser - 取得用戶資料
// @Summary 取得用戶資料
// @Description 取得用戶資料
// @Tags 用戶
// @Accept  json
// @Produce  json
// @Success 200 {object} models.JSONResult{data=string} "desc"
// @Failure 500 {object} models.JSONResult{data=string} "desc"
// @Router /api/v1/users/{id} [get]
func GetUser(c *gin.Context) {

	id := com.StrTo(c.Param("id")).MustInt()
	result, err := models.GetUser(id)

	if err != nil || result.ID <= 0 {
		c.JSON(http.StatusInternalServerError,
			ErrBody(
				http.StatusInternalServerError,
				e.StatusEnum.String(e.GetFailed),
			))
		return
	}

	c.JSON(http.StatusOK, OkBody(result))
}

// AddUserForm - 新增用戶資料
type AddUserForm struct {
	Name     string `form:"name" json:"name" xml:"name" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
	Roles    string `form:"roles" json:"roles" xml:"roles" binding:"required"`
	RealName string `form:"real_name" json:"real_name" xml:"real_name"`
}

// AddUser - 新增用戶
// @Summary 新增用戶
// @Description 新增用戶
// @Tags 用戶
// @Accept  json
// @Produce  json
// @Param name body string true "用戶名稱"
// @Param password body string true "會員密碼"
// @Param roles body string true "會員角色"
// @Param real_name body string false "用戶真實名稱"
// @Success 200 {object} models.JSONResult{data=string} "desc"
// @Failure 500 {object} models.JSONResult{data=string} "desc"
// @Router /api/v1/users [post]
func AddUser(c *gin.Context) {
	var (
		form AddUserForm
		user models.User
	)

	// Bind
	err := c.Bind(&form)
	if err != nil {
		c.JSON(http.StatusBadRequest,
			ErrBody(
				http.StatusBadRequest,
				e.StatusEnum.String(e.AddFailed),
			))
		return
	}

	// Create User Model
	hash, err := common.HashPassword(form.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			ErrBody(
				http.StatusInternalServerError,
				e.StatusEnum.String(e.AddFailed),
			))
		return
	}

	user = models.User{
		Name:     form.Name,
		Roles:    form.Roles,
		RealName: form.RealName,
		Password: hash,
	}

	err = models.AddUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest,
			ErrBody(
				http.StatusBadRequest,
				e.StatusEnum.String(e.AddFailed),
			))
		return
	}

	c.JSON(http.StatusOK, OkBody(user))
}

// UpdateUserForm - 更新用戶資料
type UpdateUserForm struct {
	ID       int    `form:"id" binding:"required"`
	Name     string `form:"name" json:"name" xml:"name" binding:"required"`
	Roles    string `form:"roles" json:"roles" xml:"roles" binding:"required"`
	RealName string `form:"real_name" json:"real_name" xml:"real_name"`
	Password string `form:"password" json:"password" xml:"password"`
}

// UpdateUser - 更新用戶資料
// @Summary 更新用戶資料
// @Description 更新用戶資料
// @Tags 用戶
// @Accept  json
// @Produce  json
// @Param id path int true "ID"
// @Param name body string true "用戶名稱"
// @Param password body string true "會員密碼"
// @Param roles body string true "會員角色"
// @Param real_name body string false "用戶真實名稱"
// @Success 200 {object} models.JSONResult{data=string} "desc"
// @Failure 500 {object} models.JSONResult{data=string} "desc"
// @Router /api/v1/users/{id} [put]
func UpdateUser(c *gin.Context) {
	var (
		form = UpdateUserForm{ID: com.StrTo(c.Param("id")).MustInt()}
		user models.User
	)

	// Bind
	err := c.Bind(&form)
	if err != nil {
		c.JSON(http.StatusBadRequest,
			ErrBody(
				http.StatusBadRequest,
				e.StatusEnum.String(e.UpdateFailed),
			))
		return
	}

	// Update User Model
	user = models.User{
		Name:     form.Name,
		Roles:    form.Roles,
		RealName: form.RealName,
	}

	if form.Password != "" {
		hash, err := common.HashPassword(form.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError,
				ErrBody(
					http.StatusInternalServerError,
					e.StatusEnum.String(e.UpdateFailed),
				))
			return
		}
		user.Password = hash
	}

	err = models.UpdateUser(form.ID, user)
	if err != nil {
		c.JSON(http.StatusBadRequest,
			ErrBody(
				http.StatusBadRequest,
				e.StatusEnum.String(e.UpdateFailed),
			))
		return
	}

	c.JSON(http.StatusOK, OkBody(user))
}

// DeleteUser - 刪除用戶
// @Summary 刪除用戶
// @Description 刪除用戶
// @Tags 用戶
// @Accept  json
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} models.JSONResult{data=string} "desc"
// @Failure 500 {object} models.JSONResult{data=string} "desc"
// @Router /api/v1/users/{id} [delete]
func DeleteUser(c *gin.Context) {

	id := com.StrTo(c.Param("id")).MustInt()
	err := models.DeleteUser(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError,
			ErrBody(
				http.StatusInternalServerError,
				e.StatusEnum.String(e.DeleteFailed),
			))
		return
	}

	c.JSON(http.StatusOK, OkBody(nil))
}
