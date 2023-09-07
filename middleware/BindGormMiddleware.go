package middleware

import (
	"github.com/gin-gonic/gin"
	"go-web-mini/global"
)

func BindGormMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		db:=global.GetDB(c)
		c.Set("db", db)
		c.Next()
	}

}

