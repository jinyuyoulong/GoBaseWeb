#项目信息

### go 开发环境安装

1. 在官网下载pkg安装包,点击安装，安装完成后，资源被写入/usr/local/go 目录下
2. 设置profile文件，我用的zsh，修改~/.zshrc。用bash的同理，修改~/.bash_profile
       # go
       export GOROOT=/usr/local/go
       export GOPATH=~/dev/go/golib:~/dev/go/project #工作区，存放go源码文件的目录
       export GOBIN=~/dev/go/gobin #存放编译后可执行文件的目录
       export PATH=$PATH:$GOROOT/bin/:$GOBIN

3.验证，命令行执行go version返回go version go1.11.5 darwin/amd64。go语言环境配置完毕。
