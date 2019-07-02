package imgmanager

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"project-web/src/bootstrap/service"

	"github.com/gographics/imagick/imagick"
)

var config *service.Config

func init() {
	config = GetConfig()
}

// UploadedImage 接收图片
func UploadedImage(file multipart.File, fileheader *multipart.FileHeader, category string, isSave bool) (filePath string, err error) {

	var fileName string
	fileName = fileheader.Filename
	if fileName == "" {
		return "", errors.New("没有要上传的文件")
	}

	hashname := MakeImageName(fileName)

	// 创建路径
	fileDictionary := MakeImagePath(fileName)
	CreateImagePath(fileDictionary, 0)
	//fileDictionary + "/" + hashname
	filePath = filepath.Join(fileDictionary, hashname)
	err = SaveImage(filePath, file, 0)
	if err != nil {
		return "", err
	}
	return filePath, nil
}

/**
* 根据原图生成缩略图
*
* @param string scalaPath 经NGINX过滤后，客户端访问的图片路径
/carlogo/100x100/8b/2019-06-28/8b18.jpg
* @return boolean
*/
func ResizeImageByOrg(scalaPath string) error {

	category := GetImageCategoryByPath(scalaPath)

	size := GetImageSizeByPath(scalaPath)

	if category == "" || size == "" {
		return errors.New("图片路径错误")
	}

	oType := reflect.TypeOf(config.Image.ImageCategory)
	var isHaveCate bool
	var index int
	for i := 0; i < oType.NumField(); i++ {
		f := oType.Field(i)
		if f.Name == category {
			isHaveCate = true
			index = i
		}
	}

	confCategory := config.Image.ImageCategroies[index]

	if confCategory != category && !isHaveCate {
		return errors.New("图片没有当前分类")
	}

	carLogo := config.Image.ImageCategory.CarLogo

	// 判断size是否是受支持的，在配置文件中
	var ok bool
	for _, value := range carLogo.Sizes {
		if value == size {
			ok = true
			break
		}
	}
	if !ok {
		return errors.New("请求图片规格错误")
	}
	sizes := strings.Split(size, "x")
	width, _ := strconv.Atoi(sizes[0])
	height, _ := strconv.Atoi(sizes[1])

	orgPath := GetImageOrgPath(scalaPath) // 根据缩略图path 获取org path
	if orgPath == "" {
		return errors.New("org 图片为空")
	}

	// get image resize save
	image := GetImage(orgPath, scalaPath, uint(width), uint(height))
	if image != nil {
		return errors.New("获取源图失败")
	}

	// scalaImage := ResizeImage(width, height)
	//         return $this->saveImage($filePath);
	return nil
}

/**
 * 根据URI获得图片完整路径
 *
 * @param  URI // https://static.xx.com/app_images/car/org/mod/date/hax.png
 * @return string
 */
func GetImagePath(rPath string) string {
	return rPath
}

/**
 * 获得图片对象通过文件路径
 *
 * @param string fileName /car/org/mod/date/hax.png ==> ../app_images/car/org/mod/date/hax.png
 * @return NULL|mixed <NULL, GD, Imagick>
 */
func GetImage(orgPath, scalaPath string, swidth, sheight uint) error {
	orgPath = filepath.Join(config.Image.ImagePath, orgPath)
	// realPath := strings.Replace(orgPath)

	imagick.Initialize()
	// Schedule cleanup
	defer imagick.Terminate()
	var err error

	mw := imagick.NewMagickWand()

	err = mw.ReadImage(orgPath)
	if err != nil {
		return err
	}

	// mw.GetImage()

	// Get original logo size
	width := mw.GetImageWidth()
	height := mw.GetImageHeight()
	// ----------
	//  压缩image
	width = swidth
	height = sheight

	err = mw.ResizeImage(width, height, imagick.FILTER_LANCZOS)

	if err != nil {
		return err
	}
	// ---------
	scalaPath = filepath.Join(config.Image.ImagePath, scalaPath)
	scalaPaths := strings.Split(scalaPath, "/")
	scalaPath = ""
	fileName := scalaPaths[len(scalaPaths)-1]
	for i := 0; i < len(scalaPaths)-1; i++ {
		subPath := scalaPaths[i]
		scalaPath += subPath + "/"
	}

	CreateImagePath(scalaPath, 0)
	scalaFile := scalaPath + fileName
	//导出图片
	err = mw.WriteImage(scalaFile)
	if err != nil {
		return err
	}
	return nil
}

