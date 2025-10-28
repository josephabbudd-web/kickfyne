package frontend

import (
	"fmt"
	"strings"

	_screens_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens"
	_utils_ "github.com/josephabbudd-web/kickfyne/source/utils"
)

// handleScreenAddAccordion handles adding a Accordion screen package.
func handleScreenAddAccordion(
	screenPackageName string,
	panelNames []string,
	screenPackageDoc string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {

	var failureMessage string
	var successMessage string
	defer func() {
		if err != nil {
			err = fmt.Errorf("frontend.handleScreenAddAccordion: %w", err)
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
	allPanelNames, localPanelNames, remotePanelNames := separatePanelNames(panelNames)
	if err = _screens_.BuildAccordionPackage(
		screenPackageName,
		panelNames,
		allPanelNames, localPanelNames, remotePanelNames,
		screenPackageDoc,
		importPrefix,
		folderPaths,
	); err != nil {
		return
	}

	successMessage = successMessageScreenAddAccordion(screenPackageName, panelNames, folderPaths)
	return
}

// handleScreenAddAppTabs handles adding a AppTabs screen package.
func handleScreenAddAppTabs(
	screenPackageName string,
	panelNames []string,
	screenPackageDoc string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {

	var failureMessage string
	var successMessage string
	defer func() {
		if err != nil {
			err = fmt.Errorf("frontend.handleScreenAddAppTabs: %w", err)
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
	allPanelNames, localPanelNames, remotePanelNames := separatePanelNames(panelNames)
	if err = _screens_.BuildAppTabsPackage(
		screenPackageName,
		panelNames,
		allPanelNames, localPanelNames, remotePanelNames,
		screenPackageDoc,
		importPrefix,
		folderPaths,
	); err != nil {
		return
	}

	successMessage = successMessageScreenAddAppTabs(screenPackageName, panelNames, folderPaths)
	return
}

// handleScreenAddDocTabs handles adding a DocTabs screen package.
func handleScreenAddDocTabs(
	screenPackageName string,
	panelNames []string,
	screenPackageDoc string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {

	var failureMessage string
	var successMessage string
	defer func() {
		if err != nil {
			err = fmt.Errorf("frontend.handleScreenAddDocTabs: %w", err)
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
	allPanelNames, localPanelNames, remotePanelNames := separatePanelNames(panelNames)
	if err = _screens_.BuildDocTabsPackage(
		screenPackageName,
		panelNames,
		allPanelNames, localPanelNames, remotePanelNames,
		screenPackageDoc,
		importPrefix,
		folderPaths,
	); err != nil {
		return
	}

	successMessage = successMessageScreenAddDocTabs(screenPackageName, panelNames, folderPaths)
	return
}

// handleScreenAddSimple handles adding a simple screen package.
func handleScreenAddSimple(
	screenPackageName string,
	panelNames []string,
	screenPackageDoc string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {

	var failureMessage string
	var successMessage string
	defer func() {
		if err != nil {
			err = fmt.Errorf("frontend.handleScreenAdd: %w", err)
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
	if err = _screens_.BuildSimplePackage(
		screenPackageName,
		panelNames,
		screenPackageDoc,
		importPrefix,
		folderPaths,
	); err != nil {
		return
	}
	successMessage = successMessageScreenAddSimple(screenPackageName, panelNames, folderPaths)
	return
}

func successMessageScreenAddSimple(
	screenPackageName string,
	localPanelNames []string,
	folderPaths *_utils_.FolderPaths,
) (successMessage string) {
	docFilePath := _utils_.ScreenDocFilePath(screenPackageName, folderPaths)
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("Added the Simple screen package named %q.\n", screenPackageName))
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

func successMessageScreenAddAppTabs(
	screenPackageName string,
	localPanelNames []string,
	folderPaths *_utils_.FolderPaths,
) (successMessage string) {
	docFilePath := _utils_.ScreenDocFilePath(screenPackageName, folderPaths)
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("Added the AppTabs screen package named %q.\n", screenPackageName))
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

func successMessageScreenAddDocTabs(
	screenPackageName string,
	localPanelNames []string,
	folderPaths *_utils_.FolderPaths,
) (successMessage string) {
	docFilePath := _utils_.ScreenDocFilePath(screenPackageName, folderPaths)
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("Added the DocTabs screen package named %q.\n", screenPackageName))
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

func successMessageScreenAddAccordion(
	screenPackageName string,
	localPanelNames []string,
	folderPaths *_utils_.FolderPaths,
) (successMessage string) {
	docFilePath := _utils_.ScreenDocFilePath(screenPackageName, folderPaths)
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("Added the Accordion screen package named %q.\n", screenPackageName))
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
