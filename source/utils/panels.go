package utils

import (
	"fmt"
	"path/filepath"
	"strings"
)

// PanelNames returns each of the current panel names.
func PanelNames(screenPackageFolderPath string) (panelNames []string, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("panelNames: %w", err)
		}
	}()

	var panelsFolderPath string = filepath.Join(screenPackageFolderPath, FolderNamePanels)
	var fileNames []string
	fileNames, err = FileNames(panelsFolderPath)
	panelNames = make([]string, 0, len(fileNames))
	for _, fileName := range fileNames {
		if strings.HasSuffix(fileName, panelFileSuffix) {
			panelName := strings.TrimSuffix(fileName, panelFileSuffix)
			panelNames = append(panelNames, panelName)
		}
	}
	return
}

// ValidateCurrentScreenPanelName validates that panel names are PascalCase and unique.
func ValidatePanelNames(
	panelNames []string,
) (isValid bool, failureMessage string) {

	failureMessages := make([]string, 0, 2*len(panelNames))
	defer func() {
		if len(failureMessages) > 0 {
			failureMessage = strings.Join(failureMessages, "\n")
			isValid = false
		} else {
			failureMessage = ""
			isValid = true
		}
	}()

	// Must be PascalCase.
	for _, panelName := range panelNames {
		// Valid. Not title case.
		if isValid, failureMessage = ValidatePascalCase(panelName, "panel"); !isValid {
			failureMessages = append(failureMessages, failureMessage)
		}
	}
	// Each panel name must be unique.
	last := len(panelNames) - 1
	for i, panelName := range panelNames {
		if i == last {
			break
		}
		for _, testName := range panelNames[i+1:] {
			if testName == panelName {
				failureMessage = fmt.Sprintf("Each panel name must be unique but %q is not.", panelName)
				failureMessages = append(failureMessages, failureMessage)
			}
		}
	}
	return
}
