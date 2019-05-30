grpc-wrapper 版本 无release
依赖 github.com/opentracing/opentracing-go v1.1.0

详细实现文档见 docs

grpc的封装扩展,集成通用的组件,形成一个微服务通讯框架.

1.支持的扩展
服务注册与发现
etcd [OK]

结构化日志 
zap [OK]

服务调用链追踪
zipkin [OK] jaeger [OK]

服务指标监控
falcon-plus [TODO]