package base

import (
	"fmt"
	"time"
)

func add1(x,y int)  {
	z:= x+y
	fmt.Println("加法计算结果：",z)
}

func Go_apply1()  {
	for i := 0;i<10;i++{
		go add1(i,i) // 每次执行的结果都不一样，但是都可以全部被执行
	}
	time.Sleep(1)
}
