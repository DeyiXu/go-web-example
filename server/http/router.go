package http

import (
	"github.com/DeyiXu/go-web-example/controller/auth"
	"github.com/gin-gonic/gin"
)

// setRouter 设置路由
func setRouter(handler *gin.Engine) {
	handler.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	handler.POST("/auth/login", auth.Post)
}

// setRouter V2
func setWechatRouter(handler *gin.Engine) {
	wechat := handler.Group("/v2")
	{
		wechat.GET("/info", func(c *gin.Context) {
			c.String(200, "da keng!")
		})
	}
}
