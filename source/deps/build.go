package deps

import (
	"fmt"
	"path/filepath"

	_metadata_ "github.com/josephabbudd-web/kickfyne/source/deps/metadata"
	_paths_ "github.com/josephabbudd-web/kickfyne/source/deps/paths"
	_thread_ "github.com/josephabbudd-web/kickfyne/source/deps/thread"
	_utils_ "github.com/josephabbudd-web/kickfyne/source/utils"
)

const (
	FolderName = "deps"
)

// CreateFramework creates the deps/ files.
func CreateFramework(
	appName string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("deps.CreateFramework: %w", err)
		}
	}()

	var path string
	var data any

	// deps/deps.go
	path = filepath.Join(folderPaths.Deps, depsFileName)
	data = depsTemplateData{
		ImportPrefix: importPrefix,
	}
	if err = _utils_.ProcessTemplate(depsFileName, path, depsTemplate, data); err != nil {
		return
	}

	// deps/metadata/
	if err = _metadata_.CreateFramework(importPrefix, folderPaths); err != nil {
		return
	}

	// deps/paths/
	if err = _paths_.CreateFramework(appName, folderPaths); err != nil {
		return
	}

	// deps/thread/
	if err = _thread_.CreateFramework(importPrefix, folderPaths); err != nil {
		return
	}

	return
}
