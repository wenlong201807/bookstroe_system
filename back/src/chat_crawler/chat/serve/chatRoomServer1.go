package main

import (
	"net"
	"fmt"
)

// 服务端的部分
func CheckError(err error) {
	if err != nil {
		//fmt.Println("Error : %s",err.Error())
		//os.Exit(1)
		panic(err)
	}
}

func ProcesssInfo(conn net.Conn) {
	buf := make([]byte, 1024)
	defer conn.Close()

	for {
		numOfBytes, err := conn.Read(buf)
		//CheckError(err)
		if err != nil {
			continue
		}

		// 接收到的字节数量大于0
		if numOfBytes != 0 { // 接收到客户端的消息时进入
			remoteAddr := conn.RemoteAddr() // 获取客户端的地址
			fmt.Printf("%v 说: %s",remoteAddr, string(buf[0:numOfBytes]))
			fmt.Println("") //换行使用
		}

	}
}

func main() {
	listen_socket, err := net.Listen("tcp", "127.0.0.1:8080")
	CheckError(err)
	defer listen_socket.Close() // 避免资源浪费

	fmt.Println("Server is waiting...")
	for { // 死循环
		conn, err := listen_socket.Accept()
		CheckError(err)
		go ProcesssInfo(conn) // 具体连接在协程中进行，比线程轻量
	}

}
