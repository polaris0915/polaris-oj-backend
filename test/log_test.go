package test

import (
	"errors"
	"testing"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
)

func inner() error {
	return errors.New("seems we have an error here")
}

func middle() error {
	err := inner()
	if err != nil {
		return err
	}
	return nil
}

func outer() error {
	err := middle()
	if err != nil {
		return err
	}
	return nil
}

func TestZap(t *testing.T) {
	// zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	// 使用 lumberjack 实现日志文件的切分
	log.Logger = zerolog.New(&lumberjack.Logger{
		Filename:   "app.log",
		MaxSize:    10,   // 每个日志文件的最大尺寸 (MB)
		MaxBackups: 5,    // 保留旧日志文件的最大数量
		MaxAge:     30,   // 保留旧日志文件的最大天数
		Compress:   true, // 是否压缩旧日志
	}).With().Caller().Timestamp().Logger()

	err := outer()
	log.Error().Err(err).Msg("是的")
}
