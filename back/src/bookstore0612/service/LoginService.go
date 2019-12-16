package service

import (
	"bookstore0612/commons"
	"fmt"
)

func LoginService(un,pwd string)(er commons.StoreResult)  {
	//u := SelByUnPwdDao(un,pwd)
	// u []{username:'aa',pwd:666}
	fmt.Println("LoginService--u:",u)
	if u !=nil{
		er.Status = 200
		er.Msg = "登录成功"

	}else {
		er.Status = 400
	}
	return
}