package middlewares

import (
	"log"

	"github.com/gin-gonic/gin"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Processing request
		c.Next()

		log.Printf(
			"Request method: %s, URL: %s, Status code: %d",
			c.Request.Method,
			c.Request.RequestURI,
			c.Writer.Status(),
		)
	}
}
