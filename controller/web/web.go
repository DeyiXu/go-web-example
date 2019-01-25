package web

import "github.com/nilorg/pkg/gin/route"

var (
	Auth  = NewAuthController()
	Home  = NewHomeController()
	Posts = NewPostsController()
	Error = NewErrorController()
)

func Router() []route.Router {
	routerList := make([]route.Router, 0)
	routerList = append(routerList, Auth)
	routerList = append(routerList, Home)
	routerList = append(routerList, Posts)
	return routerList
}
