# 项目信息

### go 开发环境安装

1. 在官网下载pkg安装包,点击安装，安装完成后，资源被写入/usr/local/go 目录下
2. 设置profile文件，我用的zsh，修改~/.zshrc。用bash的同理，修改~/.bash_profile
```
# go
export GOROOT=/usr/local/go
export GOPATH=~/dev/go/golib:~/dev/go/project #工作区，存放go源码文件的目录
export GOBIN=~/dev/go/gobin #存放编译后可执行文件的目录
export PATH=$PATH:$GOROOT/bin/:$GOBIN
# 添加代理 https://athens.azurefd.net
export GOPROXY="https://goproxy.io"
```
3. 验证，命令行执行go version返回go version go1.11.5 darwin/amd64。go语言环境配置完毕。

---

go iris web 框架模板

### 项目目录结构规范

PROJECT_NAME
├── README.md 介绍软件及文档入口
├── bin 编译好的二进制文件,执行./start.sh自动生成，该目录也用于程序打包
├── build.sh 自动编译的脚本
├── doc 该项目的文档
├── public 公共文件/静态文件
├── views html模板文件
├── lib 项目工具包
├── vendor 存放第三方库，go mod vendor 自动解析
└── src 该项目的源代码
    ├── main 项目主函数
    ├── test 测试
    ├── app 项目代码
    └── vendor 存放第三方库
        ├── github.com/xxx 第三方库
        └── xxx
