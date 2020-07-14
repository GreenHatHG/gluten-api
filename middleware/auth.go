package middleware

import (
	"github.com/gin-gonic/gin"
	"gluten/util"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		auth := context.Request.Header.Get("Authorization")
		if len(auth) == 0 {
			context.Abort()
			util.StatusUnauthorized(context)
		}
		// 校验token
		id, ok := util.ParseToken(auth)
		if !ok {
			context.Abort()
			util.StatusUnauthorized(context)
		} else {
			context.Set("id", id)
			println("token 正确", id)
		}
		context.Next()
	}
}
