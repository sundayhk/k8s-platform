package cluster

import (
	"k8s-platform/utils/logs"

	"github.com/gin-gonic/gin"
)

func Add(r *gin.Context) {
	logs.Debug(nil, "添加集群信息")
	addOrUpdate(r, "Create")
}
