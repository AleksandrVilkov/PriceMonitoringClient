package message

import (
	"PriceMonitoringClient/internal/model/user"
	"PriceMonitoringClient/internal/priceMonitoringConnector"
	"strings"
)

func HandleMessage(testMessage string, status user.Status, userID int) (string, user.Status) {

	//Проверка основных команд
	if strings.Compare(strings.ToLower(testMessage), START) == 0 {
		return WELCOME_MSG, user.START
	}

	if strings.Compare(strings.ToLower(testMessage), ADD_NEW_PRODUCT) == 0 {
		return ADD_PRODUCT_MSG, user.ADD_PRODUCT
	}

	if strings.Compare(strings.ToLower(testMessage), GET_DYNAMIC) == 0 {
		return "выбрана команда Показать динамику", user.NONE
	}

	if strings.Compare(strings.ToLower(testMessage), GET_ALL_URLS) == 0 {
		var name priceMonitoringConnector.MonitoringList
		var urls = name.GetAllUserUrls(string(userID))
		var result = ""
		for _, url := range urls {
			result = result + "\n" + url
		}
		return GET_ALL_URLS_MSG + result, user.NONE
	}
	if strings.Compare(strings.ToLower(testMessage), DELETE_URL) == 0 {
		//userStatuses[userID] = user.DELETE_PRODUCT
		return "выбрана команда  удалить url", user.DELETE_PRODUCT
	}

	//Проверка на основе статусов

	if status == user.START {

	}
	if status == user.NONE {

	}

	if status == user.ADD_PRODUCT {

	}

	if status == user.DELETE_PRODUCT {

	}

	return "Не известная команда", user.NONE
}
