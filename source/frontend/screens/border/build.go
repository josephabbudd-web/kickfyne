package border

import (
	"fmt"
	"os"
	"path/filepath"

	_data_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/border/data"
	_border_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/border/deps/border"
	_layout_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/border/deps/layout"
	_producer_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/border/deps/producer"
	_misc_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/border/misc"
	_panels_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/border/panels"
	_panel_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/border/panels/panel"
	_presetting_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/border/presetting"

	_utils_ "github.com/josephabbudd-web/kickfyne/source/utils"
)

// Build builds a type Border screen package.
func Build(
	packageName string,
	rawPanelNames []string, // Ex: "Top Bottom Left Right Center=*Edit"
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("border.Build: %w", err)
		}
	}()

	generalTemplateData := _data_.New(packageName, rawPanelNames, importPrefix)

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

	// frontend/screens/simple/«screen-package-name»/presetting
	packagePresettingPath := filepath.Join(packagePath, _utils_.FolderNamePresetting)
	if err = os.Mkdir(packagePresettingPath, _utils_.DMode); err != nil {
		return
	}

	// frontend/screens/«screen-package-name»/panels
	var packagePanelsPath string
	if generalTemplateData.UsesLocalContent {
		packagePanelsPath = filepath.Join(packagePath, _utils_.FolderNamePanels)
		if err = os.Mkdir(packagePanelsPath, _utils_.DMode); err != nil {
			return
		}
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

	// frontend/screens/«screen-package-name»/deps/border
	packageBorderPath := filepath.Join(packageDepsPath, _utils_.FolderNameBorder)
	if err = os.Mkdir(packageBorderPath, _utils_.DMode); err != nil {
		return
	}

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
	if err = _utils_.ProcessTemplate(apiFileName, fPath, apiNoBETemplate, generalTemplateData); err != nil {
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
	if generalTemplateData.UsesLocalContent {
		// frontend/screens/«screen-package-name»/panels.log
		fPath = filepath.Join(packagePanelsPath, _panels_.LogFileName)
		content := _panels_.LogContent(packageName, generalTemplateData.Areas, folderPaths)
		if err = _utils_.WriteFile(fPath, []byte(content)); err != nil {
			return
		}

		// Add each panel's file and sub folder.
		for _, area := range generalTemplateData.Areas {
			if !area.IsLocal {
				continue
			}
			data = &_panels_.PanelTemplateData{
				PackageName:  packageName,
				PanelName:    area.ItemName,
				Area:         area.Area,
				ImportPrefix: importPrefix,
				Funcs:        funcs,
			}
			// Panel file.
			// frontend/screens/«screen-package-name»/panels/«panel-name»Panel/
			fileName = area.ItemName + _panels_.PanelFileNameSuffix
			fPath = filepath.Join(packagePanelsPath, fileName)
			if err = _utils_.ProcessTemplate(fileName, fPath, _panels_.PanelTemplate, data); err != nil {
				return
			}

			// Panel sub folder holding content.go, state.go and preset.go.
			// frontend/screens/«screen-package-name»/panels/«panel-name»Panel/
			panelFolderName := area.ItemName + "Panel"
			panelFolderPath := filepath.Join(packagePanelsPath, panelFolderName)
			if err = os.Mkdir(panelFolderPath, _utils_.DMode); err != nil {
				return
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
	}

	// frontend/screens/«screen-package-name»/deps/

	// frontend/screens/«screen-package-name»/deps/layout/layout.go
	fPath = filepath.Join(packageLayoutPath, _layout_.LayoutFileName)
	if err = _utils_.ProcessTemplate(_layout_.LayoutFileName, fPath, _layout_.LayoutTemplate, generalTemplateData); err != nil {
		return
	}

	// producer folder.

	// frontend/screens/«screen-package-name»/deps/producer/border.go
	fPath = filepath.Join(packageProducerPath, _producer_.BorderContentProducerFileName)
	if err = _utils_.ProcessTemplate(_producer_.BorderContentProducerFileName, fPath, _producer_.BorderContentProducerTemplate, generalTemplateData); err != nil {
		return
	}

	// frontend/screens/«screen-package-name»/deps/producer/area.go
	fPath = filepath.Join(packageProducerPath, _producer_.BorderAreaContentProducerFileName)
	if err = _utils_.ProcessTemplate(_producer_.BorderAreaContentProducerTemplate, fPath, _producer_.BorderAreaContentProducerTemplate, generalTemplateData); err != nil {
		return
	}

	// frontend/screens/«screen-package-name»/deps/border/border.go
	fPath = filepath.Join(packageBorderPath, _border_.FileName)
	if err = _utils_.ProcessTemplate(_border_.FileName, fPath, _border_.Template, generalTemplateData); err != nil {
		return
	}

	return
}
