package main

import (
	"github.com/kataras/iris"
	"ielpm.cn/projectweb/src/app/bootstrap/route"
	"ielpm.cn/projectweb/src/app/config"
	"ielpm.cn/projectweb/src/app/config/diserver"
)

func main() {
	app := iris.New()
	diserver.NewServices(app)
	route.Configure(app)
	port := new(config.Config).New().Get("app.port").(string)
	app.Run(iris.Addr(port))
}
