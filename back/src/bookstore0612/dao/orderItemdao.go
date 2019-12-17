package dao

import (
	"bookstore0612/model"
	"bookstore0612/utils"
)

// 向数据库中插入订单向
func AddOrderItem(orderItem *model.OrderItem) error {
	// 写sql
	sqlStr := "insert into order_items(count,amount,title,author,price,img_path,order_id) values(?,?,?,?,?,?,?)"
	// 执行
	_, err := utils.Db.Exec(sqlStr, orderItem.Count,orderItem.Amount,orderItem.Title,orderItem.Author,orderItem.Price,orderItem.ImgPath,orderItem.OrderID)
	if err != nil {
		return err
	}
	return nil
}
