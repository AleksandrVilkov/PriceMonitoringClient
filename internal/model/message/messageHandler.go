package message

import (
	"PriceMonitoringClient/internal/model/user"
	"strings"
)

func HandleMessage(testMessage string, status user.Status) (string, user.Status) {

	if strings.Compare(strings.ToLower(testMessage), START) == 0 {
		return "выбрана команда старт", user.START
	}
	if strings.Compare(strings.ToLower(testMessage), GET_DYNAMIC) == 0 {
		return "выбрана команда Показать динамику", user.NONE
	}
	if strings.Compare(strings.ToLower(testMessage), ADD_NEW_PRODUCT) == 0 {
		//userStatuses[userID] = user.ADD_PRODUCT
		return "выбрана команда  Добавить продукт", user.ADD_PRODUCT
	}
	if strings.Compare(strings.ToLower(testMessage), GET_ALL_URLS) == 0 {
		//userStatuses[userID] = user.NONE
		return "выбрана команда  Показать все url", user.NONE
	}
	if strings.Compare(strings.ToLower(testMessage), DELETE_URL) == 0 {
		//userStatuses[userID] = user.DELETE_PRODUCT
		return "выбрана команда  удалить url", user.DELETE_PRODUCT
	}
	return "Не известная команда", user.NONE
}
