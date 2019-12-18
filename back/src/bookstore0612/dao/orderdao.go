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
func GetOrders() (*model.OrderPage, error) {
//func GetOrders() ([]*model.Order, error) {
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
	//// 创建order
	//var order = &model.Order{
	//	OrderID    :  orders., //string    //订单号
	//	CreateTime :  v.CreateTime, //string // 生成订单的时间
	//	TotalCount : v.TotalCount , //int64     // 订单中图书的总数量
	//	TotalAmount: v.TotalAmount , //float64   // 订单中图书的总金额
	//	State      : v.State , //int64     // 订单的状态  0 未发货  1 已发货 2 交易完成
	//	UserID     :v.UserID  , //int64     // 订单所属的用户
	//}

	// 创建page
	var OrderPage = &model.OrderPage{
		Orders:        orders,
		//Total:      len(orders),

	}
	return OrderPage, nil
	//return orders, nil
}

/*
// 获取数据库中所有订单***正常的
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
*/

