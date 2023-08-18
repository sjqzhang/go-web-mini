package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"go-web-mini/global"
	"strconv"
	"sync"
	"time"
)

type bodyWriter struct {
	gin.ResponseWriter
	bodyCache *bytes.Buffer
}

// rewrite Write()
func (w bodyWriter) Write(b []byte) (int, error) {
	w.bodyCache.Write(b)
	return w.ResponseWriter.Write(b)
}

func CacheMiddleware(paramMap sync.Map) gin.HandlerFunc {

	var ttl int64

	if v, ok := paramMap.Load("ttl"); ok {
		ttl, _ = strconv.ParseInt(v.(string), 10, 64)
	} else {
		ttl = 10
	}

	//cache,err:=lru.New(100)
	//if err!=nil{
	//	global.Log.Error(err)
	//}

	return func(c *gin.Context) {



		if val, err := redis.String(global.Redis().Get(c, c.Request.RequestURI).Result()); err == nil {
			c.Writer.Write([]byte(val))
			c.Abort()
			return
		}
		c.Writer = &bodyWriter{bodyCache: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Next()
		if c.Writer.Status() < 300 {
			global.Redis().Set(c, c.Request.RequestURI, c.Writer.(*bodyWriter).bodyCache.String(), time.Second*time.Duration(ttl))
		}
	}

}
