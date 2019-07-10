// file: controller/user_controller.go

package controller

import (
	"fmt"
	"project-web/src/library/session"
	"project-web/src/models"

	"github.com/kataras/iris"
	"github.com/kataras/iris/hero"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
)

// IndexController index controller
type IndexController struct {
	Ctx iris.Context
	Di  iris.Context
}

// 访问API数据
// func httpClient(url string) {
// 	url = "http://localhost:8081/api"
// 	client := &http.Client{}
// 	request, err := http.NewRequest("GET", url, nil)
// 	if err != nil {
// 		panic(err)
// 	}
// 	response, _ := client.Do(request)

// 	stdout := os.Stdout
// 	_, err = io.Copy(stdout, response.Body) // 输出打印
// 	println(response.StatusCode)
// }

// Get url: /
func (c *IndexController) Get() mvc.Result {
	datalist := models.CreateStrInfo().GetAll()

	return mvc.View{
		Name: "index/index.html",
		Data: iris.Map{
			"Title":    commonTitle,
			"Datalist": datalist,
		},
	}
}

// GetBy url: /{id}
func (c *IndexController) GetBy(id int) mvc.Result {
	starinfo := models.CreateStrInfo()
	if id < 0 {
		return mvc.Response{
			Path: "/",
		}
	}
	data := starinfo.GetStarInfoInfo(id)
	value := c.Ctx.GetCookie("name")
	println("cookie :", value)

	return mvc.View{
		Name: "index/info.html",
		Data: iris.Map{
			"Title": commonTitle,
			"info":  data,
		},
	}
}

// GetIndexHandler url: /setroute
func (c *IndexController) GetIndexHandler(context iris.Context) {
	context.Writef("Hello from method: %s and path: %s", context.Method(), context.Path())
}

// Redis Session

// GetSet /set set session in redis
func (c *IndexController) GetSet(context iris.Context) {
	name := c.Ctx.URLParam("name")
	if name == "" {
		name = "iris"
	}
	session.Set(context, "name", name)
	context.Writef("写入 session 数据 name:%s", name)
}

// GetSession /session 获取session
func (c *IndexController) GetSession(context iris.Context, session *sessions.Session) {
	visits := session.Increment("visits", 1)

	// name := session.SessionGet(context, "name")
	context.Writef("%d visit(s) from my current session", visits)
}

// Native Session

//Hello 默认 hero
func Hello() hero.Result {
	return hero.View{
		Name: "index/hello.html",
		Data: map[string]interface{}{
			"Title":     "Hello Page",
			"MyMessage": "Welcome to my awesome website",
		},
	}
}

// Session 默认session
func Session(context iris.Context, session *sessions.Session) string {
	visits := session.Increment("visits", 1)
	return fmt.Sprintf("%d visit(s) from my current session", visits)
}
