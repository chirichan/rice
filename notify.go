package rice

import (
	"fmt"
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/nikoksr/notify"
	"github.com/nikoksr/notify/service/telegram"
)

func NewTelegramBotNotifier(token, proxy string, chatID ...int64) (*notify.Notify, error) {
	telegramService, err := NewTelegramNotifier(token, proxy)
	if err != nil {
		return nil, err
	}
	telegramService.AddReceivers(chatID...)
	notifier := notify.New()
	notifier.UseServices(telegramService)
	return notifier, nil
}

func NewTelegramNotifier(token, proxy string) (*telegram.Telegram, error) {
	telegramBot, err := NewTelegramBot(token, proxy)
	if err != nil {
		return nil, err
	}
	telegramNotifier := &telegram.Telegram{}
	telegramNotifier.SetClient(telegramBot)
	return telegramNotifier, nil
}

func NewTelegramBot(token, proxy string) (*tgbotapi.BotAPI, error) {
	var connAttempts = 10
	var err error
	var botAPI *tgbotapi.BotAPI
	for connAttempts > 0 {
		botAPI, err = tgbotapi.NewBotAPIWithClient(token, newHttpClient(httpProxy(proxy)))
		if err == nil {
			break
		}
		log.Printf("telegram notifier: trying connect, left: %d\n", connAttempts)
		time.Sleep(1 * time.Second)
		connAttempts--
	}
	if botAPI == nil {
		return nil, fmt.Errorf("telegram bot init error: %w", err)
	}
	return botAPI, nil
}
