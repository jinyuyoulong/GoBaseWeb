package main

import (
	"project-web/src/bootstrap/route"
	"project-web/src/bootstrap/service"

	"github.com/kataras/iris"
	"github.com/kataras/iris/view"
)

func main() {
	app := iris.New()

	setApplication(app)
}

func setApplication(app *iris.Application) {

	route.SetRoute(app)

	var port string

	di := service.GetDi()
	container := di.Container
	container.Invoke(func(config *service.Config) {
		port = config.App.Port
	})
	container.Invoke(func(viewEngine *view.HTMLEngine) {
		app.RegisterView(viewEngine)
	})
	app.Run(iris.Addr(port))

}
