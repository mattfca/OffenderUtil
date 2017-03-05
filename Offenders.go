package OffenderUtil

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
