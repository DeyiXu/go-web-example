package http

import (
	"github.com/DeyiXu/go-web-example/controller/api"
	"github.com/DeyiXu/go-web-example/controller/web"
	"github.com/gin-gonic/gin"
	"github.com/nilorg/pkg/gin/route"
)

// setRouter 设置路由
func setRouter(router *gin.RouterGroup, authRouter *gin.RouterGroup, controllers ...route.Router) {
	for _, controller := range controllers {
		for _, route := range controller.Route() {
			if route.Auth {
				authRouter.Handle(route.Method, route.RelativePath, route.HandlerFunc)
			} else {
				router.Handle(route.Method, route.RelativePath, route.HandlerFunc)
			}
		}
	}
}

// setAPIRouter 设置路由
func setAPIRouter(router *gin.RouterGroup) {
	authRouter := router.Group("/")
	authRouter.Use(AuthRequired)
	setRouter(router, authRouter, api.Router()...)
}

func setWebRouter(router *gin.RouterGroup) {
	// auth
	authRouter := router.Group("/")
	authRouter.Use(AuthRequired)
	setRouter(router, authRouter, web.Router()...)
}
