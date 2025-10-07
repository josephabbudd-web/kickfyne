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
	Cusomize the map PresetScreenInitData.
*/

{{- if ne (len .RemotePanelNames) 0 }}
import(
{{ range $panelName := .RemotePanelNames }}
	_{{ call $DOT.Funcs.LowerCase $panelName }}startup_ "{{ $DOT.ImportPrefix }}/frontend/screens/{{ $panelName }}/presets"
{{ end }}
)
{{- end }}

type ScreenInitData struct {
{{- range $i, $panelName := .LocalPanelNames }}
	{{ $panelName }}Panel *{{ $panelName }}PanelInitData
{{- end }}
{{- range $i, $panelName := .RemotePanelNames }}
	{{ $panelName }}Screen *{{ $panelName }}ScreenInitData
{{- end }}
}

// NewScreenInitData constructs a new ScreenInitData.
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
{{- range $panel := .AllPanels }}
 {{- if $panel.IsLocal }}
		New{{ $panel.Name }}PanelInitData(
			"{{ $panel.Name }}", // accordionItemTitle string
			"This is the {{ $panel.Name }} panel heading.", // heading string
			"This is the {{ $panel.Name }} panel description.", // description string
		),
 {{- else }}
		New{{ $panel.Name }}ScreenInitData(
			"{{ $panel.Name }}", // accordionItemTitle string
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
	AccordionItemTitle string
	Heading string
	Description string
}

func New{{ $panelName }}PanelInitData(accordionItemTitle string, heading string, description string) (new{{ $panelName }}PanelInitData *{{ $panelName }}PanelInitData) {
	new{{ $panelName }}PanelInitData = &{{ $panelName }}PanelInitData{
		AccordionItemTitle: accordionItemTitle,
		Heading: heading,
		Description: description,
	}
	return
}
{{- end }}

{{- range $i, $panelName := .RemotePanelNames }}

type {{ $panelName }}ScreenInitData struct {
 	// Remote {{ $panelName }} screen startup data.
	AccordionItemTitle string
	ScreenInitData *_{{ call $DOT.Funcs.LowerCase $panelName }}startup_.ScreenInitData
}

func New{{ $panelName }}ScreenInitData(accordionItemTitle string, {{ call $DOT.Funcs.LowerCase $panelName }}ScreenInitData *_{{ call $DOT.Funcs.LowerCase $panelName }}startup_.ScreenInitData) (new{{ $panelName }}ScreenInitData *{{ $panelName }}ScreenInitData) {
	new{{ $panelName }}ScreenInitData = &{{ $panelName }}ScreenInitData{
		AccordionItemTitle: accordionItemTitle,
		ScreenInitData: {{ call $DOT.Funcs.LowerCase $panelName }}ScreenInitData,
	}
	return
}
{{- end }}
`
)
