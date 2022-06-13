package message

import (
	"PriceMonitoringClient/internal/model/user"
	"PriceMonitoringClient/internal/priceMonitoringConnector"
	"strconv"
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
		var res = priceMonitoringConnector.GetDynamic(strconv.Itoa(userID))
		var result = ""
		for key, value := range res {
			result = result + "🛒 Товар: " + key + ":\n"
			result = result + getStringUniqueProducts(value) + "\n"

		}
		return result, user.NONE
	}

	if strings.Compare(strings.ToLower(testMessage), GET_ALL_URLS) == 0 {

		var result = ""
		for _, url := range priceMonitoringConnector.GetAllUserUrls(strconv.Itoa(userID)) {
			result = result + "\n\n" + url
		}
		return GET_ALL_URLS_MSG + result, user.NONE
	}

	if strings.Compare(strings.ToLower(testMessage), DELETE_URL) == 0 {
		return "выбрана команда  удалить url", user.DELETE_PRODUCT
	}

	//Проверка на основе статусов

	if status == user.START {

	}
	if status == user.NONE {

	}

	if status == user.ADD_PRODUCT {
		var result = priceMonitoringConnector.SaveUrlForMonitoring(testMessage, strconv.Itoa(userID))
		return result, user.NONE
	}

	if status == user.DELETE_PRODUCT {

	}

	return "Не известная команда", user.NONE
}
