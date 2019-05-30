package config

import (
	"fmt"

	"github.com/pelletier/go-toml"
	"xxx.com/projectweb/src/app/library/helper"
)

type Config struct{}

func NewConfig() *Config {
	return &Config{}
}

// New 初始化toml配置服务
func (c *Config) New() *toml.Tree {
	rootPath := new(helper.Helper).GetRootDirectory()
	configPath := rootPath + "/config/config.toml"

	// configPath := "../config/config.toml"
	config, err := toml.LoadFile(configPath)
	if err != nil {
		fmt.Println("Toml Error!", err.Error())
	}
	return config
}
