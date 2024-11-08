package test

import (
	"net/http"
	"testing"

	"go.uber.org/zap"
)

var Logger *zap.SugaredLogger

func InitLogger() *zap.SugaredLogger {
	logger, _ := zap.NewDevelopment()
	return logger.Sugar()
}

func SimpleFunc(url string) {
	res, err := http.Get(url)
	if err != nil {
		Logger.Error(
			"http get failed...",
			zap.String("url: ", url),
			zap.Error(err),
		)
	} else {
		Logger.Info(
			"get success",
			zap.String("status: ", res.Status),
			zap.String("url: ", url),
		)
		res.Body.Close()
	}
}

func TestZap(t *testing.T) {
	Logger = InitLogger()
	defer Logger.Sync()
	SimpleFunc("http://www.baidudssd.com")
}
