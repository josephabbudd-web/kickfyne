package tabs

import (
	_utils_ "github.com/josephabbudd-web/kickfyne/source/utils"
)

type TemplateData struct {
	PackageName      string
	LocalPanelNames  []string
	RemotePanelNames []string
	ImportPrefix     string
	UseConfigTab     bool
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

{{- range $i, $panelName := .LocalPanelNames }}
 {{- if eq $i 0 }}

	_{{ call $DOT.Funcs.LowerCase $panelName }}panel_ "{{ $DOT.ImportPrefix }}/frontend/screens/{{ $DOT.PackageName }}/panels/{{ $panelName }}Panel"
 {{- else }}
	_{{ call $DOT.Funcs.LowerCase $panelName }}panel_ "{{ $DOT.ImportPrefix }}/frontend/screens/{{ $DOT.PackageName }}/panels/{{ $panelName }}Panel"
 {{- end }}
{{- end }}

{{ if ne (len .RemotePanelNames) 0 }}
	_screenmap_ "{{ .ImportPrefix }}/frontend/deps/screenmap"
{{ end }}
	_misc_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/misc"
	_panels_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/panels"
	_presetting_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/presetting"
	_types_ "{{ .ImportPrefix }}/frontend/deps/types"
)

func SetInitialTabItems(screen *_misc_.Miscellaneous, screenPreset *_presetting_.Preset) (err error) {
{{- range $panelName := .LocalPanelNames}}
	if err = Open{{ $panelName }}TabItem(screen, screenPreset.{{ $panelName }}Panel); err != nil {
		return
	}
{{- end }}

{{ range $panelName := .RemotePanelNames}}
	if err = Open{{ $panelName }}TabItem(screen, screenPreset.{{ $panelName }}Screen); err != nil {
		return
	}
{{- end }}
	return
}

{{- range $i, $panelName := .LocalPanelNames}}

// Open{{ $panelName }}TabItem constructs a {{ $panelName }}TabItem.
// The {{ $panelName }}Tab uses the local {{ $panelName }} panel for content.
func Open{{ $panelName }}TabItem(screen *_misc_.Miscellaneous, preset *_{{ call $DOT.Funcs.LowerCase $panelName }}panel_.Preset) (err error) {
	var icon fyne.Resource
	if len(preset.TabItemIconName) > 0 {
		icon = screen.APP.Settings().Theme().Icon(preset.TabItemIconName)
	}
	tabItem := container.NewTabItemWithIcon(preset.TabItemLabel, icon, widget.NewLabel("This is the {{ $panelName }} panel."))
	tabItemContentConsumer := _types_.NewDocTabsTabItemContentConsumer(screen.Layout.Tabbar(), tabItem)
	// tabItemContentConsumer := _types_.NewDocTabsTabItemContentConsumer(screen.Layout.TabbarConsumer(), screen.Layout.Tabbar(), tabItem, spawned)
	// The {{ $panelName }} panel.
	var panel *_panels_.{{ $panelName }}Panel
	if panel, err = _panels_.New{{ $panelName }}Panel(screen, tabItemContentConsumer, tabItem, preset); err != nil {
		return
	}
	panelProducer := panel.Producer()
	tabItemContentConsumer.Bind(panelProducer)
	// Add the tab to the layout.
 {{- if and (eq $i 0) $DOT.UseConfigTab }}
	screen.Layout.AddConfigurationTabItemConsumer(panel, tabItem, tabItemContentConsumer)
 {{- else }}
	screen.Layout.AddPanelerTabItemConsumer(panel, tabItem, tabItemContentConsumer)
 {{- end }}
	return
}
{{- end }}

{{- range $panelName := .RemotePanelNames}}

// Open{{ $panelName }}TabItem constructs a {{ $panelName }}TabItem.
// The {{ $panelName }}Tab uses the {{ $panelName }} screen for content.
func Open{{ $panelName }}TabItem(screen *_misc_.Miscellaneous, preset *_presetting_.{{ $panelName }}ScreenPreset) (err error) {
	var icon fyne.Resource
	if len(preset.TabItemIconName) > 0 {
		icon = screen.APP.Settings().Theme().Icon(preset.TabItemIconName)
	}
	tabItem := container.NewTabItemWithIcon(preset.TabItemLabel, icon, widget.NewLabel("This is the {{ $panelName }} panel."))
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
		preset.Preset,
	); err != nil {
		return
	}
	screen.Layout.AddRemoteScreenIDTabItemConsumer({{ call $DOT.Funcs.DeCap $panelName }}ScreenID, tabItem, docTabsTabItemContentConsumer)
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
}
`
)
