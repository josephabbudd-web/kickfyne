package frontend

import (
	_screens_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens"
	_utils_ "github.com/josephabbudd-web/kickfyne/source/utils"
)

// handleScreenRemoveAccordionItems handles removing of TabItems/Panels from a Accordion screen package.
func handleScreenRemoveAccordionItems(
	screenPackageName string,
	itemNames []string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {
	err = _screens_.RemoveItemsFromAccordionPackage(
		screenPackageName,
		itemNames,
		importPrefix,
		folderPaths,
	)
	return
}

// handleScreenRemoveAppTabsItems handles removing of TabItems/Panels from a AppTabs screen package.
func handleScreenRemoveAppTabsItems(
	screenPackageName string,
	itemNames []string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {
	err = _screens_.RemoveItemsFromAppTabsPackage(
		screenPackageName,
		itemNames,
		importPrefix,
		folderPaths,
	)
	return
}

// handleScreenRemoveDocTabsItems handles removing of TabItems/Panels from a DocTabs screen package.
func handleScreenRemoveDocTabsItems(
	screenPackageName string,
	itemNames []string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {
	err = _screens_.RemoveItemsFromDocTabsPackage(
		screenPackageName,
		itemNames,
		importPrefix,
		folderPaths,
	)
	return
}

// handleScreenRemoveSimplePanels handles removing panels from a Simple screen package.
func handleScreenRemoveSimplePanels(
	screenPackageName string,
	panelNames []string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {
	err = _screens_.RemovePanelsFromSimplePackage(
		screenPackageName,
		panelNames,
		importPrefix,
		folderPaths,
	)
	return
}

// handleScreenRemoveBorderAreas handles removing panels from a Simple screen package.
func handleScreenRemoveBorderAreas(
	screenPackageName string,
	areaNames []string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {
	err = _screens_.RemoveAreasFromBorderPackage(
		screenPackageName,
		areaNames,
		importPrefix,
		folderPaths,
	)
	return
}
