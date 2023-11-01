package logger

import (
	"os"

	"github.com/anggunpermata/patreon-clone/configs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	InitLogger()

	Debug(args ...interface{})
	Debugf(template string, args ...interface{})
	Debugw(msg string, keysAndValues ...interface{})

	Info(args ...interface{})
	Infof(template string, args ...interface{})
	Infow(msg string, keysAndValues ...interface{})

	Warn(args ...interface{})
	Warnf(template string, args ...interface{})
	Warnw(msg string, keysAndValues ...interface{})

	Error(args ...interface{})
	Errorf(template string, args ...interface{})
	Errorw(msg string, keysAndValues ...interface{})

	DPanic(args ...interface{})
	DPanicf(template string, args ...interface{})
	DPanicw(msg string, keysAndValues ...interface{})

	Panic(args ...interface{})
	Panicf(template string, args ...interface{})
	Panicw(msg string, keysAndValues ...interface{})

	Fatal(args ...interface{})
	Fatalf(template string, args ...interface{})
	Fatalw(msg string, keysAndValues ...interface{})
}

type loggerConfig struct {
	cfg         *configs.Config
	sugarLogger *zap.SugaredLogger
}

func NewLoggerConfig(cfg *configs.Config) *loggerConfig {
	return &loggerConfig{
		cfg: cfg,
	}
}

// Mapping various log level
var loggerLevelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

// Get logger level based on the input parameter
func (l *loggerConfig) GetLoggerLevel(cfg *configs.Config) (loggerLevel zapcore.Level) {
	loggerLevel, exist := loggerLevelMap[cfg.LoggerLevel]
	if !exist {
		return loggerLevelMap["debug"]
	}
	return
}

func (l *loggerConfig) InitLogger() {
	logLevel := l.GetLoggerLevel(l.cfg)
	var encoderCfg zapcore.EncoderConfig
	if l.cfg.Mode == "PRODUCTION" {
		encoderCfg = zap.NewProductionEncoderConfig()
	} else {
		encoderCfg = zap.NewDevelopmentEncoderConfig()
	}

	encoderCfg.TimeKey = "time"
	encoderCfg.LevelKey = "level"
	encoderCfg.NameKey = "name"
	encoderCfg.CallerKey = "caller"
	encoderCfg.MessageKey = "message"
	encoderCfg.StacktraceKey = "stacktrace"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.AddSync(os.Stdout),
		zap.NewAtomicLevelAt(logLevel),
	)

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	l.sugarLogger = logger.Sugar()
	l.sugarLogger.Sync()
}

// Logger methods
func (l *loggerConfig) Debug(args ...interface{}) {
	l.sugarLogger.Debug(args...)
}

func (l *loggerConfig) Debugf(template string, args ...interface{}) {
	l.sugarLogger.Debugf(template, args...)
}

func (l *loggerConfig) Debugw(msg string, keysAndValues ...interface{}) {
	l.sugarLogger.Debugw(msg, keysAndValues...)
}

func (l *loggerConfig) Info(args ...interface{}) {
	l.sugarLogger.Info(args...)
}

func (l *loggerConfig) Infof(template string, args ...interface{}) {
	l.sugarLogger.Infof(template, args...)
}

func (l *loggerConfig) Infow(msg string, keysAndValues ...interface{}) {
	l.sugarLogger.Infow(msg, keysAndValues...)
}

func (l *loggerConfig) Warn(args ...interface{}) {
	l.sugarLogger.Warn(args...)
}

func (l *loggerConfig) Warnf(template string, args ...interface{}) {
	l.sugarLogger.Warnf(template, args...)
}

func (l *loggerConfig) Warnw(msg string, keysAndValues ...interface{}) {
	l.sugarLogger.Warnw(msg, keysAndValues...)
}

func (l *loggerConfig) Error(args ...interface{}) {
	l.sugarLogger.Error(args...)
}

func (l *loggerConfig) Errorf(template string, args ...interface{}) {
	l.sugarLogger.Errorf(template, args...)
}

func (l *loggerConfig) Errorw(msg string, keysAndValues ...interface{}) {
	l.sugarLogger.Errorw(msg, keysAndValues...)
}

func (l *loggerConfig) DPanic(args ...interface{}) {
	l.sugarLogger.DPanic(args...)
}

func (l *loggerConfig) DPanicf(template string, args ...interface{}) {
	l.sugarLogger.DPanicf(template, args...)
}

func (l *loggerConfig) DPanicw(msg string, keysAndValues ...interface{}) {
	l.sugarLogger.DPanicw(msg, keysAndValues...)
}

func (l *loggerConfig) Panic(args ...interface{}) {
	l.sugarLogger.Panic(args...)
}

func (l *loggerConfig) Panicf(template string, args ...interface{}) {
	l.sugarLogger.Panicf(template, args...)
}

func (l *loggerConfig) Panicw(msg string, keysAndValues ...interface{}) {
	l.sugarLogger.Panicw(msg, keysAndValues...)
}

func (l *loggerConfig) Fatal(args ...interface{}) {
	l.sugarLogger.Fatal(args...)
}

func (l *loggerConfig) Fatalf(template string, args ...interface{}) {
	l.sugarLogger.Fatalf(template, args...)
}

func (l *loggerConfig) Fatalw(msg string, keysAndValues ...interface{}) {
	l.sugarLogger.Fatalw(msg, keysAndValues...)
}
