package base

import (
	"fmt"
)

// 不需要睡眠等待时间，更快，同步操作

func Add2(x, y int, quit chan int) {
	z := x + y
	fmt.Println("两个数相加的和是：", z)
	quit <- 1 // 写入
}

//func Read2(ch chan int)  {
//	value := <- ch
//	fmt.Println("value",strconv.Itoa(value))
//
//}
//func Write2(ch chan int)  {
//	ch <- 10
//}

func Go_channel2() {
	// 必须要先获取一个内存空间
	//ch := make(chan int)
	//go Read(ch)
	//go Write(ch)
	//time.Sleep(1)

	// 声明长度为10的数组协程
	chs := make([]chan int, 10) // 10为缓冲去大小，默认为0
	for i := 0; i < 10; i++ {
		// 声明****没懂？？？
		chs[i] = make(chan int)
		go Add2(i, i, chs[i])
	}
	//没懂
	for _, v := range chs {
		<-v
		//对应上面的写入的读取***有一个写入操作，必须要有一个读取操作
		// 否则会阻塞，除非有缓冲区
	}
}
