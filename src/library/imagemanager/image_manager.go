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
	"github.com/kataras/iris/v12"
)

var config *service.Config

func init() {
	config = getConfig()
}

// UploadedImage 接收图片
func UploadedImage(ctx iris.Context, file multipart.File, fileheader *multipart.FileHeader, category string, isSave bool) (result []map[string]string, err error) {
	var isHaveCategory bool
	for _, item := range config.Image.ImageCategroies {
		if item == category {
			isHaveCategory = true
		}
	}

	if !isHaveCategory {
		return nil, errors.New("无效的图片分类")
	}

	var fileName string
	fileName = fileheader.Filename
	if fileName == "" {
		return nil, errors.New("没有要上传的文件")
	}

	basePath := config.Image.ImagePath

	savePath := basePath

	if !isSave { // 如果不保存，则放到tmp目录下
		savePath = filepath.Join(savePath, config.Image.ImageTmp)
		// basepath/tmp/
	}

	// savePath += category + "/" + config.Image.ImageOrg + "/"
	// basepath/category/org/
	savePath = filepath.Join(savePath, category, config.Image.ImageOrg)

	// // Print the real file names and their sizes
	key := 0
	result = []map[string]string{map[string]string{"upload_name": ""}}

	result[key]["upload_name"] = fileName
	contentType := fileheader.Header.Get("Content-Type")
	fileType, _, _ := mime.ParseMediaType(contentType)
	if fileType == "" {
		// 	orginal.jpg
		i := strings.LastIndex(fileName, ".")
		if i == -1 {
			i = 0
		}
		fileType = fileName[i:]
	}
	extendName := getExtendName(fileType)
	newName := makeImageName(fileName) // hash name
	newPath := makeImagePath(newName)  // /10/time/dd34.png
	filePath := filepath.Join(savePath, newPath)
	err = createImagePath(filePath, 0)
	if err != nil {
		result[key]["error"] = "保存失败"
	}

	filePath = filepath.Join(filePath, newName+"."+extendName)

	err = moveImage(ctx, filePath, file)
	if err != nil {
		result[key]["error"] = "保存失败"
	}

	imageURL := strings.Replace(filePath, basePath, config.Image.ImageURL, 1)
	imagePath := strings.Replace(filePath, basePath, "", 1)
	result[key]["file_url"] = imageURL
	result[key]["file_path"] = imagePath

	// TODO get image size
	// orgPath := filePath
	// imagick.Initialize()
	// defer imagick.Terminate()
	// mw := imagick.NewMagickWand()
	// mw.ReadImage(orgPath)
	// result[key]["width"] = strconv.FormatUint(uint64(mw.GetImageWidth()), 10)
	// result[key]["height"] = strconv.FormatUint(uint64(mw.GetImageHeight()), 10)

	// TODO BeanStack is_save = true

	return result, nil
}

/**
* 根据原图生成缩略图
*
* @param string scalaPath 经NGINX过滤后，客户端访问的图片路径
/carlogo/100x100/8b/2019-06-28/8b18.jpg
* @return boolean
*/
func ResizeImageByOrg(scalaPath string) bool {
	orgPath := getImageOrgPath(scalaPath) // 根据缩略图path 获取org path
	orgPath = filepath.Join(config.Image.ImagePath, orgPath)
	ise, err := pathExists(orgPath)
	if !ise {
		return false
	}
	category := getImageCategoryByPath(scalaPath)
	size := getImageSizeByPath(scalaPath)

	if category == "" || size == "" {
		return false
	}
	var isHaveCategory bool
	for _, item := range config.Image.ImageCategroies {
		if item == category {
			isHaveCategory = true
		}
	}

	if !isHaveCategory {
		return false
	}
	var csizes []string
	if category == "carlogo" {
		carLogo := config.Image.ImageCategory.CarLogo
		csizes = carLogo.GetSizes()
	}
	// 判断size是否是受支持的，在配置文件中
	var isSize bool
	for _, value := range csizes {
		if value == size {
			isSize = true
			break
		}
	}
	if !isSize {
		return false
	}
	sizes := strings.Split(size, "x")
	if sizes[0] == "" {
		return false
	}
	width, _ := strconv.Atoi(sizes[0])
	height, _ := strconv.Atoi(sizes[1])

	// get image resize save
	imagick.Initialize()
	defer imagick.Terminate()
	mw := imagick.NewMagickWand()
	err = mw.ReadImage(orgPath)
	if err != nil {
		return false
	}

	err = mw.ResizeImage(uint(width), uint(height), imagick.FILTER_LANCZOS)
	if err != nil {
		return false
	}
	scalaPath = filepath.Join(config.Image.ImagePath, scalaPath)
	scalaDir := filepath.Dir(scalaPath)

	createImagePath(scalaDir, 0)
	err = mw.WriteImage(scalaPath)
	if err != nil {
		return false
	}

	// 不能在拆分逻辑，因为 mw 对象声明周期当前一个函数中
	// mw := getImage(orgPath)
	// isResize := resizeImage(mw, uint(width), uint(height))
	// imagickSaveImage(scalaPath)

	return true
}

