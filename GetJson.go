package OffenderUtil

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJSON(params string, target interface{}) error {
	apiKey := "YOUR-KEY-HERE"

	url := fmt.Sprintf("http://services.familywatchdog.us/rest/json.asp?key=%s&lite=0%s", apiKey, params)

	r, err := myClient.Get(url)

	if err != nil {
		return err
	}

	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
