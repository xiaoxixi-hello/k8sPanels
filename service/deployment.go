package service

import (
	"github.com/gin-gonic/gin"
)

// GetDeploymentList
// @Tags 私有方法
// @Summary deployment列表
// @Param name query string false "name"
// @Param namespace query string false "namespace"
// @Success 200 {string} json "{"code":200,"data":""}"
// @Router /deployment-list [get]
func GetDeploymentList(c *gin.Context) {
	name := c.Query("name")
	namespace := c.Query("namespace")

}
