package commons

// 客户端与服务器段数据交互模板
type StoreResult struct {
	Status int  // 状态，为200表示成功，其他为失败
	Data interface{} // 返回的数据***重点学习
	Msg string // 返回的消息 成功/失败
	Content string // 新增返回给前端的字符串形式的内容
}
