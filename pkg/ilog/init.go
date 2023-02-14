package ilog

import (
	"fmt"
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var zapLog *zap.Logger

var debugEnable = false

func Start(path, name string, debug bool) {
	debugEnable = debug
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "linenum",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}

	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	writer := getWriter(fmt.Sprintf("%s/%s", path, name))

	atomicLevel := zap.NewAtomicLevel()

	level := zap.InfoLevel
	atomicLevel.SetLevel(level)

	var zapCore zapcore.Core
	if debugEnable {
		atomicLevel.SetLevel(zap.DebugLevel)
		zapCore = zapcore.NewCore(encoder, zapcore.Lock(os.Stdout), atomicLevel)
	} else {
		zapCore = zapcore.NewCore(encoder, writer, atomicLevel)

	}
	zapLog = zap.New(zapCore)
}

func getWriter(filename string) zapcore.WriteSyncer {
	writeCfg := &lumberjack.Logger{
		Filename: filename,
		MaxSize:  100,
	}
	writer := zapcore.AddSync(writeCfg)

	return writer
}

// Flush 数据flush
func Flush() {
	zapLog.Sync()
}
