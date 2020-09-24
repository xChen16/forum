package main

import (
	"log"
	"net/http"

	"github.com/forum/config"
	"github.com/forum/routes"
)

func main() {
	startWebServer()
}

// 通过指定端口启动 Web 服务器
func startWebServer() {
	config := config.LoadConfig()
	r := routes.NewRouter()
	// 处理静态资源文件
	assets := http.FileServer(http.Dir(config.App.Static))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", assets))

	http.Handle("/", r) // 通过 router.go 中定义的路由器来分发请求

	log.Println("Starting HTTP service at " + config.App.Address)
	err := http.ListenAndServe(config.App.Address, nil) // 启动协程监听请求

	if err != nil {
		log.Println("An error occured starting HTTP listener at " + config.App.Address)
		log.Println("Error: " + err.Error())
	}
}
