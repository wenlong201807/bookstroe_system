package dao

import (
	"bookstore0612/model"
	"bookstore0612/utils"
)

// 获取数据库中所有图书
func GetBooks() ([]*model.Book, error) {
	// 写sql
	sqlStr := "select id,title,author,price,sales,stock,img_path from books"
	// 执行
	rows, err := utils.Db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	// 定义切片，存储数据使用
	var books []*model.Book
	for rows.Next() {
		//var book *model.Book // 犯错误的地方
		book := &model.Book{}
		// 给book中的字段赋值
		rows.Scan(&book.ID, &book.Title, &book.Author,  &book.Price, &book.Sales, &book.Stock,&book.ImgPath)
		// 将book添加到books中
		books = append(books, book)
	}
	// 返回
	return books, nil
}

// 添加一本图书到数据库中
func AddBook(b *model.Book) error  {
	// 写sql
	sqlStr := "insert into books(stock,sales,img_path,author,price,title) values(?,?,?,?,?,?)"
	// 执行
	_,err := utils.Db.Exec(sqlStr,b.Stock,b.Sales,b.ImgPath,b.Author,b.Price,b.Title,b.ID)
	if err != nil{
		return err
	}
	return nil
}
