package priceMonitoringConnector

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type MonitoringList struct {
}

//TODO не работает 400
func (monitoringList *MonitoringList) SaveUrlForMonitoring(url string, userID string) string {

	message := map[string]interface{}{
		"url": "dfsadfsasdffdsdfs",
	}

	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
	}
	reqUrl := GetStartPath() + "/api/monitoring/add/" + userID
	response, err := http.Post(reqUrl, "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		return ""
	}

	body, err := ioutil.ReadAll(response.Body)
	f := string(body)
	if err != nil {
		log.Fatalln(err, f)
	}
	return string(body)
}

func (monitoringList *MonitoringList) GetAllUserUrls(userID string) []string {
	resp, err := http.Get(GetStartPath() + "/api/monitoring/getAllURLs/" + userID)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(body))

	var result []string
	err = json.Unmarshal(body, &result)
	if err != nil {
		panic(err)
		return result
	}
	return result
}
