#设置工作区
export GOPATH=/go-workspace-git
export GOBIN=$GOPATH/bin

PATH=$PATH:$GOPATH:$GOBIN
export PATH

#添加代理https://athens.azurefd.net
export GOPROXY="https://goproxy.io" 
#使工作区生效
source ~/.bash_profile

go iris MVC框架 web 模板
### 项目目录结构规范

PROJECT_NAME
├── README.md 介绍软件及文档入口
├── bin 编译好的二进制文件,执行./start.sh自动生成，该目录也用于程序打包
├── build.sh 自动编译的脚本
├── doc 该项目的文档
├── public 公共文件/静态文件
├── views html模板文件
├── lib 项目工具包
└── src 该项目的源代码
    ├── main 项目主函数
    ├── test 测试
    ├── app 项目代码
    └── vendor 存放go的库
        ├── github.com/xxx 第三方库
        └── xxx
