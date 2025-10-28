package utils

import (
	"fmt"
)

// ScreenPackageNames returns the names of the screen packages.
func ScreenPackageNames(folderPaths *FolderPaths) (screenNames []string, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("utils.ScreenNames: %w", err)
		}
	}()

	screenNames, err = FolderNames(folderPaths.FrontendScreens)
	return
}
