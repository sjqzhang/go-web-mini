package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)



// 返回前端
func Response(c *gin.Context, httpStatus int, code int, data interface{}, message string) {

	switch data.(type) {

	case gin.H:
		c.JSON(httpStatus, gin.H{"code": code, "data": data, "message": message})
	default:

		c.JSON(httpStatus, gin.H{"code": code, "data": data, "message": message})

	}

}

// 返回前端-成功
func Success(c *gin.Context, data interface{}, message string) {
	Response(c, http.StatusOK, 200, data, message)
}

// 返回前端-失败
func Fail(c *gin.Context, data interface{}, message string) {
	Response(c, http.StatusBadRequest, 400, data, message)
}
