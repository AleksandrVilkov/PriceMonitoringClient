package message

import (
	"PriceMonitoringClient/internal/priceMonitoringConnector"
	"PriceMonitoringClient/internal/priceMonitoringConnector/parsers"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

func SendMessage(bot *tgbotapi.BotAPI, chatID int64, messageID int, messageText string) {
	var msg tgbotapi.MessageConfig
	msg = tgbotapi.NewMessage(chatID, messageText)
	msg.ReplyToMessageID = messageID
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}

func getStringUniqueProducts(products []priceMonitoringConnector.Product) string {
	type informationOnProduct struct {
		Year  string
		Month string
		Day   string
		Shop  string
		Price string
	}
	uniqueProducts := make(map[informationOnProduct]bool)
	for _, product := range products {
		year, month, day := parsers.ParseDate(product.Date).Date()
		var informationOnProduct informationOnProduct
		informationOnProduct.Day = strconv.Itoa(day)
		informationOnProduct.Month = month.String()
		informationOnProduct.Year = strconv.Itoa(year)
		informationOnProduct.Shop = product.Shop
		informationOnProduct.Price = product.Price
		uniqueProducts[informationOnProduct] = true
	}
	var result = ""
	for product, _ := range uniqueProducts {
		result = result + "🗓 Дата: " + product.Day + " " + product.Month + " " + product.Year + "\n"
		result = result + "💰 Стоимость в магазине " + product.Shop + ": " + product.Price + "\n"
	}
	return result
}
