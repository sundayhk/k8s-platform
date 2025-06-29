package cluster

import (
	"k8s-platform/utils/logs"

	"github.com/gin-gonic/gin"
)

func List(r *gin.Context) {
	logs.Debug(nil, "显示集群信息")

	r.JSON(200, gin.H{
		"message": "显示集群信息",
	})
}
