package message

type monitoringListInterface interface {
	SaveUrlForMonitoring(url string, userID string) string

	GetAllUserUrls(userID string) []string
}
