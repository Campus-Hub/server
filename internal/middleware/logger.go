package middleware

import (
	"fmt"
	"github.com/Campus-Hub/server/pkg/errno"
	"github.com/Campus-Hub/server/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"math"
	"time"
)

// LoggerMiddleware Log with important fields.
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		startTime := time.Now()
		// Call the remaining handlers for this request
		c.Next()
		stopTime := time.Since(startTime)
		spendTime := fmt.Sprintf("%d ms", int(math.Ceil(float64(stopTime.Nanoseconds()/1000000))))

		statusCode := c.Writer.Status()
		dataSize := c.Writer.Size()
		if dataSize < 0 {
			dataSize = 0
		}
		method := c.Request.Method
		uri := c.Request.RequestURI

		Log := logger.Logger.WithFields(logrus.Fields{
			"spend_time": spendTime,
			"path":       uri,
			"method":     method,
			"status":     statusCode,
		})

		// create inner error
		if len(c.Errors) > 0 {
			Log.Error(c.Errors.ByType(gin.ErrorTypePrivate))
		}
		if statusCode >= errno.Error {
			Log.Error()
		} else if statusCode >= errno.InvalidParams {
			Log.Warn()
		} else {
			Log.Info()
		}

	}
}
