package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"web01_db/model"
)

// 参考：https://www.cnblogs.com/liuhe688/p/11063945.html
// 创建处理器函数
func handler(w http.ResponseWriter, r *http.Request) {
	// 获取请求参数
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	fmt.Println("post3", username, password)
}

func testJsonRes(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	// 创建user
	user := model.User{
		ID:1,
		Username:"ASDF",
		Password:"sdgdsd",
		Email:"sdagfads",
	}
	// 将user转换成json形式
	json,_ := json.Marshal(user)
	// 返回客户端
	w.Write(json)
}
func testredirect(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Location","https://www.cnblogs.com/liuhe688/p/11063945.html")
	w.WriteHeader(302) // 设置重定向// 注意顺序

}

func main() {
	fmt.Println("test",time.Now())
	http.HandleFunc("/test", handler)
	http.HandleFunc("/testJson ", testJsonRes)
	http.HandleFunc("/testRedirect ", testredirect)
	//http.ListenAndServe("192.168.43.148:8080", nil)
	http.ListenAndServe(":8080", nil)
}
