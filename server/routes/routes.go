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
		// Restaurants
		ver1.GET("/restaurants", v1.GetRestaurants)
		ver1.GET("/restaurants/:id", v1.GetRestaurant)
		ver1.GET("/restaurants-type/:type/regional/:regional", v1.GetRestaurantsByTypeAndRegional)

		// Pics
		ver1.GET("/pics", v1.GetPics)
		ver1.GET("/pics/:id", v1.GetPic)
		ver1.GET("/pics-type/:type/regional/:regional", v1.GetPicsByTypeAndRegional)

		// Chatfuel
		ver1.GET("/chatfuel/restaurants/type/:type/regional/:regional", v1.GetChatfuelRestaurantByTypeAndRegional)
		ver1.GET("/chatfuel/pics/type/:type/regional/:regional", v1.GetChatfurlPicsByTypeAndRegional)
	}

	return r
}
