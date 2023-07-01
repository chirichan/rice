package rice

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/natefinch/lumberjack"
	"github.com/rs/zerolog"
)

const (
	_defaultMaxSize    = 512  // Mb
	_defaultMaxAge     = 30   // days
	_defaultMaxBackups = 30   // backups
	_defaultCompress   = true // compress
)

// Remote service writer
type RSWriter struct {
	url    string
	client *http.Client
}

func NewRSWriter(url string) *RSWriter {
	return &RSWriter{
		url:    url,
		client: http.DefaultClient,
	}
}

func (w *RSWriter) Write(p []byte) (n int, err error) {
	req, err := http.NewRequest("POST", w.url, bytes.NewReader(p))
	if err != nil {
		return 0, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := w.client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("remote service returned status code %d", resp.StatusCode)
	}
	return len(p), nil
}

// TODO Elasticsearch writer
type ESWriter struct{}

// TODO MQ writer
type MQWriter struct{}

func NewConsoleWriter() zerolog.ConsoleWriter {
	return zerolog.NewConsoleWriter(
		func(w *zerolog.ConsoleWriter) {
			w.FieldsExclude = append(w.FieldsExclude, []string{"user_agent", "git_revision", "go_version"}...)
		},
		func(w *zerolog.ConsoleWriter) {
			w.TimeFormat = consoleDefaultTimeFormat
		},
	)
}

func NewLumberjackLogger(fileName string, opts ...LumberjackOption) io.Writer {
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
