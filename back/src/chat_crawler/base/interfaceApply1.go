package base

import "fmt"

// 目标： 接口赋值
// //将对象实例赋值给接口
// 将一个接口赋值给另一个接口

type Animal1 interface {
	Fly()
	Run()
}
type Animal2 interface {
	Fly()
}

type Bird1 struct {
	// 定义一个类** 只是go里面叫做结构体而已
}

func (bird Bird1) Fly() {
	fmt.Println("Bird is flying...")
}
func (bird Bird1) Run() {
	fmt.Println("Bird is runing...")
}

func InterfaceApply1() {
	/*
	    //将对象实例赋值给接口
		var animal1 Animal1 // 实例化一个接口
		var animal2 Animal2 // 实例化一个接口
		bird1 := new(Bird1) // 实例化一个结构体 // 也叫实例化了一个
			animal1 = bird1 // 将结构体赋值给接口****难点
			animal2 = bird1 // 将结构体赋值给接口****难点

			animal1.Fly()
			animal1.Run()

			animal2.Fly()
	*/

	// 将一个接口赋值给另一个接口
	var animal1 Animal1 // 实例化一个接口
	bird1 := new(Bird1) // 实例化一个结构体 // 也叫实例化了一个
	animal1 = bird1
	//animal1 接口内容多的可以赋值给接口内容少的animal2
	animal2 := animal1 // 一个接口可以赋值给另一个接口的含义
	animal2.Fly()

}
