package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gookit/slog"
	"github.com/ylinyang/k8sPanels/client"
	"github.com/ylinyang/k8sPanels/models"
	"net/http"
)

// ClusterCreate
// @Tags 私有方法
// @Summary 集群创建
// @param name formData string true "name"
// @Param master_url formData string true "master_url"
// @Param config formData string true "config"
// @Success 200 {string} json "{"code":200,"data":""}"
// @Router /cluster-create [post]
func ClusterCreate(c *gin.Context) {
	name := c.PostForm("name")
	masterUrl := c.PostForm("master_url")
	config := c.PostForm("config")
	fmt.Println(config)
	if name == "" || config == "" || masterUrl == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "集群名字或kubeConfig不能为空",
		})
		return
	}
	data := &models.Cluster{
		Name:      name,
		Config:    config,
		MasterUrl: masterUrl,
	}
	if models.GetClusterInfo(name).Name == name {
		slog.Info("集群已经存在")
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "集群已经存在",
		})
		return
	}
	err := models.DB.Create(data).Error
	if err != nil {
		slog.Errorf("cluster信息写入数据库失败,", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "写入集群信息失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "SUCCESS",
	})
}

// GetNodes
// @Tags 私有方法
// @Summary 获取node节点
// @param name query string true "name"
// @Success 200 {string} json "{"code":200,"data":""}"
// @Router /cluster-node [get]
func GetNodes(c *gin.Context) {
	name := c.Query("name")
	node, err := client.GetNode(name)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "获取node失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": node,
	})
}
