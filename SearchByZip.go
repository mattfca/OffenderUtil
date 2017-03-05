package OffenderUtil

import (
	"fmt"
)

// SearchByZip returns sex offenders near zip
func SearchByZip(zip string) Offenders {
	params := fmt.Sprintf("&type=searchbyzipcode&zipcode=%s", zip)
	records := Offenders{}

	// getJSON will return the offenders JSON
	getJSON(params, &records)

	return Offenders(records)
}
