package cluster

import (
	"context"
	"k8s-platform/config"
	"k8s-platform/utils/logs"

	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Get(r *gin.Context) {
	logs.Debug(nil, "查询集群")
	clusterId := r.Query("cluster_id")
	clusterSecret, err := config.InClusterClientSet.CoreV1().Secrets(config.MetadataNamespace).Get(context.TODO(), clusterId, metav1.GetOptions{})
	returnData := config.ReturnData{}
	if err != nil {
		logs.Error(map[string]interface{}{"id": clusterId, "msg": err.Error()}, "查询集群失败")
		returnData.Status = 400
		returnData.Message = "查询集群失败：" + err.Error()
	} else {
		returnData.Status = 200
		returnData.Message = "查询成功"
		returnData.Data = make(map[string]interface{})
		clusterConfigMap := clusterSecret.Annotations
		clusterConfigMap["kubeconfig"] = string(clusterSecret.Data["kubeconfig"])
		returnData.Data["items"] = clusterConfigMap
	}
	r.JSON(200, returnData)
}
