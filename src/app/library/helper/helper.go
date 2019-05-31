package helper

import (
	"fmt"

	"github.com/pelletier/go-toml"
)

// Helper 工具类
type Helper struct{}

// NewHelper 初始化类
func NewHelper() *Helper {
	return &Helper{}
}

// NewConfig 初始化toml配置服务
func (h *Helper) NewConfig() *toml.Tree {
	rootPath := new(Helper).GetRootDirectory()
	configPath := rootPath + "/config/config.toml"

	// configPath := "../config/config.toml"
	config, err := toml.LoadFile(configPath)
	if err != nil {
		fmt.Println("Toml Error!", err.Error())
	}
	return config
}
