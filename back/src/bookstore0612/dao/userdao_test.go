package dao

import (
	"bookstore0612/model"
	"fmt"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	//fmt.Println("测试搜友大写Test开头中的方法")
	m.Run()
}

func TestUser(t *testing.T) {
	//fmt.Println("测试userdao中的函数")
	// 以下三个测试通过
	//t.Run("验证用户名和密码", testLogin)
	//t.Run("验证用户名", testRegist)
	//t.Run("保存用户名", testSave)
}

func testLogin(t *testing.T) {
	user, _ := CheckUserNameAndPassword("admin", "123456")
	fmt.Println("获取用户信息是：", user)
}
func testRegist(t *testing.T) {
	user, _ := CheckUserName("admin")
	fmt.Println("获取当前用户是：", user)
}
func testSave(t *testing.T) {
	SaveUser("admin2", "123456", "zhuweenlong")
	fmt.Println("保存用户数据成功")
}

func TestBook(t *testing.T) {
	//fmt.Println("测试bookdao中的相关函数")
	//t.Run("测试获取所有图书",testGetBooks)
	//t.Run("测试添加一本图书", testAddBook)
	//t.Run("测试删除一本图书", testDeleteBook)
	//t.Run("测试查询一本图书", testGetBook)
	//t.Run("测试修改一本图书", testUpdateBook)
	//t.Run("测试分页", testGetPageBooks)
	//t.Run("测试分页，并且有价格范围的", testGetPageBooksByPrice)
}

func testGetBooks(t *testing.T) {
	books, _ := GetBooks()
	// 遍历得到每一本图书
	for k, v := range books {
		fmt.Printf("第%v本图书的信息是：%v\n", k+1, v)
	}
}
func testAddBook(t *testing.T) {
	book := &model.Book{
		Title:   "三国演义66",
		Author:  "罗贯中66",
		Price:   66.88,
		Sales:   60,
		Stock:   60,
		ImgPath: "/static/img/default.jpg",
	}
	// 调用添加图书的函数
	AddBook(book)
	fmt.Println("添加图书成功")
}
func testDeleteBook(t *testing.T) {
	// 调用删除图书的函数
	DeleteBook("33")
	fmt.Println("删除图书成功")
}
func testGetBook(t *testing.T) {
	// 调用查询图书的函数
	book, _ := GetBookByID("38")
	fmt.Println("查询一本图书成功", book)
}
func testUpdateBook(t *testing.T) {
	book := &model.Book{
		ID:      36,
		Title:   "三66",
		Author:  "罗66",
		Price:   66.88,
		Sales:   60,
		Stock:   60,
		ImgPath: "/static/img/default.jpg",
	}
	// 调用更新图书的函数
	UpdateBook(book)
	fmt.Println("修改图书成功")
}
func testGetPageBooks(t *testing.T) {
	page, _ := GetPageBooks("9")
	fmt.Println("当前页是：", page.PageNo)
	fmt.Println("总页数是：", page.TotalPageNo)
	fmt.Println("总记录数是：", page.TotalRecord)
	fmt.Println("当前页中图书有 ")
	for _, v := range page.Book {
		fmt.Println("图书的信息是：", v)
	}
}
func testGetPageBooksByPrice(t *testing.T) {
	page, _ := GetPageBooksByPrice("5", "10", "30")
	fmt.Println("带价格范围的当前页是：", page.PageNo)
	fmt.Println("带价格范围的总页数是：", page.TotalPageNo)
	fmt.Println("带价格范围的总记录数是：", page.TotalRecord)
	fmt.Println("带价格范围的当前页中图书有 ")
	for _, v := range page.Book {
		fmt.Println("带价格范围的图书的信息是：", v)
	}
}

func TestSession(t *testing.T) {
	//fmt.Println("测试session相关函数")
	//t.Run("测试添加session", testAddSession)
	//t.Run("测试删除session", testDeleteSession)
	//t.Run("测试获取session", testGetSession)
}

func testAddSession(t *testing.T) {
	sess := &model.Session{
		SessionID: "123456987",
		UserName:  "666",
		UserID:    1,
	}
	AddSession(sess)
}
func testDeleteSession(t *testing.T) {
	DeleteSession("123456987")
}
func testGetSession(t *testing.T) {
	sess, _ := GetSession("ce4bac61-69d3-4f40-6d2a-6211d3933676")
	fmt.Println("当前session的内容是：", sess)
}

func TestCart(t *testing.T) {
	//fmt.Println("测试购物车的相关函数")
	//t.Run("测试添加到购物车", testAddCart)
	//t.Run("测试根据图书的id和购物车id获取对应的购物项", testGetCartItemByBookIDAndCartID)
	//t.Run("测试根据购物车的id获取所有购物的购物项", testGetCartItemByCartID)
	//t.Run("测试根据用户的id获取对应的购物车", testGetCartByUserID)
	//t.Run("测试根据图书的id和购物车中的id以及图书的数量，更新购物项中图书的数量", testUpdateBookCount)
	//t.Run("测试先删除购物项，然后删除购物车", testDeleteCartByCartID)
	//t.Run("测试删除购物项", testDeleteCartItemByID)
}

