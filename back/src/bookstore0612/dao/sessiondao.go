package dao

import (
	"bookstore0612/model"
	"bookstore0612/utils"
	"net/http"
)

// 向数据库中添加session
func AddSession(sess *model.Session) error {
	// 写sql
	sqlStr := "insert into sessions values(?,?,?)"
	// 执行
	_, err := utils.Db.Exec(sqlStr, sess.SessionID, sess.UserName, sess.UserID)
	if err != nil {
		return err
	}
	return nil
}

// 删除数据库中的session
func DeleteSession(sessID string) error {
	// 写sql
	sqlStr := "delete from sessions where session_id=?"
	// 执行
	_, err := utils.Db.Exec(sqlStr, sessID)
	if err != nil {
		return err
	}
	return nil
}

// 根据session的id值，从数据库总查询session
func GetSession(sessID string) (*model.Session, error) {
	// 写sql
	sqlStr := "select session_id ,username,user_id from sessions where session_id =?"
	// 预编译
	inStmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		return nil, err
	}
	//执行
	row := inStmt.QueryRow(sessID)
	// 创建session
	sess := &model.Session{}
	// 扫描数据库中的字段值为session的字段，并赋值
	row.Scan(&sess.SessionID, &sess.UserName, &sess.UserID)
	return sess, nil
}

// IsLogin 判断用户是否已经登录，false没有登录，true已经登录
func IsLogin(r *http.Request) (bool, string) {
	// 根据cookie的name获取cookie
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		// 获取cookie的value
		cookieValue := cookie.Value
		// 去数据库中根据cookieValue查询对应的session是否存在
		session, _ := GetSession(cookieValue)
		if session.UserID > 0 {
			// 已经登录
			return true, session.UserName
		}
	}

	// 没有登录
	return false, ""
}
