package diserver

import (
	"errors"
	"sync"

	"project-web/src/app/library/helper"

	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/sessions"
	"github.com/kataras/iris/view"
	"go.uber.org/dig"
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

	container.Provide(helper.NewHelper)
	container.Provide(engineFunc) // 注入
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
	container.Invoke(func(helper *helper.Helper) {
		appName = helper.NewConfig().Get("app.name").(string)
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
