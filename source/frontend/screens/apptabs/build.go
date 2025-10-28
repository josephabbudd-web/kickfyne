package apptabs

import (
	"fmt"
	"os"
	"path/filepath"

	_layout_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/apptabs/deps/layout"
	_producer_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/apptabs/deps/producer"
	_tabs_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/apptabs/deps/tabItems"
	_layoutitems_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/apptabs/layoutTabItems"

	_misc_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/apptabs/misc"
	_panels_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/apptabs/panels"
	_panel_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/apptabs/panels/panel"
	_presetting_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/apptabs/presetting"

	_utils_ "github.com/josephabbudd-web/kickfyne/source/utils"
)

// Build builds a type Simple screen package.
func Build(
	packageName string,
	rawPanelNames []string,
	allPanelNames, localPanelNames, remotePanelNames []string,
	packageDoc string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("doctabs.Build: %w", err)
		}
	}()

	// Create the folder paths in this package.

	// frontend/screens/«screen-package-name»
	packagePath := filepath.Join(folderPaths.FrontendScreens, packageName)
	if err = os.Mkdir(packagePath, _utils_.DMode); err != nil {
		return
	}

	// frontend/screens/«screen-package-name»/layoutTabItems
	packageLayoutItemsPath := filepath.Join(packagePath, _utils_.FolderNameLayoutTabItems)
	if err = os.Mkdir(packageLayoutItemsPath, _utils_.DMode); err != nil {
		return
	}

	// frontend/screens/«screen-package-name»/deps
	packageDepsPath := filepath.Join(packagePath, _utils_.FolderNameDeps)
	if err = os.Mkdir(packageDepsPath, _utils_.DMode); err != nil {
		return
	}

	// frontend/screens/«screen-package-name»/misc
	packageMiscPath := filepath.Join(packagePath, _utils_.FolderNameMisc)
	if err = os.Mkdir(packageMiscPath, _utils_.DMode); err != nil {
		return
	}

	// frontend/screens/simple/«screen-package-name»/presetting
	packagePresettingPath := filepath.Join(packagePath, _utils_.FolderNamePresetting)
	if err = os.Mkdir(packagePresettingPath, _utils_.DMode); err != nil {
		return
	}

	// frontend/screens/«screen-package-name»/panels
	packagePanelsPath := filepath.Join(packagePath, _utils_.FolderNamePanels)
	if err = os.Mkdir(packagePanelsPath, _utils_.DMode); err != nil {
		return
	}

	// frontend/screens/«screen-package-name»/deps/layout
	packageLayoutPath := filepath.Join(packageDepsPath, _utils_.FolderNameLayout)
	if err = os.Mkdir(packageLayoutPath, _utils_.DMode); err != nil {
		return
	}

	// frontend/screens/«screen-package-name»/deps/producer
	packageProducerPath := filepath.Join(packageDepsPath, _utils_.FolderNameProducer)
	if err = os.Mkdir(packageProducerPath, _utils_.DMode); err != nil {
		return
	}

	// frontend/screens/«screen-package-name»/deps/tabItems
	packageTabsPath := filepath.Join(packageDepsPath, _utils_.FolderNameTabItems)
	if err = os.Mkdir(packageTabsPath, _utils_.DMode); err != nil {
		return
	}

	var fPath string
	var data any
	var fileName string
	funcs := _utils_.GetFuncs()
	useConfigTab := (localPanelNames[0] == _utils_.ConfigTabName)

	// Add files to the package folder.

	// frontend/screens/«screen-package-name»/doc.go
	fPath = filepath.Join(packagePath, docFileName)
	files := files(packageName, localPanelNames, folderPaths)
	data = &docTemplateData{
		PackageName:  packageName,
		PackageDoc:   packageDoc,
		Files:        files,
		UseConfigTab: useConfigTab,
		Funcs:        funcs,
	}
	if err = _utils_.ProcessTemplate(docFileName, fPath, docTemplate, data); err != nil {
		return
	}

	// frontend/screens/«screen-package-name»/api.go
	fPath = filepath.Join(packagePath, apiFileName)
	data = &apiTemplateData{
		PackageName:  packageName,
		ImportPrefix: importPrefix,
		Funcs:        funcs,
	}
	if err = _utils_.ProcessTemplate(apiFileName, fPath, apiNoBETemplate, data); err != nil {
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
		AllItemNames: rawPanelNames,
		UseConfigTab: useConfigTab,
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
		LocalPanelNames:  localPanelNames,
		RemotePanelNames: remotePanelNames,
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
	if len(remotePanelNames) > 0 {
		fPath = filepath.Join(packagePresettingPath, _utils_.RemotePresetFileName)
		if err = _utils_.ProcessTemplate(_utils_.RemotePresetFileName, fPath, _presetting_.RemotePresetTemplate, data); err != nil {
			return
		}
	}

	// panels folder.

	// frontend/screens/«screen-package-name»/panels.log
	fPath = filepath.Join(packagePanelsPath, _panels_.LogFileName)
	fmt.Printf("fPath is %s\n", fPath)
	content := _panels_.LogContent(packageName, localPanelNames, folderPaths)
	if err = _utils_.WriteFile(fPath, []byte(content)); err != nil {
		return
	}

	// Add each panel's file and sub folder.
	fmt.Printf("localPanelNames is %+v", localPanelNames)
	for _, panelName := range localPanelNames {
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
		fmt.Printf("making panel folder %s\n", panelFolderName)
		panelFolderPath := filepath.Join(packagePanelsPath, panelFolderName)
		if err = os.Mkdir(panelFolderPath, _utils_.DMode); err != nil {
			return
		}
		data = &_panel_.TemplateData{
			PackageName:     packageName,
			PanelName:       panelName,
			LocalPanelNames: localPanelNames,
			ImportPrefix:    importPrefix,
			Funcs:           funcs,
		}
		// content.go
		fileName = _panel_.ContentFileName
		fPath = filepath.Join(panelFolderPath, fileName)
		if panelName == _utils_.ConfigTabName {
			if err = _utils_.ProcessTemplate(fileName, fPath, _panel_.ConfigContentTemplate, data); err != nil {
				return
			}
		} else {
			if err = _utils_.ProcessTemplate(fileName, fPath, _panel_.ContentTemplate, data); err != nil {
				return
			}
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
		if panelName == _utils_.ConfigTabName {
			if err = _utils_.ProcessTemplate(fileName, fPath, _panel_.ConfigPresetTemplate, data); err != nil {
				return
			}
		} else {
			if err = _utils_.ProcessTemplate(fileName, fPath, _panel_.PresetTemplate, data); err != nil {
				return
			}
		}
	}

	// frontend/screens/«screen-package-name»/deps/

	// frontend/screens/«screen-package-name»/deps/layout/layout.go
	fPath = filepath.Join(packageLayoutPath, _layout_.LayoutFileName)
	data = &_layout_.LayoutTemplateData{
		PackageName:  packageName,
		ImportPrefix: importPrefix,
		UseConfigTab: useConfigTab,
	}
	if err = _utils_.ProcessTemplate(_layout_.LayoutFileName, fPath, _layout_.LayoutTemplate, data); err != nil {
		return
	}

	// producer folder.

	// frontend/screens/«screen-package-name»/deps/producer/appTabs.go
	fPath = filepath.Join(packageProducerPath, _producer_.AppTabsContentProducerFileName)
	data = &_producer_.AppTabsContentProducerTemplateData{
		ImportPrefix: importPrefix,
	}
	if err = _utils_.ProcessTemplate(_producer_.AppTabsContentProducerFileName, fPath, _producer_.AppTabsContentProducerTemplate, data); err != nil {
		return
	}

	// frontend/screens/«screen-package-name»/deps/producer/tabItem.go
	fPath = filepath.Join(packageProducerPath, _producer_.TabItemContentProducerFileName)
	data = &_producer_.TabItemContentProducerTemplateData{
		ImportPrefix: importPrefix,
	}
	if err = _utils_.ProcessTemplate(_producer_.TabItemContentProducerFileName, fPath, _producer_.TabItemContentProducerTemplate, data); err != nil {
		return
	}

	// frontend/screens/«screen-package-name»/deps/tabItems/tabs.go
	fPath = filepath.Join(packageTabsPath, _tabs_.FileName)
	data = &_tabs_.TemplateData{
		PackageName:      packageName,
		ImportPrefix:     importPrefix,
		UseConfigTab:     useConfigTab,
		Funcs:            funcs,
		LocalPanelNames:  localPanelNames,
		RemotePanelNames: remotePanelNames,
	}
	if err = _utils_.ProcessTemplate(_tabs_.FileName, fPath, _tabs_.NoBETemplate, data); err != nil {
		return
	}

	return
}
