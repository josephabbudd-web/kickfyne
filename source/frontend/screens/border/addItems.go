package border

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	_data_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/border/data"
	_border_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/border/deps/border"
	_layout_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/border/deps/layout"
	_misc_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/border/misc"
	_panels_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/border/panels"
	_panel_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/border/panels/panel"
	_presetting_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/border/presetting"

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

// AddAreas add areas to a Border container.
func AddAreas(
	packageName string,
	addItemNames []string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("apptabs.AddAreas: %w", err)
		}
	}()

	manifest := _manifest_.NewAgain()
	var infoCopy *_manifest_.Info
	if infoCopy = manifest.InfoCopy(packageName); infoCopy == nil {
		return
	}
	infoCopy.AddItems(addItemNames...)
	manifestAllItemNames, _, _ := infoCopy.GetItems()
	addLocalItemNames := make([]string, 0, len(addItemNames))
	for _, itemName := range addItemNames {
		if !strings.Contains(itemName, "=") {
			addLocalItemNames = append(addLocalItemNames, itemName)
		}
	}

	generalTemplateData := _data_.New(packageName, manifestAllItemNames, importPrefix)

	// Create the folder paths in this package.

	// frontend/screens/«screen-package-name»
	packagePath := filepath.Join(folderPaths.FrontendScreens, packageName)

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

	// frontend/screens/«screen-package-name»/deps/border
	packageBorderPath := filepath.Join(packageDepsPath, _utils_.FolderNameBorder)

	var fPath string
	var data any
	var fileName string
	funcs := _utils_.GetFuncs()

	// Add files to the package folder.

	// frontend/screens/«screen-package-name»/doc.go
	fPath = filepath.Join(packagePath, docFileName)
	files := files(generalTemplateData, folderPaths)
	data = &docTemplateData{
		Data:  generalTemplateData,
		Files: files,
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

	// frontend/screens/«screen-package-name»/misc/miscellaneous.go
	fPath = filepath.Join(packageMiscPath, _misc_.MiscellaneousFileName)
	if err = _utils_.ProcessTemplate(_misc_.MiscellaneousFileName, fPath, _misc_.MiscellaneousTemplate, generalTemplateData); err != nil {
		return
	}

	// frontend/screens/simple/«screen-package-name»/presetting folder.
	fPath = filepath.Join(packagePresettingPath, _utils_.APIFileName)
	if err = _utils_.ProcessTemplate(_utils_.APIFileName, fPath, _presetting_.APITemplate, generalTemplateData); err != nil {
		return
	}
	fPath = filepath.Join(packagePresettingPath, _utils_.DefaultPresetFileName)
	if err = _utils_.ProcessTemplate(_utils_.DefaultPresetFileName, fPath, _presetting_.DefaultPresetTemplate, generalTemplateData); err != nil {
		return
	}
	fPath = filepath.Join(packagePresettingPath, _utils_.RemotePresetFileName)
	if generalTemplateData.UsesRemoteContent {
		if err = _utils_.ProcessTemplate(_utils_.RemotePresetFileName, fPath, _presetting_.RemotePresetTemplate, generalTemplateData); err != nil {
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
	content := _panels_.DocContent(packageName, generalTemplateData.Areas, folderPaths)
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
		if err = _utils_.ProcessTemplate(fileName, fPath, _panels_.PanelTemplate, data); err != nil {
			return
		}

		// Panel sub folder holding content.go, state.go and preset.go.
		// frontend/screens/«screen-package-name»/panels/«panel-name»Panel/
		panelFolderName := panelName + "Panel"
		panelFolderPath := filepath.Join(packagePanelsPath, panelFolderName)
		if err = os.Mkdir(panelFolderPath, _utils_.DMode); err != nil {
			return
		}
		data = &_panels_.PanelTemplateData{
			PackageName:  packageName,
			PanelName:    panelName,
			Area:         panelName,
			ImportPrefix: importPrefix,
			Funcs:        funcs,
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
	if err = _utils_.ProcessTemplate(_layout_.LayoutFileName, fPath, _layout_.LayoutTemplate, generalTemplateData); err != nil {
		return
	}

	// frontend/screens/«screen-package-name»/deps/border/border.go
	fPath = filepath.Join(packageBorderPath, _border_.FileName)
	if err = _utils_.ProcessTemplate(_border_.FileName, fPath, _border_.Template, generalTemplateData); err != nil {
		return
	}

	return
}
