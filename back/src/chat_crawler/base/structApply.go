package base

import "fmt"

// 类型和作用在它上面定义的方法必须在同一个包里面定义，
// 这就是为什么不能在int，float 或者类似这些的类型上定义方法
// 解决方法：使用结构体
// 高级用法，可以在任何类型上添加任何方法

type Integer struct {
	value int
}

func compareOld(a, b int) bool {
	return a < b
}

// 面向对象的写法
// a,b 都是结构体Integer的接收者，类似Integer指向的this，给了a，b一个方法叫做compare
func (a Integer) compare(b Integer) bool {
	return a.value < b.value
}

func StructApply() {
	a := Integer{1}
	b := Integer{2}
	fmt.Printf("%v", a.compare(b)) // 这个写法怎么这么陌生？？？
	fmt.Println("-----------")
	fmt.Println("b身上的方法", b.compare(b))
	// 老的写法
	c, d := 3, 4
	fmt.Println("老的写法", compareOld(c, d))
}

//  面向对象方式***模式操作
type Point struct {
	px float32
	py float32
}

// 因为都是动态获取的数据，需要使用指针，否则会变成默认值，不可变动
func (point *Point) setXY(px, py float32) {
	point.px = px
	point.py = py
}
func (point *Point) getXY() (float32, float32) {
	x := point.px + 1
	y := point.py + 2
	return x, y
}

func OopPoint() {
	// 实例化一个点的对象
	point := new(Point)// 第一种方式
	// 给这个实例化的对象设置自定义的数据
	point.setXY(1.23, 4.56)
	// 将对象里面的数据（经过函数处理过的数据）获取出来
	px, py := point.getXY()
	fmt.Println("获取自定义传入的数据", px, py)
}

/*
类的初始化有四种方式
//有三种是指针变量****重点
point := new(Point)
point := &Point{}
point := &Point{px:100,py:100}

// 有一种是定义一个空的实例
point := Point{}
*/
