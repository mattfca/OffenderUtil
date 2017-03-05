package OffenderUtil

import (
	"fmt"
)

// SearchByLatLong returns sex offenders with name in state
func SearchByLatLong(minLat string, maxLat string, minLong string, maxLong string) Offenders {
	params := fmt.Sprintf("&type=searchbylatlong&minlat=%s&maxlat=%s&minlong=%s&maxlong=%s", minLat, maxLat, minLong, maxLong)
	records := Offenders{}

	// getJSON will return the offenders JSON
	getJSON(params, &records)

	return Offenders(records)
}
