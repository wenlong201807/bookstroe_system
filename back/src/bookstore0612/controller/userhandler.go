package controller

import (
	"bookstore0612/dao"
	"fmt"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	// 获取用户名和密码
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	fmt.Println("post==username, password。。。", username, password)

	// 调用userdao中验证用户名和密码的方法
	user, _ := dao.CheckUserNameAndPassword(username, password)
	if user.ID > 0 {
		// 用户名和密码正确，执行这里
		fmt.Println("用户名和密码正确，执行这里user", user)
	} else {
		// 用户名和密码有一个不正确，执行这里
		fmt.Println("用户名和密码有一个不正确，执行这里user", user)
	}
}

func Regist(w http.ResponseWriter, r *http.Request) {
	// 获取用户名和密码
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	email := r.PostFormValue("email")
	fmt.Println("post==username, password。。。", username, password, email)

	// 调用userdao中验证用户名和密码的方法
	user, _ := dao.CheckUserName(username)
	if user.ID > 0 {
		// 用户名和密码已经存在
		fmt.Println("用户名和密码正确，执行这里user", user)
	} else {
		// 用户名和密码可用，将接收到的信息存入数据库中
		fmt.Println("用户名和密码有一个不正确，执行这里user", user)
		dao.SaveUser(username, password, email)
		// 返回前端的数据格式再此定义
		//w.Write(json.Marshal())
	}
}

// 验证用户名是否已经存在
func CheckUserName(w http.ResponseWriter, r *http.Request) {
	// 获取用户名
	username := r.PostFormValue("username")
	fmt.Println("post==username。。", username)
	// 调用userdao中验证用户名和密码的方法
	user, _ := dao.CheckUserName(username)
	if user.ID > 0 {
		// 用户名已经存在
		fmt.Println("用户名已经存在user", user)
		w.Write([]byte("用户名和密码已经存在"))
	} else {
		// 用户名可用，可以继续注册剩余信息
		fmt.Println("用户名可用，可以继续注册剩余信息user", user)
		// 返回前端的数据格式再此定义
		w.Write([]byte("用户名可用，可以继续注册剩余信息"))
	}
}
