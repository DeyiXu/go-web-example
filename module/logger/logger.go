package logger

import (
	"os"

	"github.com/spf13/viper"

	"github.com/sirupsen/logrus"
)

// Log Create a new instance of the logger. You can have any number of instances.
var log *logrus.Logger

// Init 初始化
func Init() {
	log = logrus.New()
	log.Out = os.Stdout
	// 只记录警告严重程度或以上。
	level, err := logrus.ParseLevel(viper.GetString("log.level"))
	if err != nil {
		log.Panicln(err)
		level, _ = logrus.ParseLevel("debug")
	}
	log.SetLevel(level)
}

// Fatalf 致命
func Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}

// Warningf 警告
func Warningf(format string, args ...interface{}) {
	log.Warningf(format, args...)
}

// Errorf 错误
func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

// Infof 信息
func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

// Printf 打印
func Printf(format string, args ...interface{}) {
	log.Printf(format, args...)
}

// Debugln 测试
func Debugln(args ...interface{}) {
	log.Debugln(args...)
}

// Debugf 测试
func Debugf(format string, args ...interface{}) {
	log.Debugf(format, args...)
}

// Infoln 消息
func Infoln(args ...interface{}) {
	log.Infoln(args...)
}

// Println 打印
func Println(args ...interface{}) {
	log.Println(args...)
}

// Warnln 警告
func Warnln(args ...interface{}) {
	log.Warnln(args...)
}

// Warningln 警告
func Warningln(args ...interface{}) {
	log.Warningln(args...)
}

// Errorln 错误
func Errorln(args ...interface{}) {
	log.Errorln(args...)
}

// Fatalln 严重
func Fatalln(args ...interface{}) {
	log.Fatalln(args...)
}

// Panicln 恐慌
func Panicln(args ...interface{}) {
	log.Panicln(args...)
}
