package http

import (
	"io/ioutil"
	"log"
	"net/http"
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