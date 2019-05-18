package controllers

import (
	"crypto/md5"
	"fmt"
	"go-web/Libs"
	"html/template"
	"io"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

func LoginIndex(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()

	if r.Method == "GET" {
		// 生成时间戳
		crutime := time.Now().Unix()
		h := md5.New()
		// md5 加密
		_, _ = io.WriteString(h, strconv.FormatInt(crutime, 10))

		token := fmt.Sprintf("%x",h.Sum(nil))

		v, _ := template.ParseFiles("views/login.html")
		_ = v.Execute(w, token)
	} else {
		fmt.Println(r.Form)
		_ = r.ParseForm()
		token := r.Form.Get("token")
		if token != "" {

		} else {

		}


		//fruit := r.Form.Get("fruit")
		//
		//fruits := []interface{}{"apple","pear","banana"}
		//if m := Libs.In_slice(fruit,fruits); !m{
		//	ReturnJson(400, "水果选择有误", w)
		//	return
		//}
		t, err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
		err = t.ExecuteTemplate(w, "T", template.HTML("<script>alert('you have been pwned')</script>"))
		if err != nil {
			log.Print(err)
		}
		fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username"))) //输出到服务器端
		fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
		template.HTMLEscape(w, []byte(r.Form.Get("username"))) //输出到客户端
		if len(r.Form["username"][0]) == 0 {
			Libs.ReturnJson(400, "用户名必填", w)
			return
		}

		//if m,_ := regexp.MatchString("^\\p{Han}+$",r.Form.Get("username")); !m {
		//	ReturnJson(400,"用户名必须是中文",w)
		//	return
		//}
		if m,_ := regexp.MatchString(`^[\w\.\_]{2,10}@(\w{1,})\.([a-z]{2,4})$`,r.Form.Get("username")); !m{
			Libs.ReturnJson(400,"请填写正确的邮箱地址",w)
			return
		}


		if m, _ := regexp.MatchString("^[0-9]+$", r.Form.Get("age")); !m {
			Libs.ReturnJson(400,"年龄必须为数字",w)
			return
		}

		age, _ := strconv.Atoi(r.Form.Get("age"))

		if age > 100 {
			Libs.ReturnJson(400, "年龄太大", w)
			return
		}

	}
}
