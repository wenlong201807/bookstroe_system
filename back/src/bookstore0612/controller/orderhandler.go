package controller

import (
	"net/http"
	"bookstore0612/dao"
	"bookstore0612/utils"
	"bookstore0612/model"
	"time"
	"fmt"
	"encoding/json"
	"bookstore0612/commons"
)

// Checkout 去结账
func Checkout(w http.ResponseWriter, r *http.Request) {
	// 获取session
	_, session := dao.IsLogin(r)
	// 获取用户id
	userID := session.UserID
	// 获取购物车
	cart, _ := dao.GetCartByUserID(userID)
	// 生成订单号
	orderID := utils.CreateUUID()
	fmt.Println("生成的订单号orderid：", orderID)
	// 创建生成的订单时间
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	// 创建order
	order := &model.Order{
		OrderID:     orderID,
		CreateTime:  timeStr,
		TotalCount:  cart.TotalCount,
		TotalAmount: cart.TotalAmount,
		State:       0,
		UserID:      int64(userID),
	}
	// 将订单保存到数据库中
	dao.AddOrder(order)
	// 保存订单向

	// 获取购物车中的购物项
	cartItems := cart.CartItems
	// 遍历得到每一个购物项
	for _, v := range cartItems {
		// 创建orderitem
		orderItem := &model.OrderItem{
			Count:   v.Count,        // int64   // 数量
			Amount:  v.Amount,       // float64 // 金额小计
			Title:   v.Book.Title,   // string  // 书名
			Author:  v.Book.Author,  // string  // 作者
			Price:   v.Book.Price,   // float64 // 价格
			ImgPath: v.Book.ImgPath, // string  // 封面
			OrderID: orderID,        // string  // 订单行所属的订单
		}
		// 将购物项保存到数据库
		dao.AddOrderItem(orderItem)
		// 更新当前购物车中图书的库存和销量***不可忽略
		book := v.Book
		book.Sales = book.Sales + int(v.Count)
		book.Stock = book.Stock - int(v.Count)
		// 更新图书的信息*** 把销量库存的数据更新
		dao.UpdateBook(book)

	}
	// 清空购物车
	dao.DeleteCartByCartID(cart.CartID)
	// 将订单号设置到session中
	session.OrderID = orderID
	fmt.Println("将数据返回前端****")
	// 将数据返回前端
	//***
	//自定义返回数据格式**如何定义结构体**暂时未使用同意结构返回前端
	//var er *commons.StoreResult
	//er.Content = orderID
	//er.Msg = "结账后获取生成的订单号"
	//er.Status = 19
	// 把结构体转换为json数据
	b, _ := json.Marshal(orderID)
	// 设置响应内容为json
	w.Header().Set(commons.HEADER_CONTENT_TYPE, commons.JSON_HEADER)
	w.Write(b)
}

// 获取所有订单
func GetOrders(w http.ResponseWriter, r *http.Request) {

	fmt.Println("获取当前用户cookie：", r.Cookies())
	//fmt.Println("获取当前用户名username",r.FormValue("username"))
	// 调用dao中获取所有订单中的函数
	orders, _ := dao.GetOrders()
	// 发送到页面展示
	fmt.Println("获取到的所有订单信息orders：", orders) // 订单索引而已,前端自己会识别
	//for k,v := range orders{
	//	fmt.Println("订单信息是",k+1,v)
	//}

	// 自定义返回数据格式**如何定义结构体
	//var er *commons.StoreResult
	//er.Data = orders
	//er.Msg = "所有订单信息"
	//er.Status = 18
	// 把结构体转换为json数据
	b, _ := json.Marshal(orders)
	// 设置响应内容为json
	w.Header().Set(commons.HEADER_CONTENT_TYPE, commons.JSON_HEADER)
	w.Write(b)
}

// 获取订单详情
func GetOrderItemsInfo(w http.ResponseWriter, r *http.Request) {
	// 获取订单号
	orderID := r.FormValue("orderId")
	// 根据订单号调用dao中获取所有订单向的函数
	orderItems, _ := dao.GetOrderItemsByOrderID(orderID)
	fmt.Println("订单详情：", orderItems)
	// 返回页面展示
}

// 获取我的订单
func GetMyOrders(w http.ResponseWriter, r *http.Request) {
	// 获取session
	_, session := dao.IsLogin(r)
	// 获取用户id
	userID := session.UserID
	// 调用dao中获取用户的所有订单的函数
	orders, _ := dao.GetMyOrders(userID)
	// 返回给前端展示
	fmt.Println("获取我的所有的订单orders：", orders)
}

// 发货
func SendOrder(w http.ResponseWriter, r *http.Request)  {
	// 获取要发货的订单号
	orderID := r.PostFormValue("orderId")
	fmt.Println("获取订单号：order ID：",orderID)
	// 调用dao中的更新订单的状态的函数
	dao.UpdateOrderState(orderID,1) // 1为发货2为收货
	// 返回前端信息
	// 获取所有订单的信息

}
// 收货
func ReceiveOrder(w http.ResponseWriter, r *http.Request)  {
	// 获取要收货的订单号
	orderID := r.PostFormValue("orderId")
	fmt.Println("获取订单号：order ID：",orderID)
	// 调用dao中的更新订单的状态的函数
	dao.UpdateOrderState(orderID,2) // 1为发货2为收货
	// 返回前端信息
	// 获取我的订单信息

}