package main

import (
	"project-web/src/app/bootstrap/diserver"
	"project-web/src/app/bootstrap/route"
	"project-web/src/app/library/helper"

	"github.com/kataras/iris"
)

func main() {
	app := iris.New()
	diserver.NewServices(app)
	route.SetRoute(app)
	port := new(helper.Helper).NewConfig().Get("app.port").(string)
	app.Run(iris.Addr(port))
}
