package cluster

import (
	"k8s-platform/controllers/cluster"

	"github.com/gin-gonic/gin"
)

func add(r *gin.RouterGroup) {
	r.POST("/add", cluster.Add)
}

func update(r *gin.RouterGroup) {
	r.POST("/update", cluster.Update)
}

func delete(r *gin.RouterGroup) {
	r.GET("/delete", cluster.Delete)
}

func get(r *gin.RouterGroup) {
	r.GET("/get", cluster.Get)
}

func list(r *gin.RouterGroup) {
	r.GET("/list", cluster.List)
}

func RegisterSubRouters(r *gin.RouterGroup) {
	clusterGroup := r.Group("/cluster")
	add(clusterGroup)
	update(clusterGroup)
	delete(clusterGroup)
	get(clusterGroup)
	list(clusterGroup)
}
