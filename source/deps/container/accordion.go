package container

import (
	_utils_ "github.com/josephabbudd-web/kickfyne/source/utils"
)

type accordionTemplateData struct {
	PackageName   string
	Funcs         _utils_.Funcs
	AllPanelNames []string
}

const (
	accordionTemplate = `
package {{ call .Funcs.LowerCase .PackageName }}

import (
	"slices"
)

// AccordionItem.
type AccordionItem uint

const (
{{- range $i, $panelName := .AllPanelNames }}
 {{- if eq $i 0 }}
	{{ $panelName }}Item AccordionItem = iota
 {{- else }}
	{{ $panelName }}Item
 {{- end }}
{{- end }}
)

type AccordionItemInfo struct {
	ID            string
	AccordionItem AccordionItem
}

var IDInfos = make(map[string][]AccordionItemInfo)

func AddAccordion(accordionScreenID string) {
	if _, found := IDInfos[accordionScreenID]; !found {
		IDInfos[accordionScreenID] = make([]AccordionItemInfo, 0, 5)
	}
}

func AddAccordionItem(accordionScreenID, accordionItemPanelID string, accordionItem AccordionItem) {
	var accordionItemInfos []AccordionItemInfo
	var found bool
	if accordionItemInfos, found = IDInfos[accordionScreenID]; !found {
		accordionItemInfos = make([]AccordionItemInfo, 0, 5)
	}
	accordionItemInfo := AccordionItemInfo{
		ID:            accordionItemPanelID,
		AccordionItem: accordionItem,
	}
	IDInfos[accordionScreenID] = append(accordionItemInfos, accordionItemInfo)
}

func RemoveAccordion(accordionScreenID string) {
	delete(IDInfos, accordionScreenID)
}

func RemoveAccordionItem(accordionScreenID, accordionItemPanelID string) {
	accordionItemInfos := IDInfos[accordionScreenID]
	for index, accordionItemInfo := range accordionItemInfos {
		if accordionItemInfo.ID == accordionItemPanelID {
			IDInfos[accordionScreenID] = slices.Delete(accordionItemInfos, index, index+1)
		}
	}
}
`
)
