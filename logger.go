package rice

import (
	"io"
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	_defaultMaxSize    = 512  // Mb
	_defaultMaxAge     = 30   // days
	_defaultMaxBackups = 30   // backups
	_defaultCompress   = true // compress
)

var Logger zerolog.Logger
var levelMap = map[string]zerolog.Level{
	"debug":    zerolog.DebugLevel,
	"info":     zerolog.InfoLevel,
	"warn":     zerolog.WarnLevel,
	"error":    zerolog.ErrorLevel,
	"fatal":    zerolog.FatalLevel,
	"panic":    zerolog.PanicLevel,
	"no":       zerolog.NoLevel,
	"disabled": zerolog.Disabled,
	"trace":    zerolog.TraceLevel,
	"":         zerolog.InfoLevel,
}

func InitLogger(level string, w ...io.Writer) {
	zerolog.SetGlobalLevel(levelMap[strings.ToLower(level)])
	zerolog.TimeFieldFormat = "2006/01/02 15:04:05"
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	var writers = make([]io.Writer, 0)
	if len(w) == 0 {
		w = append(w, zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "2006/01/02 15:04:05"})
	}
	writers = append(writers, w...)
	multi := zerolog.MultiLevelWriter(writers...)
	Logger = zerolog.New(multi).With().Timestamp().Caller().Stack().Logger()
}

func NewLumberjackLogger(fileName string, opts ...lumberjackOption) io.Writer {
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

type lumberjackOption func(*lumberjack.Logger)

func MaxSize(i int) lumberjackOption {
	return func(logger *lumberjack.Logger) {
		logger.MaxSize = i
	}
}

func MaxAge(i int) lumberjackOption {
	return func(logger *lumberjack.Logger) {
		logger.MaxAge = i
	}
}

func MaxBackups(i int) lumberjackOption {
	return func(logger *lumberjack.Logger) {
		logger.MaxBackups = i
	}
}

func LocalTime(b bool) lumberjackOption {
	return func(logger *lumberjack.Logger) {
		logger.LocalTime = b
	}
}

func Compress(b bool) lumberjackOption {
	return func(logger *lumberjack.Logger) {
		logger.Compress = b
	}
}
