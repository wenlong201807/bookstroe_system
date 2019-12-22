package base

import "fmt"

type Animal interface {
	Fly()
	Run()
}

type Bird struct {
	// 定义一个类** 只是go里面叫做结构体而已
}

func (bird Bird) Fly() {
	fmt.Println("Bird is flying...")
}
func (bird Bird) Run() {
	fmt.Println("Bird is runing...")
}

func InterfaceMod() {

	var animal Animal // 实例化一个接口
	bird := new(Bird) // 实例化一个结构体 // 也叫实例化了一个类

	animal = bird // 将结构体赋值给接口****难点
	animal.Fly()
	animal.Run()
}

// 在go语言中，一个类只需要实现了接口要求的所有函数，
// 我们就说这个类实现了该接口
// 这种模式叫做非侵入式接口** 与其他语言不通的地方，也是牛逼的地方
// go语言中，结构体就是类
