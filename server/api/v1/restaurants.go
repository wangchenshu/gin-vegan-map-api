package v1

import (
	e "gin-vegan-map-api/server/enum"
	"gin-vegan-map-api/server/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// GetRestaurants - 取得餐廳列表
// @Summary 取得餐廳列表
// @Description 取得餐廳列表
// @Tags 餐廳
// @Accept  json
// @Produce  json
// @Success 200 {object} models.JSONResult{data=string} "desc"
// @Failure 500 {object} models.JSONResult{data=string} "desc"
// @Router /api/v1/restaurants [get]
func GetRestaurants(c *gin.Context) {
	result, err := models.GetRestaurants()

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

// GetRestaurant - 取得餐廳資料
// @Summary 取得餐廳資料
// @Description 取得餐廳資料
// @Tags 餐廳
// @Accept  json
// @Produce  json
// @Success 200 {object} models.JSONResult{data=string} "desc"
// @Failure 500 {object} models.JSONResult{data=string} "desc"
// @Router /api/v1/restaurants/{id} [get]
func GetRestaurant(c *gin.Context) {

	id := com.StrTo(c.Param("id")).MustInt()
	result, err := models.GetRestaurant(id)

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

// GetRestaurantsByTypeAndRegional - 以分類及地區取得餐廳資料
// @Summary 以分類及取得餐廳資料
// @Description 以分類及取得餐廳資料
// @Tags 餐廳
// @Accept  json
// @Produce  json
// @Success 200 {object} models.JSONResult{data=string} "desc"
// @Failure 500 {object} models.JSONResult{data=string} "desc"
// @Router /api/v1/restaurants-type/{type}/{regional} [get]
func GetRestaurantsByTypeAndRegional(c *gin.Context) {
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
