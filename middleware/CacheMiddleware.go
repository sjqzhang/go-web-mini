package middleware

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-web-mini/global"
	"time"
)

const CACHE_KEY = "cache-key"

// CacheMiddleware 是缓存中间件
func CacheMiddleware(expiration time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method != "GET" {
			c.Next()
		} else {
			cacheKey := generateCacheKey(c)

			c.Set(CACHE_KEY, cacheKey)

			c.Next()
		}

		// 如果需要缓存响应数据，可以在处理响应之后设置缓存
		// response := c.Writer.Body.String()
		// _ = setCache(cacheKey, response, expiration)
	}
}

func generateCacheKey(c *gin.Context) string {
	return fmt.Sprintf("cache:%s:%s:%v", c.FullPath(), c.Request.Method, c.Request.URL.Query())
}

func setCache(key string, value interface{}, expiration time.Duration) error {
	ctx := context.Background()
	return global.Redis().Set(ctx, key, value, expiration).Err()
}

func getCache(key string) (string, error) {
	ctx := context.Background()
	return global.Redis().Get(ctx, key).Result()
}
