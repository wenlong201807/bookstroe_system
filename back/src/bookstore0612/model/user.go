package model

// User结构体
type User struct {
	ID int
	Username string
	Password string
	Email string
}

// UserPage结构体
type UserPage struct {
	Users        []*User
}
