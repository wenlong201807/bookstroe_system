package dao

import (
	"bookstore0612/model"
	"bookstore0612/utils"
	"fmt"
)

// 向购物车表中插入购物车
func AddCart(cart *model.Cart) error {
	// 写sql
	sqlStr := "insert into carts(id,total_count,total_amount,user_id) values(?,?,?,?)"
	// 执行
	_, err := utils.Db.Exec(sqlStr, cart.CartID, cart.GetTotalCount(), cart.GetTotalAmount(), cart.UserID)
	if err != nil {
		return err
	}
	// 获取购物车中所有购物项
	cartItems := cart.CartItems
	// 遍历得到每一个购物项
	for _, cartItem := range cartItems {
		// 将购物项插入到数据库总
		fmt.Println("AddCart---cartItem:",cartItem)
		AddCartItem(cartItem)
	}
	return nil
}
