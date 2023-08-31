package errcode

import (
	"net/http"
)

var (
	OK                 = New(http.StatusOK, "ok")
	RequestIllegal     = New(http.StatusBadRequest, "request is illegal")
	CommandIsAvailable = New(http.StatusConflict, "command is available")
	LockFail           = New(http.StatusConflict, "lock fail")
	DBError            = New(http.StatusInternalServerError, "db error")
	ServiceError       = New(http.StatusInternalServerError, "service error")
	EngineError        = New(http.StatusInternalServerError, "engine error")
)

type exception struct {
	errCode int
	message string
}

type Exception interface {
	ErrCode() int
	Message() string
}

func (e exception) ErrCode() int {
	return e.errCode
}

func (e exception) Message() string {
	return e.message
}

func New(errCode int, message string) Exception {
	return exception{
		errCode: errCode,
		message: message,
	}
}

func Wrap(exception Exception, err error) Exception {
	return New(exception.ErrCode(), exception.Message()+": "+err.Error())
}
