package namespace

import (
	"k8s-platform/controllers/namespace"

	"github.com/gin-gonic/gin"
)

func create(r *gin.RouterGroup) {
	r.POST("/create", namespace.Create)
}

func update(r *gin.RouterGroup) {
	r.POST("/update", namespace.Update)
}

func delete(r *gin.RouterGroup) {
	r.GET("/delete", namespace.Delete)
}

func get(r *gin.RouterGroup) {
	r.GET("/get", namespace.Get)
}

func list(r *gin.RouterGroup) {
	r.GET("/list", namespace.List)
}

func RegisterSubRouters(r *gin.RouterGroup) {
	namespaceGroup := r.Group("/namespace")
	create(namespaceGroup)
	update(namespaceGroup)
	delete(namespaceGroup)
	get(namespaceGroup)
	list(namespaceGroup)
}
