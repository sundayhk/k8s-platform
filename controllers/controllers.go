package controllers

import (
	"errors"
	"k8s-platform/config"
	"k8s-platform/utils/logs"

	"github.com/gin-gonic/gin"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// 定义全局数据结构
type BasicInfo struct {
	ClusterId string      `json:"cluster_id" form:"cluster_id"` // json通过post传入 form通过params传入
	Namespace string      `json:"namespace" form:"namespace"`
	Name      string      `json:"name" form:"name"`
	Item      interface{} `json:"item"` // 更新namespace时传入
	// DeleteList []string `json:""delete_list` // 支持删除多个pod
}

func BasicInit(r *gin.Context, item interface{}) (client *kubernetes.Clientset, basicInfo BasicInfo, err error) {
	basicInfo = BasicInfo{}
	basicInfo.Item = item
	requestMethod := r.Request.Method // 请求类型
	if requestMethod == "GET" {
		err = r.ShouldBindQuery(&basicInfo) // 绑定查询参数
	} else if requestMethod == "POST" {
		err = r.ShouldBindJSON(&basicInfo) // 绑定json参数
	} else {
		err = errors.New("不支持的请求类型")
	}
	logs.Debug(map[string]interface{}{"basicInfo": basicInfo}, "数据绑定结果")
	if err != nil {
		msg := "请求出错：" + err.Error()
		return nil, basicInfo, errors.New(msg)
	}
	if basicInfo.Namespace == "" {
		basicInfo.Namespace = "default"
	}
	kubeconfig := config.Clusterkubeconfig[basicInfo.ClusterId] // 获取kubeconfig文件
	restConfig, err := clientcmd.RESTConfigFromKubeConfig([]byte(kubeconfig))
	if err != nil {
		msg := "解析kubeconfig出错：" + err.Error()
		return nil, basicInfo, errors.New(msg)
	}
	clientset, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		msg := "解析clientset失败：" + err.Error()
		return nil, basicInfo, errors.New(msg)
	}
	return clientset, basicInfo, nil
}
