package rice

import (
	"context"
	"time"

	"github.com/nikoksr/notify"
	"github.com/rs/zerolog"
)

type TelegramHook struct {
	Notifier *notify.Notify
}

func NewTelegramHook(token, proxy string, chatID ...int64) (*TelegramHook, error) {
	notifier, err := NewTelegramBotNotifier(token, proxy, chatID...)
	return &TelegramHook{Notifier: notifier}, err
}

func (t *TelegramHook) Run(e *zerolog.Event, level zerolog.Level, message string) {
	if level > zerolog.DebugLevel {
		go func() {
			_ = t.send(level.String(), message)
		}()
	}
}

func (t *TelegramHook) send(title, msg string) error {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		30*time.Second,
	)
	defer cancel()
	return t.Notifier.Send(ctx, title, msg)
}
