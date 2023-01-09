package helper

import (
	"os"
	"time"

	"go.uber.org/zap/zapcore"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"

	"go.uber.org/zap"
)

// Level Log level
type Level string

const (
	//DebugLevel has verbose message
	DebugLevel Level = "debug"
	//InfoLevel is default log level
	InfoLevel Level = "info"
	//WarnLevel is for logging messages about possible issues
	WarnLevel Level = "warn"
	//ErrorLevel is for logging errors
	ErrorLevel Level = "error"
	//FatalLevel is for logging fatal messages. The system shutdown after logging the message.
	FatalLevel Level = "fatal"
)

// LogConfig ...
type LogConfig struct {
	//  是否开启控制台日志输出
	ConsoleStdoutEnable bool
	//  控制台日志是否是 JSON 格式
	ConsoleStdoutIsJSONFormat bool
	// Level 控制台日志等级
	ConsoleStdoutLevel Level

	// 是否开启控制台日志输出
	FileStdoutEnable bool
	// 控制台日志是否是 JSON 格式
	FileStdoutIsJSONFormat bool
	// Level 控制台日志等级
	FileStdoutLevel Level
	// 写入文件位置
	FileStdoutFileLocation string
	FileStdoutLogMaxSize   int
	FileStdoutCompress     bool
	FileStdoutLogMaxAge    int

	// 业务日志
	// 是否开启业务日志
	BusinessStdoutEnable bool
	// 写入文件位置
	BusinessStdoutFileLocation string
	BusinessStdoutLogMaxSize   int
	BusinessStdoutCompress     bool
	BusinessStdoutLogMaxAge    int
}

// zapLogger ...
type zapLogger struct {
	cfg                LogConfig
	sugaredLogger      *zap.SugaredLogger
	fileAtomicLevel    *zap.AtomicLevel
	consoleAtomicLevel *zap.AtomicLevel
	businessLogger     *zap.SugaredLogger
}

var url = "https://github.com/uber-go/zap"

// QuickStart ...
func QuickStart() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", url)

	QuickStart2()
}

// QuickStart2 ...
func QuickStart2() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	logger.Info("failed to fetch URL",
		// Structured context as strongly typed Field values.
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}

//---------------以下为库代码----------------------

// InitLogger ...
func InitLogger() (LoggerInterface, error) {
	cores := make([]zapcore.Core, 0)

	logger := &zapLogger{}

	//encoder
	encoder := getEncoder(true)
	// encoder := getBusinessEncoder()

	//log level
	logLevel := newLogLevel("debug")

	//console core
	writer := zapcore.Lock(os.Stdout)
	consoleCore := zapcore.NewCore(encoder, writer, logLevel)
	cores = append(cores, consoleCore)

	//filer core
	fileWriter := getLogWriter()
	fileCore := zapcore.NewCore(encoder, fileWriter, logLevel)
	cores = append(cores, fileCore)

	mixCore := zapcore.NewTee(cores...)

	logger.sugaredLogger = zap.New(mixCore, zap.AddStacktrace(zapcore.ErrorLevel), zap.AddCaller(), zap.AddCallerSkip(1)).Sugar()
	return logger, nil
}

func getEncoder(isJSON bool) zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	if isJSON {
		return zapcore.NewJSONEncoder(encoderConfig)
	}
	// 开启 level 染色
	encoderConfig.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getBusinessEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = ""
	encoderConfig.LevelKey = ""
	encoderConfig.CallerKey = ""
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// 仅支持 json 格式
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getZapLevel(level Level) zapcore.Level {
	switch level {
	case InfoLevel:
		return zapcore.InfoLevel
	case WarnLevel:
		return zapcore.WarnLevel
	case DebugLevel:
		return zapcore.DebugLevel
	case ErrorLevel:
		return zapcore.ErrorLevel
	case FatalLevel:
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}

func newLogLevel(level Level) zap.AtomicLevel {
	logLevel := zap.NewAtomicLevel()
	logLevel.SetLevel(getZapLevel(level))
	return logLevel
}

func getLogWriter() zapcore.WriteSyncer {
	// file, _ := os.Create("./logs/test.log")
	// return zapcore.AddSync(file)
	writer := &lumberjack.Logger{

		// Filename is the file to write logs to.  Backup log files will be retained
		// in the same directory.  It uses <processname>-lumberjack.log in
		// os.TempDir() if empty.
		Filename: "./logs/server.log",

		// MaxSize is the maximum size in megabytes of the log file before it gets
		// rotated. It defaults to 100 megabytes.
		MaxSize: 500, //mb

		// Compress determines if the rotated log files should be compressed
		// using gzip. The default is not to perform compression.
		Compress: false,

		// MaxAge is the maximum number of days to retain old log files based on the
		// timestamp encoded in their filename.  Note that a day is defined as 24
		// hours and may not exactly correspond to calendar days due to daylight
		// savings, leap seconds, etc. The default is not to remove old log files
		// based on age.
		MaxAge: 28,
	}
	return zapcore.AddSync(writer)
}

// TestLogger ...
func TestLogger() {
	logger, _ := InitLogger()
	defer logger.Sync()
	logger.Debug("hello Debug")
	logger.Info("hello Info")
	logger.Error("hello Error")
}

func (log *zapLogger) Sync() {
	log.sugaredLogger.Sync()
}

func (log *zapLogger) Debug(args ...interface{}) {
	log.sugaredLogger.Debug(args...)
}

func (log *zapLogger) Debugf(template string, args ...interface{}) {
	log.sugaredLogger.Debugf(template, args...)
}

func (log *zapLogger) Info(args ...interface{}) {
	log.sugaredLogger.Info(args...)
}

func (log *zapLogger) Infof(template string, args ...interface{}) {
	log.sugaredLogger.Infof(template, args...)
}

func (log *zapLogger) Warn(args ...interface{}) {
	log.sugaredLogger.Warn(args...)
}

func (log *zapLogger) Warnf(template string, args ...interface{}) {
	log.sugaredLogger.Warnf(template, args...)
}

func (log *zapLogger) Error(args ...interface{}) {
	log.sugaredLogger.Error(args...)
}

func (log *zapLogger) Errorf(template string, args ...interface{}) {
	log.sugaredLogger.Errorf(template, args...)
}
