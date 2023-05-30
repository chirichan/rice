package rice

import (
	"fmt"
	"time"

	"github.com/chirichan/rice/telegram"
	"github.com/nikoksr/notify"
)

func NewTelegramBotNotifier(token, proxy string, chatID ...int64) (*notify.Notify, error) {
	telegramService, err := NewTelegramService(token, proxy)
	if err != nil {
		return nil, err
	}
	telegramService.AddReceivers(chatID...)
	notifier := notify.New()
	notifier.UseServices(telegramService)
	return notifier, nil
}

func NewTelegramService(token, proxy string) (*telegram.Telegram, error) {
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
	return telegramService, nil
}
