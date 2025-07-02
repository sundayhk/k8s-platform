package main

import (
	_ "k8s-platform/config"
	"k8s-platform/middlewares"
	"k8s-platform/routers"

	// "k8s-platform/utils/jwtutil"
	_ "k8s-platform/controllers/initcontroller"
	"k8s-platform/utils/logs"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middlewares.JwtAuth)
	logs.Info(nil, "服务启动成功")
	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "Hello, World!",
	// 	})
	// })

	// // 测试生成 jwt token
	// token, _ := jwtutil.GetToken("aaa")
	// fmt.Println("token:", token)

	// // 验证jwt token
	// _, err := jwtutil.ParseToken(token)
	// if err != nil {
	// 	fmt.Println("token 验证失败")
	// } else {
	// 	fmt.Println("token 验证成功")
	// }

	routers.RegisterRouters(r)

	// r.NoRoute(func(c *gin.Context) {
	// 	c.JSON(404, gin.H{"error": "未找到匹配的路由"})
	// })

	// r.Run(config.Port)
	r.Run("localhost:8080") // 避免windows防火墙弹窗
}
