package service

import (
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris"
	"github.com/kataras/iris/view"
	"github.com/pelletier/go-toml"
	"go.uber.org/dig"
	"github.com/kataras/iris/sessions"
	"github.com/kataras/iris/sessions/sessiondb/redis"
	"github.com/kataras/iris/sessions/sessiondb/redis/service"
)

type Di struct {
	Container *dig.Container
}

var di *Di

// GetDI get
func GetDi() *Di {
	if di == nil {
		di = &Di{
			Container: dig.New(),
		}
	}
	return di
}

func AppConfig() *toml.Tree {
	file := "../config/config.toml"
	config, err := toml.LoadFile(file)
	if err != nil {
		fmt.Println("Toml Error!", err.Error())
	}
	return config
}

func viewEngine() *view.HTMLEngine {
	viewPath := "../view"
	layoutPath := "layout/layout.html"

	var htmlEngine *view.HTMLEngine
	htmlEngine = iris.HTML(viewPath, ".html").Layout(layoutPath)

	return htmlEngine.Reload(true)
}

func db() *xorm.Engine {
	//  读取配置文件的数据
	tomlC := AppConfig()
	driver := tomlC.Get("database.dirver").(string)
	configTree := tomlC.Get(driver).(*toml.Tree)
	userName := configTree.Get("username").(string)
	password := configTree.Get("password").(string)
	dbname := configTree.Get("dbname").(string)

	connet := fmt.Sprintf("%s:%s%s", userName, password, dbname)
	println(connet)
	engine, err := xorm.NewEngine(driver, connet)
	if err != nil {
		log.Fatal("database connet failed : %s", err)
	}
	return engine
}

func createSessions() *sessions.Sessions{
	db := redis.New(service.Config{
		Network:     "tcp",
		Addr:        "127.0.0.1:6379",
		Password:    "",
		Database:    "",
		MaxIdle:     0,
		MaxActive:   10,
		IdleTimeout: service.DefaultRedisIdleTimeout,
		Prefix:      ""}) // optionally configure the bridge between your redis server

	// close connection when control+C/cmd+C
	iris.RegisterOnInterrupt(func() {
		db.Close()
	})

	sess := sessions.New(sessions.Config{Cookie: "sessionscookieid", Expires: 45 * time.Minute})
	sess.UseDatabase(db)
	return sess
}

// BuildContainer 容器创建&注入
func BuildContainer() {
	container := GetDi().Container

	container.Provide(AppConfig)
	container.Provide(viewEngine)
	container.Provide(db)
	container.Provide(createSessions)
}
