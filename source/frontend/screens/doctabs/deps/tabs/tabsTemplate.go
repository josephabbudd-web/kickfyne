package tabs

import (
	_utils_ "github.com/josephabbudd-web/kickfyne/source/utils"
)

type TemplateData struct {
	PackageName      string
	LocalPanelNames  []string
	RemotePanelNames []string
	ImportPrefix     string
	Funcs            _utils_.Funcs
}

const (
	FileName = "tabs.go"

	NoBETemplate = `{{ $DOT := . -}}
package tabs

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

{{ if ne (len .RemotePanelNames) 0 }}
	_screenmap_ "{{ .ImportPrefix }}/frontend/screenmap"
{{ end }}
	_misc_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/misc"
	_panels_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/panels"
	_presets_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/presets"
	_types_ "{{ .ImportPrefix }}/frontend/types"
	_ids_ "{{ .ImportPrefix }}/deps/container/{{ .PackageName }}"

{{- range $i, $panelName := .RemotePanelNames }}
 {{- if eq $i 0 }}

	_{{ call $DOT.Funcs.LowerCase $panelName }}startup_ "{{ $DOT.ImportPrefix }}/frontend/screens/{{ $panelName }}/presets"
 {{- else }}
	_{{ call $DOT.Funcs.LowerCase $panelName }}startup_ "{{ $DOT.ImportPrefix }}/frontend/screens/{{ $panelName }}/presets"
 {{- end }}
{{- end }}
)

func SetInitialTabItems(screen *_misc_.Miscellaneous, screenInitData *_presets_.ScreenInitData) (err error) {
{{- range $panelName := .LocalPanelNames}}
	if err = Open{{ $panelName }}Tab(screen, screenInitData.{{ $panelName }}Panel.TabItemIcon, screenInitData.{{ $panelName }}Panel.TabItemLabel, screenInitData.{{ $panelName }}Panel); err != nil {
		return
	}
{{- end }}

{{ range $panelName := .RemotePanelNames}}
	if err = Open{{ $panelName }}Tab(screen, screenInitData.{{ $panelName }}Screen.TabItemIcon, screenInitData.{{ $panelName }}Screen.TabItemLabel, screenInitData.{{ $panelName }}Screen.ScreenInitData); err != nil {
		return
	}
{{- end }}
	return
}

{{- range $panelName := .LocalPanelNames}}

// Open{{ $panelName }}Tab constructs a {{ $panelName }}TabItem.
// The {{ $panelName }}Tab uses the local {{ $panelName }} panel for content.
func Open{{ $panelName }}Tab(screen *_misc_.Miscellaneous, tabIcon fyne.Resource, tabLabel string, startupData *_presets_.{{ $panelName }}PanelInitData) (err error) {
	tabItem := container.NewTabItemWithIcon(tabLabel, tabIcon, widget.NewLabel("This is the {{ $panelName }} panel."))
	tabItemContentConsumer := _types_.NewDocTabsTabItemContentConsumer(screen.Layout.Tabbar(), tabItem)
	// tabItemContentConsumer := _types_.NewDocTabsTabItemContentConsumer(screen.Layout.TabbarConsumer(), screen.Layout.Tabbar(), tabItem, spawned)
	// The {{ $panelName }} panel.
	var panel *_panels_.{{ $panelName }}Panel
	if panel, err = _panels_.New{{ $panelName }}Panel(screen, tabItemContentConsumer, tabItem, startupData); err != nil {
		return
	}
	panelProducer := panel.Producer()
	tabItemContentConsumer.Bind(panelProducer)
	// Add the tab to the layout.
	screen.Layout.AddPanelerTabItemConsumer(panel, tabItem, tabItemContentConsumer)
	// Update deps.
	_ids_.AddTabItem(screen.ScreenID, panel.ID(), _ids_.{{ $panelName }}Item)
	return
}
{{- end }}

{{- range $panelName := .RemotePanelNames}}

// Open{{ $panelName }}Tab constructs a {{ $panelName }}TabItem.
// The {{ $panelName }}Tab uses the {{ $panelName }} screen for content.
func Open{{ $panelName }}Tab(screen *_misc_.Miscellaneous, tabIcon fyne.Resource, tabLabel string, {{ call $DOT.Funcs.DeCap $panelName }}ScreenInitData *_{{ call $DOT.Funcs.LowerCase $panelName }}startup_.ScreenInitData) (err error) {
	tabItem := container.NewTabItemWithIcon(tabLabel, tabIcon, widget.NewLabel("This is the {{ $panelName }} panel."))
	api := _screenmap_.Map["{{ $panelName }}"]
	var docTabsTabItemContentConsumer *_types_.DocTabsTabItemContentConsumer
	var {{ call $DOT.Funcs.DeCap $panelName }}ScreenID string

	if docTabsTabItemContentConsumer, {{ call $DOT.Funcs.DeCap $panelName }}ScreenID, err = api.NewDocTabsTabItemContentConsumer(
		screen.CTX,
		screen.CTXCancel,
		screen.APP,
		screen.Window,
		screen.Layout.Tabbar(),
		tabItem,
		{{ call $DOT.Funcs.DeCap $panelName }}ScreenInitData,
	); err != nil {
		return
	}
	screen.Layout.AddRemoteScreenIDTabItemConsumer({{ call $DOT.Funcs.DeCap $panelName }}ScreenID, tabItem, docTabsTabItemContentConsumer)
	// Update deps.
	_ids_.AddTabbar(screen.ScreenID)
	_ids_.AddTabItem(screen.ScreenID, {{ call $DOT.Funcs.DeCap $panelName }}ScreenID, _ids_.{{ $panelName }}Item)
	return
}
{{- end }}

func CloseTabItem(screen *_misc_.Miscellaneous, tabItem *container.TabItem) {
	if tabItemID, found := screen.Layout.TabID(tabItem); found {
		CloseTabID(screen, tabItemID)
	}
}

func CloseTabID(screen *_misc_.Miscellaneous, tabItemID string) {
	screen.Layout.RemoveID(tabItemID)
	// Update deps.tabItemID
	_ids_.RemoveTabItem(screen.ScreenID, tabItemID)
}
`
)
