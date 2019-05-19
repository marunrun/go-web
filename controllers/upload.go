package controllers

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

func UploadIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		_, _ = io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("views/upload.html")
		_ = t.Execute(w, token)
	} else {
		// 这里一定要记得 r.ParseMultipartForm(), 否则 r.MultipartForm 是空的
		// 调用 r.FormFile() 的时候会自动执行 r.ParseMultipartForm()
		_ = r.ParseMultipartForm(32 << 20)
		// 写明缓冲的大小。如果超过缓冲，文件内容会被放在临时目录中，而不是内存。过大可能较多占用内存，过小可能增加硬盘 I/O
		// FormFile() 时调用 ParseMultipartForm() 使用的大小是 32 << 20，32MB
		file, fileHeader, err := r.FormFile("file") // file 是上传表单域的名字
		if err!= nil {
			fmt.Println("get upload file fail:",err)
			w.WriteHeader(500)
			return
		}
		defer file.Close() // 此时上传内容的 IO 已经打开，需要手动关闭！！
		_, _ = fmt.Fprintf(w, "%v", fileHeader.Header)
		// 打开目标地址，把上传的内容存进去
		f,err := os.OpenFile("./public/upload/"+fileHeader.Filename,os.O_WRONLY|os.O_CREATE,0666)
		if err != nil {
			fmt.Println("save upload file fail:", err)
			w.WriteHeader(500)
			return
		}
		defer f.Close()
		_, _ = io.Copy(f, file)
		_, _ = w.Write([]byte("upload file:" + fileHeader.Filename + " - saveto : ./public/upload/"))
	}
}
