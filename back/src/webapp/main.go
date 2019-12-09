package main

import (
	"fmt"
	"net/http"
)

// 第一种方式
//	创建处理器函数
func handler(w http.ResponseWriter,r *http.Request)  {
	fmt.Fprintln(w,"hellow999",r.URL.Path) // 获取路径
}
func main()  {
	http.HandleFunc("/http",handler)
	//创建路由
	http.ListenAndServe(":8080",nil)
}

/*
// 第一种方式基础上添加多路复用器
//	创建处理器函数
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "自己创建多路复用器", r.URL.Path) // 获取路径
}
func main() {
	// 创建多路复用器
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	//创建路由
	http.ListenAndServe(":8080", mux)
}
*/
/*
// 第一种方式
//	创建处理器函数
func handler(w http.ResponseWriter,r *http.Request)  {
	fmt.Fprintln(w,"hellow",r.URL.Path) // 获取路径
}
func main()  {
	http.HandleFunc("/",handler)
	//创建路由
	http.ListenAndServe(":8087",nil)
}
*/

/*第二种
type MyHandler struct {

}

func (m *MyHandler) ServeHTTP(w http.ResponseWriter,r *http.Request)  {
	fmt.Fprintln(w,"这是底层实现，自己实现的888",r.URL.Path) // 获取路径
}

func main()  {
	MyHandler := MyHandler{}
	http.Handle("/myHandlerr",&MyHandler)
	http.ListenAndServe(":8089",nil)
}
*/

/*第三种
type MyHandler struct {

}

func (m *MyHandler) ServeHTTP(w http.ResponseWriter,r *http.Request)  {
	fmt.Fprintln(w,"详细配置http请求get",r.URL.Path) // 获取路径
}

func main()  {
	MyHandler := MyHandler{}
	server := http.Server{
		Addr:              ":8086",
		Handler:           &MyHandler,
		//TLSConfig:         nil,
		ReadTimeout:       2*time.Second,
		//ReadHeaderTimeout: 0,
		//WriteTimeout:      0,
		//IdleTimeout:       0,
		//MaxHeaderBytes:    0,
		//TLSNextProto:      nil,
		//ConnState:         nil,
		//ErrorLog:          nil,
		//BaseContext:       nil,
		//ConnContext:       nil,
	}
	server.ListenAndServe()
}
*/
