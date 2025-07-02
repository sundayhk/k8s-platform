package cluster

import (
	"k8s-platform/utils/logs"

	"github.com/gin-gonic/gin"
)

func List(r *gin.Context) {
	logs.Debug(nil, "获取集群详情")

	r.JSON(200, gin.H{
		"message": "获取集群详情",
	})
}
