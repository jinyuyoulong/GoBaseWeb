// file: controller/user_controller.go

package controller

import (
	"fmt"
	"mime/multipart"

	"project-web/src/bootstrap/service"
	"project-web/src/library/imgmanager"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

// ImageController index controller
type ImageController struct {
	Ctx iris.Context
	Di  iris.Context
}

func (i *ImageController) Get() {
	path := i.Ctx.Path()
	host := i.Ctx.Host()
	println("get")
	println(host)
	fmt.Printf("%v\n", path)
	// imgmanager.ResizeImageByOrg(path)
	i.Ctx.Writef("%s", path)
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

	filePath := imgmanager.UploadedImage(file, fileHeader, conf.Image.ImageCategroies[0], true)

	var msg string
	if filePath != "" {
		msg = "图片上传成功"
	} else {
		msg = "图片上传失败"
	}
	println(filePath)
	ctx.JSON(iris.Map{
		"status":  true,
		"message": msg,
	})
}
func (c *ImageController) GetUpload2(ctx iris.Context) mvc.Result {
	fmt.Println("get Upload2")
	return mvc.View{
		Name: "image/index.html",
		Data: iris.Map{
			"Title": commonTitle,
		},
	}
}

// 多文件上传，求模运算的时候会将多文件看为一个文件 进行计算。计算结果将和单文件上传的地址不一样。不建议使用
func (c *ImageController) PostUpload2(ctx iris.Context) {
	fmt.Println("PostUpload2 只能浏览器测试，不能用 postman")
	// 5MB
	const maxSize = 5 << 20
	iris.LimitRequestBodySize(maxSize + 1<<20)

	var imgPath string
	var conf *service.Config
	service.GetDi().Container.Invoke(func(config *service.Config) {
		conf = config
	})
	imgPath = conf.Image.ImagePath
	// _, fileHeader, _ := ctx.FormFile("uploadfile")

	filePath := imgmanager.CreateImagePath(imgPath, 0777)
	println(filePath)
	// iris 保存图片
	ctx.UploadFormFiles(filePath, func(ctx iris.Context, file *multipart.FileHeader) {
		hashname := imgmanager.MakeImageName(file.Filename)
		file.Filename = hashname
		fileName := filePath + "/" + file.Filename
		println(fileName)
		// 这里的路径要和上面填的保持一致
		ctx.Writef("%s", fileName)
	})
}

func (i *ImageController) GetCreateimage() {
	path := i.Ctx.Path()
	host := i.Ctx.Host()
	rPath := i.Ctx.RequestPath(true)
	println("CreateResizeOrgImage")
	println(host)
	fmt.Printf("%v\n", path)
	fmt.Printf("requestPath %v\n", rPath)
	// imgmanager.ResizeImageByOrg(path)
	i.Ctx.Writef("%s", path)
}
