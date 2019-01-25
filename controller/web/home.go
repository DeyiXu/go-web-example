package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ngin "github.com/nilorg/pkg/gin"
	"github.com/nilorg/pkg/gin/route"
)

// HomeController 主控制器
type HomeController struct {
}

// NewHomeController 创建主控制器
func NewHomeController() *HomeController {
	return &HomeController{}
}

// Route ... 路由
func (home *HomeController) Route() []route.Route {
	return []route.Route{
		{
			Name:         "首页",
			Method:       http.MethodGet,
			RelativePath: "/",
			Auth:         true,
			HandlerFunc:  ngin.WebControllerFunc(home.Index, "index"),
		},
		{
			Name:         "首页",
			Method:       http.MethodGet,
			RelativePath: "/index.html",
			Auth:         true,
			HandlerFunc:  ngin.WebControllerFunc(home.Index, "index"),
		},
		{
			Name:         "测试页面",
			Method:       http.MethodGet,
			RelativePath: "/test",
			Auth:         true,
			HandlerFunc:  ngin.WebControllerFunc(home.Test, "test"),
		},
	}
}

// Index ...
func (*HomeController) Index(ctx *ngin.WebContext) {
	ctx.RenderPage(gin.H{
		"title": "index...",
	})
}

// Test ...
func (*HomeController) Test(ctx *ngin.WebContext) {
	ctx.RenderPage(gin.H{
		"title": "test...",
	})
}
