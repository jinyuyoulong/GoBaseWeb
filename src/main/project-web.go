package main

import (
	"project-web/src/bootstrap/route"
	"project-web/src/bootstrap/service"
	"project-web/src/library/session"

	"github.com/kataras/iris"
	"github.com/kataras/iris/view"
	"github.com/pelletier/go-toml"
)

func main() {
	app := iris.New()

	setApplication(app)
}

func setApplication(app *iris.Application) {

	route.SetRoute(app)

	var port string

	service.BuildContainer()

	di := service.GetDi()
	container := di.Container
	container.Invoke(func(config *toml.Tree) {
		port = config.Get("app.port").(string)
	})
	container.Invoke(func(viewEngine *view.HTMLEngine) {
		app.RegisterView(viewEngine)
	})
	session.SetSessionWithRedis(app)
	app.Run(iris.Addr(port))

	
}
