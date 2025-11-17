package border

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	_data_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/border/data"
	_border_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/border/deps/border"
	_layout_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/border/deps/layout"
	_misc_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/border/misc"
	_panels_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/border/panels"
	_presetting_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/border/presetting"

	_manifest_ "github.com/josephabbudd-web/kickfyne/manifest"
	_utils_ "github.com/josephabbudd-web/kickfyne/source/utils"
)

// RemoveAreas removes areas from a Border screen package.
func RemoveAreas(
	packageName string,
	removeItemNames []string, // Ex: "Left Bottom *Center=SomeScreen"
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("border.RemoveItems: %w", err)
		}
	}()

	manifest := _manifest_.NewAgain()
	var infoCopy *_manifest_.Info
	if infoCopy = manifest.InfoCopy(packageName); infoCopy == nil {
		return
	}
	infoCopy.RemoveItems(removeItemNames...)

	finalAllItemNames, _, _ := infoCopy.GetItems()
	log.Printf("finalAllItemNames is %#v", finalAllItemNames)
	removeLocalItemNames := make([]string, 0, len(removeItemNames))
	for _, itemName := range removeItemNames {
		if !strings.Contains(itemName, "=") {
			removeLocalItemNames = append(removeLocalItemNames, itemName)
		}
	}

	generalTemplateData := _data_.New(packageName, finalAllItemNames, importPrefix)

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

	// Remove files from the package folder.

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

	// frontend/screens/«screen-package-name»/panels.log
	fPath = filepath.Join(packagePanelsPath, _panels_.LogFileName)
	content := _panels_.LogContent(packageName, generalTemplateData.Areas, folderPaths)
	if err = _utils_.WriteFile(fPath, []byte(content)); err != nil {
		return
	}

	// Remove each panel's file and sub folder.
	for _, panelName := range removeLocalItemNames {
		// Panel file.
		// frontend/screens/«screen-package-name»/panels/«panel-name»Panel.go
		fileName = panelName + _panels_.PanelFileNameSuffix
		fPath = filepath.Join(packagePanelsPath, fileName)
		if err = os.Remove(fPath); err != nil {
			if !os.IsNotExist(err) {
				return
			}
			err = nil
		}
		// frontend/screens/«screen-package-name»/panels/«panel-name»Panel/
		folderName := panelName + "Panel"
		fPath = filepath.Join(packagePanelsPath, folderName)
		if err = os.RemoveAll(fPath); err != nil {
			if !os.IsNotExist(err) {
				return
			}
			err = nil
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
