package screens

import (
	"fmt"

	_manifest_ "github.com/josephabbudd-web/kickfyne/manifest"
	_accordion_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/accordion"
	_apptabs_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/apptabs"
	_border_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/border"
	_doctabs_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/doctabs"
	_simple_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/simple"
	_split_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/split"
	_utils_ "github.com/josephabbudd-web/kickfyne/source/utils"
)

// CreateFramework creates the HelloWorld simple screen so that the app will run.
func CreateFramework(
	manifest _manifest_.Manifest,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("frontend.CreateFramework: %w", err)
		}
	}()
	screenName := "HelloWorld"
	panelsNames := []string{"Hello", "HelloAgain"}
	if err = BuildSimplePackage(
		screenName,
		panelsNames,
		importPrefix,
		folderPaths,
	); err != nil {
		return
	}

	// Update the manifest.
	manifest.AddScreen(screenName, _manifest_.SimpleScreenInfoKind, panelsNames...)
	return
}

// BuildAccordionPackage builds a type Accordion screen package.
func BuildAccordionPackage(
	packageName string,
	rawPanelNames []string,
	allPanelNames, localPanelNames, remotePanelNames []string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {
	return _accordion_.Build(
		packageName,
		rawPanelNames,
		allPanelNames, localPanelNames, remotePanelNames,
		importPrefix,
		folderPaths,
	)
}

// BuildAppTabsPackage builds a type AppTabs screen package.
func BuildAppTabsPackage(
	packageName string,
	rawPanelNames []string,
	allPanelNames, localPanelNames, remotePanelNames []string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {
	return _apptabs_.Build(
		packageName,
		rawPanelNames,
		allPanelNames, localPanelNames, remotePanelNames,
		importPrefix,
		folderPaths,
	)
}

// BuildDocTabsPackage builds a type DocTabs screen package.
func BuildDocTabsPackage(
	packageName string,
	rawPanelNames []string,
	allPanelNames, localPanelNames, remotePanelNames []string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {
	return _doctabs_.Build(
		packageName,
		rawPanelNames,
		allPanelNames, localPanelNames, remotePanelNames,
		importPrefix,
		folderPaths,
	)
}

// BuildDocTabsPackage builds a type DocTabs screen package.
func BuildBorderPackage(
	packageName string,
	rawPanelNames []string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {
	return _border_.Build(
		packageName,
		rawPanelNames,
		importPrefix,
		folderPaths,
	)
}

// BuildSimplePackage builds a type Simple screen package.
func BuildSimplePackage(
	packageName string,
	panelNames []string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {
	return _simple_.Build(
		packageName,
		panelNames,
		importPrefix,
		folderPaths,
	)
}

// BuildDocTabsPackage builds a type DocTabs screen package.
func BuildSplitPackage(
	packageName string,
	rawPanelNames []string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {
	return _split_.Build(
		packageName,
		rawPanelNames,
		importPrefix,
		folderPaths,
	)
}
