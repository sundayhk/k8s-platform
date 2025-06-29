package cluster

import (
	"k8s-platform/utils/logs"

	"github.com/gin-gonic/gin"
)

func Delete(r *gin.Context) {
	logs.Debug(nil, "删除集群信息")

	r.JSON(200, gin.H{
		"message": "删除集群信息",
	})
}
