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

func NewTelegramHook() *TelegramHook {
	return &TelegramHook{Notifier: notify.New()}
}

func (t *TelegramHook) AddTelegramBot(token, proxy string, chatID ...int64) *TelegramHook {
	telegramService, err := NewTelegramService(token, proxy)
	if err != nil {
		return t
	}
	telegramService.AddReceivers(chatID...)
	t.Notifier.UseServices(telegramService)
	return t
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
