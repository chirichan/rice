package rice

import "gopkg.in/natefinch/lumberjack.v2"

const (
	_defaultFileName   = "default.log"
	_defaultMaxSize    = 10   // Mb
	_defaultMaxAge     = 30   // days
	_defaultMaxBackups = 3    // backups
	_defaultCompress   = true // compress
)

func NewLumberjackLogger(fileName string, opts ...LumberjackOption) *lumberjack.Logger {

	lumberjackLogger := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    _defaultMaxSize,
		MaxAge:     _defaultMaxAge,
		MaxBackups: _defaultMaxBackups,
		Compress:   _defaultCompress,
	}

	for _, opt := range opts {
		opt(lumberjackLogger)
	}

	return lumberjackLogger
}

type LumberjackOption func(*lumberjack.Logger)

func MaxSize(i int) LumberjackOption {
	return func(logger *lumberjack.Logger) {
		logger.MaxSize = i
	}
}

func MaxAge(i int) LumberjackOption {
	return func(logger *lumberjack.Logger) {
		logger.MaxAge = i
	}
}

func MaxBackups(i int) LumberjackOption {
	return func(logger *lumberjack.Logger) {
		logger.MaxBackups = i
	}
}

func LocalTime(b bool) LumberjackOption {
	return func(logger *lumberjack.Logger) {
		logger.LocalTime = b
	}
}

func Compress(b bool) LumberjackOption {
	return func(logger *lumberjack.Logger) {
		logger.Compress = b
	}
}
