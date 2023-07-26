package view

import (
	"blog/internal/modules/user/helpers"
	"blog/pkg/converters"
	"blog/pkg/sessions"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func WithGlobalData(ctx *gin.Context, data gin.H) gin.H {
	data["APP_NAME"] = viper.Get("App.Name")
	data["ERRORS"] = converters.StringToMap(sessions.Flash(ctx, "errors"))
	data["OLD"] = converters.StringToUrlValues(sessions.Flash(ctx, "old"))
	data["AUTH"] = helpers.Auth(ctx)
	return data
}
