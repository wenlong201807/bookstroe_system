package base

import (
	"fmt"
	"strconv"
	"time"
)

// 协程与线程质的不同
// 协程的特点
//1.该任务的业务代码主动要求切换，即：主动让出执行权
//2.发生了io，导致执行阻塞

// 效率高的原因
// 1.多个协程之间切换代价低（内部处理机制使然）
// 2.可以充分利用cpu

var ch7 chan int

func Go_corountine() {

	ch7 = make(chan int)

	// 协程1
	go func() {
		for i := 0; i < 100; i++ {
			if i == 10 {
				//runtime.Gosched() // 1.主动要求让出CPU
				<-ch7 // i等于10的时候，发生阻塞
			}
			fmt.Println("corountine 1:" + strconv.Itoa(i))
		}
	}()

	// 协程2
	go func() {
		for i := 100; i < 200; i++ {
			fmt.Println("corountine 2____:" + strconv.Itoa(i))
		}
		ch7 <- 1 // 这边循环结束之后，才消除阻塞
	}()

	// 主进程最后执行
	time.Sleep(time.Second)
}
