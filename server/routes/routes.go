package routes

import (
	v1 "gin-vegan-map-api/server/api/v1"

	_ "gin-vegan-map-api/docs"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// Engine - engine
func Engine() *gin.Engine {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	ver1 := r.Group("/api/v1")
	{
		// chatfuel
		ver1.GET("/chatfuel/restaurants/type/:type/regional/:regional", v1.GetRestaurantByTypeAndRegional)
		ver1.GET("/chatfuel/pics/type/:type/regional/:regional", v1.GetPicsByTypeAndRegional)
	}

	return r
}
