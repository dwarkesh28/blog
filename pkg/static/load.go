package static

import (
	"github.com/gin-gonic/gin"
)

func LoadStatic(router *gin.Engine) {
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.Static("/assets", "./assets")
	router.Static("/articles/assets", "./assets")
}
