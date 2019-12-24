package base

import (
	_ "github.com/go-sql-driver/mysql" // 必须要引入的驱动
)

// beego https://beego.me/docs/intro/
// CRUD https://beego.me/docs/mvc/model/object.md

func MySql()  {
	// 连接数据库
	//db,err := sql.Open("mysql","root:admin@tcp(127.0.0.1:3306)/test?charset=utf8")
	//if err != nil{
	//	panic(err)
	//}
	// 以上为公共部分

	// 添加

	//// 预编译
	//stmt,err := db.Prepare("insert user_info set username=?,departname=?,creat_time=?")
	//res,err := stmt.Exec("zhu33","技术部33","2019-01-02")
	//id,err := res.LastInsertId() // 需要有自增id，就会有返回值
	//if err != nil{
	//	panic(err)
	//}
	//fmt.Println(id)


	// 修改 id必须要是原来存在的 ，并且有被修改的内容，否则返回值为0
	// 预编译
	//stmt,err := db.Prepare("update user_info set username=?,departname=?,creat_time=? where id=?")
	//res,err := stmt.Exec("zhu66","技术部66","2019-01-02",2)
	//row,err := res.RowsAffected() // 有被修改的行的信息，就会有返回值 修改的行个数
	//if err != nil{
	//	panic(err)
	//}
	//fmt.Println(row) // 修改一行，返回数字 1

	// 查询一条数据***失败
	// 预编译
	//row,err := db.Query("select * from user_info  where id=2")
	//rows,err := db.Query("select * from user_info")
	//if err != nil{
	//	panic(err)
	//}
	//for rows.Next(){
	//	var id int
	//	var username string
	//	var departname string
	//	var creat_time string
	//	err = rows.Scan(&username,&departname,&creat_time,&id)
	//	fmt.Println("查询结果为：",id,username,departname,creat_time)
	//}


	// 删除
	//stmt,err := db.Prepare("delete user_info where id=3")
	//res,err = stmt.Exec(id)
	//fmt.Println(res)

}
