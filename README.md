go iris 框架 web 模板

本文档继承于 projectapi 项目的readme
其他内容请参考 projectapi 项目的readme

### 项目目录结构规范

```
PROJECT_NAME
├── README.md 介绍软件及文档入口
├── bin 编译好的二进制文件,执行./build.sh自动生成，该目录也用于程序打包
├── build.sh 自动编译的脚本
├── doc 该项目的文档
├── public 公共文件/静态文件
├── views html模板文件
├── lib 第三方包
└── src 该项目的源代码
    ├── main 项目主函数
    ├── test 测试
    ├── app 项目代码
    └── vendor 存放go的库
        ├── github.com/xxx 第三方库
        └── xxx.com/abc 公司内部的公共库
```