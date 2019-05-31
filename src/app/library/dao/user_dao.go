package dao

// data access object，DAO
// 数据库访问控制层
import (
	"log"

	"project-web/src/app/models"

	"github.com/go-xorm/xorm"
)

type UserDao struct {
	engine *xorm.Engine
}

func NewUserDao(engine *xorm.Engine) *UserDao {

	return &UserDao{
		engine: engine,
	}
}

func (s *UserDao) Get(id int) *models.User {
	data := &models.User{Id: id}
	ok, err := s.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		data.Id = 0
		return data
	}
}

func (s *UserDao) GetAll() []models.User {
	// 集合的两种创建方式
	datalist := []models.User{}
	err := s.engine.Desc("id").Find(&datalist)
	if err != nil {
		log.Println(err)
		return datalist
		// return nil 也可以
	}
	return datalist
}

func (s *UserDao) Delete(id int) error {
	// 删除
	data := &models.User{Id: id}
	_, err := s.engine.Id(data.Id).Delete(data)

	return err
}

// columns 判断强制更新
func (s *UserDao) Update(data *models.User, columns []string) error {
	_, err := s.engine.Id(data.Id).MustCols(columns...).Update(data)
	// 用到 MustCols 方法
	return err
}

func (s *UserDao) Create(data *models.User) error {
	_, err := s.engine.Insert(data)
	return err
}
