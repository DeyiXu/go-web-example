package posts

import (
	"github.com/gin-gonic/gin"
	ngin "github.com/nilorg/pkg/gin"
)

// GetEdit ...
func GetEdit(ctx *ngin.WebContext) {
	ctx.RenderPage(gin.H{
		"title": "Edit...",
	})
}
