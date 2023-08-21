package controllers

import (
	"blog/internal/modules/article/requests/articles"
	ArticleService "blog/internal/modules/article/services"
	"blog/internal/modules/user/helpers"
	"blog/pkg/converters"
	"blog/pkg/errors"
	"blog/pkg/html"
	"blog/pkg/old"
	"blog/pkg/sessions"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

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

func (controller *Controllers) Show(ctx *gin.Context) {
	// Get id from url

	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		html.Render(
			ctx,
			http.StatusInternalServerError,
			"templates/errors/html/500",
			gin.H{
				"title":   "Server Error",
				"message": "error converting the id"})
		return
	}

	article, err := controller.articleService.Find(id)

	if err != nil {
		html.Render(
			ctx,
			http.StatusNotFound,
			"templates/errors/html/404",
			gin.H{
				"title":   "Page not found",
				"message": err.Error()})
		return
	}

	html.Render(
		ctx,
		http.StatusOK,
		"modules/article/html/show",
		gin.H{
			"title":   "show article",
			"article": article})

}

func (controller *Controllers) Create(ctx *gin.Context) {
	html.Render(ctx, http.StatusOK, "modules/article/html/create", gin.H{
		"title": "Create article",
	})
}

func (controller *Controllers) Store(ctx *gin.Context) {
	file, err := ctx.FormFile("image")

	if err != nil {
		log.Fatal("failed to upload image")
		ctx.Redirect(http.StatusFound, "/")
		return
	}

	var request articles.StoreRequest
	// This will infer what binder to use depending on the content-type header.
	if err := ctx.ShouldBind(&request); err != nil {

		log.Print("can not bind")
		errors.Init()
		errors.SetFromErrors(err)
		sessions.Set(ctx, "errors", converters.MapToString(errors.Get()))

		old.Init()
		old.Set(ctx)
		sessions.Set(ctx, "old", converters.UrlValuesToString(old.Get()))

		ctx.Redirect(http.StatusFound, "/articles/create")
		return
	}
	user := helpers.Auth(ctx)
	// fullFolderPath := fmt.Sprintf("./assets/img/users/%s/", user.Name)
	pathForDatabase := fmt.Sprintf("/assets/img/users/%s/", user.Name)

	// log.Print(fullFolderPath)

	// folderPath := fmt.Sprintf("user/%s/%s", user.Name, file.Filename)

	if _, err := os.Stat("./" + pathForDatabase); os.IsNotExist(err) {
		// Folder doesn't exist, so create it
		err := os.MkdirAll("./"+pathForDatabase, os.ModePerm)
		if err != nil {
			fmt.Println("Error creating folder:", err)
		}
	}
	err = ctx.SaveUploadedFile(file, "./"+pathForDatabase+file.Filename)

	if err != nil {
		log.Print("can not upload image", err)
		return
	}

	// log.Print(request.Image)cls

	article, err := controller.articleService.StoreAsUser(request, user, pathForDatabase+file.Filename)

	if err != nil {
		log.Fatal("can not store article")
		ctx.Redirect(http.StatusFound, "/articles/create")
		return
	}

	ctx.Redirect(http.StatusFound, fmt.Sprintf("/articles/%d", article.ID))
}
