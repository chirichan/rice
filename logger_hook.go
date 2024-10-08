package rice

import (
	"bytes"
	"context"
	"time"

	"github.com/nikoksr/notify"
	"github.com/rs/zerolog"
)

type NotifyHook struct {
	Notifier *notify.Notify
}

func NewNotifyHook() *NotifyHook {
	return &NotifyHook{Notifier: notify.New()}
}

func (t *NotifyHook) Run(e *zerolog.Event, level zerolog.Level, message string) {
	if level > zerolog.DebugLevel {
		go func() {
			_ = t.send(level.String(), message)
		}()
	}
}

func (t *NotifyHook) Write(p []byte) (n int, err error) {
	go func() {
		_ = t.send("", "<pre><code>"+bytes.NewBuffer(p).String()+"</code></pre>")
	}()
	return len(p), nil
}

func (t *NotifyHook) send(title, msg string) error {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		30*time.Second,
	)
	defer cancel()
	return t.Notifier.Send(ctx, title, msg)
}
