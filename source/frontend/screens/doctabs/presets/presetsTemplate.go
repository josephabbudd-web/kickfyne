package startup

import (
	_utils_ "github.com/josephabbudd-web/kickfyne/source/utils"
)

type Panel struct {
	Name    string
	IsLocal bool
}
type PresetsTemplateData struct {
	ImportPrefix     string
	AllPanels        []Panel
	LocalPanelNames  []string
	RemotePanelNames []string
	Funcs            _utils_.Funcs
}

const (
	PresetsTemplate = `{{ $DOT := . -}}
package startup

/* KICKFYNE TODO:
	Cusomize each panel's init data struct.
	Customize func DefaultScreenInitData().
	Add any required custom ScreenInitData() funcs for displaying different content.
	Cusomize func PresetScreenInitDataByName(name string).
*/

import(
	"fyne.io/fyne/v2"
	// "fyne.io/fyne/v2/theme"
{{ range $i, $panelName := .RemotePanelNames }}
 {{- if eq $i 0 }}

 	_{{ call $DOT.Funcs.LowerCase $panelName }}startup_ "{{ $DOT.ImportPrefix }}/frontend/screens/{{ $panelName }}/presets"
 {{- else }}
 	_{{ call $DOT.Funcs.LowerCase $panelName }}startup_ "{{ $DOT.ImportPrefix }}/frontend/screens/{{ $panelName }}/presets"
 {{- end }}
{{- end }}
)

type ScreenInitData struct {
{{- range $i, $panelName := .LocalPanelNames }}
	{{ $panelName }}Panel *{{ $panelName }}PanelInitData
{{- end }}
{{- range $i, $panelName := .RemotePanelNames }}
	{{ $panelName }}Screen *{{ $panelName }}ScreenInitData
{{- end }}
}

// Param anyInitData should be nil if there are not initData for panels or screens.
func NewScreenInitData(anyInitData ...any) (screenInitData *ScreenInitData) {
	screenInitData = &ScreenInitData{}
	for _, anyInitDatum := range anyInitData {
		if anyInitDatum == nil {
			return
		}
		switch anyInitDatum := anyInitDatum.(type) {
{{- range $i, $panelName := .LocalPanelNames }}
		case *{{ $panelName }}PanelInitData:
			screenInitData.{{ $panelName }}Panel = anyInitDatum
{{- end }}
{{- range $i, $panelName := .RemotePanelNames }}
		case *{{ $panelName }}ScreenInitData:
			screenInitData.{{ $panelName }}Screen = anyInitDatum
{{- end }}
		}
	}
	return
}

// See func Presets() (presets map[string]any) in this package's api.go.
var Presets = map[string]*ScreenInitData {
	"Default" : DefaultScreenInitData(),
}

// DefaultScreenInitData() constructs the default default ScreenInitData.
func DefaultScreenInitData() (screenInitData *ScreenInitData) {
	screenInitData = NewScreenInitData(
{{- range $i, $panel := .AllPanels }}
 {{- if $panel.IsLocal }}
		New{{ $panel.Name }}PanelInitData(
			nil, // theme.Icon(theme.IconNameContentAdd) 
			"{{ $panel.Name }}", // tabItemLabel string
			"This is the {{ $panel.Name }} panel heading.", // heading string
			"This is the {{ $panel.Name }} panel description.", // description string
		),
 {{- else }}
		New{{ $panel.Name }}ScreenInitData(
			nil, // theme.Icon(theme.IconNameContentAdd) 
			"{{ $panel.Name }}", // tabItemLabel string
			_{{ call $DOT.Funcs.LowerCase $panel.Name }}startup_.DefaultScreenInitData(),
		),
 {{- end }}
{{- end }}
	)
	return
}

{{- range $i, $panelName := .LocalPanelNames }}

type {{ $panelName }}PanelInitData struct {
	// Local {{ $panelName }}Panel startup data.
	TabItemIcon fyne.Resource
	TabItemLabel string
	Heading string
	Description string
}

func New{{ $panelName }}PanelInitData(tabItemIcon fyne.Resource, tabItemLabel string, heading string, description string) (new{{ $panelName }}PanelInitData *{{ $panelName }}PanelInitData) {
	new{{ $panelName }}PanelInitData = &{{ $panelName }}PanelInitData{
		TabItemIcon: tabItemIcon,
		TabItemLabel: tabItemLabel,
		Heading: heading,
		Description: description,
	}
	return
}
{{- end }}

{{- range $i, $panelName := .RemotePanelNames }}

type {{ $panelName }}ScreenInitData struct {
 	// Remote {{ $panelName }} screen startup data.
	TabItemIcon fyne.Resource
	TabItemLabel string
	ScreenInitData *_{{ call $DOT.Funcs.LowerCase $panelName }}startup_.ScreenInitData
}

func New{{ $panelName }}ScreenInitData(tabItemIcon fyne.Resource, tabItemLabel string, {{ call $DOT.Funcs.LowerCase $panelName }}ScreenInitData *_{{ call $DOT.Funcs.LowerCase $panelName }}startup_.ScreenInitData) (new{{ $panelName }}ScreenInitData *{{ $panelName }}ScreenInitData) {
	new{{ $panelName }}ScreenInitData = &{{ $panelName }}ScreenInitData{
		TabItemIcon: tabItemIcon,
		TabItemLabel: tabItemLabel,
		ScreenInitData: {{ call $DOT.Funcs.LowerCase $panelName }}ScreenInitData,
	}
	return
}
{{- end }}
`
)
