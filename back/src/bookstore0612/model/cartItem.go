package model

// 购物项结构体
type CartItem struct {
	// 犯错CartItemID **错写成CartItem，测试的时候一直不能自增id
	CartItemID int64   //购物项的id
	Book     *Book   //购物项中的图书信息
	Count    int64   //购物项中的数量
	Amount   float64 // 购物项中金额小计，通过计算得到
	CartID   string  // 当前购物项属于哪一个购物车的
}

/*
CREATE TABLE `bookstore0612`.`cart_items`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `COUNT` int(11) NOT NULL,
  `amount` double(11, 2) NOT NULL,
  `book_id` int(11) NOT NULL,
  `cart_id` varchar(100) NOT NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`book_id`) REFERENCES `bookstore0612`.`books` (`id`),
  FOREIGN KEY (`cart_id`) REFERENCES `bookstore0612`.`carts` (`id`)
);
*/

//获取购物项中图书的金额小计，有图书的价格和图书的数量计算得到
func (cartItem *CartItem) GetAmount() float64 {
	// 获取当前购物车中图书的价格
	price := cartItem.Book.Price
	return float64(cartItem.Count) * price
}
