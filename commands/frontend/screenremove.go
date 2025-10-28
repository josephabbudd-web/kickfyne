package frontend

import (
	"fmt"
	"os"
	"path/filepath"

	_utils_ "github.com/josephabbudd-web/kickfyne/source/utils"
)

// handleScreenRemove handles the removal of a screen package.
func handleScreenRemove(
	screenPackageName string,
	folderPaths *_utils_.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("frontend.handleScreenRemove: %w", err)
		}
	}()

	// Remove the frontend/screen/«screenPackageName» folder.
	packageFolderPath := filepath.Join(folderPaths.FrontendScreens, screenPackageName)
	if err = os.RemoveAll(packageFolderPath); err != nil {
		if os.IsNotExist(err) {
			err = nil
		} else {
			fmt.Printf("Failure:\nUnable to remove the %[1]s screen's folder at frontend/screens/%[1]s.\n", screenPackageName)
		}
		return
	}
	fmt.Printf("Success:\nRemoved the %[1]s screen's folder at frontend/screens/%[1]s.\n", screenPackageName)

	return
}
