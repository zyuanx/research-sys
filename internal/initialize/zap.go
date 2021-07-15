package initialize

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

const (
	LogFormatJson    = "json"
	LogFormatConsole = "console"

	TimeKey       = "time"
	LevelKey      = "level"
	NameKey       = "logger"
	CallerKey     = "caller"
	MessageKey    = "msg"
	StacktraceKey = "stacktrace"

	MaxSize    = 1
	MaxBackups = 5
	MaxAge     = 7
)

func Zap() {
	SetLogs(zap.DebugLevel, LogFormatConsole)
}

// SetLogs 设置日志级别、输出格式和日志文件的路径
func SetLogs(logLevel zapcore.Level, logFormat string) {

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        TimeKey,
		LevelKey:       LevelKey,
		NameKey:        NameKey,
		CallerKey:      CallerKey,
		MessageKey:     MessageKey,
		StacktraceKey:  StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,    // 大写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.ShortCallerEncoder,     // 短路径编码器(相对路径+行号)
		EncodeName:     zapcore.FullNameEncoder,
	}

	// 设置日志输出格式
	var encoder zapcore.Encoder
	switch logFormat {
	case LogFormatJson:
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	default:
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	// 添加日志切割归档功能
	fileName := fmt.Sprintf("%s/%04d-%02d-%02d.log", "./logs/zap",
		time.Now().Year(), time.Now().Month(), time.Now().Day())
	hook := lumberjack.Logger{
		Filename:   fileName,   // 日志文件路径
		MaxSize:    MaxSize,    // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: MaxBackups, // 日志文件最多保存多少个备份
		MaxAge:     MaxAge,     // 文件最多保存多少天
		Compress:   true,       // 是否压缩
	}

	core := zapcore.NewCore(
		encoder,                                                                         // 编码器配置
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stderr), zapcore.AddSync(&hook)), // 打印到控制台和文件
		zap.NewAtomicLevelAt(logLevel),                                                  // 日志级别
	)

	// 开启文件及行号
	caller := zap.AddCaller()
	// 开启开发模式，堆栈跟踪
	development := zap.Development()
	// 构造日志
	logger := zap.New(core, caller, development)

	// 将自定义的logger替换为全局的logger
	zap.ReplaceGlobals(logger)
}
