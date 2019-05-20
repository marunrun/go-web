package main

import (
	"github.com/astaxie/beego/orm"
	"go-web/Libs"
	"go-web/models"
	"go-web/routes"
	"log"
	"net/http"
)
var globalSessions *Libs.Manager

func init() {

	globalSessions, _ = Libs.NewManager("memory","gosessionid",3600)


	// 注册驱动
	_ = orm.RegisterDriver("mysql", orm.DRMySQL)
	// 设置默认数据库
	_ = orm.RegisterDataBase("default", "mysql", "homestead:secret@tcp(192.168.10.10:3306)/go-web?charset=utf8", 30)
	// 注册定义的model
	orm.RegisterModel(new(models.Userinfo), new(models.User), new(models.Profile), new(models.Tag), new(models.Post))

	// 创建table 将第三个参数改成true  会自动建表
	_ = orm.RunSyncdb("default", false, false)
	orm.Debug = true

}

func main() {

	//models.InsertUser()
	//models.CreatedUsers()
	//models.UpdateUser(1)
	//search := map[string]string{}
	//
	//search["username"] = "mr"
	//search["departname"] = "zxx"
	//models.SelectUser(search)
		//加载路由
		routes.Route()
		// 监听9090 端口
		err := http.ListenAndServe(":9090", nil)
		if err != nil {
			log.Fatal("ListenAndServe", err)
		}

}
