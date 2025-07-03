package cluster

import (
	"context"
	"k8s-platform/config"
	"k8s-platform/utils/logs"

	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Delete(r *gin.Context) {
	logs.Debug(nil, "删除集群")
	clusterId := r.Query("cluster_id")
	err := config.InClusterClientSet.CoreV1().Secrets(config.MetadataNamespace).Delete(context.TODO(), clusterId, metav1.DeleteOptions{})
	returnData := config.ReturnData{}
	if err != nil {
		logs.Error(map[string]interface{}{"id": clusterId}, "删除失败")
		returnData.Status = 400
		returnData.Message = "集群删除失败: " + err.Error()
	} else {
		logs.Warning(map[string]interface{}{"id": clusterId}, "删除成功")
		returnData.Status = 200
		returnData.Message = "删除成功"
		delete(config.Clusterkubeconfig, clusterId) // 从ClusterKubeconfig map中删除
	}
	r.JSON(200, returnData)
}
