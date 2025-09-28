package metadata

import (
	"fmt"
	"path/filepath"

	_utils_ "github.com/josephabbudd-web/kickfyne/source/utils"
)

// CreateFramework creates the deps/ files.
func CreateFramework(
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("metadata.CreateFramework: %w", err)
		}
	}()

	oPath := filepath.Join(folderPaths.DepsMetaData, metadataFileName)
	data := metadataTemplateData{
		ImportPrefix: importPrefix,
	}
	err = _utils_.ProcessTemplate(metadataFileName, oPath, metadataTemplate, data)
	return
}

// AppMetaDataFilePath returns the path to the app's meta data file.
func AppMetaDataFilePath(folderPaths *_utils_.FolderPaths) (path string) {
	path = filepath.Join(folderPaths.DepsMetaData, metadataFileName)
	return
}
