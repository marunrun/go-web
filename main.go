package main

import (
	"go-web/controllers"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
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

	arith := new(controllers.Arith)
	rpc.Register(arith)


	tcpAddress, err := net.ResolveTCPAddr("tcp",":1234")
	controllers.CheckError(err)

	listener, err := net.ListenTCP("tcp",tcpAddress)
	controllers.CheckError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		jsonrpc.ServeConn(conn)
	}

/*	//加载路由
	routes.Route()
	// 监听9090 端口
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
*/
}
