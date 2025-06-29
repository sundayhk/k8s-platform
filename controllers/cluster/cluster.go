package cluster

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// ClusterInfo 集群信息
type ClusterInfo struct {
	Id          string `json:"id"`
	DisplayName string `json:"displayName"`
	City        string `json:"city"`
	District    string `json:"district"`
}

// ClusterStatus 集群状态
type ClusterStatus struct {
	ClusterInfo
	Version string `json:"version"`
	Status  string `json:"status"`
}

// ClusterConfig 集群配置
type ClusterConfig struct {
	ClusterInfo
	KubeConfig string `json:"kubeConfig"`
}

func (c *ClusterConfig) getClusterStatus() (ClusterStatus, error) {
	clusterStatus := ClusterStatus{}
	clusterStatus.ClusterInfo = c.ClusterInfo
	// 读取kubeconfig
	restConfig, err := clientcmd.RESTConfigFromKubeConfig([]byte(c.KubeConfig))
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
	clusterStatus.Version = serverVersion.String()
	clusterStatus.Status = "Active"
	return clusterStatus, nil
}
