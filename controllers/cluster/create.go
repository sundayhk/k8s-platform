package cluster

import (
	"k8s-platform/utils/logs"

	"github.com/gin-gonic/gin"
)

func Create(r *gin.Context) {
	logs.Debug(nil, "添加集群信息")
	createOrUpdate(r, "Create")
}
