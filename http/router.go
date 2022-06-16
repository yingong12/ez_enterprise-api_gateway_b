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
	//头部添加中间件
	router.Use(middleware.HeaderInjector())
	//访问日志
	router.Use(middleware.RequestLogger())
	//业务错误日志(controller最终抛出)
	router.Use(middleware.ControllerErrorLogger())
	//routes
	auth := router.Group("auth")
	{
		auth.Any("*url", controller.ForwardAccountService)
		auth.Any("", controller.ForwardAccountService)
	}
	//
	{
		//企业机构服务
		router.Use(middleware.Auth())
		//B端限制中间中件
		router.Use(middleware.GuardDog)
		{
			router.Any("enterprise", controller.ForwardCompanyService)
			router.Any("group", controller.ForwardCompanyService)
			router.Any("audit", controller.ForwardCompanyService)
			router.Any("valuate", controller.ForwardCompanyService)
			router.Any("group/*url", controller.ForwardCompanyService)
			router.Any("enterprise/*url", controller.ForwardCompanyService)
			router.Any("audit/*url", controller.ForwardCompanyService)
			router.Any("valuate/*url", controller.ForwardCompanyService)
		}
		//账号服务
		{
			router.Any("account", controller.ForwardAccountService)
			router.Any("sms", controller.ForwardAccountService)
			router.Any("account/*url", controller.ForwardAccountService)
			router.Any("sms/*url", controller.ForwardAccountService)
		}
	}
	//404
	router.NoRoute(func(ctx *gin.Context) {
		ctx.Writer.WriteString("gateway not found")
	})

	return
}
