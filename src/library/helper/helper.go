package helper

import (
	"encoding/json"
	"io/ioutil"
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

type ConfigObject struct {
}

func Config(configPath string, object *interface{}) {
	//ReadFile函数会读取文件的全部内容，并将结果以[]byte类型返回
	data, _ := ioutil.ReadFile(configPath)
	//读取的数据为json格式，需要进行解码
	_ = json.Unmarshal(data, object)
}
