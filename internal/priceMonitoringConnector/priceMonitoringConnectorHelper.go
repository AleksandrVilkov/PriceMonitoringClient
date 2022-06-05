package priceMonitoringConnector

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func MakeRequest(url *string) string {
	resp, err := http.Get(*url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return string(body)
}

func GetStartPath() string {

	type Authorization struct {
		Token string `json:"start-path"`
	}

	f, err := os.Open("configs/PriceMonitorAPIconfig.json")
	if err != nil {
		panic(err)
	}
	dec := json.NewDecoder(f)
	var auth Authorization
	dec.Decode(&auth)
	defer f.Close()
	return auth.Token

}
