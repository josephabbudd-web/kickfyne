package frontend

import (
	_screens_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens"
	_utils_ "github.com/josephabbudd-web/kickfyne/source/utils"
)

// handleScreenAddAccordion handles adding a Accordion screen package.
func handleScreenAddAccordionItems(
	screenPackageName string,
	newItemNames []string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {
	err = _screens_.AddItemsToAccordionPackage(
		screenPackageName,
		newItemNames,
		importPrefix,
		folderPaths,
	)
	return
}

// handleScreenAddAppTabs handles adding a AppTabs screen package.
func handleScreenAddAppTabsItems(
	screenPackageName string,
	newItemNames []string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {
	err = _screens_.AddItemsToAppTabsPackage(
		screenPackageName,
		newItemNames,
		importPrefix,
		folderPaths,
	)
	return
}

// handleScreenAddDocTabsItems handles adding a DocTabs screen package.
func handleScreenAddDocTabsItems(
	screenPackageName string,
	newItemNames []string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {
	err = _screens_.AddItemsToDocTabsPackage(
		screenPackageName,
		newItemNames,
		importPrefix,
		folderPaths,
	)
	return
}

// handleScreenAddSimplePanels handles adding panels to a simple screen package.
func handleScreenAddSimplePanels(
	screenPackageName string,
	newPanelNames []string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {
	err = _screens_.AddPanelsToSimplePackage(
		screenPackageName,
		newPanelNames,
		importPrefix,
		folderPaths,
	)
	return
}

// handleScreenAddBorderAreas handles adding areas to a border screen package.
func handleScreenAddBorderAreas(
	screenPackageName string,
	newAreaNames []string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {
	err = _screens_.AddAreasToBorderPackage(
		screenPackageName,
		newAreaNames,
		importPrefix,
		folderPaths,
	)
	return
}
