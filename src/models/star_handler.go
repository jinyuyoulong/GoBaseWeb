package models

import (
	"log"
)

var signleStarInfo StarInfo

// CreateStrInfo 创建StarInfo 用于外面调用
func CreateStrInfo() StarInfo {
	if (StarInfo{}) != signleStarInfo { // 只创建一个实例
		signleStarInfo = StarInfo{}
	}
	return signleStarInfo
}

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
