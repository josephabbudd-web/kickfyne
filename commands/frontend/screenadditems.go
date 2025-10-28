package frontend

import (
	"fmt"
	"strings"

	_screens_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens"
	_utils_ "github.com/josephabbudd-web/kickfyne/source/utils"
)

// handleScreenAddAccordion handles adding a Accordion screen package.
func handleScreenAddAccordionItems(
	screenPackageName string,
	newItemNames []string,
	screenPackageDoc string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {

	var failureMessage string
	var successMessage string
	defer func() {
		if err != nil {
			err = fmt.Errorf("frontend.handleScreenAddAccordionItems: %w", err)
			return
		}
		switch {
		case len(failureMessage) > 0:
			fmt.Println("Failure:")
			fmt.Println(failureMessage)
		case len(successMessage) > 0:
			fmt.Println("Success:")
			fmt.Println(successMessage)
		}
	}()

	// Add the screen package.
	if err = _screens_.AddItemsToAccordionPackage(
		screenPackageName,
		newItemNames,
		screenPackageDoc,
		importPrefix,
		folderPaths,
	); err != nil {
		return
	}

	successMessage = successMessageScreenAddAccordionItems(screenPackageName, newItemNames, folderPaths)
	return
}

// handleScreenAddAppTabs handles adding a AppTabs screen package.
func handleScreenAddAppTabsItems(
	screenPackageName string,
	newItemNames []string,
	screenPackageDoc string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {

	var failureMessage string
	var successMessage string
	defer func() {
		if err != nil {
			err = fmt.Errorf("frontend.handleScreenAddAppTabsItems: %w", err)
			return
		}
		switch {
		case len(failureMessage) > 0:
			fmt.Println("Failure:")
			fmt.Println(failureMessage)
		case len(successMessage) > 0:
			fmt.Println("Success:")
			fmt.Println(successMessage)
		}
	}()

	// Add the screen package.
	if err = _screens_.AddItemsToAppTabsPackage(
		screenPackageName,
		newItemNames,
		screenPackageDoc,
		importPrefix,
		folderPaths,
	); err != nil {
		return
	}

	successMessage = successMessageScreenAddAppTabsItems(screenPackageName, newItemNames, folderPaths)
	return
}

// handleScreenAddDocTabsItems handles adding a DocTabs screen package.
func handleScreenAddDocTabsItems(
	screenPackageName string,
	newItemNames []string,
	screenPackageDoc string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {

	var failureMessage string
	var successMessage string
	defer func() {
		if err != nil {
			err = fmt.Errorf("frontend.handleScreenAddDocTabsItems: %w", err)
			return
		}
		switch {
		case len(failureMessage) > 0:
			fmt.Println("Failure:")
			fmt.Println(failureMessage)
		case len(successMessage) > 0:
			fmt.Println("Success:")
			fmt.Println(successMessage)
		}
	}()

	// Add the screen package.
	if err = _screens_.AddItemsToDocTabsPackage(
		screenPackageName,
		newItemNames,
		screenPackageDoc,
		importPrefix,
		folderPaths,
	); err != nil {
		return
	}

	successMessage = successMessageScreenAddDocTabsItems(screenPackageName, newItemNames, folderPaths)
	return
}

// handleScreenAddSimplePanels handles adding panels to a simple screen package.
func handleScreenAddSimplePanels(
	screenPackageName string,
	newPanelNames []string,
	screenPackageDoc,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {

	var failureMessage string
	var successMessage string
	defer func() {
		if err != nil {
			err = fmt.Errorf("frontend.handleScreenAddSimplePanels: %w", err)
			return
		}
		switch {
		case len(failureMessage) > 0:
			fmt.Println("Failure:")
			fmt.Println(failureMessage)
		case len(successMessage) > 0:
			fmt.Println("Success:")
			fmt.Println(successMessage)
		}
	}()

	// Create the package folder with no panels.
	if err = _screens_.AddPanelsToSimplePackage(
		screenPackageName,
		newPanelNames,
		screenPackageDoc,
		importPrefix,
		folderPaths,
	); err != nil {
		return
	}
	successMessage = successMessageScreenAddSimplePanels(screenPackageName, newPanelNames, folderPaths)
	return
}

func successMessageScreenAddAppTabsItems(
	screenPackageName string,
	localPanelNames []string,
	folderPaths *_utils_.FolderPaths,
) (successMessage string) {
	docFilePath := _utils_.ScreenDocFilePath(screenPackageName, folderPaths)
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("Added the new TabItems to the AppTabs screen package named %q.\n", screenPackageName))
	if len(localPanelNames) == 0 {
		builder.WriteString("The package's docs may require editing.\n")
	} else {
		builder.WriteString("The package's docs and each local panel's content.go, state.go and preset.go require editing.\n")
	}
	builder.WriteString("Package docs: " + _utils_.Clickable(docFilePath) + "\n")
	for _, panelName := range localPanelNames {
		contentPath := _utils_.PanelContentFilePath(screenPackageName, panelName, folderPaths)
		statePath := _utils_.PanelStateFilePath(screenPackageName, panelName, folderPaths)
		presetPath := _utils_.PanelPresetFilePath(screenPackageName, panelName, folderPaths)
		builder.WriteString(panelName + "Panel\n")
		builder.WriteString(fmt.Sprintf("  Content:   %s.\n", _utils_.Clickable(contentPath)))
		builder.WriteString(fmt.Sprintf("  State:     %s.\n", _utils_.Clickable(statePath)))
		builder.WriteString(fmt.Sprintf("  Preset:    %s.\n", _utils_.Clickable(presetPath)))
	}
	presettingAPIFildPath := _utils_.ScreenPresettingAPIFilePath(screenPackageName, folderPaths)
	builder.WriteString("The package's presettings require editing.\n")
	builder.WriteString("Package presettings api.go: " + _utils_.Clickable(presettingAPIFildPath) + "\n")
	successMessage = builder.String()
	return
}

func successMessageScreenAddDocTabsItems(
	screenPackageName string,
	localPanelNames []string,
	folderPaths *_utils_.FolderPaths,
) (successMessage string) {
	docFilePath := _utils_.ScreenDocFilePath(screenPackageName, folderPaths)
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("Added the new TabItems to the DocTabs screen package named %q.\n", screenPackageName))
	if len(localPanelNames) == 0 {
		builder.WriteString("The package's docs may require editing.\n")
	} else {
		builder.WriteString("The package's docs and each local panel's content.go, state.go and preset.go require editing.\n")
	}
	builder.WriteString("Package docs: " + _utils_.Clickable(docFilePath) + "\n")
	for _, panelName := range localPanelNames {
		contentPath := _utils_.PanelContentFilePath(screenPackageName, panelName, folderPaths)
		statePath := _utils_.PanelStateFilePath(screenPackageName, panelName, folderPaths)
		presetPath := _utils_.PanelPresetFilePath(screenPackageName, panelName, folderPaths)
		builder.WriteString(panelName + "Panel\n")
		builder.WriteString(fmt.Sprintf("  Content:   %s.\n", _utils_.Clickable(contentPath)))
		builder.WriteString(fmt.Sprintf("  State:     %s.\n", _utils_.Clickable(statePath)))
		builder.WriteString(fmt.Sprintf("  Preset:    %s.\n", _utils_.Clickable(presetPath)))
	}
	presettingAPIFildPath := _utils_.ScreenPresettingAPIFilePath(screenPackageName, folderPaths)
	builder.WriteString("The package's presettings require editing.\n")
	builder.WriteString("Package presettings api.go: " + _utils_.Clickable(presettingAPIFildPath) + "\n")
	successMessage = builder.String()
	return
}

func successMessageScreenAddAccordionItems(
	screenPackageName string,
	localPanelNames []string,
	folderPaths *_utils_.FolderPaths,
) (successMessage string) {
	docFilePath := _utils_.ScreenDocFilePath(screenPackageName, folderPaths)
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("Added the AccordionItems to the Accordion screen package named %q.\n", screenPackageName))
	if len(localPanelNames) == 0 {
		builder.WriteString("The package's docs may require editing.\n")
	} else {
		builder.WriteString("The package's docs and each local panel's content.go, state.go and preset.go require editing.\n")
	}
	builder.WriteString("Package docs: " + _utils_.Clickable(docFilePath) + "\n")
	for _, panelName := range localPanelNames {
		contentPath := _utils_.PanelContentFilePath(screenPackageName, panelName, folderPaths)
		statePath := _utils_.PanelStateFilePath(screenPackageName, panelName, folderPaths)
		presetPath := _utils_.PanelPresetFilePath(screenPackageName, panelName, folderPaths)
		builder.WriteString(panelName + "Panel\n")
		builder.WriteString(fmt.Sprintf("  Content:   %s.\n", _utils_.Clickable(contentPath)))
		builder.WriteString(fmt.Sprintf("  State:     %s.\n", _utils_.Clickable(statePath)))
		builder.WriteString(fmt.Sprintf("  Preset:    %s.\n", _utils_.Clickable(presetPath)))
	}
	presettingAPIFildPath := _utils_.ScreenPresettingAPIFilePath(screenPackageName, folderPaths)
	builder.WriteString("The package's presettings require editing.\n")
	builder.WriteString("Package presettings api.go: " + _utils_.Clickable(presettingAPIFildPath) + "\n")
	successMessage = builder.String()
	return
}

func successMessageScreenAddSimplePanels(
	screenPackageName string,
	localPanelNames []string,
	folderPaths *_utils_.FolderPaths,
) (successMessage string) {
	docFilePath := _utils_.ScreenDocFilePath(screenPackageName, folderPaths)
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("Added the new panels to the Simple screen package named %q.\n", screenPackageName))
	builder.WriteString("The package's docs and each panel's content.go, state.go and preset.go require editing.\n")
	builder.WriteString("Package docs: " + _utils_.Clickable(docFilePath) + "\n")
	for _, panelName := range localPanelNames {
		contentPath := _utils_.PanelContentFilePath(screenPackageName, panelName, folderPaths)
		statePath := _utils_.PanelStateFilePath(screenPackageName, panelName, folderPaths)
		presetPath := _utils_.PanelPresetFilePath(screenPackageName, panelName, folderPaths)
		builder.WriteString(panelName + "Panel\n")
		builder.WriteString(fmt.Sprintf("  Content:   %s.\n", _utils_.Clickable(contentPath)))
		builder.WriteString(fmt.Sprintf("  State:     %s.\n", _utils_.Clickable(statePath)))
		builder.WriteString(fmt.Sprintf("  Preset:    %s.\n", _utils_.Clickable(presetPath)))
	}
	presettingAPIFildPath := _utils_.ScreenPresettingAPIFilePath(screenPackageName, folderPaths)
	builder.WriteString("The package's presettings require editing.\n")
	builder.WriteString("Package presettings api.go: " + _utils_.Clickable(presettingAPIFildPath) + "\n")
	successMessage = builder.String()
	return
}
