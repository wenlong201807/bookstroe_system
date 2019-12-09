package model

import (
	"fmt"
	"testing"
)

//终端下操作
// go test  查看测试对错
// go test -v 查看测试过程细节


// TestMain 函数可以在测试之前执行一些事情
func TestMain(m *testing.M) {
	fmt.Println("测试函数开始前执行的函数")
	// 通过t.run() 来执行子测试函数
	m.Run()
}

func TestUser(t *testing.T) {
	fmt.Println("开始测试user中的相关方法")
	// 然后再调用其他相关的自测试方法 t.Run 可以同时写多个 ，执行则删改各种操作
	//t.Run(" 测试添加用户", testAddUser)
	//t.Run(" 通过id测试获取用户：", testGetUserById)
	t.Run(" 测试获取所有用户信息：", testGetUsers)
}

// 如果函数名不是以Test开头，那么该函数默认不执行，我们可以将它设置为一个子测试函数
func testAddUser(t *testing.T) {
	//func TestAddUser(t *testing.T)  {
	fmt.Println("测试添加用户--子测试程序")
	//user := &User{}
	//// 调用添加用户的方法
	//user.AddUser()
	//user.AddUser2()
}

// 测试获取一个user的自测试方法
func testGetUserById(t *testing.T)  {
	fmt.Println("测试查询一条记录")
	user := User{
		ID:       1,
		//Username: "",
		//Password: "",
		//Email:    "",
	}
	// 调用获取User的方法
	u,_ := user.GetUserById()
	fmt.Println("得到的User的信息是： ",u)
}

// 测试获取所有user的方法
func testGetUsers(t *testing.T)  {
	fmt.Println("测试获取所有user的记录")
	// 定义一个切片，作为结果返回的变量
	user := &User{}
	// 调用获取所有user的方法
	us,_ := user.GetUsers()
	// 遍历切片
	for k,v := range us {
		fmt.Printf("第%v个用户是%v: \n",k+1,v)
	}

}