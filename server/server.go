package server

import (
	"github.com/DeyiXu/go-web-example/server/http"
)

// Start 启动
func Start() {
	http.Start()

}

// Close 关闭
func Close() {
	http.Close()
}
