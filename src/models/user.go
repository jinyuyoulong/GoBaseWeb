package models

import (
	"log"
)

type User struct {
	Age       int    `xorm:"not null default 0 comment('年龄') TINYINT(1)"`
	Email     string `xorm:"not null comment('邮箱') VARCHAR(50)"`
	Id        int    `xorm:"not null pk autoincr comment('主键ID') INT(10)"`
	Pwd       string `xorm:"not null comment('密码') VARCHAR(50)"`
	Username  string `xorm:"not null comment('用户名') VARCHAR(50)"`
	SysStatus int    `xorm:"not null comment('数据状态') TINYINT(1)"`
}

var signleUser User

func CreateUser() User {
	Initialize()
	if (User{}) != signleUser {
		signleUser = User{}
	}
	return signleUser
}

func (User) TableName() string {
	return "user"
}

func (User) GetSequence() string {
	return "user"
}

func (User) CreateUser(data *User) error {
	_, err := engine.Insert(data)
	return err
}

// columns 判断强制更新
func (User) UpdateUser(data *User, columns []string) error {
	_, err := engine.Id(data.Id).MustCols(columns...).Update(data)
	// 用到 MustCols 方法
	return err
}

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

func (User) Delete(id int) error {
	// 假删除
	data := &User{Id: id, SysStatus: 1}
	_, err := engine.Id(data.Id).Update(data)

	return err
}