/**
 * 压缩图片
 *
 * @param unknown width
 * @param unknown height
 * @return boolean
 */
// func ResizeImage(image *imagick.MagickWand, width, height uint) bool {

// 	return false
// }

/**
 * 剪切图片
 *
 * @param unknown width
 * @param unknown height
 * @param unknown offsetX
 * @param unknown offsetY
 * @return boolean
 */
func cropImage(width, height uint, offsetX, offsetY int) error {
	mw := imagick.NewMagickWand()
	err := mw.ReadImage("header:")
	if err != nil {
		panic(err)
	}
	err = mw.CropImage(width, height, offsetX, offsetY)
	if err != nil {
		return err
	}

	return nil
}

/**
 * 保存图片
 *
 * @param unknown file
 * @param number quality = 90 默认画质
 * @return boolean
 */
func SaveImage(filePath string, file multipart.File, permission os.FileMode) error {
	// 打开文件
	//  default 0666
	if permission == 0 {
		permission = 0666
	}
	out, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, permission)
	if err != nil {
		return err
	}

	defer out.Close()

	// 写入文件
	_, err = io.Copy(out, file)
	if err != nil {
		return err
	}
	return nil
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
func MakeImagePath(imageName string) string {
	basePath := config.Image.ImagePath
	categoryPath := config.Image.ImageCategroies[0]
	orgPath := config.Image.ImageOrg
	now := time.Now().Format("2006-01-02")
	hashName := MakeImageName(imageName)

	// 取前两个字符作为 mod 目录
	hexMod := hashName[:2]
	// path format : upload/car/org/4a/20190923/4a3f.jpg
	path := filepath.Join(basePath, categoryPath, orgPath, hexMod, now)

	return path
}

/**
 * 创建文件路径
 *
 * @param string path
 * @param number permission 默认 0755
 * @return boolean|unknown
 */
func CreateImagePath(mpath string, permission os.FileMode) string {
	if permission == 0 {
		permission = 0777
	}
	os.MkdirAll(mpath, permission)

	return mpath
}

/**
 * 根据缩略图路径获得原图路径
 * @param string path ./category/200x200/2019-06-24/19/19a34.jpg
 * @return string
 */
func GetImageOrgPath(scalaPath string) string {
	if scalaPath == "" {
		return ""
	}

	imageOrg := config.Image.ImageOrg
	imageSize := GetImageSizeByPath(scalaPath)

	// 返回将s中前n个不重叠old子串都替换为new的新字符串，如果n<0会替换所有old子串。
	orgPath := strings.Replace(scalaPath, imageSize, imageOrg, 1)
	// path.Split(fpath)
	return orgPath
}

/**
 * 根据缩略图路径获得缩略图规格
 * @param unknown path 200x200/2019-06-24/19/19a34.jpg
 * @return string|unknown
 */
func GetImageSizeByPath(path string) string {
	// over
	strs := regexp.MustCompile(`[a-zA-Z0-9]+x[0-9]+`).FindAllString(path, -1)
	if len(strs) == 0 {
		return ""
	}

	return strs[0]
}

/**
 * 根据图片路径获得图片分类
 * @param string path http://static.xxx.com/image/category/200x200/2019-06-24/19/19a34.jpg
 */
func GetImageCategoryByPath(mpath string) string {
	mpath = strings.Replace(mpath, "", config.Image.ImagePath, 0)
	// over
	categorys := strings.Split(mpath, "/")
	// for i, cate := range categorys {
	// 	fmt.Printf("categorys %d %v\n", i, cate)
	// }

	if len(categorys) > 2 {
		return categorys[1]
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
