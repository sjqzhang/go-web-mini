package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go-web-mini/common"
	"go-web-mini/response"
	"reflect"
)

// CORS跨域中间件
func BinderMiddleware(method reflect.Value) gin.HandlerFunc {

	fmt.Println("%v", method)
	return func(c *gin.Context) {

		if method.Type().NumIn() != 2 {
			c.JSON(400, gin.H{"message": "请求参数错误"})
			return
		}
		reqType := method.Type().In(1)
		var err error

		req := reflect.New(reqType.Elem()).Interface()
		// 获取请求参数 Get
		if c.Request.Method == "GET" {
			err := c.ShouldBindQuery(req)
			if err != nil {
				fmt.Println(req)
				c.JSON(400, gin.H{"message": "请求参数错误"})
				return
			}
		}
		// 绑定参数

		if (c.Request.Method == "POST" || c.Request.Method == "PUT" || c.Request.Method == "DELETE" ||
			c.Request.Method == "PATCH") && c.Request.ContentLength > 0 {
			err := c.ShouldBindJSON(req)
			if err != nil {
				fmt.Println(req)
				c.JSON(400, gin.H{"message": "请求参数错误"+err.Error()})
				return
			}
		}
		//if strings.Index(c.FullPath(), ":") > 0 {
		//	err = c.ShouldBindUri(req)
		//	if err != nil {
		//		fmt.Println(req)
		//		c.JSON(400, gin.H{"message": "请求参数错误"})
		//		return
		//	}
		//}
		err = common.Validate.Struct(req)
		if err != nil {
			errStr := err.(validator.ValidationErrors)[0].Translate(common.Trans)
			response.Fail(c, nil, errStr)
			return
		}

		results := method.Call([]reflect.Value{reflect.ValueOf(c), reflect.ValueOf(req)})
		if len(results) > 0 {
			if len(results) == 2 {
				if results[1].Interface() != nil {
					response.Fail(c, nil, results[1].Interface().(string))
					return
				}
			}
			if results[0].Interface() != nil {
				response.Response(c, 200, 0, results[0].Interface(), "Success")
				return
			}
		}
		c.Next()

	}
}
