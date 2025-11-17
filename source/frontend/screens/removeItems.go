package screens

import (
	_accordion_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/accordion"
	_apptabs_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/apptabs"
	_border_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/border"
	_doctabs_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/doctabs"
	_simple_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/simple"
	_utils_ "github.com/josephabbudd-web/kickfyne/source/utils"
)

// RemoveItemsFromAccordionPackage _accordion_.RemoveItems.
func RemoveItemsFromAccordionPackage(
	packageName string,
	newItemNames []string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {
	return _accordion_.RemoveItems(
		packageName,
		newItemNames,
		importPrefix,
		folderPaths,
	)
}

// RemoveItemsFromAppTabsPackage calls _apptabs_.RemoveItems.
func RemoveItemsFromAppTabsPackage(
	packageName string,
	newItemNames []string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {
	return _apptabs_.RemoveItems(
		packageName,
		newItemNames,
		importPrefix,
		folderPaths,
	)
}

// RemoveItemsFromDocTabsPackage calls _doctabs_.RemoveItems.
func RemoveItemsFromDocTabsPackage(
	packageName string,
	newItemNames []string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {
	return _doctabs_.RemoveItems(
		packageName,
		newItemNames,
		importPrefix,
		folderPaths,
	)
}

// RemovePanelsFromSimplePackage calls _simple_.RemovePanels.
func RemovePanelsFromSimplePackage(
	packageName string,
	panelNames []string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {
	return _simple_.RemovePanels(
		packageName,
		panelNames,
		importPrefix,
		folderPaths,
	)
}

// RemoveAreasFromBorderPackage calls _border_.RemoveAreas.
func RemoveAreasFromBorderPackage(
	packageName string,
	areaNames []string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {
	return _border_.RemoveAreas(
		packageName,
		areaNames,
		importPrefix,
		folderPaths,
	)
}
