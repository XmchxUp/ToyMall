package middleware

import (
	"time"
	"xm-mall/pkg/e"
	"xm-mall/pkg/utils"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var code int
		var data interface{}
		code = e.SUCCESS
		token := ctx.GetHeader("Authorization")
		if token == "" {
			code = 404
		} else {
			claims, err := utils.ParseToken(token)
			if err != nil {
				code = e.ErrorAuthCheckTokenFail
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ErrorAuthCheckTokenTimeout
			}
		}
		if code != e.SUCCESS {
			ctx.JSON(e.SUCCESS, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
				"data":   data,
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
