package config

import (
	"log"

	"github.com/nilorg/pkg/logger"
	"github.com/spf13/viper"
)

// Init 初始化配置文件
func Init(in string) {
	initConfigFile(in)
	initLog()
}

func initConfigFile(in string) {
	viper.SetConfigFile(in)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("读取配置文件错误：%s", err.Error())
	}
}
func initLog() {
	// 日志初始化
	logger.Init()
	// logger.RegisterSentry()
}
