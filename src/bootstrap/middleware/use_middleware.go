package middleware

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
)

// RegistMiddleware 注册中间件
func RegistMiddleware(a *iris.Application) {
	a.Use(logger.New())
}
