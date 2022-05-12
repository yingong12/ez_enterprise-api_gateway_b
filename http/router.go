package http

import (
	"account_service/http/controller"

	"github.com/gin-gonic/gin"
)

func loadRouter() (router *gin.Engine) {
	gin.SetMode(gin.DebugMode)
	router = gin.New()
	//routes
	router.POST("healthy", controller.Healthy)
	//swagger
	// router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // register swagger
	// 企业模块
	enterprise := router.Group("/enterprises")
	{
		enterprise.GET("")         //获取企业信息
		enterprise.PUT("/:en_id")  //更新企业信息
		enterprise.POST("/create") //新建企业 用于O端(zy要求)
		//TODO:单独写状态接口因为查询状态较为频繁，减少网络请求数据量. 初期可以不使用
		enterprise.GET("/state/:en_id") //获取企业状态
		enterprise.PUT("/state/:en_id") //更新企业状态
	}
	//机构模块
	group := router.Group("/groups")
	{
		enterprise.GET("")                                  //获取企业信息
		group.GET("/enterprises/all", controller.GetAssets) //
	}
	//审核模块
	audit := router.Group("audits")
	{
		audit.POST("")        //提交审核 （涉及图片上传）
		audit.GET("")         //搜索审核,分页
		audit.POST("confirm") //审核通过，打回. 同步修改企业状态
	}
	//评估模块
	valuate := router.Group("valuate")
	{
		valuate.POST("")       //提交估值
		valuate.GET("")        //获取估值结果
		valuate.POST("export") //导出 同步异步？
	}
	return
}
