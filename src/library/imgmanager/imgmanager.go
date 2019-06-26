package imgmanager

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path"
	"regexp"
	"strings"
	"time"

	"project-web/src/bootstrap/service"

	"github.com/gographics/imagick/imagick"
	"github.com/kataras/iris"
)

// ============

func UploadedImage(file multipart.File, fileheader *multipart.FileHeader, category string, isSave bool) bool {

	var fileName string
	fileName = fileheader.Filename
	if fileName == "" {
		fmt.Println("没有要上传的文件")
	}
	var basePath string
	var configEng *service.Config
	service.GetDi().Container.Invoke(func(config *service.Config) {
		configEng = config
	})

	basePath = configEng.Image.ImagePath

	hashname := MakeImageName(fileName)
	filePath := CreateImagePath(basePath, fileName, category)
	out, err := os.OpenFile(filePath+"/"+hashname, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err.Error())
		return false
	}

	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		return false
	}
	return true
}

/**
 * 根据原图生成缩略图
 *
 * @param string filePath
 * @return boolean
 */
func resizeImageByOrg(filePath string) bool {
	return true
}

/**
 * 根据URI获得图片完整路径
 *
 * @param unknown file
 * @return string
 */
func getImagePath(file string) (filePath string) {
	filePath = ""
	return filePath
}

/**
 * 获得图片对象通过文件路径
 *
 * @param string fileName
 * @return NULL|mixed <NULL, GD, Imagick>
 */
func getImage(fileName string) *imagick.MagickWand {
	mw := imagick.NewMagickWand()
	err := mw.ReadImage("header:")

	if err != nil {
		panic(err)
	}
	image := mw.GetImage()
	return image
}

/**
 * 压缩图片
 *
 * @param unknown width
 * @param unknown height
 * @return boolean
 */
func resizeImage(width, height float32) bool {
	return false
}

/**
 * 剪切图片
 *
 * @param unknown width
 * @param unknown height
 * @param unknown offsetX
 * @param unknown offsetY
 * @return boolean
 */
func cropImage(width, height uint, offsetX, offsetY int) bool {
	mw := imagick.NewMagickWand()
	err := mw.ReadImage("header:")
	if err != nil {
		panic(err)
	}
	err = mw.CropImage(width, height, offsetX, offsetY)
	if err != nil {
		fmt.Println(err.Error)
		return false
	}

	return true
}

/**
 * 保存图片
 *
 * @param unknown file
 * @param number quality = 90 默认画质
 * @return boolean
 */
func saveImage(ctx iris.Context, fileName string, quality int) {
	fmt.Println(fileName)

	ctx.Writef("|%s", "/uploads/"+fileName)
}

/**
 * 根据哈希生成文件名
 *
 * @param unknown imageName
 * @return unknown
 */
func MakeImageName(imageName string) (hashName string) {

	// 	orginal.jpg
	index := strings.LastIndex(imageName, ".")
	if index == -1 {
		index = 0
	}

	substr := imageName[:index]
	suffix := imageName[index:]
	cryptomd5 := md5.New()
	cryptomd5.Write([]byte(substr))
	hashNameByte := cryptomd5.Sum(nil)
	hashName = hex.EncodeToString(hashNameByte) // 转成16进制字符串
	return hashName + suffix
}

/**
 * 根据配置生成路径
 *
 * @param string imageName
 * @param unknown pathConfig
 * @return string
 */
func makeImagePath(imageName, categoryPath string) string {
	var path string
	println("path= ", path)
	return path
}
func makeImagePaths(imageName, pathConfig []string) string {
	path := ""
	return path
}

/**
 * 创建文件路径
 *
 * @param string path
 * @param number permission 默认 0755
 * @return boolean|unknown
 */
func CreateImagePath(basePath, imageName, categoryPath string) string {
	var path string
	path = "org/"
	now := time.Now().Format("2006-01-02")
	hashName := MakeImageName(imageName)

	// 取模运算
	// hashNameByte, _ := hex.DecodeString(hashName)
	// var modResult []byte
	// for _, mbyte := range hashNameByte {
	// 	mod := mbyte % 2
	// 	modResult = append(modResult, mod)
	// }
	// hexMod := hex.EncodeToString(modResult[0:4])

	// 取前两个字符作为 mod 目录
	hexMod := hashName[:2]

	path = basePath + "/" + categoryPath + "/" + now + "/" + hexMod

	os.MkdirAll(path, 0777)

	return path
}

/**
 * 根据mime类型获取扩展名
 * @param string fileType
 * @return string
 */
func getExtendName(fileType string) string {
	realType := "jpg"
	return realType
}

/**
 * 根据缩略图路径获得原图路径
 * @param string path /org/category/2019-12-09/a3/a38b.jpg
 * @return string
 */
func getImageOrgPath(fpath string) string {
	subPaths, _ := path.Split(fpath)

	nativePath := subPaths
	return nativePath
}

/**
 * 根据缩略图路径获得缩略图规格
 * @param unknown path
 * @return string|unknown
 */
func GetImageSizeByPath(path string) string {
	println("GetImageSizeByPath")
	strs := regexp.MustCompile(`/[a-zA-Z0-9]+x[0-9]+/`).FindAllString(path, -1)
	if len(strs) == 0 {
		return ""
	}

	return strs[0]
}

/**
 * 根据图片路径获得图片分类
 * @param unknown path
 */
func GetImageCategoryByPath(path string) string {
	// /org/category/2019-12-09/a3/a38b.jpg
	categorys := strings.Split(path, "/")
	if len(categorys) > 2 {
		return categorys[2]
	}
	return ""
}
