package controller

const (
	commonTitle string = "测试资料库"
)

type APIJson struct {
	Status bool        `json:"status"`
	Msg    interface{} `json:"msg"`
	Data   interface{} `json:"data"`
}

// APIResult 用户API 项目 忽略
func APIResult(status bool, object interface{}, msg string) (apijson *APIJson) {
	// apijson 已经在返回值处 声明了，不用重复声明。
	apijson = &APIJson{Status: status, Data: object, Msg: msg}
	return apijson
}
