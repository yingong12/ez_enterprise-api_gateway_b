package controller

import "github.com/gin-gonic/gin"

//BindPhone 绑定手机号
func BindPhone(ctx *gin.Context) {
	uid, _ := ctx.Params.Get("uid")
	ctx.Writer.Write([]byte(uid))
}

//BindUserName 绑定用户名
func BindUsername(ctx *gin.Context) {

}
func LockUnlock(ctx *gin.Context) {

}

func UpdatePswd(ctx *gin.Context) {

}

func GetAssets(ctx *gin.Context) {

}
