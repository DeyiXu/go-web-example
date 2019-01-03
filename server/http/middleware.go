package http

import (
	"net/http"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// CloseCache 关闭缓存
func CloseCache(c *gin.Context) {
	c.Next()
	c.Header("Pragma", "no-cahce")
	c.Header("Cache-Control", "no-cahce")
}

// AuthRequired 身份验证
func AuthRequired(ctx *gin.Context) {
	session := sessions.Default(ctx)
	currentAccount := session.Get("current_account")
	if currentAccount == nil {
		curl := ctx.Request.RequestURI
		if strings.Contains(curl, "logout.html") {
			curl = "/index.html"
		}
		ctx.Redirect(http.StatusSeeOther, "/login.html?redirect_url="+curl)
		ctx.Abort()
		return
	}
	ctx.Next()
}
