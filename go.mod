module xxx.com/projectweb

go 1.12

replace (
	golang.org/x/crypto@v0.0.0-20180820150726-614d502a4dac => github.com/golang/crypto v0.0.0-20190313024323-a1f597ede03a
	golang.org/x/lint@v0.0.0 => github.com/golang/lint v0.0.0-20190313153728-d0100b6bd8b3
	golang.org/x/sys@v0.3.0 => github.com/golang/sys v0.3.0
	golang.org/x/tools@v0.0.0 => github.com/golang/tools v0.0.0-20190315214010-f0bfdbff1f9c
)

replace golang.org/x/net v0.0.0 => github.com/golang/net v0.0.0-20190313220215-9f648a60d977

require (
	cloud.google.com/go v0.37.4 // indirect
	github.com/360EntSecGroup-Skylar/excelize v1.3.0
	github.com/Joker/jade v1.0.0 // indirect
	github.com/Shopify/goreferrer v0.0.0-20181106222321-ec9c9a553398 // indirect
	github.com/ajg/form v1.5.1 // indirect
	github.com/aymerick/raymond v2.0.2+incompatible // indirect
	github.com/beanstalkd/go-beanstalk v0.0.0-20190515041346-390b03b3064a
	github.com/codahale/hdrhistogram v0.0.0-20161010025455-3a0bb77429bd // indirect
	github.com/coreos/etcd v3.3.13+incompatible
	github.com/eknkc/amber v0.0.0-20171010120322-cdade1c07385 // indirect
	github.com/fatih/structs v1.1.0 // indirect
	github.com/flosch/pongo2 v0.0.0-20181225140029-79872a7b2769 // indirect
	github.com/g4zhuj/go-metrics-falcon v0.0.0-20180427054158-5159ced4eafb
	github.com/g4zhuj/grpc-wrapper v0.0.0-20190508092021-ced55bb6c5d6
	github.com/gavv/monotime v0.0.0-20171021193802-6f8212e8d10d // indirect
	github.com/go-redis/redis v6.15.2+incompatible
	github.com/go-sql-driver/mysql v1.4.1
	github.com/go-xorm/xorm v0.7.1
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/gorilla/securecookie v1.1.1
	github.com/gorilla/websocket v1.4.0 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.0.0
	github.com/imkira/go-interpol v1.1.0 // indirect
	github.com/iris-contrib/blackfriday v2.0.0+incompatible // indirect
	github.com/iris-contrib/formBinder v0.0.0-20190104093907-fbd5963f41e1 // indirect
	github.com/iris-contrib/go.uuid v2.0.0+incompatible // indirect
	github.com/iris-contrib/httpexpect v0.0.0-20180314041918-ebe99fcebbce // indirect
	github.com/json-iterator/go v1.1.6 // indirect
	github.com/k0kubun/colorstring v0.0.0-20150214042306-9440f1994b88 // indirect
	github.com/kataras/golog v0.0.0-20180321173939-03be10146386 // indirect
	github.com/kataras/iris v11.1.1+incompatible
	github.com/kataras/pio v0.0.0-20190103105442-ea782b38602d // indirect
	github.com/lukehoban/go-outline v0.0.0-20161011150102-e78556874252 // indirect
	github.com/mattn/go-colorable v0.1.1 // indirect
	github.com/mattn/go-oci8 v0.0.0-20190517005234-99a826511cf1
	github.com/microcosm-cc/bluemonday v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826 // indirect
	github.com/moul/http2curl v1.0.0 // indirect
	github.com/newhook/go-symbols v0.0.0-20151212134530-b75dfefa0d23 // indirect
	github.com/nsf/gocode v0.0.0-20190302080247-5bee97b48836 // indirect
	github.com/onsi/ginkgo v1.7.0
	github.com/onsi/gomega v1.4.3
	github.com/opentracing/opentracing-go v1.1.0
	github.com/pelletier/go-toml v1.2.0
	github.com/pkg/errors v0.8.1
	github.com/rcrowley/go-metrics v0.0.0-20181016184325-3113b8401b8a
	github.com/rogpeppe/godef v1.1.1 // indirect
	github.com/ryanuber/columnize v2.1.0+incompatible // indirect
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	github.com/smartystreets/goconvey v0.0.0-20190330032615-68dc04aab96a // indirect
	github.com/stretchr/testify v1.3.0
	github.com/tgulacsi/go v0.4.7
	github.com/tpng/gopkgs v0.0.0-20180428091733-81e90e22e204 // indirect
	github.com/uber/jaeger-client-go v2.16.0+incompatible
	github.com/uber/jaeger-lib v2.0.0+incompatible // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/xeipuuv/gojsonpointer v0.0.0-20180127040702-4e3ac2762d5f // indirect
	github.com/xeipuuv/gojsonreference v0.0.0-20180127040603-bd5ef7bd5415 // indirect
	github.com/xeipuuv/gojsonschema v1.1.0 // indirect
	github.com/yalp/jsonpath v0.0.0-20180802001716-5cc68e5049a0 // indirect
	github.com/yudai/gojsondiff v1.0.0 // indirect
	github.com/yudai/golcs v0.0.0-20170316035057-ecda9a501e82 // indirect
	github.com/yudai/pp v2.0.1+incompatible // indirect
	github.com/zmb3/gogetdoc v0.0.0-20190228002656-b37376c5da6a // indirect
	go.uber.org/atomic v1.4.0 // indirect
	go.uber.org/dig v1.7.0
	go.uber.org/multierr v1.1.0 // indirect
	go.uber.org/zap v1.10.0
	golang.org/x/crypto v0.0.0-20190513172903-22d7a77e9e5f // indirect
	golang.org/x/net v0.0.0-20190514140710-3ec191127204 // indirect
	golang.org/x/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys v0.0.0-20190516110030-61b9204099cb // indirect
	golang.org/x/text v0.3.2
	golang.org/x/tools v0.0.0-20190517183331-d88f79806bbd // indirect
	google.golang.org/appengine v1.5.0 // indirect
	google.golang.org/grpc v1.19.0
	gopkg.in/errgo.v1 v1.0.1
	gopkg.in/gographics/imagick.v2 v2.5.0
	gopkg.in/inconshreveable/log15.v2 v2.0.0-20180818164646-67afb5ed74ec
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	gopkg.in/rana/ora.v4 v4.1.15
	sourcegraph.com/sqs/goreturns v0.0.0-20181028201513-538ac6014518 // indirect
)
