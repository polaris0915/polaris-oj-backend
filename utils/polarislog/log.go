package polaris_log

import (
	"io"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.SugaredLogger = InitLogger()

// 使用自定义的zap logger
func InitLogger() *zap.SugaredLogger {
	// 编码器
	encoder := getEncoder()

	// 定义核心 Core
	core1 := zapcore.NewCore(encoder, getwriteSyncer("./utils/polarislog/log_all.log"), zapcore.DebugLevel)
	core2 := zapcore.NewCore(encoder, getwriteSyncer("./utils/polarislog/log.err.log"), zapcore.ErrorLevel)

	// 使用 filteredCore 包裹双日志
	c := zapcore.NewTee(core1, core2)

	// 创建 logger
	logger := zap.New(
		c,
		zap.AddCaller(),
		zap.AddCallerSkip(2),
	)
	return logger.Sugar()
}

// 设置日志编译器，什么类型的日志
func getEncoder() zapcore.Encoder {
	//encoder配置
	encoderConfig := zap.NewProductionEncoderConfig()
	//设置时间格式为2024-9-1-12.32
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	//json格式
	// jsonencoder := zapcore.NewJSONEncoder(encoderConfig)
	//终端形式
	ConsoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)
	return ConsoleEncoder
}

// 设置输出位置
func getwriteSyncer(logfilename string) zapcore.WriteSyncer {
	//日志文件
	logfile, _ := os.OpenFile(logfilename, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	//只输出到日志文件
	// return zapcore.AddSync(logfile)
	//也输出到终端
	wc := io.MultiWriter(logfile /* , os.Stdout */)
	return zapcore.AddSync(wc)
}
