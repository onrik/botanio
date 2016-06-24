package botanio

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

const (
	API_URL = "https://api.botan.io/track"

	STATUS_ACCEPTED = "accepted"
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
func (botan *Botan) Track(userId int, name string, message interface{}) (*Answer, error) {
	body := new(bytes.Buffer)
	if err := json.NewEncoder(body).Encode(message); err != nil {
		return nil, err
	}

	values := url.Values{
		"token": {botan.Token},
		"name":  {name},
		"uid":   {strconv.Itoa(userId)},
	}

	url := fmt.Sprintf("%s?%s", API_URL, values.Encode())
	response, err := http.Post(url, "application/json", body)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	answer := new(Answer)
	if err := json.NewDecoder(response.Body).Decode(answer); err != nil {
		return nil, err
	} else {
		return answer, nil
	}
}
