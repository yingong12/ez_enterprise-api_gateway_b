package response

import (
	"api_gateway_b/http/buz_code"

	"github.com/gin-gonic/gin"
)

func SetResponse(ctx *gin.Context, httpCode int, buzCode buz_code.Code, msg string, data interface{}) {
	ctx.JSON(httpCode, gin.H{
		"code": buzCode,
		"msg":  msg,
		"data": data,
	})
}
