package logs

import "github.com/sirupsen/logrus"

// 打印Info级别的日志
func Info(fields map[string]interface{}, msg string) {
	logrus.WithFields(fields).Info(msg)
}

// 打印Warning级别的日志
func Warning(fields map[string]interface{}, msg string) {
	logrus.WithFields(fields).Warning(msg)
}

// 打印Error级别的日志
func Error(fields map[string]interface{}, msg string) {
	logrus.WithFields(fields).Error(msg)
}

// 打印Debug级别的日志
func Debug(fields map[string]interface{}, msg string) {
	logrus.WithFields(fields).Debug(msg)
}
