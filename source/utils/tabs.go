package utils

import (
	"fmt"
	"strings"
)

// ValidateTabNames validates each new tab name.
func ValidateTabNames(
	tabNames []string,
) (isValid bool, failureMessage string) {

	failureMessages := make([]string, 0, 2*len(tabNames))
	defer func() {
		if len(failureMessages) > 0 {
			failureMessage = strings.Join(failureMessages, "\n")
			isValid = false
		} else {
			failureMessage = ""
			isValid = true
		}
	}()

	var cleanTabNames = make([]string, 0, len(tabNames))

	// Must be TitleCase.
	for _, tabName := range tabNames {
		var cleanTabName string
		if tabName[:1] == "*" {
			cleanTabName = tabName[1:]
		} else {
			cleanTabName = tabName
		}
		// Valid. Not title case.
		if isValid, failureMessage = validateScreenTabName(cleanTabName); !isValid {
			failureMessages = append(failureMessages, failureMessage)
		}
		cleanTabNames = append(cleanTabNames, cleanTabName)
	}
	// Each tab name must be unique.
	last := len(cleanTabNames) - 1
	for i, cleanTabName := range cleanTabNames {
		if i == last {
			break
		}
		for _, testName := range cleanTabNames[i+1:] {
			if testName == cleanTabName {
				failureMessage = fmt.Sprintf("Each tab name must be unique but %q is not.", cleanTabName)
				failureMessages = append(failureMessages, failureMessage)
			}
		}
	}
	return
}
