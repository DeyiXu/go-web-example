package http

import (
	"github.com/gin-gonic/gin"
)

// CloseCache 关闭缓存
func CloseCache(c *gin.Context) {
	c.Next()
	c.Header("Pragma", "no-cahce")
	c.Header("Cache-Control", "no-cahce")
}
