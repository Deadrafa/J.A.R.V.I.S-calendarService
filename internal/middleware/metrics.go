package middleware

import (
	"strconv"
	"time"

	"github.com/Deadrafa/J.A.R.V.I.S-calendarService/pkg/repository"
	"github.com/gin-gonic/gin"
)

func MetricsMiddleware(metrics repository.Collector) gin.HandlerFunc {
	return func(c *gin.Context) {

		start := time.Now()

		c.Next()

		duration := time.Since(start).Seconds()

		status := c.Writer.Status()
		statusStr := strconv.Itoa(status)

		metrics.RecordHTTPRequest(
			c.Request.Method,
			c.Request.URL.Path,
			statusStr,
		)
		metrics.RecordHTTPDuration(
			c.Request.URL.Path,
			duration,
		)
	}
}
