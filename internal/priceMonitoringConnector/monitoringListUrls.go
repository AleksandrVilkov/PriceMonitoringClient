package priceMonitoringConnector

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Message struct {
	Status string `json:"status"`
	Text   string `json:"messageText"`
}

type Product struct {
	Date  string `json:"date"`
	Price string `json:"price"`
	Shop  string `json:"shop"`
}

func SaveUrlForMonitoring(url string, userID string) string {

	message := map[string]interface{}{
		"url": url,
	}

	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
	}
	reqUrl := GetStartPath() + "/api/monitoring/add/" + userID + "?url=" + url
	response, err := http.Post(reqUrl, "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		return "Произошла неизвестная ошибка"
	}

	body, err := ioutil.ReadAll(response.Body)
	var res Message
	err = json.Unmarshal(body, &res)
	if strings.Contains(res.Status, "SUCCESS") {
		return "Все прошло отлично. " + res.Text
	} else {
		return "Ошибочка вышла. " + res.Text
	}
}

func GetAllUserUrls(userID string) []string {
	resp, err := http.Get(GetStartPath() + "/api/monitoring/getAllURLs/" + userID)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var result []string
	err = json.Unmarshal(body, &result)
	if err != nil {
		panic(err)
		return result
	}
	return result
}

func GetDynamic(userID string) map[string][]Product {
	resp, err := http.Get(GetStartPath() + "/api/analytics/getDynamicPrice/" + userID)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	var f = string(body)
	fmt.Println(f)
	if err != nil {
		log.Fatalln(err)
	}
	var response map[string][]Product
	err = json.Unmarshal(body, &response)

	return response
}
