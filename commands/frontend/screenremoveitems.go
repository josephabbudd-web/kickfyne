package frontend

import (
	"fmt"
	"strings"

	_screens_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens"
	_utils_ "github.com/josephabbudd-web/kickfyne/source/utils"
)

// handleScreenRemoveAccordionItems handles removing of TabItems/Panels from a Accordion screen package.
func handleScreenRemoveAccordionItems(
	screenPackageName string,
	itemNames []string,
	screenPackageApp string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {

	var failureMessage string
	var successMessage string
	defer func() {
		if err != nil {
			err = fmt.Errorf("frontend.handleScreenRemoveAccordionItems: %w", err)
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
	if err = _screens_.RemoveItemsFromAccordionPackage(
		screenPackageName,
		itemNames,
		screenPackageApp,
		importPrefix,
		folderPaths,
	); err != nil {
		return
	}

	successMessage = successMessageScreenRemoveAccordionItems(screenPackageName, itemNames, folderPaths)
	return
}

// handleScreenRemoveAppTabsItems handles removing of TabItems/Panels from a AppTabs screen package.
func handleScreenRemoveAppTabsItems(
	screenPackageName string,
	itemNames []string,
	screenPackageApp string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {

	var failureMessage string
	var successMessage string
	defer func() {
		if err != nil {
			err = fmt.Errorf("frontend.handleScreenRemoveAppTabsItems: %w", err)
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
	if err = _screens_.RemoveItemsFromAppTabsPackage(
		screenPackageName,
		itemNames,
		screenPackageApp,
		importPrefix,
		folderPaths,
	); err != nil {
		return
	}

	successMessage = successMessageScreenRemoveAppTabsItems(screenPackageName, itemNames, folderPaths)
	return
}

// handleScreenRemoveDocTabsItems handles removing of TabItems/Panels from a DocTabs screen package.
func handleScreenRemoveDocTabsItems(
	screenPackageName string,
	itemNames []string,
	screenPackageDoc string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {

	var failureMessage string
	var successMessage string
	defer func() {
		if err != nil {
			err = fmt.Errorf("frontend.handleScreenRemoveDocTabsItems: %w", err)
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
	if err = _screens_.RemoveItemsFromDocTabsPackage(
		screenPackageName,
		itemNames,
		screenPackageDoc,
		importPrefix,
		folderPaths,
	); err != nil {
		return
	}

	successMessage = successMessageScreenRemoveDocTabsItems(screenPackageName, itemNames, folderPaths)
	return
}

// handleScreenRemoveSimplePanels handles removing panels from a Simple screen package.
func handleScreenRemoveSimplePanels(
	screenPackageName string,
	panelNames []string,
	screenPackageDoc,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {

	var failureMessage string
	var successMessage string
	defer func() {
		if err != nil {
			err = fmt.Errorf("frontend.handleScreenRemoveSimplePanels: %w", err)
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
	if err = _screens_.RemovePanelsFromSimplePackage(
		screenPackageName,
		panelNames,
		screenPackageDoc,
		importPrefix,
		folderPaths,
	); err != nil {
		return
	}
	successMessage = successMessageScreenRemoveSimplePanels(screenPackageName, panelNames, folderPaths)
	return
}

func successMessageScreenRemoveAppTabsItems(
	screenPackageName string,
	itemNames []string,
	folderPaths *_utils_.FolderPaths,
) (successMessage string) {
	docFilePath := _utils_.ScreenDocFilePath(screenPackageName, folderPaths)
	builder := strings.Builder{}
	for _, itemName := range itemNames {
		builder.WriteString(fmt.Sprintf("Removed the TabItem %q from the AppTabs screen package named %q.\n", itemName, screenPackageName))
	}
	builder.WriteString("The package's docs may require editing.\n")
	builder.WriteString("Package docs: " + _utils_.Clickable(docFilePath) + "\n")
	presettingAPIFildPath := _utils_.ScreenPresettingAPIFilePath(screenPackageName, folderPaths)
	builder.WriteString("The package's presettings require editing.\n")
	builder.WriteString("Package presettings api.go: " + _utils_.Clickable(presettingAPIFildPath) + "\n")
	successMessage = builder.String()
	return
}

func successMessageScreenRemoveDocTabsItems(
	screenPackageName string,
	itemNames []string,
	folderPaths *_utils_.FolderPaths,
) (successMessage string) {
	docFilePath := _utils_.ScreenDocFilePath(screenPackageName, folderPaths)
	builder := strings.Builder{}
	for _, itemName := range itemNames {
		builder.WriteString(fmt.Sprintf("Removed the TabItem %q from the DocTabs screen package named %q.\n", itemName, screenPackageName))
	}
	builder.WriteString("The package's docs may require editing.\n")
	builder.WriteString("Package docs: " + _utils_.Clickable(docFilePath) + "\n")
	presettingAPIFildPath := _utils_.ScreenPresettingAPIFilePath(screenPackageName, folderPaths)
	builder.WriteString("The package's presettings require editing.\n")
	builder.WriteString("Package presettings api.go: " + _utils_.Clickable(presettingAPIFildPath) + "\n")
	successMessage = builder.String()
	return
}

func successMessageScreenRemoveAccordionItems(
	screenPackageName string,
	itemNames []string,
	folderPaths *_utils_.FolderPaths,
) (successMessage string) {
	docFilePath := _utils_.ScreenDocFilePath(screenPackageName, folderPaths)
	builder := strings.Builder{}
	for _, itemName := range itemNames {
		builder.WriteString(fmt.Sprintf("Removed the AccordionItem %q from the Accordion screen package named %q.\n", itemName, screenPackageName))
	}
	builder.WriteString("The package's docs may require editing.\n")
	builder.WriteString("Package docs: " + _utils_.Clickable(docFilePath) + "\n")
	presettingAPIFildPath := _utils_.ScreenPresettingAPIFilePath(screenPackageName, folderPaths)
	builder.WriteString("The package's presettings require editing.\n")
	builder.WriteString("Package presettings api.go: " + _utils_.Clickable(presettingAPIFildPath) + "\n")
	successMessage = builder.String()
	return
}

func successMessageScreenRemoveSimplePanels(
	screenPackageName string,
	localPanelNames []string,
	folderPaths *_utils_.FolderPaths,
) (successMessage string) {
	docFilePath := _utils_.ScreenDocFilePath(screenPackageName, folderPaths)
	builder := strings.Builder{}
	for _, panelName := range localPanelNames {
		builder.WriteString(fmt.Sprintf("Removed the Panel %q from the Simple screen package named %q.\n", panelName, screenPackageName))
	}
	builder.WriteString("The package's docs may require editing.\n")
	builder.WriteString("Package docs: " + _utils_.Clickable(docFilePath) + "\n")
	presettingAPIFildPath := _utils_.ScreenPresettingAPIFilePath(screenPackageName, folderPaths)
	builder.WriteString("The package's presettings require editing.\n")
	builder.WriteString("Package presettings api.go: " + _utils_.Clickable(presettingAPIFildPath) + "\n")
	successMessage = builder.String()
	return
}
