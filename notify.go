package rice

import (
	"fmt"
	"time"

	"github.com/chirichan/rice/telegram"
	"github.com/nikoksr/notify"
)

var Notifier = notify.New()

func InitTelegramBotNotifier(token, proxy string, chatID ...int64) error {
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
		return fmt.Errorf("telegramService init fail: %v", err)
	}
	Notifier.UseServices(telegramService)
	return nil
}
