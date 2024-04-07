package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var lg *zap.Logger

func InitLogger() (err error) {
	// 配置日志级别
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder   // 设置时间格式
	config.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder // 大写字母形式的日志级别

	// 创建文件输出
	file, err := os.Create("app.log")
	if err != nil {
		return err
	}
	// 将文件输出添加到日志记录器
	config.OutputPaths = append(config.OutputPaths, file.Name())

	// 创建Logger实例
	lg, err = config.Build()
	if err != nil {
		return err
	}
	return
}
