package imgmanager

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"project-web/src/bootstrap/service"

	"github.com/gographics/imagick/imagick"
	"github.com/kataras/iris"
	"github.com/pelletier/go-toml"
)

// var jsonConfig = "{}"
type Mimage struct {
	image_lib      string
	image_path     string
	image_url      string
	image_org      string
	image_tmp      string
	image_types    string
	water_mark     string
	image_categroy map[string]ImageCategroy
}

type ImageCategroy struct {
	paths []string
	sizes []string
}

func newImg() {
	//ReadFile函数会读取文件的全部内容，并将结果以[]byte类型返回
	data, _ := ioutil.ReadFile("imageConfig.json")

	v := Mimage{}
	//读取的数据为json格式，需要进行解码
	_ = json.Unmarshal(data, v)

	fmt.Printf("%v", v)

	// image := map[string]interface{

	// },
	// }
}

// ============

func UploadedImage(ctx iris.Context, files string, category string, isSave bool) {

	imgPath := makeImagePath(files, category)
	fmt.Println(imgPath)
	var fileName string
	fileName = files

	categories := Mimage{}
	// categories := map[string]int{"jpg": 1, "png": 1, "jpeg": 1, "gif": 1}

	if files == "" {
		fmt.Println("没有要上传的文件")
	}

	imageCategory := categories.image_categroy

	if len(imageCategory) == 0 {
		fmt.Println("无效的图片分类")
	}

	var basePath string
	service.GetDi().Container.Invoke(func(config *toml.Tree) {
		basePath = config.Get("image.basePath").(string)
	})

	// savePath = basePath
	// if !isSave {
	// 	savePath = savePath + categories.image_tmp + "/"
	// }
	// savePath = savePath + categroy + "/" + categories.image_org + "/"

	// Print the real file names and their sizes
	// var result []map[string]string
	// key := 0
	// for _, file := range files {

	// fileType = file->getRealType();
	// if (empty(fileType)) {
	//     fileType = file->getType();
	// }
	// extendName = this->getExtendName(fileType);
	// newName = this->makeImageName(file->getName());

	// if (len(imageCategory.paths) == 1) {
	// 	filePath = makeImagePath(newName, imageCategory.paths[0]);
	// } else {
	// 	filePath = makeImagePaths(newName, imageCategory.paths);
	// }

	// if (! this->createImagePath(filePath)) {
	//     result[key]["error"] = "保存失败";
	// }
	// filePath = filePath . newName . "." . extendName;

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
	// key ++;
	// }

	// // TODO BeanStack is_save = true
	ctx.Writef("|%s", "/uploads/"+fileName)

	// return result;
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
 * 移动临时文件到真正的目录
 * 返回真正目录路径
 *
 * @param string tmpFile
 * @return string
 */
func moveTmpFileToPath(tmpFile string) string {
	distFile := ""
	return distFile
}

/**
 *
 * @param unknown uploadFile
 * @param unknown distPath
 * @return boolean
 */
func moveImage(uploadFile, distPath string) (ok bool) {
	return ok
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
 * @param number quality = 90 默认
 * @return boolean
 */
func saveImage(file string, quality int) {

}

/**
 * 根据哈希生成文件名
 *
 * @param unknown imageName
 * @return unknown
 */
func makeImageName(imageName string) string {

	// path 结构 org/category/time/hash.jpg
	// org/category/2019-06-19/8b18e22f914e64a1a933541ed0e97ae0.jpg
	files := imageName
	index := strings.LastIndex(files, ".")
	if index == -1 {
		index = 0
	}
	laststr := files[index:]
	substr := files[:index]

	cryptomd5 := md5.New()
	cryptomd5.Write([]byte(substr))
	hashNameByte := cryptomd5.Sum(nil)
	hashName := hex.EncodeToString(hashNameByte)
	hashNameSub := hashName[len(hashName)-4:]

	return (hashNameSub + "/" + hashName + laststr)
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
	path = "org/"
	path += categoryPath

	now := time.Now().Format("2006-01-02")
	fmt.Println(now)
	path = path + "/" + now
	hashName := makeImageName(imageName)

	path = path + "/" + hashName
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
func createImagePath(path string, permission int) {
	// 	var path string
	// 	path = "org/"
	// 	path += categoryPath

	// 	now := time.Now().Format("2006-01-02")
	// 	fmt.Println(now)
	// 	path = path + "/" + now
	// 	path = path + "/" + makeImageName(imageName)
	// 	return path
}

/**
 * 根据mime类型获取扩展名
 * @param unknown fileType
 * @return string
 */
func getExtendName(fileType string) string {
	realType := "jpg"
	return realType
}

/**
 * 根据缩略图路径获得原图路径
 * @param unknown path
 * @return string
 */
func getImageOrgPath(path string) string {
	nativePath := ""
	return nativePath
}

/**
 * 根据缩略图路径获得缩略图规格
 * @param unknown path
 * @return string|unknown
 */
func getImageSizeByPath(path string) string {
	size := ""
	return size
}

/**
 * 根据图片路径获得图片分类
 * @param unknown path
 */
func getImageCategoryByPath(path string) string {
	category := "jpg"
	return category
}