/**
 * 根据URI获得图片完整路径
 *
 * @param  URI // https://static.xx.com/app_images/car/org/mod/date/hax.png
 * @return string
 */
func getImagePath(file string) string {
	targetPath := filepath.Join(config.Image.ImagePath, strings.ToLower(file))
	return targetPath
}

/**
 * 根据mine类型获取扩展名
 *
 * @param fileType image/jpeg
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
 * imagick 对象生命周期只在当前函数中，由外部捕获
 * @param string fileName /car/org/mod/date/hax.png ==> ../app_images/car/org/mod/date/hax.png
 * @return NULL|mixed <NULL, GD, Imagick>
 */
// func getImage(orgPath string) *imagick.MagickWand {

// 	if orgPath == "" {
// 		return nil
// 	}
// 	if config.Image.ImageLib == "Imagick" {
// imagick.Initialize()
// defer imagick.Terminate()
// mw := imagick.NewMagickWand()
// mw.ReadImage(orgPath)
// 		return mw
// 	}
// 	return nil
// }

/**
 * 压缩图片
 *
 * @param unknown width
 * @param unknown height
 * @return boolean
 */
func resizeImage(width, height uint) bool {
	imagick.Initialize()
	defer imagick.Terminate()
	mw := imagick.NewMagickWand()
	// mw.ReadImage("aa")
	if mw.GetImage() == nil {
		return false
	}
	err := mw.ResizeImage(width, height, imagick.FILTER_LANCZOS)
	if err != nil {
		return false
	}
	return true
}

/**
 * 剪切图片
 * @param mpath read image path
 * @param unknown width
 * @param unknown height
 * @param unknown offsetX
 * @param unknown offsetY
 * @return boolean
 */
func cropImage(filePath string, width, height uint, offsetX, offsetY int) (err error) {
	imagick.Initialize()
	defer imagick.Terminate()
	mw := imagick.NewMagickWand()
	mw.ReadImage(filePath)
	err = mw.CropImage(width, height, offsetX, offsetY)
	if err != nil {
		return err
	}

	return nil
}

/**
 * 保存图片
 *
 * @param filePath fileDictionary + "/" + hashname
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
	// suffix := imageName[index:]
	cryptomd5 := md5.New()

	cryptomd5.Write([]byte(substr))

	hashNameByte := cryptomd5.Sum(nil)
	hashName = hex.EncodeToString(hashNameByte) // 转成16进制字符串
	return hashName
}

/**
 * 根据配置生成路径
 *
 * @param string imageName
 * @param unknown pathConfig
 * @return string category/org/4a/20190923/4a3f.jpg
 */
func makeImagePath(imageName string) string {
	now := time.Now().Format("2006-01-02")
	// 取前两个字符作为 mod 目录
	hexMod := imageName[:4]
	hexBytes, _ := hex.DecodeString(hexMod)
	number := binary.BigEndian.Uint16(hexBytes)
	// path format : upload/car/org/4a/20190923/4a3f.jpg
	mod := uint16(10)
	resultMod := int(number%mod + 1)
	hexMod = strconv.Itoa(resultMod)
	path := filepath.Join(hexMod, now)
	return path
}

/**
 *
 * @param unknown uploadFile
 * @param unknown distPath
 * @return boolean
 */
func moveImage(ctx iris.Context, distPath string, file multipart.File) error {

	width, _ := ctx.PostValueInt("width")
	height, _ := ctx.PostValueInt("height")
	offsetX, _ := ctx.PostValueInt("offsetX")
	offsetY, _ := ctx.PostValueInt("offsetY")

	if (width != -1) || (height != -1) || (offsetX != -1) || (offsetY != -1) {

		tmpPath := strings.Replace(distPath, config.Image.ImageOrg, config.Image.ImageTmp, 1)
		tmpDir := filepath.Dir(tmpPath)

		createImagePath(tmpDir, 0)

		saveImage(tmpPath, file, 0)

		imagick.Initialize()
		defer imagick.Terminate()
		mw := imagick.NewMagickWand()
		mw.ReadImage(tmpPath)
		err := mw.CropImage(uint(width), uint(height), offsetX, offsetY)
		if err != nil {
			return err
		}
		distDir := filepath.Dir(distPath)
		createImagePath(distDir, 0)
		return mw.WriteImage(distPath)
	}

	distDir := filepath.Dir(distPath)
	createImagePath(distDir, 0)
	return saveImage(distPath, file, 0)
}

/**
 * 创建文件路径
 *
 * @param string path
 * @param number permission 默认 0755 上线可改为 2777 本地该权限无法写入
 * @return boolean|unknown
 */
func createImagePath(mpath string, permission os.FileMode) error {
	isExist, err := pathExists(mpath)
	if isExist {
		return err
	}
	if permission == 0 {
		// TODO
		permission = 0755
	}
	err = os.MkdirAll(mpath, permission)
	if err != nil {
		return err
	}

	return nil
}
func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
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
