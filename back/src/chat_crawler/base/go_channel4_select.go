package base

import (
	"time"
	"fmt"
)

func Go_channel4_select() {
	ch := make(chan int)

	go func(ch chan int) {
		ch <- 1
	}(ch)

	// 保证主进程最后执行
	time.Sleep(1 * time.Second)
	select {
	case <-ch:
		fmt.Println("come to read ch!")

		//case <- ch:
		//	fmt.Println("come to read ch!")
	default:
		fmt.Println("come to default!")

	}
}
