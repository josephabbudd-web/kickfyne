package mainmenu

import (
	"fmt"
	"path/filepath"

	"github.com/josephabbudd-web/kickfyne/source/utils"
)

// CreateFramework creates the framework's frontend/gui/mainmenu/ files.
func CreateFramework(
	importPrefix string,
	folderPaths *utils.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("mainmenu.Build: %w", err)
		}
	}()

	{
		// gui/mainmenu/init.go
		data := initTemplateData{
			ImportPrefix: importPrefix,
		}
		oPath := filepath.Join(folderPaths.FrontendMainMenu, initFileName)
		if err = utils.ProcessTemplate(initFileName, oPath, initTemplate, data); err != nil {
			return
		}
	}

	{
		// gui/mainmenu/mainmenu.go
		data := mainMenuTemplateData{
			ImportPrefix: importPrefix,
		}
		oPath := filepath.Join(folderPaths.FrontendMainMenu, mainMenuFileName)
		if err = utils.ProcessTemplate(mainMenuFileName, oPath, mainMenuTemplate, data); err != nil {
			return
		}
	}

	return
}
