package auth

import "github.com/gin-gonic/gin"

// Post 登录
func Post(ctx *gin.Context) {

	ctx.JSON(200, "Ok")
}
