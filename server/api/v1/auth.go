package v1

import (
	"gin-vegan-map-api/server/common"
	e "gin-vegan-map-api/server/enum"
	"gin-vegan-map-api/server/models"
	"gin-vegan-map-api/server/util"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type auth struct {
	Name     string `form:"name" json:"name" xml:"name" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

type responseAuth struct {
	Token string `json:"token"`
	ID    uint   `json:"id"`
}

// GetAuth -
// @Summary 登入驗證
// @Description 登入驗證
// @Tags 驗證
// @Accept  json
// @Produce  json
// @Param username body string true "username"
// @Param password body string true "password"
// @Success 200 {object} models.JSONResult{data=string} "desc"
// @Failure 500 {object} models.JSONResult{data=string} "desc"
// @Router /auth [get]
func GetAuth(c *gin.Context) {

	var a auth
	err := c.ShouldBindJSON(&a)
	if err != nil {
		c.JSON(http.StatusBadRequest,
			ErrBody(
				http.StatusBadRequest,
				e.StatusEnum.String(e.InvalidParams),
			))
		return
	}

	isExist, err := models.CheckAuth(a.Name, a.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized,
			ErrBody(
				http.StatusUnauthorized,
				e.AuthEnum.String(e.AuthFail),
			))
		return
	}

	if !isExist {
		c.JSON(http.StatusBadRequest,
			ErrBody(
				http.StatusBadRequest,
				e.UserEnum.String(e.UserNotExist),
			))
		return
	}

	token, err := util.GenerateToken(a.Name, a.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized,
			ErrBody(
				http.StatusUnauthorized,
				e.AuthEnum.String(e.GenerateTokenFail),
			))
		return
	}

	user, err := models.GetUserByName(a.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			ErrBody(
				http.StatusInternalServerError,
				e.StatusEnum.String(e.Failed),
			))
		return
	}

	resAuth := responseAuth{
		Token: token,
		ID:    user.ID,
	}
	c.JSON(http.StatusOK, OkBody(resAuth))
}

// GetToken -
func GetToken(c *gin.Context) string {
	token, ok := c.GetQuery("_t")
	if !ok {
		hToken := c.GetHeader("Authorization")
		if len(hToken) < common.BearerLength {
			return ""
		}
		token = strings.TrimSpace(hToken[common.BearerLength:])
	}

	return token
}
