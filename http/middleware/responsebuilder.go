package middleware

import (
	"github.com/gin-gonic/gin"
)

//后置
func ResponseBuilder() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//后置中间件
		ctx.Next()

	}
}
