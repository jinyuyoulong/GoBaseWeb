// file: controller/user_controller.go

package controller

import (
	"io"
	"net/http"
	"os"

	"github.com/kataras/iris"
	// "github.com/kataras/iris/cache/client"
	"project-web/src/bootstrap/service"
	"project-web/src/models"

	"github.com/kataras/iris/mvc"
	"github.com/pelletier/go-toml"
)

// IndexController index controller
type IndexController struct {
	Ctx iris.Context
	Di  iris.Context
}

const (
	commonTitle string = "测试资料库"
)

// 访问API数据
func httpClient(url string) {
	url = "http://localhost:8081/api"
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	response, _ := client.Do(request)

	stdout := os.Stdout
	_, err = io.Copy(stdout, response.Body) // 输出打印
	println(response.StatusCode)
}

//  使用DI 中的数据
func useContainer() {
	// 测试外部文件 获取 config 数据
	container := service.GetDi().Container
	container.Invoke(func(config *toml.Tree) {
		println("测试外部文件 获取 config 数据", config.Get("database.dirver").(string))
	})
}

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
func (c *IndexController) GetIndexHandler(ctx iris.Context) {
	ctx.Writef("Hello from method: %s and path: %s", ctx.Method(), ctx.Path())
}
