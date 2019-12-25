package main

import (
	"fmt"
	"net/http"
	"log"
)

//	创建处理器函数
func handler(w http.ResponseWriter,r *http.Request)  {
	fmt.Fprintln(w,"hellow",r.URL.Path)

}
func main()  {
	fmt.Println("i am linux 106.54.207.247:8080 ")
	http.HandleFunc("/",handler)
	//创建路由
	//http.ListenAndServe("106.54.207.247:8080",nil) // linux中的错误写法，
	err := http.ListenAndServe(":8080",nil)  // linux中的访问  http://106.54.207.247:8080/
	if err != nil{
		log.Fatal("List 8080")
	}

}