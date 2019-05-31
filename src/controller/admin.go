package controller

import (
	"log"
	"time"

	"project-web/src/models"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

// AdminController admin
type AdminController struct {
	Ctx iris.Context
}

const (
	adminTitle string = "管理后台"
)

// Get 根路径
// uri: /admin
func (c *AdminController) Get() mvc.Result {
	starinfo := models.CreateStrInfo()
	datalist := starinfo.GetAll()

	return mvc.View{
		Name: "admin/index.html",
		Data: iris.Map{
			"Title":    adminTitle,
			"Datalist": datalist,
		},
		Layout: "admin/layout.html",
	}
}

// GetEdit 编辑
//uri: /admin/edit
func (c *AdminController) GetEdit() mvc.Result {
	starinfo := models.CreateStrInfo()
	id, err := c.Ctx.URLParamInt("id")
	var data *models.StarInfo
	if err == nil {
		data = starinfo.GetStarInfoInfo(id)
	} else {
		data = &models.StarInfo{
			Id: 0,
		}
	}

	return mvc.View{
		Name: "admin/edit.html",
		Data: iris.Map{
			"Title": adminTitle,
			"info":  data,
		},
		Layout: "admin/layout.html",
	}
}

// PostSave 上传数据
// post uri: /admin/save
func (c *AdminController) PostSave() mvc.Result {
	info := models.CreateStrInfo()
	err := c.Ctx.ReadForm(&info) // 结合 models 中填写的 form 信息 使用
	if err != nil {
		log.Fatal(err)
	}

	if info.Id > 0 { // 更新
		info.SysUpdated = int(time.Now().Unix())
		info.UpdateStarInfo(&info, []string{"name_zh", "name_en", "avatar", "birthday", "height", "weight", "club", "jersy", "country", "moreinfo"})
	} else { // 创建
		info.SysCreated = int(time.Now().Unix())
		info.CreateStarInfo(&info)
	}

	return mvc.Response{
		Path: "/admin/",
	}
}

// GetDelete 删除
// get uri: /admin/delete?id=1
func (c *AdminController) GetDelete() mvc.Result {
	starinfo := models.CreateStrInfo()
	id, err := c.Ctx.URLParamInt("id")
	if err == nil {
		starinfo.Delete(id)
	}

	return mvc.Response{
		Path: "/admin/",
	}
}
