package controller

import (
	"bookstore0612/dao"
	"bookstore0612/model"
	"fmt"
	"net/http"
	"strconv"
)

// 获取所有图书
func GetBooks(w http.ResponseWriter, r *http.Request) {
	// 调用bookdao中获取所有图书的函数
	books, _ := dao.GetBooks()
	fmt.Println(books)
	// 执行

}

// 添加图书
func AddBook(w http.ResponseWriter, r *http.Request) {
	// 获取图书信息
	title := r.PostFormValue("title")
	author := r.PostFormValue("author")
	price := r.PostFormValue("price")
	sales := r.PostFormValue("sales")
	stock := r.PostFormValue("stock")
	fmt.Println("获取添加图书的信息title", title)
	// 将价格，销量和库存字符进行转换
	fPrice, _ := strconv.ParseFloat(price, 64)
	iSales, _ := strconv.ParseInt(sales, 10, 0)
	iStock, _ := strconv.ParseInt(stock, 10, 0)
	// 创建book
	book := &model.Book{
		Title:   title,
		Author:  author,
		Price:   fPrice,
		Sales:   int(iSales),
		Stock:   int(iStock),
		ImgPath: "/static/img/default.jpg",
	}
	// 调用bookdao中添加图书的函数
	dao.AddBook(book)
	// 调用GetBook处理函数，再查询一次数据库
	//	GetBooks(w,r)

}

// 删除图书
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	// 获取图书信息
	bookID := r.FormValue("bookId")
	fmt.Println("删除图书bookId", bookID)
	// 调用bookdao中添加图书的函数
	dao.DeleteBook(bookID)
	// 调用GetBook处理函数，再查询一次数据库
	//	GetBooks(w,r)
}

// 更新图书前***将原有的图书信息返回页面****可以不用这一步骤了
func ToUpdateBookPage(w http.ResponseWriter, r *http.Request) {
	// 获取图书信息
	bookID := r.PostFormValue("bookId")
	fmt.Println("更新图书bookId", bookID)
	// 调用bookdao中获取图书的函数
	book, _ := dao.GetBookByID(bookID)
	fmt.Println("将要更新图书的原有信息", book)
	// 返回给前端页面展示将要修改的信息
	//	GetBooks(w,r)
}

func UpdateBookPage(w http.ResponseWriter, r *http.Request) {
	// 获取图书信息
	bookID := r.PostFormValue("bookId")
	title := r.PostFormValue("title")
	author := r.PostFormValue("author")
	price := r.PostFormValue("price")
	sales := r.PostFormValue("sales")
	stock := r.PostFormValue("stock")
	fmt.Println("获取更新图书的信息bookID，title",bookID, title)
	// 将价格，销量和库存字符进行转换
	fPrice, _ := strconv.ParseFloat(price, 64)
	iSales, _ := strconv.ParseInt(sales, 10, 0)
	iStock, _ := strconv.ParseInt(stock, 10, 0)
	ibookID, _ := strconv.ParseInt(bookID, 10, 0)
	// 创建book
	book := &model.Book{
		ID:      int(ibookID),
		Title:   title,
		Author:  author,
		Price:   fPrice,
		Sales:   int(iSales),
		Stock:   int(iStock),
		ImgPath: "/static/img/default.jpg",
	}
	// 调用bookdao中更新图书的函数
	dao.UpdateBook(book)
	// 调用GetBook处理函数，再查询一次数据库
	//	GetBooks(w,r)
}
