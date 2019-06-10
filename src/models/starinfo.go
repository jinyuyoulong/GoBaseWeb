package models

import (
	"log"
)

// # 数据库建表遵循 全小写原则 xorm 映射--> NameZh:name_zh
type StarInfo struct {
	Id           int    `xorm:"not null pk autoincr comment('主键ID') INT(10)" form:"id"`
	NameZh       string `xorm:"not null comment('中文名') VARCHAR(50)" form:"name_zh"`
	NameEn       string `xorm:"not null comment('英文名') VARCHAR(50)" form:"name_en"`
	Avatar       string `xorm:"not null comment('头像') VARCHAR(255)" form:"avatar"`
	Birthday     string `xorm:"not null comment('出生日期') VARCHAR(50)" form:"birthday"`
	Height       int    `xorm:"not null default 0 comment('身高，单位cm') INT(10)" form:"height"`
	Weight       int    `xorm:"not null comment('体重，单位g') INT(10)" form:"weight"`
	Club         string `xorm:"not null comment('俱乐部') VARCHAR(50)" form:"club"`
	Jersy        string `xorm:"not null comment('球衣号码以及主打位置') VARCHAR(50)" form:"jersy"`
	Country      string `xorm:"not null comment('国籍') VARCHAR(50)" form:"country"`
	Birthaddress string `xorm:"not null comment('出生地') VARCHAR(255)" form:"birthaddress"`
	Feature      string `xorm:"not null comment('个人特点') VARCHAR(255)" form:"feature"`
	Moreinfo     string `xorm:"comment('更多介绍') TEXT" form:"moreinfo"`
	SysStatus    int    `xorm:"not null default 0 comment('状态，默认值 0 正常，1 删除') TINYINT(4)" form:"-"`
	SysCreated   int    `xorm:"not null default 0 comment('创建时间') INT(10)" form:"-"`
	SysUpdated   int    `xorm:"not null default 0 comment('最后修改时间') INT(10)" form:"-"`
}

var signleStarInfo StarInfo

// CreateStrInfo 创建StarInfo 用于外面调用
func CreateStrInfo() StarInfo {
	Initialize()

	if (StarInfo{}) != signleStarInfo { // 只创建一个实例
		signleStarInfo = StarInfo{}
	}
	return signleStarInfo
}

// TableName 表名重命名，有这个方法，自动匹配数据库表
func (StarInfo) TableName() string {
	return "star_info"
}

// GetSequence oracle 需要先查序列，具体实现待定
func (StarInfo) GetSequence() string {
	return "starInfo"
}

func (StarInfo) CreateStarInfo(data *StarInfo) error {
	// pk := StarInfo.GetSequence(data)

	_, err := engine.Insert(data)
	return err
}

// UpdateStarInfo 判断强制更新
func (StarInfo) UpdateStarInfo(data *StarInfo, columns []string) error {
	_, err := engine.Id(data.Id).MustCols(columns...).Update(data)
	// 用到 MustCols 方法
	return err
}

func (StarInfo) GetStarInfoInfo(id int) *StarInfo {
	data := &StarInfo{Id: id}
	ok, err := engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		data.Id = 0
		return data
	}
}

func (StarInfo) GetAll() []StarInfo {
	// 集合的两种创建方式
	// datalist := make([]model.StartInfo, 0)
	datalist := []StarInfo{}
	err := engine.Desc("id").Find(&datalist)
	if err != nil {
		log.Println(err)
		return datalist
		// return nil 也可以
	}
	return datalist
}

func (StarInfo) Delete(id int) error {
	// 假删除
	data := &StarInfo{Id: id, SysStatus: 1}
	_, err := engine.Id(data.Id).Update(data)

	return err
}
