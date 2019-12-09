package model

import (
	"fmt"
	"web01_db/utils"
)

type User struct {
	ID int
	Username string
	Password string
	Email string
}

// 添加user方法***方法一
func (user *User) AddUser() error  {
	// 1写sql
	sqlStr := "insert into users(username,password,email) values(?,?,?)"
	// 2预编译
	inStmt,err := utils.Db.Prepare(sqlStr)
	if err != nil{
		fmt.Println("预编译出现异常:",err)
		return err
	}
	// 3执行
	_,err2 := inStmt.Exec("admin","123456","157@qq.com")
	if err2 != nil{
		fmt.Println("执行出现异常:",err2)
		return err2
	}
	return  nil
}

// 添加user方法***方法二
func (user *User) AddUser2() error  {
	// 1写sql
	sqlStr := "insert into users(username,password,email) values(?,?,?)"

	// 2执行
	_,err2 := utils.Db.Exec(sqlStr,"admin2","12345622","15722@qq.com")
	if err2 != nil{
		fmt.Println("执行出现异常:",err2)
		return err2
	}
	return  nil
}

// 根据用户的id从数据中查询一条记录
func (user *User) GetUserById() (*User ,error) {
	// 写sql
	sqlStr := "select id,username,password,email from users where id = ?"
	// 执行
	row := utils.Db.QueryRow(sqlStr,user.ID)
	// 声明参数变量
	var id int
	var username string
	var password string
	var email string
	err := row.Scan(&id,&username,&password,&email)
	if err != nil{
		return nil,err
	}
	// 实例化结构体
	u := &User{
		ID:       id,
		Username: username,
		Password: password,
		Email:    email,
	}
	return u,nil
}

// 获取数据库中所有的记录 ** 参数获取切片，表示可以多条数据
func (user *User) GetUsers() ([]*User,error)  {
	// 写sql
	sqlStr := "select id,username,password,email from users"
	// 执行
	rows,err := utils.Db.Query(sqlStr)
	if err != nil{
		return nil,err
	}

	// 返回数据为切片，先定义一个切片，作为返回值
	var users []*User
	// 区别，循环遍历数据获取数据库中的数据
	for rows.Next(){
		// 声明参数变量
		var id int
		var username string
		var password string
		var email string
		err := rows.Scan(&id,&username,&password,&email)
		if err != nil{
			return nil,err
		}
		// 实例化结构体
		u := &User{
			ID:       id,
			Username: username,
			Password: password,
			Email:    email,
		}
		// 循环所有结果，并保存到切片中
		users = append(users,u)
	}
	// 最后返回出去
	return users,nil
}



