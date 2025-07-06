package namespace

import (
	"context"
	"k8s-platform/utils/logs"

	"k8s-platform/config"
	"k8s-platform/controllers"

	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Get(r *gin.Context) {
	logs.Debug(nil, "查询Namespace")
	clientset, basicInfo, err := controllers.BasicInit(r, nil)
	returnData := config.ReturnData{}
	if err != nil {
		returnData.Status = 400
		returnData.Message = err.Error()
		r.JSON(200, returnData)
		return
	}
	namespace, err := clientset.CoreV1().Namespaces().Get(context.TODO(), basicInfo.Name, metav1.GetOptions{})
	if err != nil {
		msg := "查询失败：" + err.Error()
		returnData.Status = 400
		returnData.Message = msg
	} else {
		returnData.Status = 200
		returnData.Message = "查询成功"
		returnData.Data = make(map[string]interface{})
		returnData.Data["items"] = namespace
	}
	r.JSON(200, returnData)
}
