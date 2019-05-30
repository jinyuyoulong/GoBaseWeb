package route

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"xxx.com/projectweb/src/app/controller"
)

// Configure registers the necessary routes to the app.
func Configure(a *iris.Application) {
	indexC := new(controller.IndexController)
	index := mvc.New(a.Party("/"))
	index.Handle(indexC)

	admin := mvc.New(a.Party("/admin"))
	admin.Handle(new(controller.AdminController))

	// -------------------------------------------------------
	a.Get("/setroute", indexC.GetIndexHandler)
	// a.Get("/follower/{id:long}", indexC.GetFollowerHandler)
	// b.Get("/like/{id:long}", GetLikeHandler)
}
