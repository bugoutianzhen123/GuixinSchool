package config

import (
	"github.com/spf13/viper"
)

type AppConfig struct {
	DB DBConf
}

type DBConf struct{
	Addr string 
}



func Load(cf *AppConfig, path string) error {
	v := viper.New()
	v.SetConfigFile(path)

	// 读取配置文件
	if err := v.ReadInConfig(); err != nil {
		return err
	}

	// 将配置绑定到结构体
	if err := v.Unmarshal(cf); err != nil {
		return err
	}

	return nil
}