package dao

// data access object，DAO
// 数据库访问控制层
import (
	"log"

	"github.com/go-xorm/xorm"
	"project-web/src/app/models"
)

type ProjectwebDao struct {
	engine *xorm.Engine
}

func NewProjectwebDao(engine *xorm.Engine) *ProjectwebDao {

	return &ProjectwebDao{
		engine: engine,
	}
}

func (s *ProjectwebDao) Get(id int) *models.StarInfo {
	data := &models.StarInfo{Id: id}
	ok, err := s.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		data.Id = 0
		return data
	}
}

func (s *ProjectwebDao) GetAll() []models.StarInfo {
	// 集合的两种创建方式
	// datalist := make([]model.StartInfo, 0)
	datalist := []models.StarInfo{}
	err := s.engine.Desc("id").Find(&datalist)
	if err != nil {
		log.Println(err)
		return datalist
		// return nil 也可以
	}
	return datalist
}

func (s *ProjectwebDao) Delete(id int) error {
	// 假删除
	data := &models.StarInfo{Id: id, SysStatus: 1}
	_, err := s.engine.Id(data.Id).Update(data)

	return err
}

// columns 判断强制更新
func (s *ProjectwebDao) Update(data *models.StarInfo, columns []string) error {
	_, err := s.engine.Id(data.Id).MustCols(columns...).Update(data)
	// 用到 MustCols 方法
	return err
}

func (s *ProjectwebDao) Create(data *models.StarInfo) error {
	_, err := s.engine.Insert(data)
	return err
}

func (s *ProjectwebDao) Search(country string) []models.StarInfo {
	datalist := []models.StarInfo{}
	err := s.engine.Where("country=?", country).Desc("id").Find(&datalist)
	if err != nil {
		log.Println(err)
		return datalist
	}
	return datalist
}
