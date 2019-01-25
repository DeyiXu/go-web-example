package web

import "github.com/gin-gonic/gin"

// HomeController 错误控制器
type ErrorController struct {
}

// NewErrorController 创建错误控制器
func NewErrorController() *ErrorController {
	return &ErrorController{}
}

// Error404 ...
func (*ErrorController) Error404(ctx *gin.Context) {
	ctx.HTML(404, "error_404.tmpl", nil)
}
