package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ngin "github.com/nilorg/pkg/gin"
	"github.com/nilorg/pkg/logger"
)

// GetLogin ...
func GetLogin(ctx *ngin.WebContext) {
	redirectURL := ctx.Query("redirect_url")
	logger.Debugf("GetLogin redirectURL:%s", redirectURL)
	ctx.RenderSinglePage(gin.H{
		"title":        "Login...",
		"redirect_url": redirectURL,
	})
}

// PostLogin ...
func PostLogin(ctx *ngin.WebContext) {
	userName := ctx.PostForm("username")
	pass := ctx.PostForm("pass")
	me := ctx.PostForm("me")
	redirectURL := ctx.PostForm("redirect_url")
	logger.Debugf("name:%s|pass:%s|me:%s", userName, pass, me)
	if userName == "xudeyi" && pass == "xudeyi" {

		ctx.SetCurrentAccount(gin.H{
			"username": userName,
			"name":     "德意洋洋",
			"status":   "在线",
			"cover":    "/assets/img/IMG_3476.jpg",
		})

		if redirectURL == "" {
			redirectURL = "/index.html"
		}
		// post redirect get request by http status code 303.
		ctx.Redirect(http.StatusSeeOther, redirectURL)
	} else {
		ctx.Redirect(http.StatusSeeOther, "/login.html")
	}
}

// Logout ...
func Logout(ctx *ngin.WebAPIContext) {
	ctx.DelCurrentAccount()
	ctx.Redirect(http.StatusSeeOther, "/login.html")
}

// GetRegister ...
func GetRegister(ctx *ngin.WebContext) {
	ctx.RenderSinglePage(gin.H{
		"title": "Register...",
	})
}

// PostRegister ...
func PostRegister(ctx *ngin.WebContext) {

}

// GetMenuData 获取菜单数据
func GetMenuData(value interface{}) gin.H {
	logger.Debugln("getMenuData...")
	return gin.H{
		"account": value,
		"info":    "...",
	}
}

// GetNavigationData 获取导航数据
func GetNavigationData(value interface{}) gin.H {
	logger.Debugln("GetNavigationData...")
	return gin.H{
		"account": value,
		"info":    "navigation...",
	}
}
