package model

// 购物车结构体
type Cart struct {
	CartID      string      // 购物车的id
	CartItems   []*CartItem // 购物车中所有的购物项
	TotalCount  int64       //购物车中图书的总数量，通过计算所得
	TotalAmount float64     //购物车中图书的总金额，通过计算所得
	UserID      int         //当前购物车所属的用户
	UserName     string      // 用户名
}

// 获取购物车中图书的中数量
func (cart *Cart) GetTotalCount() int64 {
	var totalCount int64
	// 遍历购物车中的购物项切片
	for _, v := range cart.CartItems {
		totalCount = totalCount + v.Count
	}
	return totalCount
}

// 获取购物车中图书的总金额
func (cart *Cart) GetTotalAmount() float64 {
	var totalAmount float64
	// 遍历购物车中的购物项切片
	for _, v := range cart.CartItems {
		totalAmount = totalAmount + v.GetAmount() // 原来默认值为0，是计算得到的数据，再存入数据库中的
	}
	return totalAmount
}
