package namespace

import (
	"context"
	"k8s-platform/config"
	"k8s-platform/controllers"
	"k8s-platform/utils/logs"

	"github.com/gin-gonic/gin"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Update(r *gin.Context) {
	logs.Debug(nil, "更新Namespace")
	var ns corev1.Namespace
	clientset, _, err := controllers.BasicInit(r, &ns)
	returnData := config.ReturnData{}
	if err != nil {
		returnData.Status = 400
		returnData.Message = err.Error()
		r.JSON(200, returnData)
		return
	}
	_, err = clientset.CoreV1().Namespaces().Update(context.TODO(), &ns, metav1.UpdateOptions{})
	if err != nil {
		msg := "更新失败：" + err.Error()
		returnData.Status = 400
		returnData.Message = msg
	} else {
		returnData.Status = 200
		returnData.Message = "更新成功"
	}
	r.JSON(200, returnData)
}
