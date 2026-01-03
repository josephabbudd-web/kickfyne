package doctabs

import (
	"fmt"
	"os"
	"path/filepath"

	_layout_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/doctabs/deps/layout"
	_tabs_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/doctabs/deps/tabItems"
	_layoutitems_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/doctabs/layoutTabItems"
	_misc_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/doctabs/misc"
	_panels_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/doctabs/panels"
	_panel_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/doctabs/panels/panel"
	_presetting_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/doctabs/presetting"

	_manifest_ "github.com/josephabbudd-web/kickfyne/manifest"
	_utils_ "github.com/josephabbudd-web/kickfyne/source/utils"
)

/*
	api.go
	panels/
	presetting/*
	layoutTabItems/*
	deps/tabs/
*/

// AddItems adds items to a DocTabs screen package.
func AddItems(
	packageName string,
	addItemNames []string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("doctabs.AddItems: %w", err)
		}
	}()

	manifest := _manifest_.NewAgain()
	var infoCopy *_manifest_.Info
	if infoCopy = manifest.InfoCopy(packageName); infoCopy == nil {
		return
	}
	infoCopy.AddItems(addItemNames...)
	manifestAllItemNames, manifestLocalItemNames, manifestRemoteItemNames := infoCopy.GetItems()
	addLocalItemNames := make([]string, 0, len(addItemNames))
	for _, itemName := range addItemNames {
		if itemName[:1] != "*" {
			addLocalItemNames = append(addLocalItemNames, itemName)
		}
	}

	// Create the folder paths in this package.

	// frontend/screens/«screen-package-name»
	packagePath := filepath.Join(folderPaths.FrontendScreens, packageName)

	// frontend/screens/«screen-package-name»/layoutTabItems
	packageLayoutItemsPath := filepath.Join(packagePath, _utils_.FolderNameLayoutTabItems)

	// frontend/screens/«screen-package-name»/deps
	packageDepsPath := filepath.Join(packagePath, _utils_.FolderNameDeps)

	// frontend/screens/«screen-package-name»/misc
	packageMiscPath := filepath.Join(packagePath, _utils_.FolderNameMisc)

	// frontend/screens/simple/«screen-package-name»/presets
	packagePresettingPath := filepath.Join(packagePath, _utils_.FolderNamePresetting)

	// frontend/screens/«screen-package-name»/panels
	packagePanelsPath := filepath.Join(packagePath, _utils_.FolderNamePanels)

	// frontend/screens/«screen-package-name»/deps/layout
	packageLayoutPath := filepath.Join(packageDepsPath, _utils_.FolderNameLayout)

	// frontend/screens/«screen-package-name»/deps/tabItems
	packageTabsPath := filepath.Join(packageDepsPath, _utils_.FolderNameTabItems)

	var fPath string
	var data any
	var fileName string
	funcs := _utils_.GetFuncs()

	// Add files to the package folder.

	// frontend/screens/«screen-package-name»/doc.go
	fPath = filepath.Join(packagePath, docFileName)
	files := files(packageName, manifestLocalItemNames, folderPaths)
	data = &docTemplateData{
		PackageName: packageName,
		Files:       files,
		Funcs:       funcs,
	}
	if err = _utils_.ProcessTemplate(docFileName, fPath, docTemplate, data); err != nil {
		return
	}

	// frontend/screens/«screen-package-name»/api.go
	fPath = filepath.Join(packagePath, aPIFileName)
	data = &aPITemplateData{
		PackageName:  packageName,
		ImportPrefix: importPrefix,
		Funcs:        funcs,
	}
	if err = _utils_.ProcessTemplate(aPIFileName, fPath, aPINoBETemplate, data); err != nil {
		return
	}

	// frontend/screens/«screen-package-name»/tabItemHandlers.go
	fPath = filepath.Join(packagePath, handlersFileName)
	data = &handlersTemplateData{
		PackageName:  packageName,
		ImportPrefix: importPrefix,
		Funcs:        funcs,
	}
	if err = _utils_.ProcessTemplate(handlersFileName, fPath, handlersTemplate, data); err != nil {
		return
	}

	// frontend/screens/«screen-package-name»/layoutTabItems folder.
	data = &_layoutitems_.TemplateData{
		PackageName:  packageName,
		ImportPrefix: importPrefix,
		AllItemNames: manifestAllItemNames,
	}
	fPath = filepath.Join(packageLayoutItemsPath, _utils_.DocFileName)
	if err = _utils_.ProcessTemplate(_utils_.DocFileName, fPath, _layoutitems_.DocsTemplate, data); err != nil {
		return
	}
	fPath = filepath.Join(packageLayoutItemsPath, _utils_.LayoutFileName)
	if err = _utils_.ProcessTemplate(_utils_.LayoutFileName, fPath, _layoutitems_.LayoutTemplate, data); err != nil {
		return
	}

	// frontend/screens/«screen-package-name»/misc/miscellaneous.go
	fPath = filepath.Join(packageMiscPath, _misc_.MiscellaneousFileName)
	data = &_misc_.MiscellaneousTemplateData{
		PackageName:  packageName,
		ImportPrefix: importPrefix,
	}
	if err = _utils_.ProcessTemplate(_misc_.MiscellaneousFileName, fPath, _misc_.MiscellaneousTemplate, data); err != nil {
		return
	}

	// frontend/screens/simple/«screen-package-name»/presetting folder.

	data = &_presetting_.TemplateData{
		PackageName:      packageName,
		ImportPrefix:     importPrefix,
		LocalPanelNames:  manifestLocalItemNames,
		RemotePanelNames: manifestRemoteItemNames,
		Funcs:            funcs,
	}
	fPath = filepath.Join(packagePresettingPath, _utils_.APIFileName)
	if err = _utils_.ProcessTemplate(_utils_.APIFileName, fPath, _presetting_.APITemplate, data); err != nil {
		return
	}
	fPath = filepath.Join(packagePresettingPath, _utils_.DefaultPresetFileName)
	if err = _utils_.ProcessTemplate(_utils_.DefaultPresetFileName, fPath, _presetting_.DefaultPresetTemplate, data); err != nil {
		return
	}
	fPath = filepath.Join(packagePresettingPath, _utils_.RemotePresetFileName)
	if len(manifestRemoteItemNames) > 0 {
		if err = _utils_.ProcessTemplate(_utils_.RemotePresetFileName, fPath, _presetting_.RemotePresetTemplate, data); err != nil {
			return
		}
	} else {
		if err = _utils_.WriteFile(fPath, []byte(_presetting_.NoRemotePresetTemplate)); err != nil {
			return
		}
	}

	// panels folder.

	// frontend/screens/«screen-package-name»/panels.log
	fPath = filepath.Join(packagePanelsPath, _panels_.LogFileName)
	content := _panels_.DocContent(packageName, manifestLocalItemNames, folderPaths)
	if err = _utils_.WriteFile(fPath, []byte(content)); err != nil {
		return
	}

	// Add each panel's file and sub folder.
	for _, panelName := range addLocalItemNames {
		// Panel file.
		// frontend/screens/«screen-package-name»/panels/«panel-name»Panel/
		fileName = panelName + _panels_.PanelFileNameSuffix
		fPath = filepath.Join(packagePanelsPath, fileName)
		data = &_panels_.PanelTemplateData{
			PackageName:  packageName,
			PanelName:    panelName,
			ImportPrefix: importPrefix,
			Funcs:        funcs,
		}
		if err = _utils_.ProcessTemplate(fileName, fPath, _panels_.PanelNoBETemplate, data); err != nil {
			return
		}

		// Panel sub folder holding content.go, state.go and preset.go.
		// frontend/screens/«screen-package-name»/panels/«panel-name»Panel/
		panelFolderName := panelName + "Panel"
		panelFolderPath := filepath.Join(packagePanelsPath, panelFolderName)
		if err = os.Mkdir(panelFolderPath, _utils_.DMode); err != nil {
			return
		}
		data = &_panel_.TemplateData{
			PackageName:     packageName,
			PanelName:       panelName,
			LocalPanelNames: manifestLocalItemNames,
			ImportPrefix:    importPrefix,
			Funcs:           funcs,
		}
		// content.go
		fileName = _panel_.ContentFileName
		fPath = filepath.Join(panelFolderPath, fileName)
		if err = _utils_.ProcessTemplate(fileName, fPath, _panel_.ContentTemplate, data); err != nil {
			return
		}
		// state.go
		fileName = _panel_.StateFileName
		fPath = filepath.Join(panelFolderPath, fileName)
		if err = _utils_.ProcessTemplate(fileName, fPath, _panel_.StateTemplate, data); err != nil {
			return
		}
		// preset.go
		fileName = _utils_.PresetFileName
		fPath = filepath.Join(panelFolderPath, fileName)
		if err = _utils_.ProcessTemplate(fileName, fPath, _panel_.PresetTemplate, data); err != nil {
			return
		}
	}

	// deps/ folder.

	// frontend/screens/«screen-package-name»/deps/layout/layout.go
	fPath = filepath.Join(packageLayoutPath, _layout_.LayoutFileName)
	data = &_layout_.LayoutTemplateData{
		PackageName:  packageName,
		ImportPrefix: importPrefix,
	}
	if err = _utils_.ProcessTemplate(_layout_.LayoutFileName, fPath, _layout_.LayoutTemplate, data); err != nil {
		return
	}

	// frontend/screens/«screen-package-name»/deps/tabItems/tabs.go
	fPath = filepath.Join(packageTabsPath, _tabs_.FileName)
	data = &_tabs_.TemplateData{
		PackageName:      packageName,
		ImportPrefix:     importPrefix,
		Funcs:            funcs,
		LocalPanelNames:  manifestLocalItemNames,
		RemotePanelNames: manifestRemoteItemNames,
	}
	if err = _utils_.ProcessTemplate(_tabs_.FileName, fPath, _tabs_.NoBETemplate, data); err != nil {
		return
	}

	return
}
