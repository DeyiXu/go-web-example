package error

import "github.com/gin-gonic/gin"

// Error404 ...
func Error404(ctx *gin.Context) {
	ctx.HTML(404, "error_404.tmpl", nil)
}
