package base

import (
	"fmt"
)

// 通过结构体实现继承
type Person1 struct {
	// 父类
	name string
	age  int
}

func (person Person1) getNameAndAge() (string, int) { // 父类的方法
	return person.name, person.age
}

type Student1 struct {
	// 子类
	Person1
	speciality string
}

func (student Student1) getSpeciality() (string) { // 子类的方法
	return student.speciality
}
func OopApply1() {

	// 实例化一个学生的对象
	student := new(Student1) // 实例化一个子类

	// 给这个学生赋值
	student.name = "zhuWenLong"
	student.age = 25
	student.speciality = "IT"
	// 在其他地方应用中，获取这个学生的信息
	name, age := student.getNameAndAge()   // 子类继承父类的方法
	speciality := student.getSpeciality() // 子类自己的方法
	fmt.Println("这个学生的信息是：子类继承父类的方法", name, age)
	fmt.Println("这个学生的信息是：子类自己的方法", speciality)
}
