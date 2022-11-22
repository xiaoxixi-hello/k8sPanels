package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Ping
// @Tags 测试方法
// @Summary ping
// @Success 200 {string} json "{"code":200,"data":""}"
// @Router /ping [get]
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "pong",
	})
}
