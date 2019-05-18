package main

import (
	"go-web/routes"
	"log"
	"net/http"
)

func main() {
	// 加载路由
	routes.Route()
	// 监听9090 端口
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
