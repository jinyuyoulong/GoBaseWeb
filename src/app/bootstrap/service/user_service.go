package service

// HTTP request 访问控制层
import (
	"project-web/src/app/library/dao"
	"project-web/src/app/library/datasource"
	"project-web/src/app/models"
)

type UserService struct {
	dao *dao.UserDao
}

func NewUserService() *UserService {
	return &UserService{
		dao: dao.NewUserDao(datasource.InstanceMaster()),
	}
}
func (s *UserService) GetAll() []models.User {
	return s.dao.GetAll()
}
func (s *UserService) Get(id int) *models.User {
	return s.dao.Get(id)
}
func (s *UserService) Delete(id int) error {
	return s.dao.Delete(id)
}
func (s *UserService) Update(user *models.User, columns []string) error {
	return s.dao.Update(user, columns)
}
func (s *UserService) Create(user *models.User) error {
	return s.dao.Create(user)
}
