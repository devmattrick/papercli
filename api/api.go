package api

import (
	"encoding/json"
	"log"
	"net/http"
)

func GetJson(url string, v interface{}) error {
	resp, err := http.Get(url)

	if err != nil {
		log.Fatalln("Unable to connect to Paper API: ")
		log.Fatalln(err)
		return err
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
		log.Fatalln("Unable to parse response JSON: ")
		log.Fatalln(err)
		return err
	}

	return nil
}
