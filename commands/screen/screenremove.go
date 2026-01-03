package frontend

import (
	"os"
	"path/filepath"

	_utils_ "github.com/josephabbudd-web/kickfyne/source/utils"
)

// handleScreenRemove handles the removal of a screen package.
func handleScreenRemove(
	screenPackageName string,
	folderPaths *_utils_.FolderPaths,
) (err error) {
	packageFolderPath := filepath.Join(folderPaths.FrontendScreens, screenPackageName)
	if err = os.RemoveAll(packageFolderPath); err != nil {
		if os.IsNotExist(err) {
			err = nil
		}
	}
	return
}
