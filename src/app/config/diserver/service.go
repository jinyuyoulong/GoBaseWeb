package diserver

// 包 bootstrap 负责都被标记为节拍请求或响应（即回复请求，但不会无限循环）。随机探测并线性扫描本地网络（单个接口）以用于其他正在运行的实例。

// 在每个扫描周期中，都会检查所有已配置的UDP端口（以防止因配置空间过大而导致速度下降）。

// UDP上的心跳，每一个都被标记为节拍请求或响应（即回复请求，但不会无限循环）。

import (
	"errors"
	"sync"

	"ielpm.cn/projectweb/src/app/library/helper"

	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/sessions"
	"github.com/kataras/iris/view"
	"go.uber.org/dig"
	"ielpm.cn/projectweb/src/app/config"
)

// DIserver is high-level compoment of this project
type DIserver struct {
	Container *dig.Container
	// App       *iris.Application
}

var (
	di   *DIserver
	once sync.Once
)

// NewDI 创建
func NewDI(app *iris.Application) *DIserver {
	once.Do(func() {
		di = &DIserver{
			Container: dig.New(),
			// App:       app,
		}
	})
	return di
}

// GetDI get
func GetDI() *DIserver {
	if di == nil {
		err := errors.New("di server not created! please new di server first")
		panic(err)
		return nil
	}
	return di
}

// BuildContainer 容器创建&注入
func BuildContainer(a *iris.Application) {
	container := NewDI(a).Container

	container.Provide(config.NewConfig)

	container.Provide(engineFunc) // 注入

	var appName string
	container.Invoke(func(config *config.Config) {
		appName = config.New().Get("app.name").(string)
	})
	container.Provide(setViewError)

}

func engineFunc() *view.HTMLEngine {
	viewsDir := new(helper.Helper).GetRootDirectory() + "/view"

	var sharedLayoutPath string
	sharedLayoutPath = "shared/layout.html"
	var htmlEngine *view.HTMLEngine
	htmlEngine = iris.HTML(viewsDir, ".html").Layout(sharedLayoutPath)
	htmlEngine.Reload(true)
	return htmlEngine
}
func setViewError(ctx iris.Context) {
	container := GetDI().Container
	var appName string
	container.Invoke(func(config *config.Config) {
		appName = config.New().Get("app.name").(string)
	})
	err := iris.Map{
		"app":     appName,
		"status":  ctx.GetStatusCode(),
		"message": ctx.Values().GetString("message"),
	}

	if jsonOutput := ctx.URLParamExists("json"); jsonOutput {
		ctx.JSON(err)
		return
	}
	ctx.ViewData("Err", err)
	ctx.ViewData("Title", "Error")
	ctx.View("share/error.html")
	// ctx.JSON(controller.ApiResult(false, nil, "404 not find"))
}

// NewServices 模块配置
func NewServices(a *iris.Application) {
	BuildContainer(a)

	dis := GetDI()

	container := dis.Container

	engineRegistFuc := func(htmlEngine *view.HTMLEngine) {
		a.RegisterView(htmlEngine)
	}

	container.Invoke(engineRegistFuc) // 召唤 使

	container.Invoke(func(handler context.Handler) {
		a.OnAnyErrorCode(handler)
	})

	setStaticFiles(a)
}

func setStaticFiles(a *iris.Application) {
	// static files
	staticassets := new(helper.Helper).GetRootDirectory() + "/public/"
	faviconPath := staticassets + "favicon.ico"
	// hero.Handler(func)
	a.Favicon(faviconPath)
	a.StaticWeb(staticassets[1:len(staticassets)-1], staticassets)
}

// 使用 session
func logout(ctx iris.Context, di *dig.Container) {
	di.Invoke(func(sessions *sessions.Sessions) {
		session := sessions.Start(ctx)
		session.Set("authenticated", false)
	})
}
