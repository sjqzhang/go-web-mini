package vo

import "go-web-mini/apps/workflow/errcode"

type CommResp struct {
	RetCode int         `json:"retcode"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewResp(e errcode.Exception, data interface{}) CommResp {
	if e == nil {
		e = errcode.OK
	}
	return CommResp{
		RetCode: e.ErrCode(),
		Message: e.Message(),
		Data: data,
	}
}
