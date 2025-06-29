package initcontroller

import (
	"context"
	"k8s-platform/config"
	"k8s-platform/utils/logs"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func metadataInit() {
	logs.Debug(nil, "初始化原数据命名空间")

	// 初始化config实例
	restConfig, err := clientcmd.BuildConfigFromFlags("", "config/meta.kubeconfig")
	if err != nil {
		logs.Debug(map[string]interface{}{"msg": err.Error()}, "inCluster kubeconfig加载失败")
		panic(err.Error())
	}

	// 创建客户端工具
	clientSet, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		logs.Debug(map[string]interface{}{"msg": err.Error()}, "inCluster客户端创建失败")
		panic(err.Error())
	}

	inClusterVersion, _ := clientSet.Discovery().ServerVersion()
	// 检查元数据命名空间是否存在
	_, err = clientSet.CoreV1().Namespaces().Get(context.TODO(), config.MetadataNamespace, metav1.GetOptions{})
	if err != nil {
		var metadataNamespace corev1.Namespace
		metadataNamespace.Name = config.MetadataNamespace
		_, err = clientSet.CoreV1().Namespaces().Create(context.TODO(), &metadataNamespace, metav1.CreateOptions{})
		if err != nil {
			logs.Debug(map[string]interface{}{"msg": err.Error()}, "元数据命名空间创建失败")
		} else {
			logs.Info(map[string]interface{}{"Namespace": config.MetadataNamespace, "incluster版本": inClusterVersion.String()}, "元数据命名空间已存在")
		}

	} else {
		logs.Info(map[string]interface{}{"Namespace": config.MetadataNamespace, "incluster版本": inClusterVersion.String()}, "元数据命名空间创建成功")
	}

}
