package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	defer fmt.Println("defer main")
	var user = os.Getenv("USER_")
	go func() {
		defer func() {
			fmt.Println("defer caller")
			if err := recover(); err != nil{
				fmt.Println("recover success.")
			}
		}()

		if user == "" {
			panic("should set user env.")
		}

		fmt.Println("after panic")

	}()

	time.Sleep(1 * time.Second)
	fmt.Printf("get result %d \r\n",user)
}