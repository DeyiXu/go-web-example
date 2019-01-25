package api

import (
	"github.com/gin-gonic/gin"
	"github.com/nilorg/pkg/gin/route"
	"github.com/nilorg/pkg/logger"
	ngin "github.com/nilorg/pkg/gin"
	"net/http"
)

// AuthController 授权控制器
type AuthController struct {

}
// NewAuthController 创建授权控制器
func NewAuthController() *AuthController {
	return &AuthController{}
}

// Route ... 路由
func (auth *AuthController) Route() []route.Route {
	return []route.Route{
		{
			Name:         "登录",
			Method:       http.MethodPost,
			RelativePath: "/login",
			Auth:         false,
			HandlerFunc:  ngin.WebControllerFunc(auth.PostLogin, "login"),
		},
		{
			Name:         "注册",
			Method:       http.MethodPost,
			RelativePath: "/register",
			Auth:         false,
			HandlerFunc:   ngin.WebControllerFunc(auth.PostRegister, "register"),
		},
	}
}


// PostLogin ...
func (*AuthController) PostLogin(ctx *ngin.WebContext) {
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


// PostRegister ...
func (*AuthController) PostRegister(ctx *ngin.WebContext) {

}

