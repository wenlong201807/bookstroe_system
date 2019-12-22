package base

import "fmt"

type Person struct {
	name string
	age int
}

type Student struct {
	Person
	age int // 重写会覆盖继承的相同的东西
	speciality string
}

func Sturct()  {
	// 实例化一个具体的学生
	student := Student{Person{"zhu",66},22,"it"}
	fmt.Printf("%v",student)
	fmt.Println("------")
	// 获取每一个参数的内容
	fmt.Println(student.age) // 自定的**重写的内容
	fmt.Println(student.speciality)
	fmt.Println(student.Person.age) // 继承的东西
	// 不可遍历的*****？？？
	//for k,v := range student.Person{
	//	fmt.Println("遍历每一个内容",k,v)
	//}
}
