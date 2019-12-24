package main

import (
	"net"
	"fmt"
	"strings"
)

// 定义，并声明变量类型，两步合在一起
var onlineConns  = make(map[string]net.Conn) // IP地址与其发送的内容做一个映射，关联起来
var messageQueue = make(chan string, 1000)
var quitChan = make(chan bool)
// 服务端的部分
func CheckError2(err error) {
	if err != nil {
		//fmt.Println("Error : %s",err.Error())
		//os.Exit(1)
		panic(err)
	}
}

func ProcesssInfo2(conn net.Conn) {
	buf := make([]byte, 1024)
	defer conn.Close()

	for {
		numOfBytes, err := conn.Read(buf)
		//CheckError2(err)
		if err != nil {
			continue
		}

		// 接收到的字节数量大于0
		if numOfBytes != 0 { // 接收到客户端的消息时进入
			//remoteAddr := conn.RemoteAddr() // 获取客户端的地址
			//fmt.Printf("%v 说: %s",remoteAddr, string(buf[0:numOfBytes]))
			//fmt.Println("") //换行使用

			message := string(buf[0:numOfBytes])
			messageQueue <- message

		}

	}
}

func ConsumeMessage() {
	for {
		select {
		case message := <-messageQueue:
			// 对消息进行解析
			doProcessMessage(message)
		case <- quitChan:
			break

		}
	}
}

func doProcessMessage(message string)  {
	// 127.0.0.1#您好
	contents := strings.Split(message,"#")
	if len(contents) > 1{
		addr := contents[0]
		sendMessage := contents[1]

		// 处理空格
		addr = strings.Trim(addr," ")

		if conn,ok := onlineConns[addr];ok{
			_,err := conn.Write([]byte(sendMessage))
			if err != nil{
				fmt.Println("online conns send failure!")
			}
		}
	}
}

func main() {
	listen_socket, err := net.Listen("tcp", "127.0.0.1:8080")
	CheckError2(err)
	defer listen_socket.Close() // 避免资源浪费

	fmt.Println("Server is waiting...")

	go ConsumeMessage() // 用来消费消息队列里面的东西
	for { // 死循环
		conn, err := listen_socket.Accept()
		CheckError2(err)

		// 将conn存储到onlineConns映射表中
		addr := fmt.Sprintf("%s",conn.RemoteAddr())
		onlineConns[addr] = conn

		for i,v := range onlineConns{
			fmt.Println("连接中的有：",i,v)
		}
		go ProcesssInfo2(conn) // 具体连接在协程中进行，比线程轻量
	}

}
