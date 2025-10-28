package presetting

import (
	_utils_ "github.com/josephabbudd-web/kickfyne/source/utils"
)

type Panel struct {
	Name    string
	IsLocal bool
}
type TemplateData struct {
	PackageName      string
	ImportPrefix     string
	LocalPanelNames  []string
	RemotePanelNames []string
	Funcs            _utils_.Funcs
}

const (
	APITemplate = `{{ $DOT := . -}}
package presetting
{{- if ne (len .LocalPanelNames) 0 }}

import(
{{- range $panelName := .LocalPanelNames }}
	_{{ call $DOT.Funcs.LowerCase $panelName }}panel_ "{{ $DOT.ImportPrefix }}/frontend/screens/{{ $DOT.PackageName }}/panels/{{ $panelName }}Panel"
{{ end }}
)
{{- end }}

type Preset struct {
{{- range $i, $panelName := .LocalPanelNames }}
	{{ $panelName }}Panel *_{{ call $DOT.Funcs.LowerCase $panelName }}panel_.Preset
{{- end }}
{{- range $i, $panelName := .RemotePanelNames }}
	{{ $panelName }}Screen *{{ $panelName }}ScreenPreset
{{- end }}
}

// new constructs a new Preset.
func new(presets ...any) (screenPreset *Preset) {
	screenPreset = &Preset{}
	for _, preset := range presets {
		switch preset := preset.(type) {
{{- range $i, $panelName := .LocalPanelNames }}
		case *_{{ call $DOT.Funcs.LowerCase $panelName }}panel_.Preset:
			screenPreset.{{ $panelName }}Panel = preset
{{- end }}
{{- range $i, $panelName := .RemotePanelNames }}
		case *{{ $panelName }}ScreenPreset:
			screenPreset.{{ $panelName }}Screen = preset
{{- end }}
		}
	}
	return
}

// See func Presets() (screenPreset map[string]any) in this package's api.go.
var Presets = map[string]*Preset {
	"Default" : newDefaultPreset(),
}
`
)
