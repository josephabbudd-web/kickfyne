package frontend

import (
	"fmt"
	"path/filepath"

	_mainmenu_ "github.com/josephabbudd-web/kickfyne/source/frontend/mainmenu"
	_screenmap_ "github.com/josephabbudd-web/kickfyne/source/frontend/screenmap"
	_screens_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens"
	_types_ "github.com/josephabbudd-web/kickfyne/source/frontend/types"
	_utils_ "github.com/josephabbudd-web/kickfyne/source/utils"
)

// CreateFramework creates the framework's frontend/ files.
func CreateFramework(
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("frontend.CreateFramework: %w", err)
		}
	}()

	// gui/mainmenu/ package
	if err = _mainmenu_.CreateFramework(importPrefix, folderPaths); err != nil {
		return
	}

	// frontend/types/
	if err = _types_.CreateFramework(importPrefix, folderPaths); err != nil {
		return
	}

	// Add the HelloWorld screen.
	if err = _screens_.CreateFramework(importPrefix, folderPaths); err != nil {
		return
	}
	// Rebuild frontend/frontend.go
	if err = RebuildFrontendGo(importPrefix, folderPaths); err != nil {
		return
	}
	// Rebuild frontend/screenmap/screenmap.go
	err = _screenmap_.CreateFramework(importPrefix, folderPaths)

	return

}

func RebuildFrontendGo(
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {
	var screenPackageNames []string
	if screenPackageNames, err = _utils_.ScreenPackageNames(folderPaths); err != nil {
		return
	}

	// frontend/frontend.go
	oPath := filepath.Join(folderPaths.Frontend, frontendFileName)
	data := frontendTemplateData{
		ImportPrefix: importPrefix,
		ScreenNames:  screenPackageNames,
	}
	err = _utils_.ProcessTemplate(frontendFileName, oPath, frontendNoBETemplate, data)
	return
}
