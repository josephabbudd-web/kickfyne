package utils

import (
	"fmt"
	"slices"
	"strings"
)

func UniqueSplitAreas(areaNames []string) (areUnique bool, failureMessage string) {
	areUnique = true
	failureMessages := make([]string, 0, len(areaNames))
	defer func() {
		if !areUnique {
			failureMessage = strings.Join(failureMessages, "\n")
		}
	}()
	// Each name must be unique.
	last := len(areaNames) - 1
	for i, areaName := range areaNames {
		if i == last {
			break
		}
		if slices.Contains(areaNames[i+1:], areaName) {
			failureMessage = fmt.Sprintf("Each Split area name must be unique but %q is not.", areaName)
			failureMessages = append(failureMessages, failureMessage)
			areUnique = false
		}
	}
	return
}

func ValidSplitAreaItem(areaScreenName string) (isValid bool, area string, screen string, failureMessage string) {
	parts := strings.Split(areaScreenName, "=")
	switch len(parts) {
	case 1:
		switch areaScreenName {
		case "Leading", "Trailing":
			isValid = true
			area = areaScreenName
		default:
			failureMessage = fmt.Sprintf("Failure: %q is not an area. The only areas are Leading and Trailing.", areaScreenName)
		}
	case 2:
		switch parts[0] {
		case "Leading", "Trailing":
			area = parts[0]
			if parts[1][:1] != "*" {
				failureMessage = fmt.Sprintf("Failure: You must prefix the %q area's content source screen name with *.", area)
				return
			}
			screen = parts[1][1:]
			isValid, failureMessage = ValidatePascalCase(parts[1][1:], "source screen")
		default:
			failureMessage = fmt.Sprintf("%q is not an area. The only areas are Leading and Trailing.", parts[0])
		}
	}
	return
}
