package main

import (
	"fmt"
	"net/http"
)

// 设置cookie到客户端 // 用途**1、广告推荐，2、免登录
func setCookie(w http.ResponseWriter, r *http.Request) {
	// 创建Cookie
	cookie := http.Cookie{
		Name:  "user",
		Value: "admin",
		//Path:       "",
		//Domain:     "",
		//Expires:    time.Time{},
		//RawExpires: "",
		MaxAge:     10,// 指定有效时间，秒为单位，不管浏览器是否关闭，到了时间才会消失 // 默认是会话级别的
		//Secure:     false,
		HttpOnly: true,
		//SameSite:   0,
		//Raw:        "",
		//Unparsed:   nil,
	}
	cookie2 := http.Cookie{
		Name:     "user22",
		Value:    "admin22",
		HttpOnly: true,
	}
	// 第一种方式
	// 将cookie发送给浏览器
	//w.Header().Set("Set-Cookie",cookie.String())
	// 添加第二个cookie
	//w.Header().Add("Set-Cookie",cookie2.String())

	// 第二种方式
	http.SetCookie(w, &cookie)
	http.SetCookie(w, &cookie2)
}

// 获取客户端的coookie
func getCookies(w http.ResponseWriter, r *http.Request) {
	// 获取请求头中所有的cookies
	cookies := r.Header["Cookie"]
	fmt.Println("得到的所有的Cookie有：", cookies)
	// 获取请求头中指定的cookies
	cookie, _ := r.Cookie("user")
	fmt.Println("得到的指定的Cookie有：", cookie)
}

func main() {
	http.HandleFunc("/setCookie", setCookie)
	http.HandleFunc("/getCookies", getCookies)
	http.ListenAndServe(":8085", nil)
}
