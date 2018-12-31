package auth

import (
	"github.com/gin-gonic/gin"
	ngin "github.com/nilorg/pkg/gin"
)

// Login ...
func Login(ctx *ngin.WebContext) {
	ctx.RenderSinglePage(gin.H{
		"title": "Login...",
	})
}

// Register ...
func Register(ctx *ngin.WebContext) {
	ctx.RenderSinglePage(gin.H{
		"title": "Register...",
	})
}
