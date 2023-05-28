package rice

import (
	"context"
	"fmt"
	"time"

	"github.com/chirichan/rice/telegram"
	"github.com/nikoksr/notify"
	"github.com/rs/zerolog"
)

type TelegramHook struct {
	TelegramService *telegram.Telegram
	Notifier        *notify.Notify
}

func NewTelegramHook(token, proxy string, chatID ...int64) (*TelegramHook, error) {
	var connAttempts = 10
	var telegramService *telegram.Telegram
	var err error
	for connAttempts > 0 {
		telegramService, err = telegram.NewWithHttpClient(token, newHttpClient(httpProxy(proxy)))
		if err == nil {
			break
		}
		time.Sleep(1 * time.Second)
		connAttempts--
	}
	if telegramService == nil {
		return nil, fmt.Errorf("telegramService init fail: %v", err)
	}
	notifier := notify.New()
	notifier.UseServices(telegramService)
	return &TelegramHook{TelegramService: telegramService, Notifier: notifier}, nil
}

func (t *TelegramHook) Run(e *zerolog.Event, level zerolog.Level, message string) {
	if level >= zerolog.DebugLevel {
		go func() {
			_ = t.send("", message)
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
