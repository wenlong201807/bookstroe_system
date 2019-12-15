package controller

import (
	"bookstore0612/dao"
	"bookstore0612/model"
	"bookstore0612/utils"
	"fmt"
	"net/http"
)

// 添加图书到购物车 post请求
func AddBook2Cart(w http.ResponseWriter, r *http.Request) {
	bookID := r.PostFormValue("bookId")
	fmt.Println("要添加图书的id是：", bookID)
	// 根据图书的id获取图书的信息
	book, _ := dao.GetBookByID(bookID)
	// 判断是否登录
	_, session := dao.IsLogin(r)
	//获取用户id
	userID := session.UserID
	// 判断数据库中是否有当前用户的购物车
	cart, _ := dao.GetCartByUserID(userID)
	if cart != nil {
		// 当前用户已经有购物车了
		fmt.Println("当前用户已经有购物车了** 再次购买的话，只需要添加数量了")


	} else {
		// 证明当前用户还没有购物车，需要创建一个购物车并添加到数据库中
		// 1.创建购物车
		// 生成购物车的id
		cartID := utils.CreateUUID()
		cart := &model.Cart{ // 里面的cartItems 是难点为第二步骤
			CartID: cartID,
			UserID: userID,
		}
		// 2.创建购物车中的购物项
		var cartItems []*model.CartItem
		cartItem := &model.CartItem{
			//CartItemID: 0, // 自动生成的
			Book:  book,
			Count: 1, // 第一次默认添加一本，之后就不走这部分逻辑了
			//Amount:     0, // 计算得到的
			CartID: cartID,
		}
		// 将购物车添加到切片
		cartItems = append(cartItems, cartItem)
		// 3.将切片设置到cart中
		cart.CartItems = cartItems
		// 4.将购物车保存到数据库中
		dao.AddCart(cart)
	}
}
