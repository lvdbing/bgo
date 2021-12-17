package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/lvdbing/bgo/internal/helper/jwthelper"
	"github.com/lvdbing/bgo/internal/model"
	"github.com/lvdbing/bgo/pkg/errcode"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string
		var respCode = errcode.Success

		if s, exist := c.GetQuery("token"); exist {
			token = s
		} else {
			token = c.GetHeader("token")
		}
		if token == "" {
			respCode = errcode.InvalidParams
		} else {
			_, err := jwthelper.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					respCode = errcode.UnauthorizedTokenTimeout
				default:
					respCode = errcode.UnauthorizedTokenError
				}
			}
		}

		if respCode != errcode.Success {
			resp := model.NewResponse(c)
			resp.SendError(respCode)
			c.Abort()
			return
		}

		c.Next()
	}
}
