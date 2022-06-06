package message

import (
	"strings"
)

func HandleMessage(testMessage string, userID int) string {

	if strings.Compare(strings.ToLower(testMessage), START) == 0 {
		return "выбрана команда старт"
	}
	if strings.Compare(strings.ToLower(testMessage), GET_DYNAMIC) == 0 {
		return "выбрана команда Показать динамику"
	}
	if strings.Compare(strings.ToLower(testMessage), ADD_NEW_PRODUCT) == 0 {
		return "выбрана команда  Добавить продукт"
	}
	if strings.Compare(strings.ToLower(testMessage), GET_ALL_URLS) == 0 {
		return "выбрана команда  Показать все url"
	}

	if strings.Compare(strings.ToLower(testMessage), DELETE_URL) == 0 {
		return "выбрана команда  удалить url"
	}
	return "Не известная команда"
	//return priceMonitoringConnector.SaveUrlForMonitoring("www.lolkek.ru", strconv.Itoa(userID))
}
