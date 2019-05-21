package Libs

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Server struct {
	ServerName string `json:"serverName"`
	ServerIp   string `json:"serverIp"`
}

type ServerSlice struct {
	Servers []Server `json:"servers"`
}

func ReturnJson(code int, msg string, w http.ResponseWriter) {
	Errors := make(map[string]interface{})

	Errors["code"] = code
	Errors["msg"] = msg
	res, _ := json.Marshal(Errors)
	_, _ = fmt.Fprint(w, string(res))

}

func JsonTest()  {
	var s ServerSlice

	s.Servers = append(s.Servers, Server{ServerName: "Shanghai_VPN", ServerIp: "127.0.0.1"})
	s.Servers = append(s.Servers, Server{ServerName: "Beijing_VPN", ServerIp: "127.0.0.2"})
	b, err := json.Marshal(s)
	if err != nil{
		fmt.Println("json err",err)
	}

	fmt.Println(string(b))

	str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	_ = json.Unmarshal([]byte(str), &s)
	fmt.Println(s)
	b = []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	var f interface{}
	_ = json.Unmarshal(b,&f)
	fmt.Println(f)

}

func JsonTag()  {
	type Server struct {
		// ID 不会导出到JSON中
		ID int `json:"-"`

		// ServerName2 的值会进行二次JSON编码
		ServerName  string `json:"serverName"`
		ServerName2 string `json:"serverName2,string"`

		// 如果 ServerIP 为空，则不输出到JSON串中
		ServerIP   string `json:"serverIP,omitempty"`
	}

	s := Server {
		ID:         3,
		ServerName:  `Go "1.0" `,
		ServerName2: `Go "1.0" `,
		ServerIP:   ``,
	}
	b, _ := json.Marshal(s)
	_, _ = os.Stdout.Write(b)
	var str Server
	_ = json.Unmarshal([]byte(b), &str)

	fmt.Println(str)

}