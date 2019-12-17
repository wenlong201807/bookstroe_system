package model

//orderItem 结构体
type OrderItem struct {
	OrderItemID int64   // 订单项的id
	Count       int64   // 数量
	Amount      float64 // 金额小计
	Title       string  // 书名
	Author      string  // 作者
	Price       float64 // 价格
	ImgPath     string  // 封面
	OrderID     string  // 订单行所属的订单
}
