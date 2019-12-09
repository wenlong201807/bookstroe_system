package dao

import (
	"bookstore0612/model"
	"bookstore0612/utils"
)

//根据用户名和密码从数据库中查询一条记录
func CheckUserNameAndPassword(username ,password string) (*model.User ,error) {
// 写sql语句
sqlStr := "select id,username,password,email from users where username=? and password=?"
// 执行
row := utils.Db.QueryRow(sqlStr,username,password)
user := &model.User{}
row.Scan(&user.ID,&user.Username,&user.Password,&user.Email)
return user ,nil
}

// 根据用户名和密码从数据库中查询一条记录
func CheckUserName(username string) (*model.User ,error)  {
	//写sql语句
	sqlStr := "select id,username,password,email from users where username=?"
	// 执行
	row := utils.Db.QueryRow(sqlStr,username)
	user := &model.User{}
	//  必须要使用指针获取
	row.Scan(&user.ID,&user.Username,&user.Password,&user.Email)
	return user ,nil
}

// 向数据库中插入用户信息
func SaveUser(username ,password,email string) error {
	//写sql语句
	sqlStr := "insert into users(username ,password,email) values(?,?,?)"
	//执行
	_,err := utils.Db.Exec(sqlStr,username ,password,email)
	if err != nil{
		return err
	}
	return nil
}