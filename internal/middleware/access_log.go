package middleware

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lvdbing/bgo/global"
	"github.com/lvdbing/bgo/pkg/logger"
)

type accessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w accessLogWriter) Write(b []byte) (int, error) {
	if n, err := w.body.Write(b); err != nil {
		return n, err
	}
	return w.ResponseWriter.Write(b)
}

func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		// OPTIONS请求不打印访问日志。
		if c.Request.Method == http.MethodOptions {
			c.Next()
			return
		}

		writer := &accessLogWriter{
			ResponseWriter: c.Writer,
			body:           bytes.NewBufferString(""),
		}
		c.Writer = writer

		var body []byte
		var req string
		if c.Request.Method != http.MethodGet && c.Request.Method != http.MethodDelete {
			var err error
			body, err = ioutil.ReadAll(c.Request.Body)
			if err != nil {
				global.Logger.Errorf(c, "read body from request err: %v", err)
			} else {
				c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
			}
			req = string(body)
		} else {
			req = c.Request.URL.Path
		}

		begin := time.Now().Unix()
		c.Next()
		end := time.Now().Unix()

		fields := logger.Fields{
			"request":  req,
			"response": writer.body.String(),
		}
		format := "access log: method: %s, status_code: %d, begin: %d, end: %d"
		global.Logger.WithFields(fields).Infof(c, format,
			c.Request.Method,
			writer.Status(),
			begin,
			end,
		)
	}
}
