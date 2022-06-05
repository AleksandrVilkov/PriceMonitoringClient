package config

import (
	"encoding/json"
	"os"
)

func GetToken(pathToJson string) string {

	type Authorization struct {
		Token string `json:"token"`
	}

	f, err := os.Open(pathToJson)
	if err != nil {
		panic(err)
	}
	dec := json.NewDecoder(f)
	var auth Authorization
	dec.Decode(&auth)
	defer f.Close()
	return auth.Token

}
