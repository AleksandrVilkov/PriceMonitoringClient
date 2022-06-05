package message

import (
	"PriceMonitoringClient/internal/priceMonitoringConnector"
	"strconv"
)

func HandleMessage(testMessage string, userID int) string {
	return priceMonitoringConnector.SaveUrlForMonitoring("www.lolkek.ru", strconv.Itoa(userID))
}
