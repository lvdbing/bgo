package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lvdbing/bgo/global"
	"github.com/lvdbing/bgo/internal/helper/email"
	"github.com/lvdbing/bgo/internal/model"
	"github.com/lvdbing/bgo/pkg/errcode"
)

func Recovery() gin.HandlerFunc {
	mailer := email.NewEmail(&email.SMTPInfo{
		From:     global.EmailSetting.From,
		Host:     global.EmailSetting.Host,
		Port:     global.EmailSetting.Port,
		Username: global.EmailSetting.Username,
		Password: global.EmailSetting.Password,
		IsSSL:    global.EmailSetting.IsSSL,
	})
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				format := "panic recover err: %v"
				global.Logger.WithCallersFrames().Panicf(c, format, err)

				err := mailer.SendMail(
					global.EmailSetting.To,
					fmt.Sprintf("Bgo发生异常 %s", time.Now().Format("2006-01-02 15:04:05")),
					fmt.Sprintf("panic err: %v", err),
				)
				if err != nil {
					global.Logger.Panicf(c, "email.SendMail err: %v", err)
				}

				model.NewResponse(c).SendError(errcode.ServerError)
				c.Abort()
			}
		}()
	}
}
