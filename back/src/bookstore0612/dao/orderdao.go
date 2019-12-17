package dao

import (
	"bookstore0612/model"
	"bookstore0612/utils"
)

// 向数据库中插入订单
func AddOrder(order *model.Order) error {
	// 写sql
	sqlStr := "insert into orders(id,create_time,total_count,total_amount,state,user_id) values(?,?,?,?,?,?)"
	// 执行
	_, err := utils.Db.Exec(sqlStr, order.OrderID, order.CreateTime, order.TotalCount, order.TotalAmount, order.State, order.UserID)
	if err != nil {
		return err
	}
	return nil
}

// 获取数据库中所有订单
func GetOrders() (ordersAll []*model.Order, err error) {
	// 写sql
	sqlStr := "select id,create_time,total_count,total_amount,state,user_id from orders"
	// 执行
	rows, err := utils.Db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	var orders []*model.Order
	for rows.Next() {
		order := &model.Order{}
		rows.Scan(&order.OrderID, &order.CreateTime, &order.TotalCount, &order.TotalAmount, &order.State, &order.UserID)
		orders = append(orders, order)
	}
	return orders, nil
}
