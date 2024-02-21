package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"

	"go-gin-server/internal/logger"
)

// RequestLogger logs incoming request.
func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Next()
		latency := time.Since(t)
		line := fmt.Sprintf("%s %s %s %s\n",
			c.Request.Method,
			c.Request.RequestURI,
			c.Request.Proto,
			latency,
		)
		logger.Log.Info("Request: ", line)
	}
}

// ResponseLogger is a response logger.
func ResponseLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("X-Content-Type-Options", "nosniff")
		c.Next()
		line := fmt.Sprintf("%d %s %s\n",
			c.Writer.Status(),
			c.Request.Method,
			c.Request.RequestURI,
		)
		logger.Log.Info("Response: ", line)
	}
}
