package controller

import "github.com/gin-gonic/gin"

//Create	登录态系统
//@Summary	登录态校验
//@Description	登录态校验
//@Tags	登录态校验
//@Produce	json
//@Param	xxx query flow_analysis.GetFlowRequest  false "字段注解"
//@Success 200 {object} flow_analysis.FlowCard
//@Router	/flow/flow_profile [post]
func Check(ctx *gin.Context) {
	ctx.Writer.Write([]byte("check.post"))
}

func SignInUsername(ctx *gin.Context) {
	ctx.Writer.Write([]byte("signin/username.post"))
}

func SignInSMS(ctx *gin.Context) {
	ctx.Writer.Write([]byte("signin/SMS.post"))
}

func SignUpUsername(ctx *gin.Context) {

}
func SignUpSMS(ctx *gin.Context) {
}
