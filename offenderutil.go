package OffenderUtil

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Offenders represents a response from FamilyWatchdog
type Offenders struct {
	Searchtype string     `json:"searchtype"`
	Vers       string     `json:"vers"`
	Status     string     `json:"status"`
	Offenders  []Offender `json:"offenders"`
}

// Aliases is offender aliases
type Aliases struct {
	Name       string `json:"name"`
	FirstName  string `json:"firstname"`
	MiddleName string `json:"middlename"`
	LastName   string `json:"lastname"`
	Suffix     string `json:"suffix"`
}

// OtherAddresses are additional addresses
type OtherAddresses struct {
	Type      string `json:"type"`
	Name      string `json:"name"`
	Street1   string `json:"street1"`
	Street2   string `json:"street2"`
	City      string `json:"city"`
	State     string `json:"state"`
	Zipcode   string `json:"zipcode"`
	MatchType string `json:"matchtype"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

// Convictions are the charges
type Convictions struct {
	Charge          string `json:"charge"`
	Statute         string `json:"statute"`
	VictimAge       string `json:"victimage"`
	VictimSex       string `json:"victimsex"`
	ConvictionState string `json:"convictionstate"`
}

// Markings are additional markings
type Markings struct {
	Description string `json:"description"`
}

// Birthdates are unknown
type Birthdates struct {
	Dob string `json:"dob"`
}

// PossibleNicknames are nicknames
type PossibleNicknames struct {
	FirstName string `json:"firstname"`
}

// Offender is a single offender
type Offender struct {
	OffenderID        string              `json:"offenderid"`
	Photo             string              `json:"photo"`
	Name              string              `json:"name"`
	FirstName         string              `json:"firstname"`
	MiddleName        string              `json:"middlename"`
	LastName          string              `json:"lastname"`
	Suffix            string              `json:"suffix"`
	Dob               string              `json:"dob"`
	Age               string              `json:"age"`
	Sex               string              `json:"sex"`
	Race              string              `json:"race"`
	Hair              string              `json:"hair"`
	Eye               string              `json:"eye"`
	Height            string              `json:"height"`
	Weight            string              `json:"weight"`
	Street1           string              `json:"street1"`
	Street2           string              `json:"street2"`
	City              string              `json:"city"`
	State             string              `json:"state"`
	Zipcode           string              `json:"zipcode"`
	County            string              `json:"county"`
	MatchType         string              `json:"matchtype"`
	Latitude          string              `json:"latitude"`
	Longitude         string              `json:"longitude"`
	Aliases           []Aliases           `json:"aliases"`
	OtherAddresses    []OtherAddresses    `json:"otheraddresses"`
	Convictions       []Convictions       `json:"convictions"`
	Markings          []Markings          `json:"markings"`
	Birthdates        []Birthdates        `json:"birthdates"`
	PossibleNicknames []PossibleNicknames `json:"possiblenicknames"`
}

var myClient = &http.Client{Timeout: 120 * time.Second}

var apiKey string

// SetAPIKey will set the api key
func SetAPIKey(a string) {
	apiKey = a
}

// getJSON will make a request and return the response as a struct supplied
func getJSON(params string, target interface{}) error {
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
	err := getJSON(params, &records)
	log.Println(err)

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
