package helper

import (
	"log"
	"os"
	"path"
)

// Helper 工具
type Helper struct{}

// NewHelper 初始化
func NewHelper() *Helper {
	return &Helper{}
}

// GetRootDirectory 返回程序开始执行(命令执行所在目录，比如 shall 执行所在的目录) 所在目录的上一级目录
// 例：/Users/xxx/dev/go/projectweb ps.execute 文件在 projectweb/bin 目录下
func (h *Helper) GetRootDirectory() string {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return path.Dir(wd)
}
