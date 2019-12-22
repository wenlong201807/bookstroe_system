package base

import (
	"fmt"
	"time"
)

// 协程与线程质的不同
// 协程的特点
//1.该任务的业务代码主动要求切换，即：主动让出执行权
//2.发生了io，导致执行阻塞

// 经典的延迟控制操作
func Go_channel6_timeout() {
	ch := make(chan int)

	// 保证主进程最后执行
	//time.Sleep(1 * time.Second)**有延迟控制，不需要睡眠控制
	select {
	case <-ch:
		fmt.Println("come to read ch!")
	case <-time.After(time.Second): //原理同5的
		fmt.Println("come to timeout!")
		//default:** 延迟控制操作不能有默认输出
		//	fmt.Println("come to default!")
	}

	fmt.Println("end of code")
}
