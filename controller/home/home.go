package home

import (
	"github.com/gin-gonic/gin"
	ngin "github.com/nilorg/pkg/gin"
)

// Index ...
func Index(ctx *ngin.WebContext) {
	ctx.SetCurrentAccount(gin.H{
		"name":   "德意洋洋",
		"status": "在线",
		"cover":  "/assets/img/IMG_3476.jpg",
	})
	ctx.RenderPage(gin.H{
		"title": "index...",
	})
}

// Test ...
func Test(ctx *ngin.WebContext) {
	ctx.RenderPage(gin.H{
		"title": "test...",
	})
}
