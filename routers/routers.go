package routers

import (
	"k8s-platform/routers/auth"
	"k8s-platform/routers/cluster"

	"github.com/gin-gonic/gin"
)

// 注册路由方法
func RegisterRouters(r *gin.Engine) {
	// /api/auth/login
	// /api/auth/logout
	apiGroup := r.Group("/api") // 定义路由分组
	auth.RegisterSubRouters(apiGroup)
	cluster.RegisterSubRouters(apiGroup)
}
