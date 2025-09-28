package utils

import (
	"fmt"
	"strings"
)

// ValidateAccordionItemNames validates each new accordionItem name.
func ValidateAccordionItemNames(
	accordionItemNames []string,
) (isValid bool, failureMessage string) {

	failureMessages := make([]string, 0, 2*len(accordionItemNames))
	defer func() {
		if len(failureMessages) > 0 {
			failureMessage = strings.Join(failureMessages, "\n")
			isValid = false
		} else {
			failureMessage = ""
			isValid = true
		}
	}()

	var cleanAccordionItemNames = make([]string, 0, len(accordionItemNames))

	// Must be TitleCase.
	for _, accordionItemName := range accordionItemNames {
		var cleanAccordionItemName string
		if accordionItemName[:1] == "*" {
			cleanAccordionItemName = accordionItemName[1:]
		} else {
			cleanAccordionItemName = accordionItemName
		}
		// Valid. Not title case.
		if isValid, failureMessage = validateScreenAccordionItemName(cleanAccordionItemName); !isValid {
			failureMessages = append(failureMessages, failureMessage)
		}
		cleanAccordionItemNames = append(cleanAccordionItemNames, cleanAccordionItemName)
	}
	// Each accordionItem name must be unique.
	last := len(cleanAccordionItemNames) - 1
	for i, cleanAccordionItemName := range cleanAccordionItemNames {
		if i == last {
			break
		}
		for _, testName := range cleanAccordionItemNames[i+1:] {
			if testName == cleanAccordionItemName {
				failureMessage = fmt.Sprintf("Each accordionItem name must be unique but %q is not.", cleanAccordionItemName)
				failureMessages = append(failureMessages, failureMessage)
			}
		}
	}
	return
}
