package frontend

import (
	"fmt"
	"path/filepath"

	_manifest_ "github.com/josephabbudd-web/kickfyne/manifest"
	_mainmenu_ "github.com/josephabbudd-web/kickfyne/source/frontend/deps/mainmenu"
	_screenmap_ "github.com/josephabbudd-web/kickfyne/source/frontend/deps/screenmap"
	_types_ "github.com/josephabbudd-web/kickfyne/source/frontend/deps/types"
	_screens_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens"
	_utils_ "github.com/josephabbudd-web/kickfyne/source/utils"
)

// CreateFramework creates the framework's frontend/ files.
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

	// frontend/frontend.go
	if err = RebuildFrontendGo([]string{"HelloWorld"}, importPrefix, folderPaths); err != nil {
		return
	}

	// frontend/settings.go
	oPath := filepath.Join(folderPaths.Frontend, settingsFileName)
	data := frontendTemplateData{

		ImportPrefix: importPrefix,
		ScreenNames:  []string{"HelloWorld"},
	}
	err = _utils_.ProcessTemplate(settingsFileName, oPath, settingsTemplate, data)

	// frontend/deps/mainmenu/ package
	if err = _mainmenu_.CreateFramework(importPrefix, folderPaths); err != nil {
		return
	}

	// frontend/deps/types/
	if err = _types_.CreateFramework(importPrefix, folderPaths); err != nil {
		return
	}

	// Add the HelloWorld screen.
	if err = _screens_.CreateFramework(manifest, importPrefix, folderPaths); err != nil {
		return
	}

	// Rebuild frontend/deps/screenmap/screenmap.go
	err = _screenmap_.CreateFramework(importPrefix, folderPaths)

	return

}

func RebuildFrontendGo(
	screenPackageNames []string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {
	// var screenPackageNames []string
	// if screenPackageNames, err = _utils_.ScreenPackageNames(folderPaths); err != nil {
	// 	return
	// }

	// frontend/frontend.go
	oPath := filepath.Join(folderPaths.Frontend, frontendFileName)
	data := frontendTemplateData{
		ImportPrefix: importPrefix,
		ScreenNames:  screenPackageNames,
		Funcs:        _utils_.GetFuncs(),
	}
	err = _utils_.ProcessTemplate(frontendFileName, oPath, frontendNoBETemplate, data)
	return
}
