package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
)

type TimingMiddleware struct{}

func NewTimingMiddleware() *TimingMiddleware {
	return &TimingMiddleware{}
}

func (m *TimingMiddleware) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Pre-processing: Start timer
		start := time.Now()
		c.Set("startedAt", start)

		// Process request
		c.Next()

		// Post-processing: Format response
		// Only wrap if a result was set by the handler
		result, exists := c.Get("result")
		if exists {
			duration := time.Since(start)
			c.JSON(c.Writer.Status(), gin.H{
				"result":          result,
				"executionTimeMs": duration.Milliseconds(),
			})
		}
	}
}
