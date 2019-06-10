package session

import (
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/sessions"
	bservice "project-web/src/bootstrap/service"
)

func SessionGet(ctx context.Context,key string) string{
	var value string
	bservice.GetDi().Container.Invoke( func(sess *sessions.Sessions){
 		value = sess.Start(ctx).GetString(key)
	})
	return value
}
func SessionSet(ctx context.Context,key string,value string){
	bservice.GetDi().Container.Invoke( func(sess *sessions.Sessions){
		sess.Start(ctx).Set(key,value)
	})
}

func SessionDelete(ctx context.Context,key string){
	bservice.GetDi().Container.Invoke( func(sess *sessions.Sessions){
 		sess.Start(ctx).Delete(key)
	})
}