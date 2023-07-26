package routes

import (
	homeCtrl "blog/internal/modules/home/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {

	homeController := homeCtrl.New()
	router.GET("/", homeController.Index)

	// router.GET("/about", func(ctx *gin.Context) {
	// 	html.Render(ctx, http.StatusOK, "modules/home/html/about", gin.H{
	// 		"title": "About Page",
	// 	})
	// })
}
