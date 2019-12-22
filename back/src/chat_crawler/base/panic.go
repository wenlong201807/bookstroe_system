package base

import "fmt"

func Panic()  {
	// deferr关闭程序使用
	defer func() {
		fmt.Println("再panic之前的的defer***在抛出异常之前执行，return之后执行")
	}()

	panic("我是panic...panic之后的任何代码不在执行")

	defer func() {
		fmt.Println("再panic之后的的defer***无法执行")
	}()



}
