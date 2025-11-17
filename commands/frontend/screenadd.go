package frontend

import (
	_screens_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens"
	_utils_ "github.com/josephabbudd-web/kickfyne/source/utils"
)

// handleScreenAddAccordion handles adding a Accordion screen package.
func handleScreenAddAccordion(
	screenPackageName string,
	panelNames []string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {
	allPanelNames, localPanelNames, remotePanelNames := separatePanelNames(panelNames)
	err = _screens_.BuildAccordionPackage(
		screenPackageName,
		panelNames,
		allPanelNames, localPanelNames, remotePanelNames,
		importPrefix,
		folderPaths,
	)
	return
}

// handleScreenAddAppTabs handles adding a AppTabs screen package.
func handleScreenAddAppTabs(
	screenPackageName string,
	panelNames []string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {
	allPanelNames, localPanelNames, remotePanelNames := separatePanelNames(panelNames)
	err = _screens_.BuildAppTabsPackage(
		screenPackageName,
		panelNames,
		allPanelNames, localPanelNames, remotePanelNames,
		importPrefix,
		folderPaths,
	)
	return
}

// handleScreenAddDocTabs handles adding a DocTabs screen package.
func handleScreenAddDocTabs(
	screenPackageName string,
	panelNames []string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {
	allPanelNames, localPanelNames, remotePanelNames := separatePanelNames(panelNames)
	err = _screens_.BuildDocTabsPackage(
		screenPackageName,
		panelNames,
		allPanelNames, localPanelNames, remotePanelNames,
		importPrefix,
		folderPaths,
	)
	return
}

// handleScreenAddDocTabs handles adding a DocTabs screen package.
func handleScreenAddBorder(
	screenPackageName string,
	rawPanelNames []string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {
	err = _screens_.BuildBorderPackage(
		screenPackageName,
		rawPanelNames,
		importPrefix,
		folderPaths,
	)
	return
}

// HandleScreenAddSimple handles adding a simple screen package.
func HandleScreenAddSimple(
	screenPackageName string,
	panelNames []string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {
	err = _screens_.BuildSimplePackage(
		screenPackageName,
		panelNames,
		importPrefix,
		folderPaths,
	)
	return
}
