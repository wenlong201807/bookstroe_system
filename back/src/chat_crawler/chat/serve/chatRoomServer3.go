package main

import (
	"net"
	"fmt"
	"strings"
	"os"
	"log"
)

// 常量
const LOG_DIRECTORY = "./test.log"

// 定义，并声明变量类型，两步合在一起
var onlineConns3 = make(map[string]net.Conn) // IP地址与其发送的内容做一个映射，关联起来
var messageQueue3 = make(chan string, 1000)
var quitChan3 = make(chan bool)
var logFile *os.File
var logger *log.Logger
// 服务端的部分
func CheckError3(err error) {
	if err != nil {
		//fmt.Println("Error : %s",err.Error())
		//os.Exit(1)
		panic(err)
	}
}

func ProcesssInfo3(conn net.Conn) {
	buf := make([]byte, 1024)
	defer func(conn net.Conn) {
		// 退出之前，删除原有连接
		addr := fmt.Sprintf("%s", conn.RemoteAddr())
		delete(onlineConns3, addr)
		conn.Close()

		for i := range onlineConns3 {
			fmt.Println("now online conns:" + i)
		}
	}(conn)

	for {
		numOfBytes, err := conn.Read(buf)
		if err != nil {
			break
		}

		// 接收到的字节数量大于0
		if numOfBytes != 0 { // 接收到客户端的消息时进入
			message := string(buf[0:numOfBytes])
			messageQueue3 <- message
		}

	}
}

func ConsumeMessage3() {
	for {
		select {
		case message := <-messageQueue3:
			// 对消息进行解析
			doProcessMessage3(message)
		case <-quitChan3:
			break

		}
	}
}

func doProcessMessage3(message string) {
	// 127.0.0.1#您好
	contents := strings.Split(message, "#")

	if len(contents) > 1 {
		addr := contents[0]
		sendMessage := strings.Join(contents[1:], "#")

		// 处理空格
		addr = strings.Trim(addr, " ")

		if conn, ok := onlineConns3[addr]; ok {
			_, err := conn.Write([]byte(sendMessage))
			if err != nil {
				fmt.Println("online conns send failure!")
			}
		}
	} else {
		contents := strings.Split(message, "*")
		if strings.ToUpper(contents[1]) == "LIST" {
			var ips string = ""
			for i := range onlineConns3 {
				ips = ips + "|" + i
			}
			if conn, ok := onlineConns3[contents[0]]; ok {
				_, err := conn.Write([]byte(ips))
				if err != nil {
					fmt.Println("online conns send failure!")
				}
			}
		}
	}
}

func main() {
	logFile, err := os.OpenFile(LOG_DIRECTORY, os.O_RDWR|os.O_CREATE, 0)
	if err != nil {
		fmt.Println("log file create failure!")
		os.Exit(-1)
	}
	defer logFile.Close()
	// 日志格式
	logger = log.New(logFile, "\r\n", log.Ldate|log.Ltime|log.Llongfile)

	listen_socket, err := net.Listen("tcp", "127.0.0.1:8080")
	CheckError3(err)
	defer listen_socket.Close() // 避免资源浪费

	fmt.Println("Server is waiting...")
	logger.Println("i  am writting the logs...")
	go ConsumeMessage3() // 用来消费消息队列里面的东西
	for { // 死循环
		conn, err := listen_socket.Accept()
		CheckError3(err)

		// 将conn存储到onlineConns映射表中
		addr := fmt.Sprintf("%s", conn.RemoteAddr())
		onlineConns3[addr] = conn

		for i, v := range onlineConns3 {
			fmt.Println("连接中的有：", i, v)
		}
		go ProcesssInfo3(conn) // 具体连接在协程中进行，比线程轻量
	}

}
