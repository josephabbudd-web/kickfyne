package mainmenu

import (
	"fmt"
	"path/filepath"

	_utils_ "github.com/josephabbudd-web/kickfyne/source/utils"
)

// CreateFramework creates the framework's frontend/gui/mainmenu/ files.
func CreateFramework(
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("mainmenu.CreateFramework: %w", err)
		}
	}()

	// gui/mainmenu/mainmenu.go
	data := mainMenuTemplateData{
		ImportPrefix: importPrefix,
	}
	oPath := filepath.Join(folderPaths.FrontendMainMenu, _utils_.MainMenuFileName)
	if err = _utils_.ProcessTemplate(_utils_.MainMenuFileName, oPath, mainMenuTemplate, data); err != nil {
		return
	}

	return
}
