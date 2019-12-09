package controller

import (
	"bookstore0612/dao"
	"fmt"
	"net/http"
)

// 获取所有图书
func GetBooks(w http.ResponseWriter, r *http.Request)  {
	// 调用bookdao中获取所有图书的函数
	books,_ := dao.GetBooks()
	fmt.Println(books)
	// 执行

}
