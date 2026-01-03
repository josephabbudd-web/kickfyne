package types

import (
	"fmt"
	"path/filepath"

	_utils_ "github.com/josephabbudd-web/kickfyne/source/utils"
)

// CreateFramework creates the framework's frontend/deps/types.
func CreateFramework(
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("types.Build: %w", err)
		}
	}()

	var oPath string
	var templateData any

	// frontend/deps/types/interfaces.go.
	oPath = filepath.Join(folderPaths.FrontendTypes, interfacesFileName)
	if err = _utils_.WriteFile(oPath, []byte(interfacesNoBETemplate)); err != nil {
		return
	}

	// frontend/deps/types/producerState.go.
	oPath = filepath.Join(folderPaths.FrontendTypes, producerStateFileName)
	if err = _utils_.WriteFile(oPath, []byte(producerStateTemplate)); err != nil {
		return
	}

	// frontend/deps/types/accordionItemContentConsumer.go
	oPath = filepath.Join(folderPaths.FrontendTypes, accordionItemContentConsumerFileName)
	if err = _utils_.WriteFile(oPath, []byte(accordionItemContentConsumerTemplate)); err != nil {
		return
	}

	// frontend/deps/types/appTabItemContentConsumer.go
	oPath = filepath.Join(folderPaths.FrontendTypes, appTabItemContentConsumerFileName)
	if err = _utils_.WriteFile(oPath, []byte(appTabItemContentConsumerTemplate)); err != nil {
		return
	}

	// frontend/deps/types/borderAreaContentConsumer.go
	oPath = filepath.Join(folderPaths.FrontendTypes, borderAreaContentConsumerFileName)
	if err = _utils_.WriteFile(oPath, []byte(borderAreaContentConsumerTemplate)); err != nil {
		return
	}

	// frontend/deps/types/docTabItemContentConsumer.go
	templateData = &docTabItemContentConsumerTemplateData{
		ImportPrefix: importPrefix,
	}
	oPath = filepath.Join(folderPaths.FrontendTypes, docTabItemContentConsumerFileName)
	if err = _utils_.ProcessTemplate(docTabItemContentConsumerFileName, oPath, docTabItemContentConsumerTemplate, templateData); err != nil {
		return
	}

	// frontend/deps/types/windowContentConsumer.go
	oPath = filepath.Join(folderPaths.FrontendTypes, windowContentConsumerFileName)
	if err = _utils_.ProcessTemplate(windowContentConsumerFileName, oPath, windowContentConsumerTemplate, nil); err != nil {
		return
	}

	// frontend/deps/types/splitContentProducer.go
	oPath = filepath.Join(folderPaths.FrontendTypes, splitContentProducerFileName)
	if err = _utils_.ProcessTemplate(splitContentProducerFileName, oPath, splitContentProducerTemplate, templateData); err != nil {
		return
	}

	// frontend/deps/types/splitAreaContentProducer.go
	oPath = filepath.Join(folderPaths.FrontendTypes, splitAreaContentProducerFileName)
	if err = _utils_.WriteFile(oPath, []byte(splitAreaContentProducerTemplate)); err != nil {
		return
	}

	// frontend/deps/types/splitAreaContentConsumer.go
	oPath = filepath.Join(folderPaths.FrontendTypes, splitAreaContentConsumerFileName)
	if err = _utils_.WriteFile(oPath, []byte(splitAreaContentConsumerTemplate)); err != nil {
		return
	}

	return
}
