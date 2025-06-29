package initcontroller

func init() {
	// logs.Debug(nil, "初始化基本数据")
	/*
		1. 通过kubeconfig创建client-go客户端
		2. 检查元数据命名空间是否创建,没有则创建
	*/
	metadataInit()
}
