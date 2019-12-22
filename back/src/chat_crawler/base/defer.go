package base

import "fmt"

func Defer()  {

	for i:=1;i<=5;i++{
		fmt.Println("正常顺序打印的",i)
		defer fmt.Println("defer部分的",i)
	}
	fmt.Println("循环体外面打印的")
}
