package controllers

import (
	"net/http"

	ArticleService "blog/internal/modules/article/services"
	"blog/pkg/html"

	"github.com/gin-gonic/gin"
)

type Controllers struct {
	articleService ArticleService.ArticleServiceInterface
}

func New() *Controllers {
	return &Controllers{
		articleService: ArticleService.New(),
	}
}

func (controller *Controllers) Index(ctx *gin.Context) {
	html.Render(ctx, http.StatusOK, "modules/home/html/home", gin.H{
		"title":    "Home Page",
		"featured": controller.articleService.GetFeaturedArticle(),
		"stories":  controller.articleService.GetStoriesArticle(),
	})

	// ctx.JSON(http.StatusOK, gin.H{
	// 	"featured": controller.articleService.GetFeaturedArticle(),
	// 	"stories":  controller.articleService.GetStoriesArticle(),
	// })
}
