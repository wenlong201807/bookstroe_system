package controller

import (
	"bookstore0612/commons"
	"bookstore0612/dao"
	"bookstore0612/model"
	"bookstore0612/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// 添加图书到购物车 post请求
func AddBook2Cart(w http.ResponseWriter, r *http.Request) {

	// 判断是否登录
	flag, session := dao.IsLogin(r)
	if flag {
		// 已经登录状态下才可以添加到购物车中
		bookID := r.PostFormValue("bookId")
		fmt.Println("要添加图书的id是：", bookID)
		// 根据图书的id获取图书的信息
		book, _ := dao.GetBookByID(bookID)

		//获取用户id
		userID := session.UserID
		// 判断数据库中是否有当前用户的购物车
		cart, _ := dao.GetCartByUserID(userID)
		fmt.Println("判断数据库中是否有当前用户的购物车cart:", cart)
		if cart != nil {
			// 当前用户已经有购物车了
			//fmt.Println("当前用户已经有购物车了** 再次购买的话，只需要添加数量了")
			//当前用户已经有购物车了，此时需要判断购物车中是否有当前这本图书
			cartItem, _ := dao.GetCartItemByBookIDAndCartID(bookID, cart.CartID)
			//fmt.Println("判断购物车中是否有当前这本图书cartItem:",cartItem)
			if cartItem != nil {
				//当前用户已经有购物车，只需要将该图书所对应的购物项中的数量+1即可
				// 1.获取购物车切片中的所有的购物项
				cts := cart.CartItems
				// 2.遍历得到每一个购物项
				for _, v := range cts {
					//fmt.Println("遍历得到每一个购物项v:",v) // &{10 <nil> 1 19.8 ef35fca1-b0b3-4bfc-7e18-cf7c27b441cd}
					fmt.Println("v.Book.ID:", v.Book.ID) // 存在的
					//fmt.Println("cartItem.Book.ID:",cartItem.Book.ID)  // 报错
					fmt.Println("cartItem.Book:", cartItem.Book) //<nil>

					//3.找到当前的购物项****进不来
					if v.Book.ID == cartItem.Book.ID { // 添加的id与数据原有的id可以匹配上
						// 将购物车中的图书的数量+1
						v.Count = v.Count + 1
						// 将数据库中该购物项的**测试成功，但是数据库没有更新相关数据
						fmt.Println("获取的三个参数：v.Count, v.Book.ID, cart.CartID", v.Count, v.Book.ID, cart.CartID)
						//dao.UpdateBookCount(v.Count, v.Book.ID, cart.CartID)
						dao.UpdateBookCount(v)

					}
				}
			} else {
				fmt.Println("当前购物车中还没有该图书对应的购物项")
				// 购物车中的购物项中还没有该图书，此时需要创建一个购物项，并添加到数据库中
				// .创建购物车中的购物项
				fmt.Println("获取的图书的信息是：", book)
				cartItem := &model.CartItem{
					//CartItemID: 0, // 自动生成的
					Book:  book,
					Count: 1, // 第一次默认添加一本，之后就不走这部分逻辑了
					//Amount:     0, // 计算得到的
					CartID: cart.CartID,
				}
				// 将购物项添加到当前cart的切片中
				cart.CartItems = append(cart.CartItems, cartItem)
				// .将新创建的购物项添加到购物车保存到数据库中
				dao.AddCartItem(cartItem)
			}
			//  不管之前有没有图书对应的购物项，都需要更新总金额与总数量
			dao.UpdateCart(cart)
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

			// 返回页面信息
			// 自定义返回数据格式**如何定义结构体
			var er commons.StoreResult
			er.Msg = "当前用户还没有购物车，需要创建一个购物车并添加到数据库中"
			er.Status = 11
			// 把结构体转换为json数据
			b, _ := json.Marshal(er)
			// 设置响应内容为json
			w.Header().Set(commons.HEADER_CONTENT_TYPE, commons.JSON_HEADER)
			w.Write(b)
		}
	} else {
		// 没有登录，需要先登录
	}
}

// 根据用户id的信息获取购物车的信息
func GetCartInfo(w http.ResponseWriter, r *http.Request) {
	_, session := dao.IsLogin(r)
	// 获取用户id
	userID := session.UserID
	// 根据用户的id从数据库总获取对应的购物车
	cart, _ := dao.GetCartByUserID(userID)
	// 设置用户名 ** 前后端不分离的时候
	cart.UserName = session.UserName
	fmt.Println("获取到的购物车内容cart：", cart)
	if cart != nil {
		// 将购物车中的内容返回页面展示
		// 设置用户名 ** 前后端不分离的时候
		cart.UserName = session.UserName

	} else {
		// 该用户还没有购物车
	}

}

// 清空购物车
func DeleteCart(w http.ResponseWriter, r *http.Request) {
	// 获取要删除购物车的id
	cartID := r.FormValue("cartId")
	//清空购物车
	dao.DeleteCartByCartID(cartID)
	// 返回页面提示，删除成功
	// 再次调用购物车查询剩余购物项
}

// 删除购物项**难点
func DeleteCartItem(w http.ResponseWriter, r *http.Request) {
	// 获取要删除购物项的id
	cartItemID := r.FormValue("cartItemId")
	// 将购物项的id转换为int
	iCartItemID, _ := strconv.ParseInt(cartItemID, 10, 64)
	// 获取session
	_, session := dao.IsLogin(r)
	// 获取用户id
	userID := session.UserID
	// 获取该用户的购物车
	cart, _ := dao.GetCartByUserID(userID)
	// 获取购物车中的所有购物项
	cartItems := cart.CartItems
	// 遍历得到每一个购物项
	for k, v := range cartItems {
		// 寻找要删除的购物项
		if v.CartItemID == iCartItemID {
			// 这个就是我们要删除的购物项
			// 将当前购物项从切片中移除
			cartItems = append(cartItems[:k], cartItems[k+1:]...)
			// 将删除购物项之后的切片再次付给购物车中的切片**不可忽略
			cart.CartItems = cartItems
			// 将当前购物项从数据库中删除
			dao.DeleteCartItemByID(cartItemID)
		}
	}
	// 更新购物车中的图书的总金额和总数量
	dao.UpdateCart(cart)
	// 调用获取购物项信息的函数再次查询购物车信息
	//GetCartInfo(w,r)
	// 返回页面信息


}







