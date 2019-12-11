package model

// page结构
type Page struct {
	Book        []*Book // 每页查询出来的图书存放的切片
	PageNo      int64   // 当前页
	PageSize    int64   // 每页显示的条数
	TotalPageNo int64   // 总页数，通过计算得到
	TotalRecord int64   // 总记录数，通过查询数据库得到
	MinPrice    string  // 总记录数，通过查询数据库得到
	MaxPrice    string  // 总记录数，通过查询数据库得到
	IsLogin     bool    // 通过session查询，判断是否登录状态,默认是false
	UserName    string  // 通过session查询，判断是否登录状态
}

// 判断是否有上一页
func (p *Page) IsHasPrev() bool {
	return p.PageNo > 1
}

// 判断是否有下一页
func (p *Page) IsHasNext() bool {
	return p.PageNo < p.TotalPageNo
}

// 获取上一页
func (p *Page) GetPrevPageNo() int64 {
	if p.IsHasPrev() {
		return p.PageNo - 1
	}
	return 1
}

// 获取下一页
func (p *Page) GetNextPageNo() int64 {
	if p.IsHasNext() {
		return p.PageNo + 1
	}
	return p.TotalPageNo // 没有下一页，就是最后一页
}
