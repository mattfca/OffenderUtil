package OffenderUtil

import (
	"fmt"
)

// SearchByNameState returns sex offenders with name in state
func SearchByNameState(fname string, lname string, state string) Offenders {
	params := fmt.Sprintf("&type=searchbynamestate&fname=%s&lname=%s&state=%s", fname, lname, state)
	records := Offenders{}

	// getJSON will return the offenders JSON
	getJSON(params, &records)

	return Offenders(records)
}
