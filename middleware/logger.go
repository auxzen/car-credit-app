package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		duration := time.Since(start)
		status := c.Writer.Status()

		log.Printf("%s - %s %s %d [%s]",
			c.ClientIP(),
			c.Request.Method,
			c.Request.URL.Path,
			status,
			duration,
		)
	}
}
