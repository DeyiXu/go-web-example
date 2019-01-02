package http

import (
	"github.com/gin-gonic/gin"
)

// setAPIRouter 设置路由
func setAPIRouter(router *gin.RouterGroup) {
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
}

// setAPI2Router V2
func setAPI2Router(router *gin.RouterGroup) {

	router.GET("/info", func(c *gin.Context) {
		c.String(200, "go web example.")
	})

}
