package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

//记录controller抛出的错误
func ControllerErrorLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		err, ok := ctx.Get("buz_err")
		if err == nil || !ok {
			return
		}
		log.Printf("[endpoint]:%s [error]:%v\n", ctx.Request.URL.Path, err)
	}
}
