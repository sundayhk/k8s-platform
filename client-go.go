// client-go测试用例
package main

// deployment需要引用依赖
// 名称空间需要引用依赖

// func main() {
// 	// 1.初始化kubeconfig实例 由于config配置文件中已经有server地址 第一个参数可以省略 ""
// 	config, err := clientcmd.BuildConfigFromFlags("", "./config/meta.kubeconfig")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	// 2. 创建客户端 clientset
// 	clientset, err := kubernetes.NewForConfig(config)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	// 3.操作集群
// 	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
// 	if err != nil {
// 		logs.Error(nil, "查询Pod列表失败")
// 	} else {
// 		fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
// 	}

// 	// 查询deployment列表
// 	deployment, _ := clientset.AppsV1().Deployments("").List(context.TODO(), metav1.ListOptions{})
// 	deploymentItems := deployment.Items
// 	for _, deploy := range deploymentItems {
// 		fmt.Printf("当前资源的名字是：%s 名称空间：%s\n", deploy.Name, deploy.Namespace)
// 	}

// 	// 查询jobs
// 	jobs, _ := clientset.BatchV1().Jobs("").List(context.TODO(), metav1.ListOptions{})
// 	for _, job := range jobs.Items {
// 		fmt.Printf("当前job：%s 名称空间：%s\n", job.Name, job.Namespace)
// 	}

// 	// 通过get方法查询单个pod详情
// 	describe, err := clientset.CoreV1().Pods("default").Get(context.TODO(), "nginx-6c77bc8877-dzgsx", metav1.GetOptions{})
// 	if err != nil {
// 		fmt.Println("查询详情失败")
// 	} else {
// 		fmt.Println("查询到详情:", describe)
// 		fmt.Println("镜像名字", describe.Spec.Containers[0].Image)
// 	}

// 	// 获取名称空间详情
// 	nsDescribe, _ := clientset.CoreV1().Namespaces().Get(context.TODO(), "default", metav1.GetOptions{})
// 	fmt.Println(nsDescribe.Status)

// 	// 获取deployment并修改 更新操作
// 	deploymentDescribe, _ := clientset.AppsV1().Deployments("default").Get(context.TODO(), "nginx", metav1.GetOptions{})
// 	fmt.Println("查询到deployment名字是：", deploymentDescribe.Name)
// 	// labels := deploymentDescribe.Labels
// 	// labels["key"] = "value" //由于类型是map引用数据不需要重新赋值给deploymentDescribe
// 	deploymentDescribe.Annotations["key"] = "value"

// 	// 修改deployment副本数
// 	newReplicas := int32(3)
// 	deploymentDescribe.Spec.Replicas = &newReplicas

// 	// 修改deployment镜像
// 	deploymentDescribe.Spec.Template.Spec.Containers[0].Image = "registry.cn-shanghai.aliyuncs.com/rushbi/tengine-vts:2.3.2"
// 	_, err = clientset.AppsV1().Deployments("default").Update(context.TODO(), deploymentDescribe, metav1.UpdateOptions{})
// 	if err != nil {
// 		fmt.Println("更新失败", err.Error())
// 	} else {
// 		fmt.Println("更新成功")
// 	}

// 	// 删除一个POD
// 	err = clientset.CoreV1().Pods("default").Delete(context.TODO(), "nginx-57d9bf8fc7-t6vfk", metav1.DeleteOptions{})
// 	if err != nil {
// 		fmt.Println("删除失败", err.Error())
// 	} else {
// 		fmt.Println("删除ok")
// 	}

// 	// 删除一个deployment
// 	err = clientset.AppsV1().Deployments("default").Delete(context.TODO(), "nginx", metav1.DeleteOptions{})
// 	if err != nil {
// 		fmt.Println("删除deployment失败", err.Error())
// 	} else {
// 		fmt.Println("删除deployment成功")
// 	}

