package base

import (
	"fmt"
	"time"
)

func test_routine()  {
	fmt.Println("this is one routine!")
}

func GoMod()  {
	// go 定义一个协程，只需要在一个函数的开头加上go关键字即可，就是go的协程
	go test_routine()

	// 协程也是进程之一，还是子进程，main函数为入口函数，是主进程，主进程死亡，子进程全部死亡，
	// 因此，子进程要在主进程死亡之前执行才有效果
	time.Sleep(1)
}