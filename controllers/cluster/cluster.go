package cluster

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// 定义结构体，描述集群集群信息
type ClusterInfo struct {
	Id          string `json:"id"`
	DisplayName string `json:"display_name"`
	// Region      string `json:"region"` // 区域
	// AZ          string `json:"az"`     // 可用区
	City     string `json:"city"`
	District string `json:"district"`
}

// 定义结构体，描述集群状态
type ClusterStatus struct {
	ClusterInfo
	Version string `json:"version"`
	Status  string `json:"status"`
}

// 定义结构体，描述集群配置信息
type ClusterConfig struct {
	ClusterInfo
	Kubeconfig string `json:"kubeconfig"`
}

/*
结构体的方法，用于判断集群是否可用
返回ClusterStatus结构体，包含集群信息、版本和状态
如果集群不可用，返回错误信息
使用示例:
clusterConfig := ClusterConfig{}
clusterConfig.getClusterStatus()
*/
func (c *ClusterConfig) getClusterStatus() (ClusterStatus, error) {
	clusterStatus := ClusterStatus{}
	clusterStatus.ClusterInfo = c.ClusterInfo
	// 创建clientset, 读取kubeconfig
	restConfig, err := clientcmd.RESTConfigFromKubeConfig([]byte(c.Kubeconfig))
	if err != nil {
		return clusterStatus, err
	}
	// 验证
	clientSet, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		return clusterStatus, err
	}
	// 获取版本
	serverVersion, err := clientSet.Discovery().ServerVersion()
	if err != nil {
		return clusterStatus, err
	}
	// 获取版本不报错，表示可用
	clusterStatus.Version = serverVersion.String()
	clusterStatus.Status = "Active"
	return clusterStatus, nil
}
