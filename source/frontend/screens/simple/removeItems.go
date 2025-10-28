package simple

import (
	"fmt"
	"os"
	"path/filepath"

	_layout_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/simple/deps/layout"
	_panelers_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/simple/deps/panelers"

	_panels_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/simple/panels"
	_presetting_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/simple/presetting"

	_manifest_ "github.com/josephabbudd-web/kickfyne/manifest"
	_utils_ "github.com/josephabbudd-web/kickfyne/source/utils"
)

/*
	api.go
	presetting/*
	panels/
	deps/layout/
*/

// Build builds a type Simple screen package.
func RemovePanels(
	packageName string,
	removePanelNames []string,
	packageDoc string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("simple.RemoveSimplePanels: %w", err)
		}
	}()

	// Create the folder paths in this package.

	// frontend/screens/simple/«screen-package-name»
	packagePath := filepath.Join(folderPaths.FrontendScreens, packageName)

	// frontend/screens/«screen-package-name»/deps
	packageDepsPath := filepath.Join(packagePath, _utils_.FolderNameDeps)

	// frontend/screens/simple/«screen-package-name»/presetting
	packagePresettingPath := filepath.Join(packagePath, _utils_.FolderNamePresetting)

	// frontend/screens/simple/«screen-package-name»/panels
	packagePanelsPath := filepath.Join(packagePath, _utils_.FolderNamePanels)

	// frontend/screens/simple/«screen-package-name»/deps/layout
	packageLayoutPath := filepath.Join(packageDepsPath, _utils_.FolderNameLayout)

	// frontend/screens/simple/«screen-package-name»/deps/panelers
	packagePanelersPath := filepath.Join(packageDepsPath, _utils_.FolderNamePanelers)

	var fPath string
	var data any
	var fileName string
	funcs := _utils_.GetFuncs()

	// Remove files to the package folder.

	// Put together all of the panel names.

	manifest := _manifest_.NewAgain()
	var infoCopy *_manifest_.Info
	if infoCopy = manifest.InfoCopy(packageName); infoCopy == nil {
		return
	}
	infoCopy.Remove(removePanelNames...)
	allPanelNames, _, _ := infoCopy.GetItems()

	// frontend/screens/simple/«screen-package-name»/api.go
	fPath = filepath.Join(packagePath, aPIFileName)
	data = &aPITemplateData{
		PackageName:      packageName,
		LocalPanelNames:  allPanelNames,
		DefaultPanelName: allPanelNames[0],
		ImportPrefix:     importPrefix,
		Funcs:            funcs,
	}
	if err = _utils_.ProcessTemplate(aPIFileName, fPath, aPINoBETemplate, data); err != nil {
		return
	}

	// frontend/screens/simple/«screen-package-name»/presetting folder.

	data = &_presetting_.TemplateData{
		PackageName:     packageName,
		ImportPrefix:    importPrefix,
		LocalPanelNames: allPanelNames,
		Funcs:           funcs,
	}
	fPath = filepath.Join(packagePresettingPath, _utils_.APIFileName)
	if err = _utils_.ProcessTemplate(_utils_.APIFileName, fPath, _presetting_.APITemplate, data); err != nil {
		return
	}
	fPath = filepath.Join(packagePresettingPath, _utils_.DefaultPresetFileName)
	if err = _utils_.ProcessTemplate(_utils_.DefaultPresetFileName, fPath, _presetting_.DefaultPresetTemplate, data); err != nil {
		return
	}

	// panels folder.

	// Remove each panel's file and sub folder.
	for _, panelName := range removePanelNames {
		// Panel file.
		// frontend/screens/«screen-package-name»/panels/«panel-name»Panel.go
		fileName = panelName + _panels_.PanelFileNameSuffix
		fPath = filepath.Join(packagePanelsPath, fileName)
		if err = os.Remove(fPath); err != nil {
			return
		}
		// frontend/screens/«screen-package-name»/panels/«panel-name»Panel/
		fPath = filepath.Join(packagePanelsPath, panelName)
		if err = os.RemoveAll(fPath); err != nil {
			return
		}
	}

	// deps/ folder.

	// frontend/screens/simple/«screen-package-name»/deps/layout/layout.go
	fPath = filepath.Join(packageLayoutPath, _layout_.LayoutFileName)
	data = &_layout_.LayoutTemplateData{
		PackageName:     packageName,
		ImportPrefix:    importPrefix,
		Funcs:           funcs,
		LocalPanelNames: allPanelNames,
	}
	if err = _utils_.ProcessTemplate(_layout_.LayoutFileName, fPath, _layout_.LayoutTemplate, data); err != nil {
		return
	}

	// frontend/screens/simple/«screen-package-name»/deps/panelers/panelers.go
	fPath = filepath.Join(packagePanelersPath, _panelers_.PanelersFileName)
	data = &_panelers_.PanelersTemplateData{
		ImportPrefix:    importPrefix,
		LocalPanelNames: allPanelNames,
	}
	if err = _utils_.ProcessTemplate(_panelers_.PanelersFileName, fPath, _panelers_.PanelersTemplate, data); err != nil {
		return
	}

	return
}
