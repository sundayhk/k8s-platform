package cluster

import (
	"k8s-platform/utils/logs"

	"github.com/gin-gonic/gin"
)

func Add(r *gin.Context) {
	logs.Debug(nil, "添加集群信息")

	r.JSON(200, gin.H{
		"message": "添加集群信息",
	})
}
