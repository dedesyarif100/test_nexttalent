package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"rest-api-golang-v3/country"
)

var Name = ""

func Timezone() (*country.Timezone, error) {
	req, err := http.NewRequest("GET", "https://timeapi.io/api/Time/current/zone?timeZone=Asia/Jakarta", nil)
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}

	data := string([]byte(body))

	result := country.Timezone{}
	err = json.Unmarshal([]byte(data), &result)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &result, nil
}
