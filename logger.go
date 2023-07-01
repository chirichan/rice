package rice

import (
	"io"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

const consoleDefaultTimeFormat = "2006/01/02 15:04:05.000"

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
	zerolog.TimeFieldFormat = consoleDefaultTimeFormat
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
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
