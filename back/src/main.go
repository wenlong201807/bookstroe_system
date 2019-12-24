package main

import (
	"fmt"
	"net/http"
)

//	创建处理器函数
func handler(w http.ResponseWriter,r *http.Request)  {
	fmt.Fprintln(w,"hellow",r.URL.Path)
}
func main()  {
	fmt.Println("i am linux 106.54.207.247 main window 192.168.43.148")
	http.HandleFunc("/",handler)
	//创建路由
	http.ListenAndServe("106.54.207.247:8080",nil)

}