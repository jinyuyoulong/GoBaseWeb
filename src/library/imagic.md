Go语言imagick 使用总结

imagick 是一个开源的 c语言图片处理库，所以需要本地安装，并且配置 CGO

Mac 安装imagic

brew install imagemagick

CGO配置

    export CGO_CFLAGS_ALLOW='-Xpreprocessor' #cgo	

Initialize() 和 Terminate

根据ImageMagick C API，Initialize()应该只调用一次来设置使用ImageMagick的资源。这通常在您main()或init()整个应用程序或库中完成。应用程序可以推迟调用以Terminate()拆除ImageMagick资源。

多次调用特殊方法，导致常见的问题，这是一个错误对于Initialize，和Terminate来说，如死机或丢失代理。除了在程序中对ImageMagick的绝对需求外，不要使用Terminate。

内存管理

由于这是一个CGO绑定，并且Go GC不管理由C API分配的内存，因此必须使用Terminate（）和Destroy（）方法。

通过New*构造函数（MagickWand，DrawingWand，PixelIterator，PixelWand，...）创建的类型由Go GC通过使用终结器进行管理。

如果使用struct literals，则应手动释放资源：
```go
package main

import "github.com/gographics/imagick/imagick"

func main() {
	imagick.Initialize()
	// Schedule cleanup
	defer imagick.Terminate()
	var err error

	mw := imagick.NewMagickWand()

	err = mw.ReadImage("header.png")
	if err != nil {
		panic(err)
	}

	// Get original logo size
	width := mw.GetImageWidth()
	height := mw.GetImageHeight()
	println("width:", width, " ", height)

}
```



