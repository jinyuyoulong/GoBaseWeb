package models

import (
	"project-web/src/bootstrap/service"

	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

func init() {
	if engine == nil {
		container := service.GetDi().Container
		container.Invoke(func(db *xorm.Engine) {
			engine = db
		})
	}
}
