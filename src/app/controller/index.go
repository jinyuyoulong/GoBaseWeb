// file: controller/user_controller.go

package controller

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/kataras/iris"
	// "github.com/kataras/iris/cache/client"

	"github.com/kataras/iris/mvc"

	"xxx.com/projectweb/src/app/bootstrap/service"
	"xxx.com/projectweb/src/app/config"
	"xxx.com/projectweb/src/app/config/diserver"
	"xxx.com/projectweb/src/app/library/datasource"
	"xxx.com/projectweb/src/app/models"
)

// IndexController index controller
type IndexController struct {
	Ctx iris.Context
}

const (
	commonTitle string = "测试资料库"
)

// 访问API数据
func httpClient(url string) {

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

// Get url: /
func (c *IndexController) Get() mvc.Result {
	httpClient("http://localhost:8081/api")

	Service := service.NewprojectapiService()
	datalist := Service.GetAll()

	// 测试外部文件 获取 config 数据
	container := diserver.GetDI().Container
	container.Invoke(func(config *config.Config) {
		println("测试外部文件 获取 config 数据", config.New().Get("database.dirver").(string))
	})

	return mvc.View{
		Name: "index" + filefix,
		Data: iris.Map{
			"Title":    commonTitle,
			"Datalist": datalist,
		},
	}
}

// GetBy url: /{id}
func (c *IndexController) GetBy(id int) mvc.Result {
	Service := service.NewprojectapiService()
	if id < 0 {
		return mvc.Response{
			Path: "/",
		}
	}
	data := Service.Get(id)
	value := c.Ctx.GetCookie("name")
	println("cookie :", value)

	return mvc.View{
		Name: "info" + filefix,
		Data: iris.Map{
			"Title": commonTitle,
			"info":  data,
		},
	}
}

// GetSearch uri: /search?country=china
func (c *IndexController) GetSearch() mvc.Result {
	Service := service.NewprojectapiService()
	country := c.Ctx.URLParam("country")
	datalist := Service.Search(country)
	return mvc.View{
		Name: "index" + filefix, // index.html
		Data: iris.Map{
			"Title":    commonTitle,
			"Datalist": datalist,
		},
	}
}

func (c *IndexController) GetIndexHandler(ctx iris.Context) {
	ctx.Writef("Hello from method: %s and path: %s", ctx.Method(), ctx.Path())
}

// GetClearcache 手动清除缓存
// uri: /clearcache
func (c *IndexController) GetClearcache() mvc.Result {
	err := datasource.InstanceMaster().ClearCache(&models.StarInfo{})
	if err != nil {
		log.Fatal(err)
	}

	return mvc.Response{
		Text: "xorm 缓存清除成功",
	}
}
