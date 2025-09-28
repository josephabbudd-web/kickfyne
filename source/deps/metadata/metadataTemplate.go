package metadata

type metadataTemplateData struct {
	ImportPrefix string
}

const (
	metadataFileName = "metadata.go"

	metadataTemplate = `package metadata

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/storage"

	toml "github.com/pelletier/go-toml"

	_paths_ "{{ .ImportPrefix }}/deps/paths"
)

var appMetaData fyne.AppMetadata
var loadErr error
var loaded bool

func Fyne() (fyneAppMetadata fyne.AppMetadata, err error) {
	if loadErr == nil && !loaded {
		loadMetaData()
	}
	if loadErr != nil {
		err = loadErr
		return
	}
	// Return the fyne.AppMetadata.
	fyneAppMetadata = appMetaData
	return
}

func loadMetaData() {
	loaded = true
	// Open.
	var rc fyne.URIReadCloser
	var err error
	if rc, err = storage.Reader(_paths_.MetaDataURI()); loadErr != nil {
		loadErr = fmt.Errorf("metadata.loadAppMetaData:storage.Reader %w", err)
		return
	}
	// Decode.
	decoder := toml.NewDecoder(rc)
	if err = decoder.Decode(&appMetaData); err != nil {
		loadErr = fmt.Errorf("metadata.loadAppMetaData.decoder.Decode: %w", err)
		return
	}
}
`
)
