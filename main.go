package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()               // 解析参数
	fmt.Println(r.Form)             //这里打印的信息是在服务端控制台输出的
	fmt.Println("path", r.URL.Path) //
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("value:", strings.Join(v, ""))
	}
	_, _ = fmt.Fprintf(w, "hello marun!")
}

func ReturnJson(code int, msg string, w http.ResponseWriter) {
	Errors := make(map[string]interface{})

	Errors["code"] = code
	Errors["msg"] = msg
	res, _ := json.Marshal(Errors)
	_, _ = fmt.Fprint(w, string(res))

}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method", r.Method)
	_ = r.ParseForm()

	v := url.Values{}
	v.Set("marun", "123")
	v.Add("friend", "mayun")
	v.Add("friend", "mahuateng")
	fmt.Println(v.Encode())

	if r.Method == "GET" {
		t, _ := template.ParseFiles("./views/login.html")
		log.Println(t.Execute(w, nil))
	} else {

		if len(r.Form["username"][0]) == 0 {
			log.Println(r.Form["username"][0])
			ReturnJson(400, "用户名必填", w)
			return
		}




		if m, _ := regexp.MatchString("^[0-9]+$", r.Form.Get("age")); !m {
			ReturnJson(400,"年龄必须为数字",w)
			return
		}

		age, _ := strconv.Atoi(r.Form.Get("age"))

		if age > 100 {
			log.Println(age)
			ReturnJson(400, "年龄太大", w)
			return
		}



	}
}

func main() {
	http.HandleFunc("/", sayHelloName) // 设置访问的路由
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
