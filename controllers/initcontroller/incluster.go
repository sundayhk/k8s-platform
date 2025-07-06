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
	logs.Debug(nil, "初始化Config定义Namespace")

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

	// 将clientSet赋值给全局变量config.InClusterClientSet，可以让项目里的其他包和函数共享该实例，避免重复创建
	config.InClusterClientSet = clientSet
	inClusterVersion, _ := clientSet.Discovery().ServerVersion()
	// 检查config定义名称空间是否存在
	_, err = clientSet.CoreV1().Namespaces().Get(context.TODO(), config.MetadataNamespace, metav1.GetOptions{})
	if err != nil {
		var metadataNamespace corev1.Namespace
		metadataNamespace.Name = config.MetadataNamespace
		_, err = clientSet.CoreV1().Namespaces().Create(context.TODO(), &metadataNamespace, metav1.CreateOptions{})
		if err != nil {
			logs.Debug(map[string]interface{}{"msg": err.Error()}, "config定义名称空间创建失败")
		} else {
			logs.Info(map[string]interface{}{"Namespace": config.MetadataNamespace, "incluster版本": inClusterVersion.String()}, "config定义名称空间已存在")
		}
	} else {
		logs.Info(map[string]interface{}{"Namespace": config.MetadataNamespace, "incluster版本": inClusterVersion.String()}, "config定义名称空间创建成功")
	}
	// 读取所有secret, 存储到全局变量config.Clusterkubeconfig
	config.Clusterkubeconfig = make(map[string]string)
	listOptions := metav1.ListOptions{
		LabelSelector: config.ClusterConfigSecretLabelKey + "=" + config.ClusterConfigSecretLabelValue,
	}
	secretList, _ := config.InClusterClientSet.CoreV1().Secrets(config.MetadataNamespace).List(context.TODO(), listOptions)
	for _, secret := range secretList.Items {
		clusterId := secret.Name
		kubeconfig := secret.Data["kubeconfig"]
		config.Clusterkubeconfig[clusterId] = string(kubeconfig)
	}
}
