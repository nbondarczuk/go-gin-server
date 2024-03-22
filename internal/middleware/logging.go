package middleware

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"

	"go-gin-server/internal/log"
)

// RequestLogger logs incoming request.
func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Next()
		latency := time.Since(t)
		line := fmt.Sprintf("%s %s %s %s",
			c.Request.Method,
			c.Request.RequestURI,
			c.Request.Proto,
			latency,
		)
		log.Logger.Info("Received", slog.String("request", line))
	}
}

// ResponseLogger is a response logger.
func ResponseLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("X-Content-Type-Options", "nosniff")
		c.Next()
		line := fmt.Sprintf("%d %s %s",
			c.Writer.Status(),
			c.Request.Method,
			c.Request.RequestURI,
		)
		log.Logger.Info("Produced", slog.String("response", line))
	}
}
