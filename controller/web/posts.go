package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ngin "github.com/nilorg/pkg/gin"
	"github.com/nilorg/pkg/gin/route"
)

// PostsController 文章制器
type PostsController struct {
}

// NewPostsController 创建文章控制器
func NewPostsController() *PostsController {
	return &PostsController{}
}

// Route ... 路由
func (posts *PostsController) Route() []route.Route {
	return []route.Route{
		{
			Name:         "编辑文章",
			Method:       http.MethodGet,
			RelativePath: "/posts/edit.html",
			Auth:         true,
			HandlerFunc:  ngin.WebControllerFunc(posts.GetEdit, "posts_edit"),
		},
	}
}

// GetEdit ...
func (*PostsController) GetEdit(ctx *ngin.WebContext) {
	ctx.RenderPage(gin.H{
		"title": "Edit...",
	})
}
