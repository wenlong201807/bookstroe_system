package base

import "fmt"

func TypeFunc()  {

	// 函数的高级用法，自定义函数
	// 声明函数类型 包括函数名字，函数参数，函数返回值
	type sum func(x,y int) int
	// 实例化一个函数，才能使用
	var f sum = func(x, y int) int {
		return x+y
	}

	//调用实例化的函数，并打印
	fmt.Println(f(3,4))
}
