package simple

import (
	"fmt"
	"os"
	"path/filepath"

	_accordionitems_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/accordion/deps/accordionItems"
	_layout_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/accordion/deps/layout"
	_layoutitems_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/accordion/layoutAccordionItems"
	_misc_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/accordion/misc"
	_panels_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/accordion/panels"
	_panel_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/accordion/panels/panel"
	_presetting_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/accordion/presetting"

	_manifest_ "github.com/josephabbudd-web/kickfyne/manifest"
	_utils_ "github.com/josephabbudd-web/kickfyne/source/utils"
)

// AddItems appends items to and Accordion screen package.
func AddItems(
	packageName string,
	addItemNames []string,
	packageDoc string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("accordion.AddItems: %w", err)
		}
	}()

	manifest := _manifest_.NewAgain()
	var infoCopy *_manifest_.Info
	if infoCopy = manifest.InfoCopy(packageName); infoCopy == nil {
		return
	}
	infoCopy.Add(addItemNames...)
	allItemNames, allLocalItemNames, allRemoteItemNames := infoCopy.GetItems()
	addLocalItemNames := make([]string, 0, len(addItemNames))
	// addRemoteItemNames := make([]string, 0, len(addItemNames))
	for _, itemName := range addItemNames {
		if itemName[:1] == "*" {
			// addRemoteItemNames = append(addRemoteItemNames, itemName[1:])
		} else {
			addLocalItemNames = append(addLocalItemNames, itemName)
		}
	}

	// Create the folder paths in this package.

	// frontend/screens/«screen-package-name»
	packagePath := filepath.Join(folderPaths.FrontendScreens, packageName)

	// frontend/screens/«screen-package-name»/layoutAccordionItems
	packageLayoutItemsPath := filepath.Join(packagePath, _utils_.FolderNameLayoutAccorionItems)

	// frontend/screens/«screen-package-name»/deps
	packageDepsPath := filepath.Join(packagePath, _utils_.FolderNameDeps)

	// frontend/screens/«screen-package-name»/misc
	packageMiscPath := filepath.Join(packagePath, _utils_.FolderNameMisc)

	// frontend/screens/«screen-package-name»/panels
	packagePanelsPath := filepath.Join(packagePath, _utils_.FolderNamePanels)

	// frontend/screens/simple/«screen-package-name»/presetting
	packagePresettingPath := filepath.Join(packagePath, _utils_.FolderNamePresetting)

	// deps/

	// frontend/screens/«screen-package-name»/deps/layout
	packageLayoutPath := filepath.Join(packageDepsPath, _utils_.FolderNameLayout)

	// frontend/screens/«screen-package-name»/deps/accordionItems
	packageAccordionItemPath := filepath.Join(packageDepsPath, _utils_.FolderNameAccordionItems)

	var fPath string
	var data any
	var fileName string
	funcs := _utils_.GetFuncs()

	// Add files to the package folder.

	// frontend/screens/«screen-package-name»/doc.go
	fPath = filepath.Join(packagePath, docFileName)
	files := files(packageName, allLocalItemNames, folderPaths)
	data = &docTemplateData{
		PackageName: packageName,
		PackageDoc:  packageDoc,
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

	// frontend/screens/«screen-package-name»/layoutAccordionItems folder.
	data = &_layoutitems_.TemplateData{
		PackageName:  packageName,
		ImportPrefix: importPrefix,
		AllItemNames: allItemNames,
	}
	fPath = filepath.Join(packageLayoutItemsPath, _utils_.DocFileName)
	if err = _utils_.ProcessTemplate(_utils_.DocFileName, fPath, _layoutitems_.DocsTemplate, data); err != nil {
		return
	}
	fPath = filepath.Join(packageLayoutItemsPath, _utils_.LayoutFileName)
	if err = _utils_.ProcessTemplate(_utils_.LayoutFileName, fPath, _layoutitems_.LayoutTemplate, data); err != nil {
		return
	}

	// frontend/screens/simple/«screen-package-name»/presetting folder.

	data = &_presetting_.TemplateData{
		PackageName:      packageName,
		ImportPrefix:     importPrefix,
		LocalPanelNames:  allLocalItemNames,
		RemotePanelNames: allRemoteItemNames,
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
	if len(allRemoteItemNames) > 0 {
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
	content := _panels_.LogContent(packageName, allLocalItemNames, folderPaths)
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
		fmt.Printf("making panel folder %s\n", panelFolderName)
		panelFolderPath := filepath.Join(packagePanelsPath, panelFolderName)
		if err = os.Mkdir(panelFolderPath, _utils_.DMode); err != nil {
			return
		}
		data = &_panel_.TemplateData{
			PackageName:     packageName,
			PanelName:       panelName,
			LocalPanelNames: allLocalItemNames,
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

	// frontend/screens/«screen-package-name»/misc/miscellaneous.go
	fPath = filepath.Join(packageMiscPath, _misc_.MiscellaneousFileName)
	data = &_misc_.MiscellaneousTemplateData{
		PackageName:  packageName,
		ImportPrefix: importPrefix,
	}
	if err = _utils_.ProcessTemplate(_misc_.MiscellaneousFileName, fPath, _misc_.MiscellaneousTemplate, data); err != nil {
		return
	}

	// frontend/screens/«screen-package-name»/deps

	// frontend/screens/«screen-package-name»/deps/accordionItems/accordionItems.go
	fPath = filepath.Join(packageAccordionItemPath, _accordionitems_.FileName)
	data = &_accordionitems_.TemplateData{
		PackageName:      packageName,
		ImportPrefix:     importPrefix,
		Funcs:            funcs,
		LocalPanelNames:  allLocalItemNames,
		RemotePanelNames: allRemoteItemNames,
	}
	if err = _utils_.ProcessTemplate(_accordionitems_.FileName, fPath, _accordionitems_.NoBETemplate, data); err != nil {
		return
	}

	// frontend/screens/«screen-package-name»/deps/layout/layout.go
	fPath = filepath.Join(packageLayoutPath, _layout_.LayoutFileName)
	data = &_layout_.LayoutTemplateData{
		PackageName:  packageName,
		ImportPrefix: importPrefix,
	}
	if err = _utils_.ProcessTemplate(_layout_.LayoutFileName, fPath, _layout_.LayoutTemplate, data); err != nil {
		return
	}

	return
}
