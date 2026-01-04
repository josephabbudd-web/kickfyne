package version

import (
	"fmt"
	"os"
)

const (
	versionAPINewBreaking  = 0
	versionAPIAddedFeature = 9
	versionAPIBugFix       = 0
	versionDate            = "Jan 3, 2026"
)

// V returns the version.
func V() (version string) {
	version = fmt.Sprintf("This is the %s version %d.%d.%d. Published %s", os.Args[0], versionAPINewBreaking, versionAPIAddedFeature, versionAPIBugFix, versionDate)
	return
}
