package routes

import (
	v1 "gin-vegan-map-api/server/api/v1"

	"github.com/gin-gonic/gin"
)

// Engine - engine
func Engine() *gin.Engine {
	r := gin.Default()

	ver1 := r.Group("/api/v1")
	{
		// chatfuel
		ver1.GET("/chatfuel/restaurants/type/:type/regional/:regional", v1.GetRestaurantByTypeAndRegional)
		ver1.GET("/chatfuel/pics/type/:type/regional/:regional", v1.GetPicsByTypeAndRegional)
	}

	return r
}
