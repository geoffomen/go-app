package zapImp

import (
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
	MonitorLocation   string
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

func (l *ZapLogger) Infof(format string, args ...interface{}) {
	l.SugaredLogger.Infof(format, args...)
}

func (l *ZapLogger) Errorf(format string, args ...interface{}) {
	l.SugaredLogger.Errorf(format, args...)
}

func (l *ZapLogger) Fatalf(format string, args ...interface{}) {
	l.SugaredLogger.Errorf(format, args...)
	os.Exit(1)
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
	case "error":
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}
