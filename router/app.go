package router

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/ylinyang/k8sPanels/docs"
	"github.com/ylinyang/k8sPanels/service"
)

func Router() *gin.Engine {
	r := gin.Default()

	// swag 页面
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// 测试页面
	r.GET("/ping", service.Ping)

	// 集群接口
	r.POST("/cluster-create", service.ClusterCreate)

	// node接口
	r.GET("/cluster-node", service.GetNodes)

	// deployment接口
	
	return r
}
