package base

import (
	"time"
	"fmt"
)

// 经典的延迟控制操作
func Go_channel5_timeout() {
	ch := make(chan int)
	timeout := make(chan int, 1)

	go func() {
		time.Sleep(1 * time.Second)
		timeout <- 1
	}()

	// 保证主进程最后执行
	//time.Sleep(1 * time.Second)**有延迟控制，不需要睡眠控制
	select {
	case <-ch:
		fmt.Println("come to read ch!")
	case <-timeout:
		fmt.Println("come to timeout!")
		//default:** 延迟控制操作不能有默认输出
		//	fmt.Println("come to default!")
	}

	fmt.Println("end of code")
}
