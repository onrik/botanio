package botanio

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

const (
	TRACK_URL     = "https://api.botan.io/track"
	SHORTENER_URL = "https://api.botan.io/s/"

	STATUS_ACCEPTED = "accepted"
)

var (
	token = ""
)

type Answer struct {
	Status string `json:"status"`
	Info   string `json:"info"`
}

type Botan struct {
	Token string
}

type Map map[string]interface{}

func New(token string) *Botan {
	return &Botan{Token: token}
}

// Track action
func (botan *Botan) Track(userID int, name string, message interface{}) (*Answer, error) {
	return track(botan.Token, userID, name, message)
}

// Short url
func (botan *Botan) Short(userID int, url string) (string, error) {
	return short(botan.Token, userID, url)
}

// Set token
func SetToken(t string) {
	token = t
}

// Track action
func Track(userID int, name string, message interface{}) (*Answer, error) {
	return track(token, userID, name, message)
}

// Short url
func Short(userID int, url string) (string, error) {
	return short(token, userID, url)
}

func track(token string, userID int, name string, message interface{}) (*Answer, error) {
	body := new(bytes.Buffer)
	if err := json.NewEncoder(body).Encode(message); err != nil {
		return nil, err
	}

	values := url.Values{
		"token": {token},
		"name":  {name},
		"uid":   {strconv.Itoa(userID)},
	}

	url := fmt.Sprintf("%s?%s", TRACK_URL, values.Encode())
	response, err := http.Post(url, "application/json", body)
	if response != nil {
		defer response.Body.Close()
	}

	if err != nil {
		return nil, err
	}

	if response.StatusCode >= http.StatusBadRequest {
		return nil, fmt.Errorf("Status: %d", response.StatusCode)
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	answer := new(Answer)
	if err := json.Unmarshal(data, answer); err != nil {
		return nil, err
	} else {
		return answer, nil
	}
}

func short(token string, userID int, u string) (string, error) {
	values := url.Values{
		"token":    {token},
		"url":      {u},
		"user_ids": {strconv.Itoa(userID)},
	}

	response, err := http.Get(fmt.Sprintf("%s?%s", SHORTENER_URL, values.Encode()))
	if response != nil {
		defer response.Body.Close()
	}

	if err != nil {
		return "", err
	}

	if response.StatusCode >= http.StatusBadRequest {
		return "", fmt.Errorf("Status: %d", response.StatusCode)
	}

	if data, err := ioutil.ReadAll(response.Body); err != nil {
		return "", err
	} else {
		return string(data), nil
	}

}
