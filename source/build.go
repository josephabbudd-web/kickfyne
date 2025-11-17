package source

import (
	"fmt"
	"os"

	_manifest_ "github.com/josephabbudd-web/kickfyne/manifest"
	_deps_ "github.com/josephabbudd-web/kickfyne/source/deps"
	_frontend_ "github.com/josephabbudd-web/kickfyne/source/frontend"
	_root_ "github.com/josephabbudd-web/kickfyne/source/root"
	_utils_ "github.com/josephabbudd-web/kickfyne/source/utils"
)

func HasAppFolder(currentWP, appName string) (hasAppFolder bool, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("source.HasAppFolder: %w", err)
		}
	}()

	var dirEntrys []os.DirEntry
	if dirEntrys, err = os.ReadDir(currentWP); err != nil {
		return
	}
	var dirEntry os.DirEntry
	for _, dirEntry = range dirEntrys {
		if dirEntry.IsDir() {
			dName := dirEntry.Name()
			if hasAppFolder = dName == appName; hasAppFolder {
				return
			}
		}
	}
	return
}

// CreateFramework builds the framework in an appName folder in this parent folder.
func CreateFramework(
	manifest _manifest_.Manifest,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("source.CreateFramework: %w", err)
		}
	}()

	appName := folderPaths.Base()

	// App folder.
	if err = _root_.CreateFramework(appName, importPrefix, folderPaths); err != nil {
		return
	}

	// Deps
	if err = _deps_.CreateFramework(appName, importPrefix, folderPaths); err != nil {
		return
	}

	// Frontend
	if err = _frontend_.CreateFramework(manifest, importPrefix, folderPaths); err != nil {
		return
	}

	// Update the manifest.
	manifest.AddFramework()

	return
}
