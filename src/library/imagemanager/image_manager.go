package imagemanager

import (
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"io"
	"mime"
	"mime/multipart"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"project-web/src/bootstrap/service"

	"github.com/gographics/imagick/imagick"
)

var config *service.Config

func init() {
	config = getConfig()
}

// UploadedImage 接收图片
func UploadedImage(file multipart.File, fileheader *multipart.FileHeader, category string, isSave bool) (filePath string, err error) {

	var isHaveCategory bool
	for _, item := range config.Image.ImageCategroies {
		if item == category {
			isHaveCategory = true
		}
	}

	if !isHaveCategory {
		return "", errors.New("无效的图片分类")
	}

	var fileName string
	fileName = fileheader.Filename
	if fileName == "" {
		return "", errors.New("没有要上传的文件")
	}

	hashname := makeImageName(fileName)

	// 创建路径
	fileDictionary := makeImagePath(fileName)
	err = createImagePath(fileDictionary, 0)
	if err != nil {
		return "", err
	}

	//fileDictionary + "/" + hashname
	filePath = filepath.Join(fileDictionary, hashname)
	err = saveImage(filePath, file, 0)
	if err != nil {
		return "", err
	}

	// ===========

	basePath := config.Image.ImagePath

	savePath := basePath
	if !isSave {
		savePath = savePath + config.Image.ImageTmp + "/"
	}

	savePath += category + "/" + config.Image.ImageOrg + "/"

	// Print the real file names and their sizes
	var result []map[string]string
	key := 0

	// for _, file := range files {
	// // TODO
	// result[key]["upload_name"] = file->getName();
	// result[key]["error"] = file->getError();

	result[key]["upload_name"] = fileName
	// MIMEHeader代表一个MIME头，将键映射为值的集合。
	// type MIMEHeader map[string][]string
	content_type := fileheader.Header.Get("Content-Type")
	fileType, _, _ := mime.ParseMediaType(content_type)
	// if fileType == "" {
	//     fileType = file->getType();
	// }
	extendName := getExtendName(fileType)
	newName := makeImageName(fileName)
	newPath := makeImagePath(newName)
	filePath = savePath + newPath
	err = createImagePath(filePath, 0)
	if err != nil {
		result[key]["error"] = "保存失败"
	}
	filePath = filePath + newName + "." + extendName

	// if (! this->moveImage(file, filePath)) {
	//     result[key]["error"] = "保存失败";
	// }
	// imageUrl = str_replace(basePath, this->config->image->image_url, filePath);
	// imagePath = str_replace(basePath, "", filePath);
	// result[key]["file_url"] = imageUrl;
	// result[key]["file_path"] = imagePath;

	// // TODO get image size
	// image = this->getImage(filePath);
	// result[key]["width"] = image->getWidth();
	// result[key]["height"] = image->getHeight();
	key++
	// }
	// TODO BeanStack is_save = true

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

	category := getImageCategoryByPath(scalaPath)

	size := getImageSizeByPath(scalaPath)

	if category == "" || size == "" {
		return errors.New("图片路径错误")
	}

	var isHaveCategory bool
	for _, item := range config.Image.ImageCategroies {
		if item == category {
			isHaveCategory = true
		}
	}

	if !isHaveCategory {
		return errors.New("无效的图片分类")
	}

	var csizes []string
	if category == "carlogo" {
		carLogo := config.Image.ImageCategory.CarLogo
		csizes = carLogo.GetSizes()
	}
	// 判断size是否是受支持的，在配置文件中
	var ok bool
	for _, value := range csizes {
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

	orgPath := getImageOrgPath(scalaPath) // 根据缩略图path 获取org path
	if orgPath == "" {
		return errors.New("org 图片为空")
	}

	// get image resize save
	image := getImage(orgPath, scalaPath, uint(width), uint(height))
	if image != nil {
		return errors.New("获取源图失败")
	}

	// scalaImage := ResizeImage(width, height)
	// return this->saveImage(filePath);
	return nil
}

/**
 * 根据URI获得图片完整路径
 *
 * @param  URI // https://static.xx.com/app_images/car/org/mod/date/hax.png
 * @return string
 */
func getImagePath(rPath string) string {
	targetPath := filepath.Join(config.Image.ImagePath, rPath)
	return targetPath
}

/**
 * 根据mine类型获取扩展名
 *
 * @param unknown fileType
 * @return string
 */
func getExtendName(fileType string) string {
	if fileType == "" {
		return ""
	}

	res := strings.Split(fileType, "/")[1]

	// // 修正ie pjpeg 问题
	res = strings.ToLower(res)
	if res == "pjpeg" {
		res = "jpg"
	}
	return res
}

/**
 * 获得图片对象通过文件路径
 *
 * @param string fileName /car/org/mod/date/hax.png ==> ../app_images/car/org/mod/date/hax.png
 * @return NULL|mixed <NULL, GD, Imagick>
 */
func getImage(orgPath, scalaPath string, swidth, sheight uint) error {
	orgPath = getImagePath(orgPath)
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
	//  压缩裁剪 image
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

	createImagePath(scalaPath, 0)
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
func cropImage(mpath string, width, height uint, offsetX, offsetY int) error {
	mpath = getImagePath(mpath)
	mw := imagick.NewMagickWand()
	err := mw.ReadImage(mpath)
	if err != nil {
		return err
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
 * @return boolean
 */
func saveImage(filePath string, file multipart.File, permission os.FileMode) error {
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
func makeImageName(imageName string) (hashName string) {

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
func makeImagePath(imageName string) string {
	basePath := config.Image.ImagePath
	categoryPath := config.Image.ImageCategroies[0]
	orgPath := config.Image.ImageOrg
	now := time.Now().Format("2006-01-02")
	hashName := makeImageName(imageName)

	// 取前两个字符作为 mod 目录
	hexMod := hashName[:4]
	hexBytes, _ := hex.DecodeString(hexMod)
	number := binary.BigEndian.Uint16(hexBytes)
	// path format : upload/car/org/4a/20190923/4a3f.jpg
	mod := uint16(10)
	resultMod := int(number%mod + 1)
	hexMod = strconv.Itoa(resultMod)
	path := filepath.Join(basePath, categoryPath, orgPath, hexMod, now)
	println(path)
	return path
}

/**
 * 创建文件路径
 *
 * @param string path
 * @param number permission 默认 0755 上线可改为 2777 本地该权限无法写入
 * @return boolean|unknown
 */
func createImagePath(mpath string, permission os.FileMode) error {
	if permission == 0 {
		// TODO
		permission = 0755
	}
	err := os.MkdirAll(mpath, permission)
	if err != nil {
		return err
	}

	return nil
}

/**
 * 根据缩略图路径获得原图路径
 * @param string path ./category/200x200/2019-06-24/19/19a34.jpg
 * @return string
 */
func getImageOrgPath(scalaPath string) string {
	if scalaPath == "" {
		return ""
	}

	imageOrg := config.Image.ImageOrg
	imageSize := getImageSizeByPath(scalaPath)

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
func getImageSizeByPath(path string) string {
	// over
	strs := regexp.MustCompile(`[a-zA-Z0-9]+x[0-9]+`).FindAllString(path, -1)
	if len(strs) == 0 {
		return ""
	}

	return strs[0]
}

/**
 * 根据图片路径获得图片分类
 * @param string path /category/200x200/2019-06-24/19/19a34.jpg
 */
func getImageCategoryByPath(mpath string) string {
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

func getConfig() *service.Config {
	var config *service.Config
	service.GetDi().Container.Invoke(func(conf *service.Config) {
		config = conf
	})
	return config
}
