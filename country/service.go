package country

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Service interface {
	FindByName(name string) (*Country, error)
	FindTimezone(timezone string) (*Timezone, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindByName(name string) (*Country, error) {
	country, err := s.repository.FindByName(name)
	if err != nil {
		return nil, err
	}
	return country, err
}

func (s *service) FindTimezone(timezone string) (*Timezone, error) {
	result, err := GetTimezone(timezone)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetTimezone(timezone string) (*Timezone, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://timeapi.io/api/Time/current/zone?timeZone=%s", timezone), nil)
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

	result := Timezone{}
	err = json.Unmarshal([]byte(data), &result)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("timezone not valid")
	}

	return &result, nil
}
