package auth

import (
	"k8s-platform/config"
	"k8s-platform/utils/jwtutil"
	"k8s-platform/utils/logs"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 登陆控制器
func Login(r *gin.Context) {
	UserInfo := UserInfo{}
	returnData := config.NewReturnData()
	if err := r.ShouldBindBodyWithJSON(&UserInfo); err != nil {
		returnData.Status = 401
		returnData.Message = err.Error()
		r.JSON(200, returnData)
		return
	}
	logs.Info(map[string]interface{}{"用户名": UserInfo.Username, "密码": UserInfo.Password}, "开始验证登陆信息")
	if UserInfo.Username == config.UserName && UserInfo.Password == config.Password {
		ss, err := jwtutil.GetToken(UserInfo.Username)
		if err != nil {
			logs.Error(map[string]interface{}{"用户名": UserInfo.Username, "错误信息": err.Error()}, "用户名密码验证成功，生成token失败")
			returnData.Status = 401
			returnData.Message = "生成token失败"
			r.JSON(200, returnData)
			return
		}

		logs.Info(map[string]interface{}{"用户名": UserInfo.Username}, "登陆成功")
		returnData.Message = "登陆成功"
		returnData.Data["token"] = ss
		r.JSON(200, returnData)
		return
	} else {
		r.JSON(200, gin.H{
			"stats":   401,
			"message": "用户名或密码错误",
		})
		return
	}
}

// 退出登陆
func Logout(r *gin.Context) {
	r.JSON(200, gin.H{
		"stats":   200,
		"message": "退出成功",
	})
	logs.Debug(nil, "退出成功")
}