func testAddCart(t *testing.T) {
	fmt.Println("测试AddCart方法")
	// 设置要买的第一本书
	book := &model.Book{
		ID: 3,
		//Title:   "",
		//Author:  "",
		Price: 27.20,
		//Sales:   0,
		//Stock:   0,
		//ImgPath: "",
	}
	fmt.Println("买的第一本书book", book)
	// 设置要买的第二本书
	book2 := &model.Book{
		ID:    4,
		Price: 23.00,
	}
	fmt.Println("买的第二本书book2", book2)
	// 创建一个购物项切片
	var cartItems []*model.CartItem

	//创建两个购物项
	cartItem := &model.CartItem{
		//CartItem: 0,
		Book:  book,
		Count: 10,
		//Amount:   0,
		CartID: "666888",
	}
	fmt.Println("购物项1==cartItem:", cartItem)
	cartItems = append(cartItems, cartItem)

	cartItem2 := &model.CartItem{
		//CartItem: 0,
		Book:  book2,
		Count: 10,
		//Amount:   0,
		CartID: "666888",
	}
	fmt.Println("购物项2==cartItem2:", cartItem2)
	cartItems = append(cartItems, cartItem2)
	fmt.Println("总共添加到购物车中的cartItems:", cartItems)
	// 创建购物车
	cart := &model.Cart{
		CartID:    "666888",
		CartItems: cartItems,
		//TotalCount:  20,
		//TotalAmount: 502.00,
		UserID: 4,
	}
	fmt.Println("cart:", cart)

	//将购物车中的商品添加到数据库中
	AddCart(cart)
}
func testGetCartItemByBookIDAndCartID(t *testing.T) {
	cartItem, _ := GetCartItemByBookIDAndCartID("3", "666888")
	fmt.Println("图书id=3,购物车id=666888的购物项的信息是：", cartItem)
}
func testGetCartItemByCartID(t *testing.T) {
	cartItems, _ := GetCartItemByCartID("666888")
	for k, v := range cartItems {
		fmt.Printf("第%v个购物项是：%v\n", k+1, v)
	}
}
func testGetCartByUserID(t *testing.T) {
	cart, _ := GetCartByUserID(3)
	fmt.Println("用户id=3的购物车的信息是", cart)
}
func testUpdateBookCount(t *testing.T) {
	//UpdateBookCount(180,4,"666888") // 参数已经修改了
	//fmt.Println("更新原有图书数量成功")
}
func testDeleteCartByCartID(t *testing.T) { // 测试成功，数据库没有更新，44---12分钟
	DeleteCartByCartID("ef41bb78-88d6-45f6-76f3-16a5052a80bb")
	fmt.Println("先删除购物项，然后删除购物车**成功")
}
func testDeleteCartItemByID(t *testing.T) { // 测试成功，数据库没有更新，44---12分钟
	DeleteCartItemByID("21")
	fmt.Println("删除其中一个购物项**成功")
}

func TestOrder(t *testing.T) {
	fmt.Println("测试订单相关的函数")
	//t.Run("测试添订单和订单向", testAddOrder)
	t.Run("测试查看所有的订单信息", testGetOrders)
}

func testAddOrder(t *testing.T) {
	// 生成订单号
	orderID := "12365478"
	// 创建生成的订单时间
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	// 创建订单
	order := &model.Order{
		OrderID:     orderID,
		CreateTime:  timeStr,
		TotalCount:  2,
		TotalAmount: 230,
		State:       0,
		UserID:      1,
	}
	// 创建订单
	orderItem := &model.OrderItem{

		Count:   1,                        // int64   // 数量
		Amount:  300,                      // float64 // 金额小计
		Title:   "三国演义",                   // string  // 书名
		Author:  " 时代的不是 ",                // string  // 作者
		Price:   300,                      // float64 // 价格
		ImgPath: "static/img/default.jpg", // string  // 封面
		OrderID: orderID,                  // string  // 订单行所属的订单
	}

	orderItem2 := &model.OrderItem{

		Count:   1,                        // int64   // 数量
		Amount:  100,                      // float64 // 金额小计
		Title:   "三国演义22",                 // string  // 书名
		Author:  " 时代的不是22 ",              // string  // 作者
		Price:   300,                      // float64 // 价格
		ImgPath: "static/img/default.jpg", // string  // 封面
		OrderID: orderID,                  // string  // 订单行所属的订单
	}

	// 保存订单
	AddOrder(order)
	// 保存订单向
	AddOrderItem(orderItem)
	AddOrderItem(orderItem2)
	fmt.Println("测试订单，订单项成功")
}
func testGetOrders(t *testing.T)  {
	orders,_:= GetOrders()
	for k,v := range orders{
		fmt.Println("订单信息是",k+1,v)
	}
}