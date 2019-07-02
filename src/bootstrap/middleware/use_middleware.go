package middleware

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
)

// RegistMiddleware 注册中间件
func RegistMiddleware(a *iris.Application) {
	a.Use(logger.New())
}
