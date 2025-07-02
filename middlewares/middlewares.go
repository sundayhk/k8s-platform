// 中间件
package middlewares

import (
	"k8s-platform/config"
	"k8s-platform/utils/jwtutil"
	"k8s-platform/utils/logs"

	"github.com/gin-gonic/gin"
)

func JwtAuth(r *gin.Context) {
	requestUrl := r.Request.URL.Path
	logs.Debug(map[string]interface{}{"请求路径": requestUrl}, "")
	// 取消验证
	if requestUrl == "/api/auth/login" || requestUrl == "/api/auth/logout" {
		logs.Debug(map[string]interface{}{"requestUrl": requestUrl}, "登陆和退出接口不验证token")
		r.Next()
		return
	}
	returnData := config.NewReturnData()

	// 必须验证
	tokenString := r.Request.Header.Get("Authorization")
	// 判断为空
	if tokenString == "" {
		returnData.Status = 401
		returnData.Message = "请求未携带token,请登录后尝试"
		r.JSON(200, returnData)
		r.Abort() // 终止后续中间件和处理函数的执行
		return
	}
	// 验证合法
	claims, err := jwtutil.ParseToken(tokenString)
	if err != nil {
		returnData.Status = 401
		returnData.Message = "Token验证未通过"
		r.JSON(200, returnData)
		r.Abort()
		return
	}
	// 验证成功
	r.Set("claims", claims)
	r.Next()
}
