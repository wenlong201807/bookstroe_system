package main

import "chat_crawler/base"

func main() {

	//base.Arr()
	//base.Map()
	//base.Range()
	//base.Func1()
	//base.FuncTwo()
	//base.NilFunc()
	//base.Defer()
	//base.Panic()
	//base.TypeFunc()
	//base.Sturct()
	//base.StructApply()
	//base.OopPoint() // 面向对象操作模式
	//base.OopApply1() // 父类与子类的属性，方法之间的关系，以及如何实现
	//base.InterfaceMod() // 接口模型初试 // 对象赋值给接口
	//base.InterfaceApply1() // // 将一个接口赋值给另一个接口
	//base.AnyInterface()
	//base.GoMod()
	//base.Go_apply1()
	//base.Go_channel() // 通道与协程初始版
	//base.Go_channel2() // 数组与协程和通道混合版
	//base.Go_channel3() // 数组与协程和通道混合版
	//base.Go_channel4_select()
	//base.Go_channel5_timeout()
	//base.Go_channel5_timeout()
	base.Go_corountine() // 协程的特点
	// 类型和作用在它上面定义的方法必须在同一个包里面定义，
	// 这就是为什么不能在int，float 或者类似这些的类型上定义方法
	// 解决方法：使用结构体

	// 在go语言中，一个类只需要实现了接口要求的所有函数，
	// 我们就说这个类实现了该接口
	// 非侵入式接口

	/*
	  // any类型***重点
	  //go语言中的任何对象实例都满足空接口interface{}
	  var v1 interface{} = 1
	  var v2 interface{} = "abc"
	  var v3 interface{} = 2.3654
	  var v4 interface{} = make(map[string]string)
	*/

	// 协程之间通信，使用chanel,即，多个函数之间公用一个数据，同时进行对一个数据操作，保证前后变化有序

	// 协程与线程质的不同
	// 协程的特点
	//1.该任务的业务代码主动要求切换，即：主动让出执行权
	//2.发生了io，导致执行阻塞
}
