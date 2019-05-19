package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"go-web/models"
)

func init() {
	// 注册驱动
	_ = orm.RegisterDriver("mysql", orm.DRMySQL)
	// 设置默认数据库
	_ = orm.RegisterDataBase("default", "mysql", "homestead:secret@tcp(192.168.10.10:3306)/go-web?charset=utf8", 30)
	// 注册定义的model
	orm.RegisterModel(new(models.Userinfo), new(models.User), new(models.Profile), new(models.Tag), new(models.Post))

	// 创建table
	_ = orm.RunSyncdb("default", false, true)
	orm.Debug = true

}

func main() {

	insertUser()
	/*	//加载路由
		routes.Route()
		// 监听9090 端口
		err := http.ListenAndServe(":9090", nil)
		if err != nil {
			log.Fatal("ListenAndServe", err)
		}*/
}

func insertUser() {
	o := orm.NewOrm()
	var user models.Userinfo
	user.Username = "mr"
	user.Departname = "zxx"

	id, err := o.Insert(&user)
	if err == nil {
		fmt.Println(id)
	}
}
