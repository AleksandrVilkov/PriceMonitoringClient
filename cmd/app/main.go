package main

import (
	"PriceMonitoringClient/internal/config"
	message2 "PriceMonitoringClient/internal/model/message"
	"PriceMonitoringClient/internal/model/user"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(config.GetToken("configs/token.json"))
	if err != nil {
		panic(err)
	}
	bot.Debug = true

	var userStatuses = make(map[int]user.Status)
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := bot.GetUpdatesChan(updateConfig)
	for update := range updates {
		if update.Message == nil {
			continue
		}

		answer, userStatus := message2.HandleMessage(update.Message.Text, userStatuses[int(update.Message.Chat.ID)])
		message2.SendMessage(bot,
			update.Message.Chat.ID,
			update.Message.MessageID, answer)
		userStatuses[int(update.Message.Chat.ID)] = userStatus
	}
}
