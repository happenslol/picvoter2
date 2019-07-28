package actions

import (
	"strings"
    "fmt"

	"github.com/gobuffalo/envy"
	"github.com/spf13/afero"
)

// M is used to construct json responses easily
type M map[string]interface{}

var fs = afero.NewOsFs()
var storagePath = envy.Get("STORAGE_LOCATION", "")
var scanLocations = getScanLocations()

func getScanLocations() []string {
	locationsString := envy.Get("SCAN_LOCATIONS", "")
	parts := strings.Split(locationsString, ",")
	for i, p := range parts {
		parts[i] = strings.Trim(p, " ")
	}

    fmt.Printf("looking in locations: %v\n", parts)

	return parts
}
