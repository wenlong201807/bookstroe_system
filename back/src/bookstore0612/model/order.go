package model

//order结构体
type Order struct {
	OrderID     string    //订单号
	CreateTime  string // 生成订单的时间
	TotalCount  int64     // 订单中图书的总数量
	TotalAmount float64   // 订单中图书的总金额
	State       int64     // 订单的状态  0 未发货  1 已发货 2 交易完成
	UserID      int64     // 订单所属的用户
}


// 前后端混合开发时，的使用方式
// 是否发货**未发货
func (order *Order) NoSend() bool  {
	return order.State == 0
}
// 是否已经收货物**已发货
func (order *Order) SendComplate() bool  {
	return order.State == 1
}
// 是否完成交易
func (order *Order) Complate() bool  {
	return order.State == 2
}
