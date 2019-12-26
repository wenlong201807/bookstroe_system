package main

import (
	"net"
	"fmt"
	"bufio"
	"os"
	"strings"
)

// 客户端的部分
func CheckError3(err error) {
	if err != nil {
		fmt.Printf("Error : %s",err.Error())
		os.Exit(1)
		//panic(err)
	}
}
func MessageSend3(conn net.Conn) {
	var input string
	for {
		reader := bufio.NewReader(os.Stdin)
		data, _, _ := reader.ReadLine()
		input = string(data)

		if strings.ToUpper(input) == "EXIT" {
			conn.Close()
			break
		}

		_, err := conn.Write([]byte(input))
		if err != nil {
			conn.Close()
			fmt.Println("client connect failure : " + err.Error())
			break
		}
	}
}
func main() {
	// 连接服务器ip并实现多个客户端可以通信
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	CheckError3(err)
	defer conn.Close()
	//fmt.Println("我自己的IP和端口是：",conn.RemoteAddr())

	//conn.Write([]byte("hello zhu"))
	go MessageSend3(conn)

	// 客户端读取数据
	buf := make([]byte, 1024)
	for {
		_, err := conn.Read(buf)
		//CheckError3(err)
		if err !=nil{
			fmt.Println("您已经退出，欢迎再次使用文龙版通信系统，祝您生活愉快！")
			os.Exit(0) // 正常退出
		}
		fmt.Println("receive server message content:" + string(buf))
	}
	fmt.Println("client program end!")
}
