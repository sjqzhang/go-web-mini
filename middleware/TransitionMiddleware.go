package middleware

import (
	"github.com/gin-gonic/gin"
	"go-web-mini/global"
	"gorm.io/gorm"
	"log"
)

// CORS跨域中间件
func TransitionMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		db := global.DB.Session(&gorm.Session{SkipDefaultTransaction: false})
		db.Transaction(func(tx *gorm.DB) error {
			defer func() {
				if err := recover(); err != nil {
					log.Printf("Panic info is: %v", err)
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
