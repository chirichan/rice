package rice

import (
	"io"
	"log/slog"
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

const consoleDefaultTimeFormat = "2006/01/02 15:04:05.000000"

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

func InitZerolog(level string, w ...io.Writer) {
	zerolog.SetGlobalLevel(levelMap[strings.ToLower(level)])
	zerolog.TimeFieldFormat = consoleDefaultTimeFormat
	// zerolog.TimestampFieldName = "timestamp"
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	InitLogger(w...)
}

func InitLogger(w ...io.Writer) {
	if len(w) == 0 {
		w = make([]io.Writer, 0)
		w = append(w, NewConsoleWriter())
	}
	multi := zerolog.MultiLevelWriter(w...)

	gitRevision, goVersion := VersionInfo()

	Logger = zerolog.New(multi).With().
		Timestamp().
		Str("git_revision", gitRevision).
		Str("go_version", goVersion).
		Caller().
		Stack().
		Logger()
}

func SetTextSlog() {
	opt := &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}
	handler := slog.NewTextHandler(os.Stdout, opt)
	newSlog := slog.New(handler)
	slog.SetDefault(newSlog)
}

func SetJsonSlog() {
	opt := &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}
	handler := slog.NewTextHandler(os.Stdout, opt)
	newSlog := slog.New(handler)
	slog.SetDefault(newSlog)
}
