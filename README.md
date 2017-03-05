# OffenderUtil

Simple golang wrapper for http://www.familywatchdog.us api

## Usage

Call OffenderUtil with the search you want

```golang
package main

import (
	"fmt"

	"github.com/mattfca/OffenderUtil"
)

func main() {
	OffenderUtil.SetAPIKey("YOUR-KEY-HERE")

	zipOffenders := OffenderUtil.SearchByZip("90210")
	fmt.Printf("Search Type: %s \nCount: %d\n\n", zipOffenders.Searchtype, len(zipOffenders.Offenders))

	nameStateOffenders := OffenderUtil.SearchByNameState("John", "Doe", "IL")
	fmt.Printf("Search Type: %s \nCount: %d\n\n", nameStateOffenders.Searchtype, len(nameStateOffenders.Offenders))

	latLongOffenders := OffenderUtil.SearchByLatLong("39.9537592", "39.9557592", "75.2694439", "-75.2654439")
	fmt.Printf("Search Type: %s \nCount: %d\n\n", latLongOffenders.Searchtype, len(latLongOffenders.Offenders))
}
```
