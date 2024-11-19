package polaris_logger

import (
	"context"
	"errors"
	"os"
	"polaris-oj-backend/config"
	"polaris-oj-backend/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

/*
	日志输入：
		1. 需要知道是谁执行的操作
			这里的是谁执行的操作也就是请求用户的基本信息，例如`用户名, 用户id`
			那问题时如何有效的记录这些用户的信息呢？
				方法一：直接从请求的 *gin.context中获取token，然后解析token拿到用户的数据（这里之后在生成token的时候就是用的用户信息生成的）
				方法二：在用户登录之后，记录用户的ip地址，然后将以ip地址为键，用户的用户名，用户id为值存入redis中（但是这个方式会不会因为用户突然改变ip地址，而导致拿不到用户的信息
		2. 需要记录基本的信息
		3. 在调用完相关的日志输出api需要返回一个原生的error

	Logger.Error()
	Logger.Info()

*/

// TODO: 配置logger的环境，开发或者生产
var Logger *zap.Logger = initLogger(true)

func initLogger(isDev bool) *zap.Logger {
	// 日志编码器配置
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "Time",
		LevelKey:       "Level",
		NameKey:        "Logger",
		CallerKey:      "Caller",
		MessageKey:     "Msg",
		StacktraceKey:  "Stacktrace",
		EncodeLevel:    zapcore.CapitalLevelEncoder,   // 日志级别大写
		EncodeTime:     zapcore.ISO8601TimeEncoder,    // 时间格式
		EncodeDuration: zapcore.StringDurationEncoder, // 持续时间格式
		EncodeCaller:   zapcore.FullCallerEncoder,     // 调用文件简短路径
	}

	// 创建日志编码器
	encoder := zapcore.NewJSONEncoder(encoderConfig)

	// 创建日志文件写入器
	logAllFile, _ := os.OpenFile(config.Log.LogPath+"/log_all.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	logErrFile, _ := os.OpenFile(config.Log.LogPath+"/log_err.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	logAllWriter := zapcore.AddSync(logAllFile) // 所有日志写入器
	logErrWriter := zapcore.AddSync(logErrFile) // 错误日志写入器

	// 创建日志级别过滤器
	allLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.InfoLevel // 记录 Info 及以上级别日志
	})
	errorLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel // 记录 Error 及以上级别日志
	})

	// 核心配置
	cores := []zapcore.Core{
		zapcore.NewCore(encoder, logAllWriter, allLevel),   // 所有日志输出到 log_all.log
		zapcore.NewCore(encoder, logErrWriter, errorLevel), // 错误日志输出到 log_err.log
	}

	// 开发模式下日志同时输出到控制台
	if isDev {
		consoleWriter := zapcore.AddSync(os.Stdout) // 控制台输出
		cores = append(cores, zapcore.NewCore(encoder, consoleWriter, allLevel))
	}

	// 合并多个核心
	combinedCore := zapcore.NewTee(cores...)

	// 创建 Logger
	logger := zap.New(combinedCore, zap.AddCaller())

	// 替换全局 Logger（可选）
	zap.ReplaceGlobals(logger)
	return logger
}

func getLogger(ctx any) *zap.Logger {
	switch c := ctx.(type) {
	case *gin.Context:
		if logger, exists := c.Get("logger"); exists {
			return logger.(*zap.Logger)
		}
		return zap.L().With(zap.String("context", "default-gin"))
	case context.Context:
		if userInfo, ok := c.Value(config.Log.LogContextKey).(*utils.Claims); ok {
			return zap.L().With(
				zap.String("userIdentity", userInfo.Identity),
				zap.String("userAccount", userInfo.UserAccount),
			)
		}
		return zap.L().With(zap.String("context", "default-standard"))
	default:
		return zap.L().With(zap.String("context", "unknow"))
	}
}

func Info(ctx any, msg string, fileds ...zap.Field) {
	getLogger(ctx).Info(msg, fileds...)
}

func Error(ctx any, msg string, fileds ...zap.Field) error {
	getLogger(ctx).Error(msg, fileds...)
	return errors.New(msg)
}
