package controller

import (
	"bookstore0612/commons"
	"bookstore0612/dao"
	"bookstore0612/model"
	"bookstore0612/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

// 推出登录
func Logout(w http.ResponseWriter, r *http.Request) {
	// 获取cookie
	cookie, _ := r.Cookie("user")
	fmt.Println("退出登录时获取前端传入的cookie：", cookie)
	if cookie != nil {
		// 获取cookie的value
		cookieValue := cookie.Value
		// 删除数据库中与之对应的session
		dao.DeleteSession(cookieValue)
		// 设置cookie失效**数据库中的失效
		cookie.MaxAge = -1 // 注意区分session的失效
		// 发给前端，使前端的cookie失效// 将cookie发送给浏览器
		http.SetCookie(w, cookie)
	}

	// 去首页，分页查看图书信息页面***或者返回登录页面
	//dao.GetPageBooksByPrice()
}

func Login(w http.ResponseWriter, r *http.Request) {
	// 判断是否已经登录
	flag, _ := dao.IsLogin(r)
	if flag {
		// 已经登录了
		fmt.Println("您已经是登录状态了，直接跳转至首页吧")
		// 直接跳转到首页去，不用再生成seesion，cookie了

		// 自定义返回数据格式**如何定义结构体
		var er commons.StoreResult
		er.Msg = "已经是登录状态"
		er.Status = 5
		// 把结构体转换为json数据
		b, _ := json.Marshal(er)
		// 设置响应内容为json
		w.Header().Set(commons.HEADER_CONTENT_TYPE, commons.JSON_HEADER)
		w.Write(b)
	} else {
		// 之前没有登录的
		// 需要第一次进行登录，再后续操作
		// 获取用户名和密码
		username := r.PostFormValue("username")
		password := r.PostFormValue("password")
		fmt.Println("post==username, password。。。", username, password)

		// 调用userdao中验证用户名和密码的方法
		user, _ := dao.CheckUserNameAndPassword(username, password)
		if user.ID > 0 {
			// 用户名和密码正确，执行这里

			// 生成UUID作为session的id
			uuid := utils.CreateUUID()
			// 创建一个session
			sess := &model.Session{
				SessionID: uuid,
				UserName:  user.Username,
				UserID:    user.ID,
			}
			// 将session保存到数据库中
			dao.AddSession(sess)
			// 创建一个cookie，让它与session相关联
			cookie := http.Cookie{
				Name:  "user",
				Value: uuid,
				//Path:       "",
				//Domain:     "",
				//Expires:    time.Time{},
				//RawExpires: "",
				//MaxAge:   0,
				//Secure:     false,
				HttpOnly: true,
				//SameSite:   0,
				//Raw:        "",
				//Unparsed:   nil,
			}
			// 将cookie发送给浏览器
			http.SetCookie(w, &cookie)

			fmt.Println("用户名和密码正确，执行这里user", user)

			// 自定义返回数据格式**如何定义结构体
			var er commons.StoreResult
			er.Msg = "登录成功"
			er.Status = 6
			// 把结构体转换为json数据
			b, _ := json.Marshal(er)
			// 设置响应内容为json
			w.Header().Set(commons.HEADER_CONTENT_TYPE, commons.JSON_HEADER)
			w.Write(b)

		} else {
			// 用户名和密码有一个不正确，执行这里
			fmt.Println("用户名和密码有一个不正确，执行这里user", user)

			// 自定义返回数据格式**如何定义结构体
			var er commons.StoreResult
			er.Msg = "登录失败"
			er.Status = 7
			// 把结构体转换为json数据
			b, _ := json.Marshal(er)
			// 设置响应内容为json
			w.Header().Set(commons.HEADER_CONTENT_TYPE, commons.JSON_HEADER)
			w.Write(b)
		}

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

		// 自定义返回数据格式**如何定义结构体
		var er commons.StoreResult
		er.Msg = "用户名和密码已经存在"
		er.Status = 1
		// 把结构体转换为json数据
		b, _ := json.Marshal(er)
		// 设置响应内容为json
		w.Header().Set(commons.HEADER_CONTENT_TYPE, commons.JSON_HEADER)
		w.Write(b)
	} else {
		// 用户名和密码可用，将接收到的信息存入数据库中
		fmt.Println("用户名和密码有一个不正确，执行这里user", user)
		dao.SaveUser(username, password, email)
		// 返回前端的数据格式再此定义
		//w.Write(json.Marshal())

		// 自定义返回数据格式**如何定义结构体
		var er commons.StoreResult
		er.Msg = "注册成功"
		er.Status = 2
		// 把结构体转换为json数据
		b, _ := json.Marshal(er)
		// 设置响应内容为json
		w.Header().Set(commons.HEADER_CONTENT_TYPE, commons.JSON_HEADER)
		w.Write(b)
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
		//w.Write([]byte("用户名和密码已经存在"))
		// 自定义返回数据格式**如何定义结构体
		var er commons.StoreResult
		er.Msg = "用户名已经存在"
		er.Status = 3
		// 把结构体转换为json数据
		b, _ := json.Marshal(er)
		// 设置响应内容为json
		w.Header().Set(commons.HEADER_CONTENT_TYPE, commons.JSON_HEADER)
		w.Write(b)
	} else {
		// 用户名可用，可以继续注册剩余信息
		fmt.Println("用户名可用，可以继续注册剩余信息user", user)
		// 返回前端的数据格式再此定义
		//w.Write([]byte("用户名可用，可以继续注册剩余信息"))
		// 自定义返回数据格式**如何定义结构体
		var er commons.StoreResult
		er.Msg = "用户名可用"
		er.Status = 4
		// 把结构体转换为json数据
		b, _ := json.Marshal(er)
		// 设置响应内容为json
		w.Header().Set(commons.HEADER_CONTENT_TYPE, commons.JSON_HEADER)
		w.Write(b)
	}
}
