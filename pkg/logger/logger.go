package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"time"
)

var logger *Logger

func InitLogger(option Option) (err error) {
	// 默认utc时间记录日志并只打印在consul, 日志格式为text， 默认日志大小100M保留7个日志副本， 最大保留天数7天， 日志切割不进行压缩日志
	writes := make([]io.Writer, 0)
	// 输出到控制台
	if option.IsConsole {
		writes = append(writes, os.Stdout)
	}
	// 配置日志分割，归档， 如果写入到目录则可以分割，否则不进行分割
	if option.IsFile {
		file := &lumberjack.Logger{
			Filename:   option.Path,
			MaxSize:    option.MaxSize,
			MaxBackups: option.MaxBackup,
			MaxAge:     option.MaxBackup,
			Compress:   option.Compress,
			LocalTime:  option.LocalTime,
		}
		writes = append(writes, file)
	}

	encoderConfig := zap.NewProductionEncoderConfig()
	// 判断是否使用本地时间
	if option.LocalTime {
		encoderConfig.EncodeTime = func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
			encoder.AppendString(time.Format("2006-01-02 15:04:05"))
		}
	} else {
		encoderConfig.EncodeTime = func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
			encoder.AppendString(time.UTC().Format("2006-01-02 15:04:05"))
		}
	}
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewConsoleEncoder(encoderConfig) // text输出日志
	level, err := zapcore.ParseLevel(option.Level)
	if err != nil {
		err = fmt.Errorf("get logger level error %s", err.Error())
		return
	}
	logger = &Logger{
		name:    option.Name,
		level:   level,
		encoder: encoder,
	}
	logger.sugarLog(writes...)
	return
}

func GetLogger() *Logger {
	if logger == nil {
		_ = InitLogger(DefaultOption())
	}
	return logger
}

type Option struct {
	Name      string
	Level     string
	Path      string
	IsConsole bool
	IsFile    bool
	LocalTime bool
	Compress  bool
	Format    string
	MaxSize   int
	MaxBackup int
	MaxAge    int
}

func DefaultOption() Option {
	return Option{
		Name:      "gin-x",
		Level:     "debug",
		Path:      "",
		IsConsole: true,
		IsFile:    false,
		LocalTime: false,
		Compress:  false,
		Format:    "text",
		MaxSize:   100,
		MaxBackup: 7,
		MaxAge:    7,
	}
}

type Logger struct {
	name        string
	level       zapcore.LevelEnabler
	encoder     zapcore.Encoder
	sugarLogger *zap.SugaredLogger
}

func (log *Logger) sugarLog(writes ...io.Writer) {
	cores := make([]zapcore.Core, 0)
	for _, w := range writes {
		cores = append(cores, zapcore.NewCore(log.encoder, zapcore.AddSync(w), log.level))
	}
	log.sugarLogger = zap.New(zapcore.NewTee(cores...), zap.AddCaller()).Named(log.name).Sugar()
	return
}

func (log *Logger) S(namespace string) *zap.SugaredLogger {
	return log.sugarLogger.Named(namespace)
}

func (log *Logger) Sync() (err error) {
	err = log.sugarLogger.Sync()
	return
}
