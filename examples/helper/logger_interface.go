package helper

// LoggerInterface is our contract for the logger
type LoggerInterface interface {
	Sync()

	Debugf(format string, args ...interface{})

	Debug(args ...interface{})

	Infof(format string, args ...interface{})

	Info(args ...interface{})

	Warnf(format string, args ...interface{})

	Warn(args ...interface{})

	Errorf(format string, args ...interface{})

	Error(args ...interface{})
}
