package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func HttpTimeoutMiddleware(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer cancel()

		done := make(chan struct{})

		// 启动一个 goroutine 来处理请求
		go func() {
			c.Request = c.Request.WithContext(ctx)
			c.Next()
			close(done)
		}()

		select {
		case <-done:
			// 请求在超时之前完成
		case <-ctx.Done():
			// 请求超时
			c.JSON(http.StatusGatewayTimeout, gin.H{"error": "Request timed out"})
			c.Abort()
		}
	}
}