// 	// 创建名称空间
// 	var createNamespace corev1.Namespace //声明一个namespace实例
// 	createNamespace.Name = "test1"
// 	_, err = clientset.CoreV1().Namespaces().Create(context.TODO(), &createNamespace, metav1.CreateOptions{})
// 	if err != nil {
// 		fmt.Println("创建namespace失败", err.Error())
// 	}
// 	// clientset.CoreV1().Namespaces().Delete(context.TODO(), "test1", metav1.DeleteOptions{}) //删除test1名称空间

// 	// 创建一个deployment
// 	var createDeploymentInstance appsv1.Deployment
// 	createDeploymentInstance.Name = "nginx"
// 	createDeploymentInstance.Namespace = "test1"
// 	label := make(map[string]string)
// 	label["app"] = "nginx" //selector label
// 	label["version"] = "v1"
// 	createDeploymentInstance.Labels = label                          // deployment 自身的labels
// 	createDeploymentInstance.Spec.Selector = &metav1.LabelSelector{} // Selector必须创建空指针否者报错
// 	createDeploymentInstance.Spec.Selector.MatchLabels = label       // deployment 匹配pod的labels
// 	createDeploymentInstance.Spec.Template.Labels = label            // pod的labels

// 	// 创建容器或者声明一个container
// 	var containers []corev1.Container
// 	// createDeploymentInstance.Spec.Template.Spec.Containers[0].Name = "nginx"
// 	// createDeploymentInstance.Spec.Template.Spec.Containers[0].Image = "registry.cn-shanghai.aliyuncs.com/rushbi/tengine-vts:2.3.2"
// 	var container corev1.Container
// 	container.Name = "nginx"
// 	container.Image = "registry.cn-shanghai.aliyuncs.com/rushbi/tengine-vts:2.3.2"
// 	containers = append(containers, container)
// 	container.Name = "redis-exporter"
// 	container.Image = "registry.cn-shanghai.aliyuncs.com/rushbi/redis_exporter:v1.50.0"
// 	containers = append(containers, container)

// 	// 将容器2 通过append追加进去
// 	// createDeploymentInstance.Spec.Template.Spec.Containers = append(createDeploymentInstance.Spec.Template.Spec.Containers, container)
// 	createDeploymentInstance.Spec.Template.Spec.Containers = containers

// 	_, err = clientset.AppsV1().Deployments("test1").Create(context.TODO(), &createDeploymentInstance, metav1.CreateOptions{})
// 	if err != nil {
// 		fmt.Println("创建deployment失败", err.Error())
// 	}

// 	// 通过json串创建资源
// 	deployJson := `{
// 		"kind": "Deployment",
// 		"apiVersion": "apps/v1",
// 		"metadata": {
// 			"name": "redis",
// 			"creationTimestamp": null,
// 			"labels": {
// 				"app": "redis"
// 			}
// 		},
// 		"spec": {
// 			"replicas": 1,
// 			"selector": {
// 				"matchLabels": {
// 					"app": "redis"
// 				}
// 			},
// 			"template": {
// 				"metadata": {
// 					"creationTimestamp": null,
// 					"labels": {
// 						"app": "redis"
// 					}
// 				},
// 				"spec": {
// 					"containers": [
// 						{
// 							"name": "redis",
// 							"image": "redis",
// 							"resources": {}
// 						}
// 					]
// 				}
// 			},
// 			"strategy": {}
// 		},
// 		"status": {}
// 	}`
// 	var newDeployment2 appsv1.Deployment
// 	// json转struct结构体
// 	err = json.Unmarshal([]byte(deployJson), &newDeployment2)
// 	if err != nil {
// 		fmt.Println("json转struct失败:", err.Error())
// 	}
// 	fmt.Println("json转struct之后的配置详情:", newDeployment2)
// 	_, err = clientset.AppsV1().Deployments("default").Create(context.TODO(), &newDeployment2, metav1.CreateOptions{})
// 	if err != nil {
// 		fmt.Println("创建deployment失败:", err.Error())
// 	}
// }
