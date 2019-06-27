package imgmanager

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"project-web/src/bootstrap/service"

	"github.com/gographics/imagick/imagick"
	"github.com/kataras/iris"
)

var config *service.Config

func init() {
	config = GetConfig()
}

// ============
//
func UploadedImage(file multipart.File, fileheader *multipart.FileHeader, category string, isSave bool) string {

	var fileName string
	fileName = fileheader.Filename
	if fileName == "" {
		fmt.Println("没有要上传的文件")
	}
	var basePath string

	basePath = config.Image.ImagePath

	hashname := MakeImageName(fileName)

	// 创建路径
	fileDictionary := CreateImagePath(basePath, fileName, category)
	filePath := fileDictionary + "/" + hashname
	// 打开文件
	out, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err.Error())
		return ""
	}

	defer out.Close()

	// 写入文件
	_, err = io.Copy(out, file)
	if err != nil {
		fmt.Printf("UploadedImage: %s", err.Error())
		return ""
	}
	return filePath
}

/**
 * 根据原图生成缩略图
 *
 * @param string filePath 经NGINX过滤后，客户端访问的图片路径
 * @return boolean
 */
func ResizeImageByOrg(filePath string) bool {
	orgPath := GetImageOrgPath(filePath) // 替换100x100 获取org path
	if orgPath == "" {
		return false
	}
	category := GetImageCategoryByPath(orgPath)
	size := GetImageSizeByPath(filePath)
	if category == "" || size == "" {
		return false
	}
	// imageCategories := config.Image.ImageCategories
	// category := imageCategories[0]
	carLogo := config.Image.ImageCategory.CarLogo
	ok := false
	for _, value := range carLogo.Sizes {
		if value == size {
			ok = true
			break
		}
	}
	println("ok ", ok)
	if !ok {
		println("图片 file path size 错误")
		return false
	}

	//         if (! isset($imageCategroy->$categroy, $imageCategroy->$categroy->sizes) && in_array($size, $imageCategroy->$categroy->sizes)) {
	//             return false;
	//         }

	//         $sizes = explode('x', $size);
	//         if (empty($sizes)) {
	//             return false;
	//         }

	//         if (! $this->getImage($orgPath)) {
	//             return false;
	//         }

	//         $this->resizeImage($sizes[0], $sizes[1]);

	//         return $this->saveImage($filePath);
	return true
}

/**
 * 根据URI获得图片完整路径
 *
 * @param unknown file
 * @return string
 */
func GetImagePath(file string) (filePath string) {
	filePath = ""
	return filePath
}

/**
 * 获得图片对象通过文件路径
 *
 * @param string fileName
 * @return NULL|mixed <NULL, GD, Imagick>
 */
func GetImage(fileName string) *imagick.MagickWand {
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
func ResizeImage(width, height float32) bool {
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
func MakeImagePath(imageName, categoryPath string) string {
	var path string
	println("path= ", path)
	return path
}
func MakeImagePaths(imageName, pathConfig []string) string {
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
	orgPath := config.Image.ImageOrg
	now := time.Now().Format("2006-01-02")
	hashName := MakeImageName(imageName)

	// 取前两个字符作为 mod 目录
	hexMod := hashName[:2]
	path := filepath.Join(basePath, categoryPath, orgPath, now, hexMod)
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
 * @param string path ./category/200x200/2019-06-24/19/19a34.jpg
 * @return string
 */
func GetImageOrgPath(fpath string) string {
	if fpath == "" {
		return ""
	}

	imageOrg := config.Image.ImageOrg
	imageSize := GetImageSizeByPath(fpath)

	// 返回将s中前n个不重叠old子串都替换为new的新字符串，如果n<0会替换所有old子串。
	strings.Replace(fpath, imageSize, imageOrg, 1)
	// path.Split(fpath)
	return fpath
}

/**
 * 根据缩略图路径获得缩略图规格
 * @param unknown path 200x200/2019-06-24/19/19a34.jpg
 * @return string|unknown
 */
func GetImageSizeByPath(path string) string {
	// over
	strs := regexp.MustCompile(`/[a-zA-Z0-9]+x[0-9]+/`).FindAllString(path, -1)
	if len(strs) == 0 {
		return ""
	}

	return strs[0]
}

/**
 * 根据图片路径获得图片分类
 * @param string path http://static.xxx.com/image/category/200x200/2019-06-24/19/19a34.jpg
 */
func GetImageCategoryByPath(path string) string {
	// over
	categorys := strings.Split(path, "/")
	if len(categorys) > 2 {
		return categorys[2]
	}
	return ""
}

func GetConfig() *service.Config {
	var config *service.Config
	service.GetDi().Container.Invoke(func(conf *service.Config) {
		config = conf
	})
	return config
}
