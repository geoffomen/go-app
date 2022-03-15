package zapimp

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

type ZapLogger struct {
	SugaredLogger *zap.SugaredLogger
	conf          Configuration
}

// Configuration stores the config for the logger
type Configuration struct {
	EnableConsole     bool
	ConsoleJSONFormat bool
	ConsoleLevel      string
	EnableFile        bool
	FileJSONFormat    bool
	FileLevel         string
	FileLocation      string
	ErrFileLevel      string
	ErrFileLocation   string
}

func New(conf Configuration) (*ZapLogger, error) {
	cores := []zapcore.Core{}

	if conf.EnableConsole {
		level := getZapLevel(conf.ConsoleLevel)
		writer := zapcore.Lock(os.Stdout)
		core := zapcore.NewCore(getEncoder(conf.ConsoleJSONFormat), writer, level)
		cores = append(cores, core)
	}

	if conf.EnableFile {
		level := getZapLevel(conf.FileLevel)
		writer := zapcore.AddSync(&lumberjack.Logger{
			Filename: conf.FileLocation,
			MaxSize:  500,
			Compress: true,
			MaxAge:   28,
		})
		core := zapcore.NewCore(getEncoder(conf.FileJSONFormat), writer, level)
		cores = append(cores, core)

		level = getZapLevel(conf.ErrFileLevel)
		writer = zapcore.AddSync(&lumberjack.Logger{
			Filename: conf.ErrFileLocation,
			MaxSize:  500,
			Compress: true,
			MaxAge:   28,
		})
		core = zapcore.NewCore(getEncoder(conf.FileJSONFormat), writer, level)
		cores = append(cores, core)
	}

	combinedCore := zapcore.NewTee(cores...)

	// AddCallerSkip skips 2 number of callers, this is important else the file that gets
	// logged will always be the wrapped file. In our case zap.go
	logger := zap.New(combinedCore,
		zap.AddCallerSkip(2),
		zap.AddCaller(),
	).Sugar()

	return &ZapLogger{
		SugaredLogger: logger,
		conf:          conf,
	}, nil
}

func (l *ZapLogger) Debugf(format string, args ...interface{}) {
	l.SugaredLogger.Debugf(format, args...)
}

// Print ..
func (l *ZapLogger) Print(args ...interface{}) {
	l.SugaredLogger.Infof("%v", args...)
}

// Println ..
func (l *ZapLogger) Println(args ...interface{}) {
	l.SugaredLogger.Infof("%v\n", args...)
}

// Printf ..
func (l *ZapLogger) Printf(format string, args ...interface{}) {
	l.SugaredLogger.Infof(format, args...)
}

func (l *ZapLogger) Infof(format string, args ...interface{}) {
	l.SugaredLogger.Infof(format, args...)
}

func (l *ZapLogger) Warnf(format string, args ...interface{}) {
	l.SugaredLogger.Warnf(format, args...)
}

func (l *ZapLogger) Errorf(format string, args ...interface{}) {
	l.SugaredLogger.Errorf(format, args...)
}

func (l *ZapLogger) Fatalf(format string, args ...interface{}) {
	l.SugaredLogger.Fatalf(format, args...)
}

func (l *ZapLogger) Panicf(format string, args ...interface{}) {
	l.SugaredLogger.Panicf(format, args...)
}

func (l *ZapLogger) WithFields(fields map[string]interface{}) *ZapLogger {
	var f = make([]interface{}, 0)
	for k, v := range fields {
		f = append(f, k)
		f = append(f, v)
	}
	newLogger := l.SugaredLogger.With(f...)
	return &ZapLogger{SugaredLogger: newLogger, conf: l.conf}
}

func getEncoder(isJSON bool) zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	if isJSON {
		return zapcore.NewJSONEncoder(encoderConfig)
	}
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getZapLevel(level string) zapcore.Level {
	switch level {
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "debug":
		return zapcore.DebugLevel
	case "error":
		return zapcore.ErrorLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		panic(fmt.Errorf("failed to init log, unknow log level: %s", level))
	}
}
