// file: controller/user_controller.go

package controller

import (
	"fmt"
	"mime/multipart"

	"project-web/src/bootstrap/service"
	"project-web/src/library/imgmanager"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/pelletier/go-toml"
)

// ImageController index controller
type ImageController struct {
	Ctx iris.Context
	Di  iris.Context
}

func (c *ImageController) GetUpload(ctx iris.Context) mvc.Result {
	var configEng *toml.Tree
	service.GetDi().Container.Invoke(func(config *toml.Tree) {
		configEng = config
	})
	category := configEng.Get("image.image_categroy").([]*toml.Tree)
	fmt.Printf("%d %T", len(category), category)
	path := category[0]

	fmt.Printf("%v %T", path, path)

	return mvc.View{
		Name: "index/upload_form.html",
		Data: iris.Map{
			"Title": commonTitle,
		},
	}
}

func (c *ImageController) PostUpload(ctx iris.Context) {
	// 处理上传请求，并保存文件到目录
	// app.Post("/upload", iris.LimitRequestBodySize(maxSize+1<<20), func(ctx iris.Context) {
	_, _, err := ctx.FormFile("uploadfile")

	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.HTML("Error while uploading: <b>" + err.Error() + "</b>")
		return
	}
	// defer file.Close()

	// fname := fileHeader.Filename
	// hashname := imgmanager.MakeImageName(fname)

	// var imgPath string
	var configEng *toml.Tree
	service.GetDi().Container.Invoke(func(config *toml.Tree) {
		configEng = config
	})
	car_logo := configEng.Get("image.image_categroy.car_logo").(*toml.Tree)
	fmt.Printf("%v %T", car_logo[0].Get("paths").(string), car_logo)
	// ok := imgmanager.UploadedImage(file, fileHeader, category[0], true)

	// // filePath := imgmanager.CreateImagePath(file, imgPath, fname, "category")

	// // out, err := os.OpenFile(filePath+"/"+hashname, os.O_WRONLY|os.O_CREATE, 0666)
	// // if err != nil {
	// // 	ctx.StatusCode(iris.StatusInternalServerError)
	// // 	ctx.HTML("Error while uploading: <b>" + err.Error() + "</b>")
	// // 	return
	// // }

	// // defer out.Close()
	// // io.Copy(out, file)
	// var msg string
	// if ok {
	// 	msg = "图片上传成功"
	// } else {
	// 	msg = "图片上传失败"
	// }
	// ctx.JSON(iris.Map{
	// 	"status":  true,
	// 	"message": msg,
	// })
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
	service.GetDi().Container.Invoke(func(config *toml.Tree) {
		imgPath = config.Get("image.imgPath").(string)
	})

	_, fileHeader, _ := ctx.FormFile("uploadfile")

	filePath := imgmanager.CreateImagePath(imgPath, fileHeader.Filename, "category")
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
