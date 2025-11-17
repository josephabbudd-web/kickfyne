package utils

import (
	"fmt"
	"strings"
)

func ValidBorderAreaItem(areaItemName string) (isValid bool, area string, screen string, failureMessage string) {
	parts := strings.Split(areaItemName, "=")
	switch len(parts) {
	case 1:
		switch areaItemName {
		case "Top", "Bottom", "Left", "Right", "Center":
			isValid = true
			area = areaItemName
		default:
			failureMessage = fmt.Sprintf("Failure: %q is not an area. The only areas are Top, Bottom, Left, Right, Center.", areaItemName)
		}
	case 2:
		switch parts[0] {
		case "Top", "Bottom", "Left", "Right":
			failureMessage = fmt.Sprintf("Failure: The %[1]q area always uses the screen's own %[1]q panel for content.\n         Only the Center area may use another screen for content.", parts[0])
			return
		case "Center":
			if parts[1][:1] != "*" {
				failureMessage = fmt.Sprintf("Failure: You must prefix the %q area's source screen name with *.", parts[0])
				return
			}
			area = parts[0]
			screen = parts[1][1:]
			isValid, failureMessage = ValidatePascalCase(screen, "borderAreaItem")
		default:
			failureMessage = fmt.Sprintf("%q is not an area. The only areas are Top, Bottom, Left, Right, Center.", parts[0])
		}
	}
	return
}
