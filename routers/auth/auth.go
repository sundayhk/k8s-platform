package auth

import (
	"k8s-platform/controllers/auth"

	"github.com/gin-gonic/gin"
)

// 登陆
func login(authGroup *gin.RouterGroup) {
	authGroup.POST("/login", auth.Login)
}

// 退出
func logout(authGroup *gin.RouterGroup) {
	authGroup.GET("/logout", auth.Logout)
}

// 注册路由
func RegisterSubRouters(g *gin.RouterGroup) {
	authGroup := g.Group("/auth")
	login(authGroup)
	logout(authGroup)
}
