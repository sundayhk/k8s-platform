package cluster

import (
	"context"
	"k8s-platform/config"
	"k8s-platform/utils"
	"k8s-platform/utils/logs"
	"net/http"

	"github.com/gin-gonic/gin"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func addOrUpdate(r *gin.Context, method string) {
	var methodText string
	if method == "Create" {
		methodText = "创建"
	} else {
		methodText = "更新"
	}

	clusterConfig := ClusterConfig{}
	returnData := config.NewReturnData()
	if err := r.ShouldBindBodyWithJSON(&clusterConfig); err != nil {
		returnData.Status = 400
		returnData.Message = "添加集群配置不完整: " + err.Error()
		r.JSON(200, returnData)
		return
	}
	// 判断集群是否正常
	clusterStatus, err := clusterConfig.getClusterStatus()
	if err != nil {
		returnData.Status = 400
		returnData.Message = "无法获取集群配置: " + err.Error()
		r.JSON(http.StatusOK, returnData)
		logs.Error(map[string]interface{}{"error": err.Error()}, methodText+"添加集群失败,无法获取集群信息")
		return
	}

	logs.Info(map[string]interface{}{"集群名称": clusterConfig.DisplayName, "集群ID": clusterConfig.Id}, "开始"+methodText+"集群")
	// 创建一个集群配置的secret对象
	var clusterConfigSecret corev1.Secret
	clusterConfigSecret.Name = clusterConfig.Id
	clusterConfigSecret.Labels = make(map[string]string)
	clusterConfigSecret.Labels[config.ClusterConfigSecretLabelKey] = config.ClusterConfigSecretLabelValue
	// 添加注释，保存集群的配置信息
	// clusterConfigSecret.Annotations = make(map[string]string)
	// clusterConfigSecret.Annotations["displayName"] = clusterConfig.DisplayName
	// clusterConfigSecret.Annotations["region"] = clusterConfig.Region
	// clusterConfigSecret.Annotations["az"] = clusterConfig.AZ
	m := utils.StructToMap(clusterStatus) // 结构体转map
	clusterConfigSecret.Annotations = m
	// 保存kubeconfig
	clusterConfigSecret.StringData = make(map[string]string) // StringData方法会自动加密data数据
	clusterConfigSecret.StringData["kubeconfig"] = clusterConfig.Kubeconfig
	// 创建secret
	// clientSet使用全局变量config.InClusterClientSet，
	// 在 initcontroller/incluster.go 文件的 metadataInit 函数里，会对 InClusterClientSet 进行初始化
	if method == "Create" {
		_, err = config.InClusterClientSet.CoreV1().Secrets(config.MetadataNamespace).Create(context.TODO(), &clusterConfigSecret, metav1.CreateOptions{})
		return
	} else {
		_, err = config.InClusterClientSet.CoreV1().Secrets(config.MetadataNamespace).Update(context.TODO(), &clusterConfigSecret, metav1.UpdateOptions{})
	}

	if err != nil {
		logs.Error(map[string]interface{}{"集群ID": clusterConfig.Id, "集群名称": clusterConfig.DisplayName, "msg": err.Error()}, "集群"+methodText+"失败")
		returnData.Status = 400
		returnData.Message = "集群" + methodText + "失败: " + err.Error()
		r.JSON(200, returnData)
		return
	} else {
		logs.Error(map[string]interface{}{"集群ID": clusterConfig.Id, "集群名称": clusterConfig.DisplayName, "msg": err.Error()}, "集群"+methodText+"成功")
		returnData.Status = 200
		returnData.Message = "集群" + methodText + "成功"
		r.JSON(200, returnData)
		return
	}
}
