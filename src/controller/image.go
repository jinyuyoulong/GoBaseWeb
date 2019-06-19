// file: controller/user_controller.go

package controller

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"

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
	file, fileHeader, err := ctx.FormFile("uploadfile")

	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.HTML("Error while uploading: <b>" + err.Error() + "</b>")
		return
	}
	defer file.Close()

	// imgmanager.UploadedImage(file, fileHeader, true)

	fname := fileHeader.Filename
	fmt.Println(fname)
	out, err := os.OpenFile("../uploads/"+fname, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.HTML("Error while uploading: <b>" + err.Error() + "</b>")
		return
	}

	defer out.Close()
	io.Copy(out, file)

	ctx.JSON(iris.Map{
		"status":  true,
		"message": "ok",
	})
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
func (c *ImageController) PostUpload2(ctx iris.Context) {
	// file, fileHeader, err := ctx.FormFile("uploadfile")
	// if err != nil {
	// 	panic(err.Error)
	// }
	// fmt.Printf("%v\n%v", file, fileHeader)
	fmt.Println("PostUpload2 只能浏览器测试，不能postman")
	// 5MB
	const maxSize = 5 << 20
	iris.LimitRequestBodySize(maxSize + 1<<20)
	var imgPath string
	service.GetDi().Container.Invoke(func(config *toml.Tree) {
		imgPath = config.Get("image.imgPath").(string)
	})

	// iris 保存图片
	ctx.UploadFormFiles(imgPath, beforeSave)

}

// 保存时文件名出来
func beforeSave(ctx iris.Context, file *multipart.FileHeader) {
	fmt.Println("beforeSave")

	// file name struct: __1-name.jpg
	// ip := ctx.RemoteAddr()
	// ip = strings.Replace(ip, ".", "_", -1)
	// ip = strings.Replace(ip, ":", "_", -1)
	// file.Filename = ip + "-" + file.Filename
	// ctx.Writef("|%s", "/uploads/"+file.Filename)

	imgmanager.UploadedImage(ctx, file.Filename, "category", true)
}
