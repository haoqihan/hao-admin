package middleware

import (
	"hao-admin/global"
	"hao-admin/pkg/logger"
	"bytes"
	"github.com/gin-gonic/gin"
	"time"
)

type AccessLogWrite struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w AccessLogWrite) Write(p []byte) (int, error) {
	if n, err := w.body.Write(p); err != nil {
		return n, err
	}
	return w.ResponseWriter.Write(p)
}

func AccessLog() gin.HandlerFunc {
	return func(context *gin.Context) {
		bodyWrite := &AccessLogWrite{body: bytes.NewBufferString(""), ResponseWriter: context.Writer}
		context.Writer = bodyWrite
		beginTime := time.Now().Unix()
		context.Next()
		endTime := time.Now().Unix()
		fields := logger.Fields{
			"request":  context.Request.PostForm.Encode(),
			"response": bodyWrite.body.String(),
		}
		s := "access log: method: %s ,statue_code: %d, begin_time: %d,end_time: %d"
		global.Logger.WithFields(fields).Infof(
			context,
			s,
			context.Request.Method,
			bodyWrite.Status(),
			beginTime,
			endTime)

	}
}
