package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ngin "github.com/nilorg/pkg/gin"
	"github.com/nilorg/pkg/gin/route"
	"github.com/nilorg/pkg/logger"
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
			Method:       http.MethodGet,
			RelativePath: "/login.html",
			Auth:         false,
			HandlerFunc:  ngin.WebControllerFunc(auth.GetLogin, "login"),
		},
		{
			Name:         "注册",
			Method:       http.MethodGet,
			RelativePath: "/register.html",
			Auth:         false,
			HandlerFunc:  ngin.WebControllerFunc(auth.GetRegister, "register"),
		},
		{
			Name:         "退出登录",
			Method:       http.MethodGet,
			RelativePath: "/logout.html",
			Auth:         true,
			HandlerFunc:  ngin.WebAPIControllerFunc(auth.Logout),
		},
	}
}

// GetLogin ...
func (*AuthController) GetLogin(ctx *ngin.WebContext) {
	redirectURL := ctx.Query("redirect_url")
	logger.Debugf("GetLogin redirectURL:%s", redirectURL)
	ctx.RenderSinglePage(gin.H{
		"title":        "Login...",
		"redirect_url": redirectURL,
	})
}

// Logout ...
func (*AuthController) Logout(ctx *ngin.WebAPIContext) {
	ctx.DelCurrentAccount()
	ctx.Redirect(http.StatusSeeOther, "/login.html")
}

// GetRegister ...
func (*AuthController) GetRegister(ctx *ngin.WebContext) {
	ctx.RenderSinglePage(gin.H{
		"title": "Register...",
	})
}

// GetMenuData 获取菜单数据
func (*AuthController) GetMenuData(value interface{}) gin.H {
	logger.Debugln("getMenuData...")
	return gin.H{
		"account": value,
		"info":    "...",
	}
}

// GetNavigationData 获取导航数据
func (*AuthController) GetNavigationData(value interface{}) gin.H {
	logger.Debugln("GetNavigationData...")
	return gin.H{
		"account": value,
		"info":    "navigation...",
	}
}
