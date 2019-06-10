package session

import (
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/sessions"
	"github.com/kataras/iris/sessions/sessiondb/redis"
	"github.com/kataras/iris/sessions/sessiondb/redis/service"
	bservice "project-web/src/bootstrap/service"
)


func SetSessionWithRedis(app *iris.Application) {
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
}

func SessionGet(ctx context.Context,key string) string{
	var value string
	container := bservice.GetDi().Container
	container.Invoke( func(sess *sessions.Sessions){
 		value = sess.Start(ctx).GetString(key)
	})
	return value
}
func SessionSet(ctx context.Context,key string,value string){
	container := bservice.GetDi().Container
	container.Invoke( func(sess *sessions.Sessions){
		sess.Start(ctx).Set(key,value)
	})
}

func SessionDelete(ctx context.Context,key string){
	container := bservice.GetDi().Container
	container.Invoke( func(sess *sessions.Sessions){
 		sess.Start(ctx).Delete(key)
	})
}