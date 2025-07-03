package cluster

import (
	"context"
	"k8s-platform/config"
	"k8s-platform/utils/logs"

	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func List(r *gin.Context) {
	logs.Debug(nil, "集群详情")
	listOptions := metav1.ListOptions{
		LabelSelector: config.ClusterConfigSecretLabelKey + "=" + config.ClusterConfigSecretLabelValue,
	}
	secretList, err := config.InClusterClientSet.CoreV1().Secrets(config.MetadataNamespace).List(context.TODO(), listOptions)
	returnData := config.ReturnData{}
	if err != nil {
		logs.Error(map[string]interface{}{"msg": err.Error()}, "查询集群列表失败")
		returnData.Status = 400
		returnData.Message = "查询失败：" + err.Error()
		r.JSON(200, returnData)
		return
	} else {
		var clusterList []map[string]string
		for _, v := range secretList.Items {
			anno := v.Annotations
			clusterList = append(clusterList, anno)
		}

		logs.Info(nil, "查询集群列表成功")
		returnData.Status = 200
		returnData.Message = "查询成功"
		returnData.Data = make(map[string]interface{})
		returnData.Data["items"] = secretList
		r.JSON(200, returnData)
		return
	}
}
