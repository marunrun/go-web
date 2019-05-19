package models

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func test() {
	db, err := sql.Open("mysql", "homestead:secret@tcp(192.168.10.10:3306)/go-web?charset=utf8")
	checkErr(err)
	// 插入数据
	/*	stmt, err := db.Prepare("INSERT INTO userinfo SET username=?, department=?, created=?")
		checkErr(err)
		res, err := stmt.Exec("marun","技术部门","2019-05-19")
		checkErr(err)
		id, err := res.LastInsertId()
		checkErr(err)
		fmt.Println(id)*/

	// 更新数据
	/*	stmt, err := db.Prepare("UPDATE userinfo set username=? where uid=?")
		checkErr(err)
		res, err := stmt.Exec("mr", 1)
		checkErr(err)
		affect, err := res.RowsAffected()
		fmt.Println(affect)
	*/

	//查询数据
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)
	defer rows.Close()
	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}

}

func checkErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
