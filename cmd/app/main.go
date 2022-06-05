package main

import (
	"PriceMonitoringClient/internal/config"
	message2 "PriceMonitoringClient/internal/model/message"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(config.GetToken("configs/token.json"))
	if err != nil {
		panic(err)
	}
	bot.Debug = true
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := bot.GetUpdatesChan(updateConfig)
	for update := range updates {
		if update.Message == nil {
			continue
		}
		message2.SendMessage(bot,
			update.Message.Chat.ID,
			update.Message.MessageID,
			message2.HandleMessage(update.Message.Text, int(update.Message.Chat.ID)))
	}
}
