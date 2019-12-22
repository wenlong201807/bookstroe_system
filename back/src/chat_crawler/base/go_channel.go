package base

import (
	"fmt"
	"strconv"
	"time"
)

// go语言在语言级别提供的goroutine间的通信方式

// 阻塞  除非有goroutine 对其进行操作
//ch <- c 写
// c:= <- ch 读

func Read(ch chan int)  {
	value := <- ch
	fmt.Println("value",strconv.Itoa(value))
}
func Write(ch chan int)  {
	ch <- 10
}

func Go_channel()  {
	// 必须要先获取一个内存空间
	ch := make(chan int)
	go Read(ch)
	go Write(ch)
	time.Sleep(1)
}