package internal

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SendMessage(bot *tgbotapi.BotAPI, chatID int64, messageID int, messageText string) {
	var msg tgbotapi.MessageConfig
	msg = tgbotapi.NewMessage(chatID, messageText)
	msg.ReplyToMessageID = messageID
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}
