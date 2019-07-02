// file: controller/user_controller.go

package controller

import (
	"project-web/src/bootstrap/service"
	"project-web/src/library/imagemanager"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

// ImageController index controller
type ImageController struct {
	Ctx iris.Context
	Di  iris.Context
}

func (c *ImageController) GetUpload(ctx iris.Context) mvc.Result {

	return mvc.View{
		Name: "index/upload_form.html",
		Data: iris.Map{
			"Title": commonTitle,
		},
	}
}

func (c *ImageController) PostUpload(ctx iris.Context) {

	file, fileHeader, err := ctx.FormFile("uploadfile")

	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.HTML("Error while uploading: <b>" + err.Error() + "</b>")
		return
	}
	defer file.Close()

	var conf *service.Config
	service.GetDi().Container.Invoke(func(config *service.Config) {
		conf = config
	})

	filePath, err := imagemanager.UploadedImage(file, fileHeader, conf.Image.ImageCategroies[0], true)

	if err != nil {
		ctx.Application().Logger().Warnf(err.Error())
	}

	ctx.Writef("%v", filePath)
}

// GetResizeimage 生成指定的缩略图
// 接受参数格式：http://static.com/image/resizeimage?path=/carlogo/100x100/8b/2019-07-01/8b18.jpg
func (i *ImageController) GetResizeimage() {
	tPath := i.Ctx.URLParam("path")
	host := i.Ctx.Host()
	err := imagemanager.ResizeImageByOrg(tPath)
	if err != nil {
		i.Ctx.Writef("%v", err.Error())
	} else {
		URL := host + tPath
		i.Ctx.Writef("%v", URL)
		// i.Ctx.Redirect(host + tPath)
	}
}
