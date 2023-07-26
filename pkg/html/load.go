package html

import "github.com/gin-gonic/gin"

func LoadHTML(router *gin.Engine) {
	//internal/modules/moduleName/html/view.tmpl
	router.LoadHTMLGlob("internal/**/**/**/*tmpl")
}
