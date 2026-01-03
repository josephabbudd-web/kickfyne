package data

import (
	// "slices"
	"log"
	"strings"

	_utils_ "github.com/josephabbudd-web/kickfyne/source/utils"
)

const (
	LeadingArea  = "Leading"
	TrailingArea = "Trailing"
)

type AreaData struct {
	ScreenName string
	IsLocal    bool
}

type TemplateData struct {
	PackageName string
	// UniqueRemoteScreenNames []string
	Leading           *AreaData
	Trailing          *AreaData
	ImportPrefix      string
	UsesRemoteContent bool
	UsesLocalContent  bool
	Funcs             _utils_.Funcs
}

func New(packageName string, rawPanelNames []string, importPrefix string) (templateData *TemplateData) {
	templateData = &TemplateData{
		PackageName:  packageName,
		ImportPrefix: importPrefix,
		Funcs:        _utils_.GetFuncs(),
		Leading:      &AreaData{},
		Trailing:     &AreaData{},
	}
	for _, rawPanelName := range rawPanelNames {
		log.Println("rawPanelName is ", rawPanelName)
		parts := strings.Split(rawPanelName, "=")
		if len(parts) == 1 {
			templateData.UsesLocalContent = true
			switch rawPanelName {
			case LeadingArea:
				templateData.Leading.IsLocal = true
			case TrailingArea:
				templateData.Trailing.IsLocal = true
			}
			continue
		}
		// 2 parts: "Center=*SomeScreen".
		templateData.UsesRemoteContent = true
		areaName := parts[0]
		screenName := parts[1][1:]
		switch areaName {
		case LeadingArea:
			templateData.Leading.ScreenName = screenName
		case TrailingArea:
			templateData.Trailing.ScreenName = screenName
		}
	}
	return
}
