package route

import (
	"project-web/src/controller"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

// SetRoute 配置路由
func SetRoute(route *iris.Application) {
	IndexRoute(route)
	AdminRoute(route)
	ImageRout(route)
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

func ImageRout(route *iris.Application) {
	imgV := mvc.New(route.Party("/image"))
	imgC := new(controller.ImageController)
	imgV.Handle(imgC)
	// url := "/carlogo/100x100/8b/2019-07-01/8b18e22f914e64a1a933541ed0e97ae0.jpg"
	// route.Get(url, imgC.GetCreateResizeOrgImage)
	// route.Get("image/create/")
}
