package config

import (
	"k8s-platform/utils/logs"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"k8s.io/client-go/kubernetes"
)

const (
	TimeFormat                    string = "2006-01-02 15:04:05"
	ClusterConfigSecretLabelKey   string = "sundayhk.com/cluster.metadata"
	ClusterConfigSecretLabelValue string = "true"
)

// 全局变量
var (
	Port          string // 端口号
	JwtSignKey    string // jwt secret
	JwtExpireTime int64  // jwt token 过期时间 单位分钟
	UserName      string
	Password      string
	// inCluster
	MetadataNamespace  string                // 名称空间
	InClusterClientSet *kubernetes.Clientset // 保存clientset类型到InClusterClientSet
	Clusterkubeconfig  map[string]string     // kubeconfig文件
)

type ReturnData struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

// 默认值构造函数
func NewReturnData() ReturnData {
	ReturnData := ReturnData{}
	ReturnData.Status = 200
	ReturnData.Data = make(map[string]interface{})
	return ReturnData
}

// func NewReturnData() *ReturnData {
// 	return &ReturnData{
// 		Status: 200,
// 		Data:   make(map[string]interface{}),
// 	}
// }

func initLogConfig(logLevel string) {
	if logLevel == "debug" {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}
	logrus.SetReportCaller(true) // 文件名行号
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: TimeFormat,
		// runtime.Frame: 帧 可用于获取调用者返回的PC值函数，文件或行信息
		CallerPrettyfier: func(f *runtime.Frame) (function string, file string) {
			fileName := path.Base(f.File)
			return f.Function, fileName
		},
	})
}

func init() {
	// 设置默认值
	viper.SetDefault("LOG_LEVEL", "debug")
	viper.SetDefault("PORT", ":8080")
	viper.SetDefault("JWT_SIGN_KEY", "sundayhk.com") // jwt secret
	viper.SetDefault("JWT_EXPIRE_TIME", 120)         // jwt token 过期时间 单位分钟
	viper.SetDefault("USERNAME", "sunday")
	viper.SetDefault("PASSWORD", "sunday")
	viper.SetDefault("METADATA_NAMESPACE", "krm")

	// 获取系统环境变量
	// viper.AutomaticEnv()                     // USERNAME冲突
	logLevel := viper.GetString("LOG_LEVEL") // 从环境变量中获取 LOG_LEVEL
	initLogConfig(logLevel)
	logs.Debug(nil, "加载默认配置")

	Port = viper.GetString("PORT")
	JwtSignKey = viper.GetString("JWT_SIGN_KEY")
	JwtExpireTime = viper.GetInt64("JWT_EXPIRE_TIME")
	UserName = viper.GetString("USERNAME")
	Password = viper.GetString("PASSWORD")
	MetadataNamespace = viper.GetString("METADATA_NAMESPACE")
	logs.Debug(map[string]interface{}{"用户名": UserName, "密码": Password}, "获取用户密码配置")
}
