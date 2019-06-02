package main

import (
	"go-web/routes"
	"log"
	"net/http"
)

/*func init() {

	// 注册驱动
	_ = orm.RegisterDriver("mysql", orm.DRMySQL)
	// 设置默认数据库
	_ = orm.RegisterDataBase("default", "mysql", "homestead:secret@tcp(192.168.10.10:3306)/go-web?charset=utf8", 30)
	// 注册定义的model
	orm.RegisterModel(new(models.Userinfo), new(models.User), new(models.Profile), new(models.Tag), new(models.Post))

	// 创建table 将第三个参数改成true  会自动建表
	_ = orm.RunSyncdb("default", false, false)
	orm.Debug = true

}*/

func main() {

	routes.Route()

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
