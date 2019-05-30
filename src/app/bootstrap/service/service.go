package service

// HTTP request 访问控制层
import (
	"ielpm.cn/projectweb/src/app/library/dao"
	"ielpm.cn/projectweb/src/app/library/datasource"
	"ielpm.cn/projectweb/src/app/models"
)

type ProjectapiService interface {
	GetAll() []models.StarInfo
	Get(id int) *models.StarInfo
	Delete(id int) error
	Update(user *models.StarInfo, columns []string) error
	Create(user *models.StarInfo) error
	Search(country string) []models.StarInfo
}
type ProjectService struct {
	dao *dao.ProjectwebDao
}

func NewprojectapiService() *ProjectService {
	return &ProjectService{
		dao: dao.NewProjectwebDao(datasource.InstanceMaster()),
	}
}
func (s *ProjectService) GetAll() []models.StarInfo {
	return s.dao.GetAll()
}
func (s *ProjectService) Get(id int) *models.StarInfo {
	return s.dao.Get(id)
}
func (s *ProjectService) Delete(id int) error {
	return s.dao.Delete(id)
}
func (s *ProjectService) Update(user *models.StarInfo, columns []string) error {
	return s.dao.Update(user, columns)
}
func (s *ProjectService) Create(user *models.StarInfo) error {
	return s.dao.Create(user)
}
func (s *ProjectService) Search(country string) []models.StarInfo {
	return s.dao.Search(country)
}
