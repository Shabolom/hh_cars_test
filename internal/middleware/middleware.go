package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hh_test_autho/internal/tools"
	"time"
)

var logInf = tools.InfoLogs()

func Timer() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		handler := c.Request.RequestURI
		method := c.Request.Method

		c.Next()

		latency := time.Since(t).Seconds()
		statusCode := c.Writer.Status()
		info := fmt.Sprintf("статус: %d | хэндлер: %s | метод: %s | время выполнения: %f сек", statusCode, handler, method, latency)
		logInf.WithField("middleware", "middleware").Info(info)
	}
}
