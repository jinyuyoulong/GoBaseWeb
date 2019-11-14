package main

import (
	"project-web/src/bootstrap/middleware"
	"project-web/src/bootstrap/route"
	"project-web/src/bootstrap/service"

	"github.com/kataras/iris/v12"

	"github.com/kataras/iris/v12/view"
)

func main() {
	app := initApplication()

	setApplication(app)
}

func initApplication() *iris.Application {
	app := iris.New()
	return app
}
func setApplication(app *iris.Application) {

	middleware.RegistMiddleware(app)

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
