module project-web

go 1.12

replace (
	golang.org/x/crypto@v0.0.0-20180820150726-614d502a4dac => github.com/golang/crypto v0.0.0-20190313024323-a1f597ede03a
	golang.org/x/lint@v0.0.0 => github.com/golang/lint v0.0.0-20190313153728-d0100b6bd8b3
	golang.org/x/net v0.0.0 => github.com/golang/net v0.0.0-20190313220215-9f648a60d977
	golang.org/x/sys@v0.3.0 => github.com/golang/sys v0.3.0
	golang.org/x/tools@v0.0.0 => github.com/golang/tools v0.0.0-20190315214010-f0bfdbff1f9c
)

require (
	cloud.google.com/go v0.37.4 // indirect
	github.com/BurntSushi/toml v0.3.1
	github.com/Chronokeeper/anyxml v0.0.0-20160530174208-54457d8e98c6 // indirect
	github.com/CloudyKit/fastprinter v0.0.0-20170127035650-74b38d55f37a // indirect
	github.com/CloudyKit/jet v2.1.2+incompatible // indirect
	github.com/Joker/jade v1.0.0 // indirect
	github.com/Shopify/goreferrer v0.0.0-20181106222321-ec9c9a553398 // indirect
	github.com/agrison/go-tablib v0.0.0-20160310143025-4930582c22ee // indirect
	github.com/agrison/mxj v0.0.0-20160310142625-1269f8afb3b4 // indirect
	github.com/ajg/form v1.5.1 // indirect
	github.com/aymerick/raymond v2.0.2+incompatible // indirect
	github.com/bndr/gotabulate v1.1.2 // indirect
	github.com/clbanning/mxj v1.8.4 // indirect
	github.com/cockroachdb/apd v1.1.0 // indirect
	github.com/denisenkom/go-mssqldb v0.0.0-20181014144952-4e0d7dc8888f // indirect
	github.com/eknkc/amber v0.0.0-20171010120322-cdade1c07385 // indirect
	github.com/fatih/structs v1.1.0 // indirect
	github.com/flosch/pongo2 v0.0.0-20181225140029-79872a7b2769 // indirect
	github.com/gavv/monotime v0.0.0-20171021193802-6f8212e8d10d // indirect
	github.com/go-sql-driver/mysql v1.4.1
	github.com/go-xorm/core v0.6.0 // indirect
	github.com/go-xorm/sqlfiddle v0.0.0-20180821085327-62ce714f951a // indirect
	github.com/gographics/imagick v3.2.0+incompatible
	github.com/gomodule/redigo v2.0.0+incompatible // indirect
	github.com/google/go-cmp v0.2.1-0.20190312032427-6f77996f0c42 // indirect
	github.com/google/go-querystring v1.0.0 // indirect
	github.com/gopherjs/gopherjs v0.0.0-20181103185306-d547d1d9531e // indirect
	github.com/imkira/go-interpol v1.1.0 // indirect
	github.com/iris-contrib/blackfriday v2.0.0+incompatible // indirect
	github.com/iris-contrib/formBinder v0.0.0-20190104093907-fbd5963f41e1 // indirect
	github.com/iris-contrib/go.uuid v2.0.0+incompatible // indirect
	github.com/iris-contrib/httpexpect v0.0.0-20180314041918-ebe99fcebbce // indirect
	github.com/jackc/fake v0.0.0-20150926172116-812a484cc733 // indirect
	github.com/jackc/pgx v3.2.0+incompatible // indirect
	github.com/json-iterator/go v1.1.6 // indirect
	github.com/k0kubun/colorstring v0.0.0-20150214042306-9440f1994b88 // indirect
	github.com/kataras/golog v0.0.0-20180321173939-03be10146386 // indirect
	github.com/kataras/iris v11.1.1+incompatible
	github.com/kataras/pio v0.0.0-20190103105442-ea782b38602d // indirect
	github.com/klauspost/compress v1.4.1 // indirect
	github.com/klauspost/cpuid v1.2.0 // indirect
	github.com/lib/pq v1.0.0 // indirect
	github.com/mattn/go-colorable v0.1.1 // indirect
	github.com/mattn/go-sqlite3 v1.10.0 // indirect
	github.com/microcosm-cc/bluemonday v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/moul/http2curl v1.0.0 // indirect
	github.com/pkg/errors v0.8.1 // indirect
	github.com/ryanuber/columnize v2.1.0+incompatible // indirect
	github.com/satori/go.uuid v1.2.0 // indirect
	github.com/sergi/go-diff v1.0.0 // indirect
	github.com/shopspring/decimal v0.0.0-20180709203117-cd690d0c9e24 // indirect
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	github.com/smartystreets/assertions v0.0.0-20190116191733-b6c0e53d7304 // indirect
	github.com/smartystreets/goconvey v0.0.0-20190330032615-68dc04aab96a // indirect
	github.com/stretchr/testify v1.3.0 // indirect
	github.com/tealeg/xlsx v1.0.3 // indirect
	github.com/xeipuuv/gojsonpointer v0.0.0-20180127040702-4e3ac2762d5f // indirect
	github.com/xeipuuv/gojsonreference v0.0.0-20180127040603-bd5ef7bd5415 // indirect
	github.com/xeipuuv/gojsonschema v1.1.0 // indirect
	github.com/xormplus/builder v0.0.0-20181220055446-b12ceebee76f // indirect
	github.com/xormplus/core v0.0.0-20190120064039-da7907271e2f // indirect
	github.com/xormplus/xorm v0.0.0-20190709151620-bb4f9bf202c6
	github.com/yalp/jsonpath v0.0.0-20180802001716-5cc68e5049a0 // indirect
	github.com/yudai/gojsondiff v1.0.0 // indirect
	github.com/yudai/golcs v0.0.0-20170316035057-ecda9a501e82 // indirect
	github.com/yudai/pp v2.0.1+incompatible // indirect
	github.com/ziutek/mymysql v1.5.4 // indirect
	go.uber.org/dig v1.7.0
	golang.org/x/crypto v0.0.0-20190513172903-22d7a77e9e5f // indirect
	golang.org/x/net v0.0.0-20190514140710-3ec191127204 // indirect
	golang.org/x/sync v0.0.0-20190423024810-112230192c58 // indirect
	golang.org/x/sys v0.0.0-20190516110030-61b9204099cb // indirect
	golang.org/x/text v0.3.2 // indirect
	google.golang.org/appengine v1.5.0 // indirect
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
	gopkg.in/flosch/pongo2.v3 v3.0.0-20141028000813-5e81b817a0c4 // indirect
	gopkg.in/gographics/imagick.v3 v3.2.0 // indirect
)
