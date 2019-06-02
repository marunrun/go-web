package routes

import (
	"go-web/controllers"
	"net/http"
)

func Route() {

	http.HandleFunc("/",controllers.HomeIndex) // 设置访问的路由
	http.HandleFunc("/login", controllers.LoginIndex)
	http.HandleFunc("/upload", controllers.UploadIndex)
	http.HandleFunc("/message", controllers.GetMessage)
	http.HandleFunc("/ws", controllers.Handle)
}