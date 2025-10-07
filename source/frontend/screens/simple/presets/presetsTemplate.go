package startup

import (
	_utils_ "github.com/josephabbudd-web/kickfyne/source/utils"
)

type PresetsTemplateData struct {
	ImportPrefix    string
	LocalPanelNames []string
	Funcs           _utils_.Funcs
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

type ScreenInitData struct {
{{- range $i, $panelName := .LocalPanelNames }}
	{{ $panelName }}Panel *{{ $panelName }}PanelInitData
{{ end }}
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
{{- range $i, $panelName := .LocalPanelNames }}
		New{{ $panelName }}PanelInitData(
			"This is the {{ $panelName }} panel heading.", // heading string
			"This is the {{ $panelName }} panel description.", // description string
		),
{{- end }}
	)
	return
}

{{- range $i, $panelName := .LocalPanelNames }}

type {{ $panelName }}PanelInitData struct {
	// Local {{ $panelName }}Panel startup data.
	Heading string
	Description string
}

func New{{ $panelName }}PanelInitData(heading string, description string) (new{{ $panelName }}PanelInitData *{{ $panelName }}PanelInitData) {
	new{{ $panelName }}PanelInitData = &{{ $panelName }}PanelInitData{
		Heading: heading,
		Description: description,
	}
	return
}
{{- end }}

`
)
