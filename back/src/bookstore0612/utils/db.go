package utils

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Db *sql.DB
	err error
)

func init()  {
	Db,err = sql.Open("mysql","root:admin@tcp(localhost:3306)/bookstore0612") // 注意写法细节
	//defer  Db.Close()
	if err != nil{
		panic(err.Error())
	}
}
