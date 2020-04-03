package v1

import (
	e "gin-vegan-map-api/server/enum"
	"gin-vegan-map-api/server/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// GetPics - 取得圖片列表
// @Summary 取得圖片列表
// @Description 取得圖片列表
// @Tags 圖片
// @Accept  json
// @Produce  json
// @Success 200 {object} models.JSONResult{data=string} "desc"
// @Failure 500 {object} models.JSONResult{data=string} "desc"
// @Router /api/v1/pics [get]
func GetPics(c *gin.Context) {
	result, err := models.GetPics()

	if err != nil {
		c.JSON(http.StatusBadRequest,
			ErrBody(
				http.StatusBadRequest,
				e.StatusEnum.String(e.GetFailed),
			))
		return
	}

	c.JSON(http.StatusOK, OkBody(result))
}

// GetPic - 取得圖片資料
// @Summary 取得圖片資料
// @Description 取得圖片資料
// @Tags 圖片
// @Accept  json
// @Produce  json
// @Success 200 {object} models.JSONResult{data=string} "desc"
// @Failure 500 {object} models.JSONResult{data=string} "desc"
// @Router /api/v1/pics/{id} [get]
func GetPic(c *gin.Context) {

	id := com.StrTo(c.Param("id")).MustInt()
	result, err := models.GetPic(id)

	if err != nil || result.ID <= 0 {
		c.JSON(http.StatusBadRequest,
			ErrBody(
				http.StatusBadRequest,
				e.StatusEnum.String(e.GetFailed),
			))
		return
	}

	c.JSON(http.StatusOK, OkBody(result))
}

// GetPicsByTypeAndRegional - 以分類及地區取得圖片資料
// @Summary 以分類及地區取得圖片資料
// @Description 以分類及地區取得圖片資料
// @Tags 圖片
// @Accept  json
// @Produce  json
// @Success 200 {object} models.JSONResult{data=string} "desc"
// @Failure 500 {object} models.JSONResult{data=string} "desc"
// @Router /api/v1/pics-type/{type}/regional/{regional} [get]
func GetPicsByTypeAndRegional(c *gin.Context) {
	typeName := c.Param("type")
	regional := c.Param("regional")

	result, err := models.GetPicsByTypeAndRegional(typeName, regional)
	if err != nil {
		c.JSON(http.StatusBadRequest,
			ErrBody(
				http.StatusBadRequest,
				e.StatusEnum.String(e.GetFailed),
			))
		return
	}

	c.JSON(http.StatusOK, OkBody(result))
}
