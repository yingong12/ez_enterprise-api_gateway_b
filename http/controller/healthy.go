package controller

import "github.com/gin-gonic/gin"

func Healthy(ctx *gin.Context) {
	ctx.Data(200, "text", []byte("ok"))
}
