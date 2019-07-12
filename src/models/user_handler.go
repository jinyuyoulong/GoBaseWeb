package models

import (
	"log"
	pageModel "project-web/src/library/page/model"
)

var signleUser User

// CreateUser User操作类
func CreateUser() User {
	if (User{}) != signleUser {
		signleUser = User{}
	}
	return signleUser
}

// TableName 表名
func (User) TableName() string {
	return "user"
}

// GetSequence get 序列号
func (User) GetSequence() string {
	return "user"
}

// InsertUser 增
func (User) InsertUser(data *User) error {
	_, err := engine.Insert(data)
	return err
}

// Delete 删
func (User) Delete(id int) error {
	// 假删除
	data := &User{Id: id, SysStatus: 1}
	_, err := engine.Id(data.Id).Update(data)

	return err
}

// UpdateUser 改
// columns 判断强制更新
func (User) UpdateUser(data *User, columns []string) error {
	_, err := engine.Id(data.Id).MustCols(columns...).Update(data)
	// 用到 MustCols 方法
	return err
}

// GetUserInfo 查
func (User) GetUserInfo(id int) *User {
	data := &User{Id: id}
	ok, err := engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		data.Id = 0
		return data
	}
}

// GetAll 查all
func (User) GetAll() []User {
	// 集合的两种创建方式
	// datalist := make([]model.StartInfo, 0)
	datalist := []User{}
	err := engine.Desc("id").Find(&datalist)
	if err != nil {
		log.Println(err)
		return datalist
		// return nil 也可以
	}
	return datalist
}

// GetByPage 分页
// limit 每页数量
// offset 页码
func (User) GetByPage(limit, offset int) *pageModel.PageInfo {

	pageinfo := new(pageModel.PageInfo)
	pageinfo.PageNum = offset
	pageinfo.PageSize = limit
	if offset <= 1 {
		pageinfo.IsFirstPage = true
		pageinfo.PrePage = 0
	}

	pagesize := pageModel.PageQueryCondition{
		PageNum:  offset,
		PageSize: limit,
	}

	total := engine.Sql("select count(*) as total from user").Query()

	pageinfo.Total = total.Result[0]["total"].(int64)
	pageinfo.Pages = int(pageinfo.Total / int64(limit))
	if int(pageinfo.Total%int64(limit)) > 0 {
		pageinfo.Pages++
	}
	dboffset := (pagesize.PageNum - 1) * pagesize.PageSize

	results, _ := engine.Sql("SELECT * FROM user LIMIT ? OFFSET ?", pagesize.PageSize, dboffset).Query().List()
	lenResult := len(results)
	pageinfo.Size = lenResult
	if lenResult < limit || pageinfo.Total == int64(limit*offset) {
		pageinfo.IsLastPage = true
		pageinfo.NextPage = 0
	}

	pageinfo.ListData = results
	return pageinfo
}
