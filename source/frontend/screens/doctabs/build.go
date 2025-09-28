package simple

import (
	"fmt"
	"os"
	"path/filepath"

	_layout_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/doctabs/deps/layout"
	_panelers_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/doctabs/deps/panelers"
	_producer_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/doctabs/deps/producer"
	_tabs_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/doctabs/deps/tabs"

	_misc_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/doctabs/misc"
	_panels_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/doctabs/panels"
	_panel_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/doctabs/panels/panel"
	_startup_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/doctabs/startup"

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

	defaultPanelName := localPanelNames[0]

	// Build allStartupPanels for template data.
	allStartupPanels := make([]_startup_.Panel, len(rawPanelNames))
	allAPIPanels := make([]apiPanel, len(rawPanelNames))
	for i, name := range rawPanelNames {
		if name[:1] == "*" {
			allStartupPanels[i] = _startup_.Panel{
				Name:    name[1:],
				IsLocal: false,
			}
			allAPIPanels[i] = apiPanel{
				Name:    name[1:],
				IsLocal: false,
			}
		} else {
			allStartupPanels[i] = _startup_.Panel{
				Name:    name,
				IsLocal: true,
			}
			allAPIPanels[i] = apiPanel{
				Name:    name,
				IsLocal: true,
			}
		}
	}

	// Create the folder paths in this package.

	// frontend/screens/«screen-package-name»
	packagePath := filepath.Join(folderPaths.FrontendScreens, packageName)
	if err = os.Mkdir(packagePath, _utils_.DMode); err != nil {
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

	// frontend/screens/simple/«screen-package-name»/startup
	packageStartupPath := filepath.Join(packagePath, _utils_.FolderNameStartup)
	if err = os.Mkdir(packageStartupPath, _utils_.DMode); err != nil {
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

	// frontend/screens/«screen-package-name»/deps/panelers
	packagePanelersPath := filepath.Join(packageDepsPath, _utils_.FolderNamePanelers)
	if err = os.Mkdir(packagePanelersPath, _utils_.DMode); err != nil {
		return
	}

	// frontend/screens/«screen-package-name»/deps/producer
	packageProducerPath := filepath.Join(packageDepsPath, _utils_.FolderNameProducer)
	if err = os.Mkdir(packageProducerPath, _utils_.DMode); err != nil {
		return
	}

	// frontend/screens/«screen-package-name»/deps/tabs
	packageTabsPath := filepath.Join(packageDepsPath, _utils_.FolderNameTabs)
	if err = os.Mkdir(packageTabsPath, _utils_.DMode); err != nil {
		return
	}

	var fPath string
	var data any
	var fileName string
	funcs := _utils_.GetFuncs()

	// Add files to the package folder.

	// frontend/screens/«screen-package-name»/doc.go
	fPath = filepath.Join(packagePath, docFileName)
	successMessage := docTemplateSuccessMessage(packageName, localPanelNames, folderPaths)
	data = &docTemplateData{
		PackageName: packageName,
		PackageDoc:  packageDoc,
		Files:       successMessage,
		Funcs:       funcs,
	}
	if err = _utils_.ProcessTemplate(docFileName, fPath, docTemplate, data); err != nil {
		return
	}

	// frontend/screens/«screen-package-name»/api.go
	fPath = filepath.Join(packagePath, aPIFileName)
	data = &aPITemplateData{
		PackageName:      packageName,
		AllPanels:        allAPIPanels,
		AllPanelNames:    allPanelNames,
		DefaultPanelName: defaultPanelName,
		ImportPrefix:     importPrefix,
		Funcs:            funcs,
	}
	if err = _utils_.ProcessTemplate(aPIFileName, fPath, aPINoBETemplate, data); err != nil {
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

	// frontend/screens/simple/«screen-package-name»/startup/startup.go
	fPath = filepath.Join(packageStartupPath, _utils_.StartupFileName)
	allPanels := make([]_startup_.Panel, len(rawPanelNames))
	for i, name := range rawPanelNames {
		if name[:1] == "*" {
			allPanels[i] = _startup_.Panel{
				Name:    name[1:],
				IsLocal: false,
			}
		} else {
			allPanels[i] = _startup_.Panel{
				Name:    name,
				IsLocal: true,
			}
		}
	}
	data = &_startup_.StartupTemplateData{
		ImportPrefix:     importPrefix,
		AllPanels:        allStartupPanels,
		LocalPanelNames:  localPanelNames,
		RemotePanelNames: remotePanelNames,
		Funcs:            funcs,
	}
	if err = _utils_.ProcessTemplate(_utils_.StartupFileName, fPath, _startup_.StartupTemplate, data); err != nil {
		return
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

		// Panel sub folder holding content and state.
		// frontend/screens/«screen-package-name»/panels/«panel-name»Panel/
		panelFolderName := panelName + "Panel"
		fmt.Printf("making panel folder %s\n", panelFolderName)
		panelFolderPath := filepath.Join(packagePanelsPath, panelFolderName)
		if err = os.Mkdir(panelFolderPath, _utils_.DMode); err != nil {
			return
		}
		// content.go
		fileName = _panel_.ContentFileName
		fPath = filepath.Join(panelFolderPath, fileName)
		data = &_panel_.ContentTemplateData{
			PackageName:     packageName,
			PanelName:       panelName,
			LocalPanelNames: localPanelNames,
			ImportPrefix:    importPrefix,
			Funcs:           funcs,
		}
		if err = _utils_.ProcessTemplate(fileName, fPath, _panel_.ContentTemplate, data); err != nil {
			return
		}
		// state.go
		fileName = _panel_.StateFileName
		fPath = filepath.Join(panelFolderPath, fileName)
		data = &_panel_.StateTemplateData{
			PackageName:  packageName,
			PanelName:    panelName,
			ImportPrefix: importPrefix,
		}
		if err = _utils_.ProcessTemplate(fileName, fPath, _panel_.StateTemplate, data); err != nil {
			return
		}
	}

	// deps/ folder.

	// frontend/screens/«screen-package-name»/deps/layout/layout.go
	fPath = filepath.Join(packageLayoutPath, _layout_.LayoutFileName)
	data = &_layout_.LayoutTemplateData{
		PackageName:      packageName,
		ImportPrefix:     importPrefix,
		Funcs:            funcs,
		LocalPanelNames:  localPanelNames,
		RemotePanelNames: remotePanelNames,
	}
	if err = _utils_.ProcessTemplate(_layout_.LayoutFileName, fPath, _layout_.LayoutTemplate, data); err != nil {
		return
	}

	// frontend/screens/«screen-package-name»/deps/tabs/tabs.go
	fPath = filepath.Join(packageTabsPath, _tabs_.FileName)
	data = &_tabs_.TemplateData{
		PackageName:      packageName,
		ImportPrefix:     importPrefix,
		Funcs:            funcs,
		LocalPanelNames:  localPanelNames,
		RemotePanelNames: remotePanelNames,
	}
	if err = _utils_.ProcessTemplate(_tabs_.FileName, fPath, _tabs_.NoBETemplate, data); err != nil {
		return
	}

	// frontend/screens/«screen-package-name»/deps/panelers/panelers.go
	fPath = filepath.Join(packagePanelersPath, _panelers_.PanelersFileName)
	data = &_panelers_.PanelersTemplateData{
		ImportPrefix:    importPrefix,
		LocalPanelNames: localPanelNames,
	}
	if err = _utils_.ProcessTemplate(_panelers_.PanelersFileName, fPath, _panelers_.PanelersTemplate, data); err != nil {
		return
	}

	// frontend/screens/«screen-package-name»/deps/producer/docTabs.go
	fPath = filepath.Join(packageProducerPath, _producer_.DocTabsContentProducerFileName)
	data = &_producer_.DocTabsContentProducerTemplateData{
		ImportPrefix: importPrefix,
	}
	if err = _utils_.ProcessTemplate(_producer_.DocTabsContentProducerFileName, fPath, _producer_.DocTabsContentProducerTemplate, data); err != nil {
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

	return
}
