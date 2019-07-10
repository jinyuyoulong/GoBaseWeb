package session

import (
	bservice "project-web/src/bootstrap/service"

	"github.com/kataras/iris/context"
	"github.com/kataras/iris/sessions"
)

var thissession *sessions.Sessions

func init() {
	bservice.GetDi().Container.Invoke(func(sess *sessions.Sessions) {
		thissession = sess
	})
}

// Set 增，改
func Set(ctx context.Context, key string, value string) {
	thissession.Start(ctx).Set(key, value)
}

// Get 查
func Get(ctx context.Context, key string) string {
	var value string
	value = thissession.Start(ctx).GetString(key)
	return value
}

// Delete 删
func Delete(ctx context.Context, key string) {
	thissession.Start(ctx).Delete(key)
}
