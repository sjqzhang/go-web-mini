package vo

type Response struct {
	Code    int         `json:"code"`    // 0:成功，非0:失败
	Data    interface{} `json:"data"`    // 返回数据
	Message string      `json:"message"` // 提示信息
}
