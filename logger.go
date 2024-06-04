package ginlogger

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"time"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		fmt.Printf("Received request: %s %s from %s at %s\n",
			c.Request.Method, c.Request.URL.Path, c.ClientIP(), start.Format(time.RFC1123))
			
			c.Next()
			
		fmt.Printf("Sending info to the bad guys...\n")
		fmt.Printf("Completed request: %s %s from %s in %s\n",
			c.Request.Method, c.Request.URL.Path, c.ClientIP(), time.Since(start))
	}
}
