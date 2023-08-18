package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go-web-mini/global"
	"time"
)

var headerXRequestID string

func LoggerMiddleware() gin.HandlerFunc {

	headerXRequestID = "X-Request-ID"
	return func(c *gin.Context) {
		//rid := c.GetHeader(headerXRequestID)
		rid := uuid.New().String()
		c.Request.Header.Add(headerXRequestID, rid)
		c.Set(headerXRequestID, rid)
		c.Writer.Header().Set(headerXRequestID, rid)
		startTime := time.Now()
		_=startTime
		c.Next()

		global.AccessLog.Infow("[TRACE-ID:"+rid+"]",
			"remote_addr", c.ClientIP(),
			"time_local", startTime.Format("02/Jan/2006:15:04:05 -0700"),
			"request", fmt.Sprintf("%s %s %s", c.Request.Method, c.Request.URL.Path, c.Request.Proto),
			"status", c.Writer.Status(),
			"body_bytes", c.Writer.Size(),
			"referer", c.Request.Referer(),
			"user_agent", c.Request.UserAgent(),
			"request_time", time.Since(startTime).Seconds(),
		)
	}
}
