package dao

import (
	"bookstore0612/model"
	"bookstore0612/utils"
)

// 向购物项表中插入购物项
func AddCartItem(cartItem *model.CartItem) error {
	// 写sql
	sqlStr := "insert into cart_items(count,amount,book_id,cart_id) values(?,?,?,?)"
	// 执行
	_, err := utils.Db.Exec(sqlStr, cartItem.Count, cartItem.GetAmount(), cartItem.Book.ID, cartItem.CartID)
	if err != nil {
		return err
	}
	return nil
}
