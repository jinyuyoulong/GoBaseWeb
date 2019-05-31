package route

import (
	"project-web/src/app/controller"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

// SetRoute 配置路由
func SetRoute(route *iris.Application) {
	IndexRoute(route)
	AdminRoute(route)
}

// IndexRoute 配置index route
func IndexRoute(route *iris.Application) {
	indexC := new(controller.IndexController)
	index := mvc.New(route.Party("/"))
	index.Handle(indexC)

	route.Get("/setroute", indexC.GetIndexHandler)
}

// AdminRoute admin route
func AdminRoute(route *iris.Application) {
	admin := mvc.New(route.Party("/admin"))
	admin.Handle(new(controller.AdminController))
}
