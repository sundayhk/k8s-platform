package cluster

import "github.com/gin-gonic/gin"

func addOrUpdate(r *gin.Context, method string) {
	var requestMethod string
	if method == "Create" {
		requestMethod = "创建"
	} else {
		requestMethod = "更新"
	}

	// clusterConfig := clusterConfig

}
