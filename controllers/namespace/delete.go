package namespace

import (
	"context"
	"k8s-platform/config"
	"k8s-platform/controllers"
	"k8s-platform/utils/logs"

	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Delete(r *gin.Context) {
	logs.Debug(nil, "删除Namespace")
	clientset, basicInfo, err := controllers.BasicInit(r, nil)
	returnData := config.ReturnData{}
	if err != nil {
		returnData.Status = 400
		returnData.Message = err.Error()
		r.JSON(200, returnData)
		return
	}
	if basicInfo.Name == "kube-system" {
		returnData.Status = 400
		returnData.Message = "kube-system命名空间不能删除"
		r.JSON(200, returnData)
		return
	}
	err = clientset.CoreV1().Namespaces().Delete(context.TODO(), basicInfo.Name, metav1.DeleteOptions{})
	if err != nil {
		returnData.Status = 400
		returnData.Message = "删除失败：" + err.Error()
	} else {
		returnData.Status = 200
		returnData.Message = "删除成功"
	}
	r.JSON(200, returnData)
}
