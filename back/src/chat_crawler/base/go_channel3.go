package base

import (
	"fmt"
	"time"
)

var ch chan int // 全局变量
func test_channel3() {

	ch <- 1 // 写入
	fmt.Println("chan 1")
	ch <- 1 // 写入
	fmt.Println("chan 2")
	ch <- 1 // 写入
	fmt.Println("chan 3")
	fmt.Println("come to end gorountine 1")
}

func Go_channel3() {
	ch = make(chan int,2) // 默认缓冲区大小为0
	//ch = make(chan int) // 默认缓冲区大小为0
	//ch = make(chan int,0) // 默认缓冲区大小为0
	go test_channel3() // 开始执行协程函数**有读写操作，一次，需要有读取结果之后才能输出执行结果
	// 此处睡眠，影响通道函数内部的代码执行，与之后的代码的执行的先后顺序
	//time.Sleep(2 * time.Second)
	fmt.Println("running end!")

	<-ch // 读取 // 读取操作完成，才执行协程对应函数的结果输出***在这一行

	fmt.Println("读写协程之后的操作在这之后或者之前，取决于是否有够大的缓冲区，同时要保证在主进程退出之前。。")
	// 保证主进程最后退出
	time.Sleep(2 * time.Second)
	fmt.Println("主进程退出")
}
