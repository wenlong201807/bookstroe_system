package base

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
		CheckError(err)

		if numOfBytes != 0 {
			fmt.Println("Has received this message: %s", string(buf))
		}

	}
}

func ChatRoom() {
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
