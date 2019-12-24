package base

import (
	"net"
	"fmt"
)

// 客户端的部分
func CheckError2(err error) {
	if err != nil {
		//fmt.Println("Error : %s",err.Error())
		//os.Exit(1)
		panic(err)
	}
}

func ChatRoomClient()  {
	conn,err := net.Dial("tcp","127.0.0.1:8080")
	CheckError2(err)
	defer conn.Close()

	conn.Write([]byte("hello zhu"))

	fmt.Println("Has sent the message")
}