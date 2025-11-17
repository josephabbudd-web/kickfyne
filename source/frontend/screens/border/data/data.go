package data

import (
	"slices"
	"strings"

	_utils_ "github.com/josephabbudd-web/kickfyne/source/utils"
)

const (
	TopArea    = "Top"
	BottomArea = "Bottom"
	LeftArea   = "Left"
	RightArea  = "Right"
	CenterArea = "Center"
)

type AreaData struct {
	Area     string
	ItemName string
	IsLocal  bool
}

type TemplateData struct {
	PackageName           string
	Areas                 []*AreaData
	UniqueRemoteItemNames []string
	Top                   *AreaData
	Bottom                *AreaData
	Left                  *AreaData
	Right                 *AreaData
	Center                *AreaData
	TopIndex              int
	BottomIndex           int
	LeftIndex             int
	RightIndex            int
	CenterIndex           int
	ImportPrefix          string
	UsesRemoteContent     bool
	UsesLocalContent      bool
	Funcs                 _utils_.Funcs
}

func New(packageName string, rawPanelNames []string, importPrefix string) (templateData *TemplateData) {
	templateData = &TemplateData{
		PackageName:  packageName,
		ImportPrefix: importPrefix,
		Funcs:        _utils_.GetFuncs(),

		UniqueRemoteItemNames: make([]string, 0, len(rawPanelNames)),
		Top:                   &AreaData{},
		Bottom:                &AreaData{},
		Left:                  &AreaData{},
		Right:                 &AreaData{},
		Center:                &AreaData{},
		TopIndex:              -1,
		BottomIndex:           -1,
		LeftIndex:             -1,
		RightIndex:            -1,
		CenterIndex:           -1,
	}
	for _, rawPanelName := range rawPanelNames {
		parts := strings.Split(rawPanelName, "=")
		if len(parts) == 1 {
			templateData.UsesLocalContent = true
			switch rawPanelName {
			case TopArea:
				templateData.Top.Area = rawPanelName
				templateData.Top.ItemName = rawPanelName
				templateData.Top.IsLocal = true
			case BottomArea:
				templateData.Bottom.Area = rawPanelName
				templateData.Bottom.ItemName = rawPanelName
				templateData.Bottom.IsLocal = true
			case LeftArea:
				templateData.Left.Area = rawPanelName
				templateData.Left.ItemName = rawPanelName
				templateData.Left.IsLocal = true
			case RightArea:
				templateData.Right.Area = rawPanelName
				templateData.Right.ItemName = rawPanelName
				templateData.Right.IsLocal = true
			case CenterArea:
				templateData.Center.Area = rawPanelName
				templateData.Center.ItemName = rawPanelName
				templateData.Center.IsLocal = true
			}
			continue
		}
		// 2 parts: "Center=*SomeScreen".
		templateData.UsesRemoteContent = true
		areaName := parts[0]
		screenName := parts[1][1:]
		switch areaName {
		case TopArea:
			templateData.Top.Area = areaName
			templateData.Top.ItemName = screenName
			templateData.Top.IsLocal = false
		case BottomArea:
			templateData.Bottom.Area = areaName
			templateData.Bottom.ItemName = screenName
			templateData.Bottom.IsLocal = false
		case LeftArea:
			templateData.Left.Area = areaName
			templateData.Left.ItemName = screenName
			templateData.Left.IsLocal = false
		case RightArea:
			templateData.Right.Area = areaName
			templateData.Right.ItemName = screenName
			templateData.Right.IsLocal = false
		case CenterArea:
			templateData.Center.Area = areaName
			templateData.Center.ItemName = screenName
			templateData.Center.IsLocal = false
		}
		if !slices.Contains(templateData.UniqueRemoteItemNames, screenName) {
			templateData.UniqueRemoteItemNames = append(templateData.UniqueRemoteItemNames, screenName)
		}
	}
	templateData.Areas = []*AreaData{
		templateData.Top,
		templateData.Bottom,
		templateData.Left,
		templateData.Right,
		templateData.Center,
	}
	// Set the area indexes.
	lastIndex := 0
	if len(templateData.Top.Area) > 0 {
		templateData.TopIndex = lastIndex
		lastIndex++
	}
	if len(templateData.Bottom.Area) > 0 {
		templateData.BottomIndex = lastIndex
		lastIndex++
	}
	if len(templateData.Left.Area) > 0 {
		templateData.LeftIndex = lastIndex
		lastIndex++
	}
	if len(templateData.Right.Area) > 0 {
		templateData.RightIndex = lastIndex
		lastIndex++
	}
	if len(templateData.Center.Area) > 0 {
		templateData.CenterIndex = lastIndex
	}
	return
}
