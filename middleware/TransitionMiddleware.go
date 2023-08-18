package middleware

import (
	"github.com/gin-gonic/gin"
	"go-web-mini/global"
	"gorm.io/gorm"
)

// CORS跨域中间件
func TransitionMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {



		db := global.DB.Session(&gorm.Session{SkipDefaultTransaction: false,Context: c})
		db.Transaction(func(tx *gorm.DB) error {
			defer func() {
				if err := recover(); err != nil || c.IsAborted() || c.Writer.Status() >= 400 {
					global.TraceLog(c, "SQL rollback URI:"+c.Request.RequestURI+" SQL:"+tx.Statement.SQL.String())
					tx.Rollback()
				} else {
					tx.Commit()
				}

			}()

			c.Set("db", tx)
			c.Next()
			return nil
		})

	}
}
