package paths

import (
	"fmt"
	"path/filepath"

	"github.com/josephabbudd-web/kickfyne/source/utils"
)

// CreateFramework creates the deps/paths/ files.
func CreateFramework(
	appName string,
	folderPaths *utils.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("paths.CreateFramework: %w", err)
		}
	}()

	// paths/paths.go
	oPath := filepath.Join(folderPaths.DepsPaths, fileName)
	data := templateData{
		AppName: appName,
		Funcs:   utils.GetFuncs(),
	}
	err = utils.ProcessTemplate(fileName, oPath, template, data)
	return
}
