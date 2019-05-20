package Libs

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Server struct {
	ServerName string
	ServerIp   string
}

type ServerSlice struct {
	Servers []Server
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
