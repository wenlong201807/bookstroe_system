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
		rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		// 将book添加到books中
		books = append(books, book)
	}
	// 返回
	return books, nil
}

// 添加一本图书到数据库中
func AddBook(b *model.Book) error {
	// 写sql
	sqlStr := "insert into books(stock,sales,img_path,author,price,title) values(?,?,?,?,?,?)"
	// 执行*** 必须要与sql语句一一对应，犯了错
	_, err := utils.Db.Exec(sqlStr, b.Stock, b.Sales, b.ImgPath, b.Author, b.Price, b.Title)
	if err != nil {
		return err
	}
	return nil
}

// 根据图书id删除一本图书
func DeleteBook(bookId string) error {
	// 写sql
	sqlStr := "delete from books where id=?"
	//执行
	_, err := utils.Db.Exec(sqlStr, bookId)
	if err != nil {
		return err
	}
	return nil
}

// 根据图书id从数据库中查询一本图书信息
func GetBookByID(bookID string) (*model.Book, error) {
	// 写sql
	sqlStr := "select id, title,author,price,sales,stock,img_path from books where id=?"
	// 执行
	//row, err := utils.Db.Query(sqlStr, bookID) //查询所有
	row := utils.Db.QueryRow(sqlStr, bookID) // 查询一条 ** 犯错的地方

	// 创建book
	book := &model.Book{}
	// 为book中的字段赋值**完整扫描**按顺序**&符号不可少
	row.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
	return book, nil
}

// 根据图书id更新图书信息
func UpdateBook(b *model.Book)error  {
	// 写sql
	sqlStr := "update books set stock=?,sales=?,img_path=?,author=?,price=?,title=? where id=?"
	// 执行*** 必须要与sql语句一一对应，犯了错
	_, err := utils.Db.Exec(sqlStr, b.Stock, b.Sales, b.ImgPath, b.Author, b.Price, b.Title,b.ID)
	if err != nil {
		return err
	}
	return nil
}