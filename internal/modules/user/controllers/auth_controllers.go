package controllers

import (
	"blog/internal/modules/user/requests/auth"
	UserService "blog/internal/modules/user/services"
	"blog/pkg/converters"
	"blog/pkg/errors"
	"blog/pkg/html"
	"blog/pkg/old"
	"blog/pkg/sessions"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	userService UserService.UserServiceInterface
}

func New() *Controller {
	return &Controller{
		userService: UserService.New(),
	}
}

func (controller *Controller) Register(ctx *gin.Context) {
	html.Render(ctx, http.StatusOK, "modules/user/html/register", gin.H{
		"title": "Register",
	})
}

func (controller *Controller) HandleRegister(ctx *gin.Context) {
	var request auth.RegisterRequest
	// This will infer what binder to use depending on the content-type header.
	if err := ctx.ShouldBind(&request); err != nil {
		errors.Init()
		errors.SetFromErrors(err)
		sessions.Set(ctx, "errors", converters.MapToString(errors.Get()))

		old.Init()
		old.Set(ctx)
		sessions.Set(ctx, "old", converters.UrlValuesToString(old.Get()))

		ctx.Redirect(http.StatusFound, "/register")
		return
	}

	if controller.userService.CheckUserExist(request.Email) {
		errors.Init()
		errors.Add("Email", "Email address already exists")
		sessions.Set(ctx, "errors", converters.MapToString(errors.Get()))

		old.Init()
		old.Set(ctx)
		sessions.Set(ctx, "old", converters.UrlValuesToString(old.Get()))

		ctx.Redirect(http.StatusFound, "/register")
		return
	}

	user, err := controller.userService.Create(request)

	if err != nil {
		ctx.Redirect(http.StatusFound, "/register")
		return
	}

	sessions.Set(ctx, "auth", strconv.Itoa(int(user.ID)))

	log.Printf("user created as %s \n", user.Name)
	ctx.Redirect(http.StatusFound, "/")
}

func (controller *Controller) Login(ctx *gin.Context) {
	html.Render(ctx, http.StatusOK, "modules/user/html/login", gin.H{
		"title": "Login",
	})
}

func (controller *Controller) HandleLogin(ctx *gin.Context) {

	var request auth.LoginRequest
	// This will infer what binder to use depending on the content-type header.
	if err := ctx.ShouldBind(&request); err != nil {
		errors.Init()
		errors.SetFromErrors(err)
		sessions.Set(ctx, "errors", converters.MapToString(errors.Get()))

		old.Init()
		old.Set(ctx)
		sessions.Set(ctx, "old", converters.UrlValuesToString(old.Get()))

		ctx.Redirect(http.StatusFound, "/login")
		return
	}
	user, err := controller.userService.HandleUserLogin(request)

	if err != nil {
		errors.Init()
		errors.Add("email", err.Error())
		sessions.Set(ctx, "errors", converters.MapToString(errors.Get()))

		old.Init()
		old.Set(ctx)
		sessions.Set(ctx, "old", converters.UrlValuesToString(old.Get()))

		ctx.Redirect(http.StatusFound, "/login")
		return
	}

	sessions.Set(ctx, "auth", strconv.Itoa(int(user.ID)))

	log.Printf("the user logged in successfully with a name %s \n", user.Name)
	ctx.Redirect(http.StatusFound, "/")
}

func (controller *Controller) HandleLogout(ctx *gin.Context) {
	sessions.Remove(ctx, "auth")

	ctx.Redirect(http.StatusFound, "/")

}
