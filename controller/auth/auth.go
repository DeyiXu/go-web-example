package auth

import (
	"github.com/gin-gonic/gin"
	ngin "github.com/nilorg/pkg/gin"
	"github.com/nilorg/pkg/logger"
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

// GetMenuData 获取菜单数据
func GetMenuData(value interface{}) gin.H {
	logger.Debugln("getMenuData...")
	logger.Debugln(value)
	return gin.H{
		"account": value,
		"info":    "...",
	}
}

// GetNavigationData 获取导航数据
func GetNavigationData(value interface{}) gin.H {
	logger.Debugln("GetNavigationData...")
	logger.Debugln(value)
	return gin.H{
		"account": value,
		"info":    "navigation...",
	}
}
