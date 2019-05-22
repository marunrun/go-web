package controllers

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strings"
	"time"
)

func IpTest() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s ip-addr\n", os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]
	addr := net.ParseIP(name)
	if addr == nil {
		fmt.Println("Invalid address")
	} else {
		fmt.Println("The address is ", addr.String())
	}
	os.Exit(0)
}

func ConnHttp() {
	if len(os.Args) != 2 {
		_, _ = fmt.Fprint(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	CheckError(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	CheckError(err)

	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	CheckError(err)

	result, err := ioutil.ReadAll(conn)

	CheckError(err)
	fmt.Println(string(result))
	os.Exit(0)
}

func ListenPort() {
	service := ":1200"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	CheckError(err)
	fmt.Println(tcpAddr)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	CheckError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	conn.SetReadDeadline(time.Now().Add(2 * time.Minute))
	request := make([]byte, 128)
	defer conn.Close()
	var str string

	for {
		read_len, err := conn.Read(request)
		if err != nil {
			fmt.Println(err)
			break
		}
		str += string(request[:read_len])
		if string(request[:read_len]) == "\r\n" && strings.TrimSpace(str) != ""{
			conn.Write([]byte(strings.TrimSpace(str)+"\r\n"))
			str = ""
		}

		if string(request[:read_len]) == "\r\n" && strings.TrimSpace(str) == "q"{
			conn.Close()
		}

		request = make([]byte, 128) // clear last read content
	}
}




func CheckError(err error) {
	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
