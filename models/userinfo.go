package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type Userinfo struct {
	Uid        int `orm:"PK"` //如果表的主键不是id，那么需要加上pk注释，显式的说这个字段是主键
	Username   string
	Departname string
	Created    time.Time
}

// 单条插入数据
func InsertUser() {
	o := orm.NewOrm()
	var user Userinfo
	user.Username = "mr"
	user.Departname = "zxx"

	id, err := o.Insert(&user)
	if err == nil {
		fmt.Println(id)
	}
}

// 批量插入数据
func CreatedUsers() {
	o := orm.NewOrm()
	users := []Userinfo{
		{Username: "marun"},
		{Username: "runrun"},
	}

	successNums, err := o.InsertMulti(5, users)
	if err == nil {
		fmt.Println(successNums)
	}
}

// 更新数据
func UpdateUser(uid int)  {
	o := orm.NewOrm()
	user := Userinfo{Uid:uid}

	if o.Read(&user) == nil{
		user.Username = "test"
		user.Created = time.Now()
		// o.Update(&user,"Field1","Field2") 只更新某几个字段
		//Where:用来设置条件，支持多个参数，第一个参数如果为整数，相当于调用了Where("主键=?",值)。

		if num, err := o.Update(&user,"Username","Created"); err == nil{
			fmt.Println(num)
		}
	}
}

func SelectUser(user Userinfo )  {

	o := orm.NewOrm()

	if uid := user.Uid; uid == 0 {
		qs := o.QueryTable(user) // 返回 QuerySeter
		_ = qs.Filter("username", user.Username).One(&user)
	}


	//o := orm.NewOrm()
	//err := o.Read(&user)
	//
	//if err == orm.ErrNoRows {
	//	fmt.Println("查询不到")
	//} else if err == orm.ErrMissPK {
	//	fmt.Println("找不到主键")
	//} else {
	//	fmt.Println(user.Uid, user.Username)
	//}

}