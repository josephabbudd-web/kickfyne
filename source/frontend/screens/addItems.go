package screens

import (
	_accordion_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/accordion"
	_apptabs_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/apptabs"
	_border_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/border"
	_doctabs_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/doctabs"
	_simple_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/simple"
	_utils_ "github.com/josephabbudd-web/kickfyne/source/utils"
)

// AddItemsToAccordionPackage builds a type Accordion screen package.
func AddItemsToAccordionPackage(
	packageName string,
	newItemNames []string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {
	return _accordion_.AddItems(
		packageName,
		newItemNames,
		importPrefix,
		folderPaths,
	)
}

// AddItemsToAppTabsPackage builds a type AppTabs screen package.
func AddItemsToAppTabsPackage(
	packageName string,
	newItemNames []string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {
	return _apptabs_.AddItems(
		packageName,
		newItemNames,
		importPrefix,
		folderPaths,
	)
}

// AddItemsToDocTabsPackage builds a type DocTabs screen package.
func AddItemsToDocTabsPackage(
	packageName string,
	newItemNames []string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {
	return _doctabs_.AddItems(
		packageName,
		newItemNames,
		importPrefix,
		folderPaths,
	)
}

// AddPanelsToSimplePackage builds a type Simple screen package.
func AddPanelsToSimplePackage(
	packageName string,
	panelNames []string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {
	return _simple_.AddPanels(
		packageName,
		panelNames,
		importPrefix,
		folderPaths,
	)
}

// AddAreasToBorderPackage builds a type Simple screen package.
func AddAreasToBorderPackage(
	packageName string,
	areaNames []string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {
	return _border_.AddAreas(
		packageName,
		areaNames,
		importPrefix,
		folderPaths,
	)
}
