package main

import (
	"os"
	"os/signal"

	"github.com/DeyiXu/go-web-example/model"
	"github.com/DeyiXu/go-web-example/module/config"
	"github.com/DeyiXu/go-web-example/module/store"
	"github.com/DeyiXu/go-web-example/server"
	"github.com/DeyiXu/go-web-example/service"
)

func init() {
	config.Init("./config.toml")
}
func start() {
	server.Start()
	store.Start()
	model.AutoMigrate()

	service.Start()
}
func close() {
	server.Close()
	store.Close()

	service.Close()
}
func main() {
	start()

	// 等待中断信号优雅地关闭服务器
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	defer close()
}
