package main

import (
	"fmt"
	"net/http"
)

// 创建处理器函数
func handler(w http.ResponseWriter, r *http.Request) {
	//http://localhost:8080/hello?aa=sdff&kk=
	// fmt.Fprintln 可以直接在浏览器中页面显示
	// fmt.Println 可以终端控制台示
	fmt.Fprintln(w, "你发送的请求的地址r.URL.Path是：", r.URL.Path)                            // /hello
	fmt.Fprintln(w, "你发送的请求的请求地址后的查询字符串是 r.URL.RawQuery是：", r.URL.RawQuery)         // aa=sdff&kk=145
	fmt.Fprintln(w, "请求投中所有的信息r.Header==", r.Header)                                // map[Accept:[text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3] Accept-Encoding:[gzip, deflate, br] Accept-Language:[zh-CN,zh;q=0.9,en;q=0.8] Cache-Control:[max-age=0] Connection:[keep-alive] Sec-Fetch-Mode:[navigate] Sec-Fetch-Site:[none] Sec-Fetch-User:[?1] Upgrade-Insecure-Requests:[1] User-Agent:[Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36]]
	fmt.Fprintln(w, "请求头中Accept-Encoding的信息", r.Header["Accept-Encoding"])          // [gzip, deflate, br]
	fmt.Fprintln(w, "请求头中Accept-Encoding的GET方式信息", r.Header.Get("Accept-Encoding")) // gzip, deflate, br
	// 获取请求体中的内容长度
	len := r.ContentLength
	// 创建byte切片
	body := make([]byte, len)
	// 将请求体中的内容读到 body中
	r.Body.Read(body)
	// 在浏览器中显示请求体中的内容
	fmt.Println(w, "请求体中的内容有： ", string(body))  // 终端打印
	fmt.Fprintln(w, "请求体中的内容有： ", string(body)) // 浏览器页面打印

	// 解析表单，在调用r.Form 之前必须执行该操作
	//r.ParseForm()
	// 获取请求参数
	//http://localhost:8080/hello?usernam=zhu&pas=wenlong
	fmt.Fprintln(w, "post请求参数r.PostForm有： ", r.PostForm)                              // post map[] // 只支持application/x-www-form-urlencoded
	fmt.Fprintln(w, "post请求参数r.Form.Get(username)有： ", r.Form.Get("username"))        // post map[]
	fmt.Fprintln(w, "请求参数r.Form有： ", r.Form)                                          // post // map[pas:[wenlong] usernam:[zhu]]
	fmt.Fprintln(w, "url请求参数r.FormValue有： ", r.FormValue("username"))                 // post // map[pas:[wenlong] usernam:[zhu]]
	fmt.Fprintln(w, "form表单中userna请求参数r.PostFormValue有： ", r.PostFormValue("userna")) // post // map[pas:[wenlong] usernam:[zhu]]
	// 参考：https://www.cnblogs.com/liuhe688/p/11063945.html
	// 第一种方式
	//username1 := r.Form["username"][0]
	//password1 := r.Form["password"][0]
	//fmt.Println("post1",username1,password1)

	// 第二种方式
	//username := r.Form.Get("username")
	//password := r.Form.Get("password")
	//fmt.Println("post2",username,password)
	// 第三种方式 直接调用即可
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	fmt.Println("post3", username, password)
}

func main() {

	fmt.Println(666)
	http.HandleFunc("/hello", handler)
	http.ListenAndServe("192.168.43.148:8080", nil)
}
