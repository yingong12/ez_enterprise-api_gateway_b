package http

import (
	"api_gateway_b/http/controller"
	"api_gateway_b/http/middleware"

	"github.com/gin-gonic/gin"
)

//鉴权，限流，转发请求到对应服务
func loadRouter() (router *gin.Engine) {
	gin.SetMode(gin.DebugMode)
	router = gin.New()
	//health check
	router.GET("health", controller.Healthy)
	//middleware
	{
		//鉴权中间件
		router.Use(middleware.Auth())
		//头部添加中间件
		router.Use(middleware.HeaderInjector())
		//访问日志
		router.Use(middleware.RequestLogger())
		//业务错误日志(controller最终抛出)
		router.Use(middleware.ControllerErrorLogger())
	}
	//routes
	//重写
	{
		router.POST("enterprise/update", controller.STDWrapperRaw(controller.UpdateEnterprise))

	}
	//直接转发
	{
		//企业机构服务
		router.Any("enterprises/*url", controller.ForwardCompanyService)
		router.Any("groups/*url", controller.ForwardCompanyService)
		router.Any("audits/*url", controller.ForwardCompanyService)
		router.Any("valuate/*url", controller.ForwardCompanyService)
		//账号服务
		router.Any("auth/*url", controller.ForwardAccountService)
		router.Any("account/*url", controller.ForwardAccountService)
		router.Any("sms/*url", controller.ForwardAccountService)
	}

	return
}
