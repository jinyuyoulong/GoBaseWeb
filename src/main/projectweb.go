package main

import (
	"github.com/kataras/iris"
	"xxx.com/projectweb/src/app/bootstrap/route"
	"xxx.com/projectweb/src/app/config"
	"xxx.com/projectweb/src/app/config/diserver"
)

func main() {
	app := iris.New()
	diserver.NewServices(app)
	route.Configure(app)
	port := new(config.Config).New().Get("app.port").(string)
	app.Run(iris.Addr(port))
}
