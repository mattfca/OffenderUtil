package OffenderUtil

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Offenders represents a response from FamilyWatchdog
type Offenders struct {
	Searchtype string `json:"searchtype"`
	Vers       string `json:"vers"`
	Status     string `json:"status"`
	Offenders  []struct {
		OffenderID string `json:"offenderid"`
		Photo      string `json:"photo"`
		Name       string `json:"name"`
		FirstName  string `json:"firstname"`
		MiddleName string `json:"middlename"`
		LastName   string `json:"lastname"`
		Suffix     string `json:"suffix"`
		Dob        string `json:"dob"`
		Age        string `json:"age"`
		Sex        string `json:"sex"`
		Race       string `json:"race"`
		Hair       string `json:"hair"`
		Eye        string `json:"eye"`
		Height     string `json:"height"`
		Weight     string `json:"weight"`
		Street1    string `json:"street1"`
		Street2    string `json:"street2"`
		City       string `json:"city"`
		State      string `json:"state"`
		Zipcode    string `json:"zipcode"`
		County     string `json:"county"`
		MatchType  string `json:"matchtype"`
		Latitude   string `json:"latitude"`
		Longitude  string `json:"longitude"`
	} `json:"offenders"`
}

var myClient = &http.Client{Timeout: 10 * time.Second}

// getJSON will make a request and return the response as a struct supplied
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

// SearchByZip returns sex offenders near zip
func SearchByZip(zip string) Offenders {
	params := fmt.Sprintf("&type=searchbyzipcode&zipcode=%s", zip)
	records := Offenders{}

	// getJSON will return the offenders JSON
	getJSON(params, &records)

	return Offenders(records)
}

// SearchByNameState returns sex offenders with name in state
func SearchByNameState(fname string, lname string, state string) Offenders {
	params := fmt.Sprintf("&type=searchbynamestate&fname=%s&lname=%s&state=%s", fname, lname, state)
	records := Offenders{}

	// getJSON will return the offenders JSON
	getJSON(params, &records)

	return Offenders(records)
}

// SearchByLatLong returns sex offenders with name in state
func SearchByLatLong(minLat string, maxLat string, minLong string, maxLong string) Offenders {
	params := fmt.Sprintf("&type=searchbylatlong&minlat=%s&maxlat=%s&minlong=%s&maxlong=%s", minLat, maxLat, minLong, maxLong)
	records := Offenders{}

	// getJSON will return the offenders JSON
	getJSON(params, &records)

	return Offenders(records)
}
