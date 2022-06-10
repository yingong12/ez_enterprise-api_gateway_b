package controller

import "github.com/gin-gonic/gin"

//BindPhone 绑定手机号
func ForwardAccountService(ctx *gin.Context) {
	// host := getAccountServiceHost()
	// fullURL := host + ctx.Request.URL.String()
	// redirect(ctx, fullURL)
}
func getAccountServiceHost() string {
	return "localhost:8686"
}
