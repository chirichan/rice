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

var (
	Logger *zerolog.Logger = &zerolog.Logger{}
)

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

func Init(level, path string) {

	if level == "" {
		level = "debug"
	}

	if path == "" {
		path = _defaultFileName
	}

	zerolog.SetGlobalLevel(levelMap[strings.ToLower(level)])
	// zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.TimeFieldFormat = "2006/01/02 15:04:05"
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	writers := make([]io.Writer, 0)

	consoleWriter := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "2006/01/02 15:04:05"}

	fileWriter := NewLumberjackLogger(path)

	writers = append(writers, fileWriter, consoleWriter)

	multi := zerolog.MultiLevelWriter(writers...)

	*Logger = zerolog.New(multi).With().Timestamp().Caller().Stack().Logger()
}
