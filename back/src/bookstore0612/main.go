package main

import (
	"bookstore0612/controller"
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter,r *http.Request)  {
	fmt.Println("测试接口。。。")
}
func main()  {
	http.HandleFunc("/main",indexHandler)
	fmt.Println("localhost:8080已经启动。。。")

	http.HandleFunc("/login",controller.Login)// 登录接口调用
	http.HandleFunc("/regist",controller.Regist) // 注册
	http.HandleFunc("/checkUserName",controller.CheckUserName) // 验证用户名是否可用**用户名不可重复
	http.HandleFunc("/getBooks",controller.GetBooks) // 获取所有图书信息


	// IP及端口的监听
	//http.ListenAndServe(":8080",nil)
	http.ListenAndServe("192.168.43.148:8080",nil)
}
