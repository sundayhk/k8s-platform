package routers

import (
	"k8s-platform/routers/auth"

	"github.com/gin-gonic/gin"
)

// 注册路由方法
func RegisterRouters(r *gin.Engine) {
	// /api/auth/login
	// /api/auth/logout
	apiGroup := r.Group("/api") // api 路由组
	auth.RegisterSubRouters(apiGroup)
}
