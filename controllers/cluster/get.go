package cluster

import (
	"k8s-platform/utils/logs"

	"github.com/gin-gonic/gin"
)

func Get(r *gin.Context) {
	logs.Debug(nil, "查询集群信息")

	r.JSON(200, gin.H{
		"message": "查询集群信息",
	})
}
