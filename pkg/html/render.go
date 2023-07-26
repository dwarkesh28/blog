package html

import (
	"blog/internal/providers/view"

	"github.com/gin-gonic/gin"
)

func Render(ctx *gin.Context, code int, name string, data gin.H) {
	data = view.WithGlobalData(ctx, data)
	ctx.HTML(code, name, data)
}
