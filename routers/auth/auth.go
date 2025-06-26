package auth

import (
	"k8s-platform/controllers/auth"

	"github.com/gin-gonic/gin"
)

// 登陆
func login(authGroup *gin.RouterGroup) {
	authGroup.POST("/login", auth.Login)
}

func logout(authGroup *gin.RouterGroup) {
	authGroup.GET("/logout", auth.Logout)
}

func RegisterSubRouters(g *gin.RouterGroup) {
	authGroup := g.Group("/auth")
	login(authGroup)
	logout(authGroup)
}
