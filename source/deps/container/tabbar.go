package container

import (
	_utils_ "github.com/josephabbudd-web/kickfyne/source/utils"
)

type tabbarTemplateData struct {
	PackageName   string
	Funcs         _utils_.Funcs
	AllPanelNames []string
}

const (
	tabbarTemplate = `{{ $DOT := . }}
package {{ call .Funcs.LowerCase .PackageName }}

import (
	"slices"
)

// Tab.
type Tab uint

const (
{{- range $i, $panelName := .AllPanelNames }}
 {{- if eq $i 0 }}
	{{ $panelName }}Item Tab = iota
 {{- else }}
	{{ $panelName }}Item
 {{- end }}
{{- end }}
)

type TabInfo struct {
	ID  string
	Tab Tab
}

var IDInfos = make(map[string][]TabInfo)

func AddTabbar(tabbarScreenID string) {
	if _, found := IDInfos[tabbarScreenID]; !found {
		IDInfos[tabbarScreenID] = make([]TabInfo, 0, 5)
	}
}

func AddTabItem(tabbarScreenID, tabItemPanelID string, tab Tab) {
	var tabInfos []TabInfo
	var found bool
	if tabInfos, found = IDInfos[tabbarScreenID]; !found {
		tabInfos = make([]TabInfo, 0, 5)
	}
	tabInfo := TabInfo{
		ID:  tabItemPanelID,
		Tab: tab,
	}
	IDInfos[tabbarScreenID] = append(tabInfos, tabInfo)
}

func RemoveTabbar(tabbarScreenID string) {
	delete(IDInfos, tabbarScreenID)
}

func RemoveTabItem(tabbarScreenID, tabItemPanelID string) {
	tabInfos := IDInfos[tabbarScreenID]
	for index, tabInfo := range tabInfos {
		if tabInfo.ID == tabItemPanelID {
			IDInfos[tabbarScreenID] = slices.Delete(tabInfos, index, index+1)
		}
	}
}
`
)
