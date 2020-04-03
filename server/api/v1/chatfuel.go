package v1

import (
	e "gin-vegan-map-api/server/enum"
	"gin-vegan-map-api/server/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetRestaurantByTypeAndRegional -
// GetEmployee - 以分類及地區取得資料
// @Summary 以分類及取得資料
// @Description 以分類及取得資料
// @Tags Chatfuel
// @Accept  json
// @Produce  json
// @Success 200 {object} models.JSONResult{data=string} "desc"
// @Failure 500 {object} models.JSONResult{data=string} "desc"
// @Router /api/v1/chatfuel/type/{type}/{regional} [get]
func GetRestaurantByTypeAndRegional(c *gin.Context) {
	typeName := c.Param("type")
	regional := c.Param("regional")
	result, err := models.GetRestaurantsByTypeAndRegional(typeName, regional)

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

// GetRestaurantByFriedAndRegional -
// GetEmployee - 以地區取得炸物資料
// @Summary 以地區取得炸物資料
// @Description 以地區取得炸物資料
// @Tags Chatfuel
// @Accept  json
// @Produce  json
// @Success 200 {object} models.JSONResult{data=string} "desc"
// @Failure 500 {object} models.JSONResult{data=string} "desc"
// @Router /api/v1/chatfuel/fried/{regional} [get]
func GetRestaurantByFriedAndRegional(c *gin.Context) {
	regional := c.Param("regional")
	result, err := models.GetRestaurantsByFriedRegional(regional)

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

// GetPicsByTypeAndRegional -
// GetEmployee - 以分類及地區取得圖片資料
// @Summary 以分類及地區取得圖片資料
// @Description 以分類及地區取得圖片資料
// @Tags Chatfuel
// @Accept  json
// @Produce  json
// @Success 200 {object} models.JSONResult{data=string} "desc"
// @Failure 500 {object} models.JSONResult{data=string} "desc"
// @Router /api/v1/chatfuel/type/{type}/regional/{regional}/pic [get]
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
