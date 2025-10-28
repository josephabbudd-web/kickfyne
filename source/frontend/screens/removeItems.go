package screens

import (
	_accordion_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/accordion"
	_apptabs_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/apptabs"
	_doctabs_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/doctabs"
	_simple_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/simple"
	_utils_ "github.com/josephabbudd-web/kickfyne/source/utils"
)

// RemoveItemsFromAccordionPackage builds a type Accordion screen package.
func RemoveItemsFromAccordionPackage(
	packageName string,
	newItemNames []string,
	packageDoc string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {
	return _accordion_.RemoveItems(
		packageName,
		newItemNames,
		packageDoc,
		importPrefix,
		folderPaths,
	)
}

// RemoveItemsFromAppTabsPackage builds a type AppTabs screen package.
func RemoveItemsFromAppTabsPackage(
	packageName string,
	newItemNames []string,
	packageDoc string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {
	return _apptabs_.RemoveItems(
		packageName,
		newItemNames,
		packageDoc,
		importPrefix,
		folderPaths,
	)
}

// RemoveItemsFromDocTabsPackage builds a type DocTabs screen package.
func RemoveItemsFromDocTabsPackage(
	packageName string,
	newItemNames []string,
	packageDoc string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {
	return _doctabs_.RemoveItems(
		packageName,
		newItemNames,
		packageDoc,
		importPrefix,
		folderPaths,
	)
}

// RemovePanelsFromSimplePackage builds a type Simple screen package.
func RemovePanelsFromSimplePackage(
	packageName string,
	panelNames []string,
	packageDoc string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {
	return _simple_.RemovePanels(
		packageName,
		panelNames,
		packageDoc,
		importPrefix,
		folderPaths,
	)
}
