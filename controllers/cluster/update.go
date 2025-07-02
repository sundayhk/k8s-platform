package cluster

import (
	"k8s-platform/utils/logs"

	"github.com/gin-gonic/gin"
)

func Update(r *gin.Context) {
	logs.Debug(nil, "更新集群信息")

	addOrUpdate(r, "Update")
}
