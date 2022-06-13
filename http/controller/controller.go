package controller

import (
	"api_gateway_b/http/buz_code"
	"net/http"

	"github.com/gin-gonic/gin"
)

type STDResponse struct {
	Code buz_code.Code `json:"code"`
	Msg  string        `json:"msg"`
	Data interface{}   `json:"data"`
}
type RawResponse []byte

//标准返回并且输出错误日志
func STDwrapperJSON(handler func(*gin.Context) (STDResponse, error)) func(*gin.Context) {
	return func(ctx *gin.Context) {
		res, err := handler(ctx)
		ctx.Set("buz_err", err)
		//http status永远200。 用buz_code去表示错误
		status := http.StatusOK
		//write json response
		ctx.JSON(status, res)
	}
}

func STDWrapperRaw(handler func(*gin.Context) (RawResponse, error)) func(*gin.Context) {
	return func(ctx *gin.Context) {
		res, err := handler(ctx)
		ctx.Set("buz_err", err)
		//永远200
		ctx.Status(http.StatusOK)
		//标记转发成功。用来区分远程错误还是网关错误
		ctx.Header("forward_succeed", "1")
		//write json response
		ctx.Writer.Write(res)
	}
}
