package model

// session结构体
type Session struct {
	SessionID string
	UserName  string
	UserID    int
	Cart      *Cart // 前后端不分离的时候需要使用
	OrderID   string
}
