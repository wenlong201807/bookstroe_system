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
	_, err := utils.Db.Exec(sqlStr, orderItem.Count, orderItem.Amount, orderItem.Title, orderItem.Author, orderItem.Price, orderItem.ImgPath, orderItem.OrderID)
	if err != nil {
		return err
	}
	return nil
}

// 根据订单号获取该订单的所有订单项
func GetOrderItemsByOrderID(orderID string) ([]*model.OrderItem, error) {
	sql := "select id,count,amount,title,author,price,img_path,order_id from order_items where order_id=?"
	rows, err := utils.Db.Query(sql, orderID)
	if err != nil {
		return nil, err
	}
	var orderItems []*model.OrderItem
	for rows.Next() {
		orderItem := &model.OrderItem{}
		rows.Scan(&orderItem.OrderItemID, &orderItem.Count, &orderItem.Amount, &orderItem.Title, &orderItem.Author, &orderItem.Price, &orderItem.ImgPath, &orderItem.OrderID)
		// 添加到切片中
		orderItems = append(orderItems, orderItem)
	}
	return orderItems, nil
}
