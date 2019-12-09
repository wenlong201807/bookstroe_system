package dao

import (
	"bookstore0612/model"
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("测试bookdao中的方法")
	m.Run()
}

func TestUser(t *testing.T) {
	fmt.Println("测试userdao中的函数")
	// 以下三个测试通过
	//t.Run("验证用户名和密码", testLogin)
	//t.Run("验证用户名", testRegist)
	//t.Run("保存用户名", testSave)
}

func testLogin(t *testing.T) {
	user, _ := CheckUserNameAndPassword("admin", "123456")
	fmt.Println("获取用户信息是：", user)
}

func testRegist(t *testing.T) {
	user, _ := CheckUserName("admin")
	fmt.Println("获取当前用户是：", user)
}

func testSave(t *testing.T) {
	SaveUser("admin2", "123456", "zhuweenlong")
	fmt.Println("保存用户数据成功")
}

func TestBook(t *testing.T) {
	fmt.Println("测试bookdao中的相关函数")
	//t.Run("测试获取所有图书",testGetBooks)
	//t.Run("测试添加一本图书", testAddBook)
	//t.Run("测试删除一本图书", testDeleteBook)
	//t.Run("测试查询一本图书", testGetBook)
	t.Run("测试修改一本图书", testUpdateBook)
}

func testGetBooks(t *testing.T) {
	books, _ := GetBooks()
	// 遍历得到每一本图书
	for k, v := range books {
		fmt.Printf("第%v本图书的信息是：%v\n", k+1, v)
	}
}

func testAddBook(t *testing.T) {
	book := &model.Book{
		Title:   "三国演义66",
		Author:  "罗贯中66",
		Price:   66.88,
		Sales:   60,
		Stock:   60,
		ImgPath: "/static/img/default.jpg",
	}
	// 调用添加图书的函数
	AddBook(book)
	fmt.Println("添加图书成功")
}
func testDeleteBook(t *testing.T) {
	// 调用删除图书的函数
	DeleteBook("33")
	fmt.Println("删除图书成功")
}
func testGetBook(t *testing.T) {
	// 调用查询图书的函数
	book, _ := GetBookByID("38")
	fmt.Println("查询一本图书成功", book)
}
func testUpdateBook(t *testing.T) {
	book := &model.Book{
		ID:      36,
		Title:   "三66",
		Author:  "罗66",
		Price:   66.88,
		Sales:   60,
		Stock:   60,
		ImgPath: "/static/img/default.jpg",
	}
	// 调用更新图书的函数
	UpdateBook(book)
	fmt.Println("修改图书成功")
}
