package routes

import (
	"blog/internal/middlewares"
	articleCtrl "blog/internal/modules/article/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {

	articleController := articleCtrl.New()
	router.GET("/articles/:id", articleController.Show)
	authGroup := router.Group("/articles")

	authGroup.Use(middlewares.IsAuth())
	{
		authGroup.GET("/create", articleController.Create)
		authGroup.POST("/store", articleController.Store)
	}

}
