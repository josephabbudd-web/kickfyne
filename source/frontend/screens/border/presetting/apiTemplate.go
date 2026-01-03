package presetting

const (
	APITemplate = `{{ $DOT := . -}}
package presetting
{{- if .UsesLocalContent }}

import(
{{- range $area := .Areas }}
 {{- if eq (len $area.Area) 0 }}{{ continue }}{{ end }}
 {{- if eq $area.IsLocal false }}{{ continue }}{{ end }}
	_{{ call $DOT.Funcs.LowerCase $area.ItemName }}panel_ "{{ $DOT.ImportPrefix }}/frontend/screens/{{ $DOT.PackageName }}/panels/{{ $area.ItemName }}Panel"
{{ end }}
)
{{- end }}

type Preset struct {
{{- range $area := .Areas }}
 {{- if eq (len $area.Area) 0 }}{{ continue }}{{ end }}
 {{- if eq $area.IsLocal false }}{{ continue }}{{ end }}
	{{ $area.ItemName }}Panel *_{{ call $DOT.Funcs.LowerCase $area.ItemName }}panel_.Preset
{{- end }}
{{- range $remoteItemName := .UniqueRemoteItemNames }}
	{{ $remoteItemName }}Screen *{{ $remoteItemName }}ScreenPreset
{{- end }}
}

// new constructs a new Preset.
func new(presets ...any) (screenPreset *Preset) {
	screenPreset = &Preset{}
	for _, preset := range presets {
		switch preset := preset.(type) {
{{- range $area := .Areas }}
 {{- if eq (len $area.Area) 0 }}{{ continue }}{{ end }}
 {{- if eq $area.IsLocal false }}{{ continue }}{{ end }}
		case *_{{ call $DOT.Funcs.LowerCase $area.ItemName }}panel_.Preset:
			screenPreset.{{ $area.ItemName }}Panel = preset
{{- end }}
{{- range $remoteItemName := .UniqueRemoteItemNames }}
		case *{{ $remoteItemName }}ScreenPreset:
			screenPreset.{{ $remoteItemName }}Screen = preset
{{- end }}
		}
	}
	return
}

// See func Presets() (screenPreset map[string]any) in this package's api.go.
var Presets = map[string]*Preset {
	"Window":        newDefaultPreset(),
	"AccordionItem": newDefaultPreset(),
	"BorderCenter":  newDefaultPreset(),
	"SplitLeading":  newDefaultPreset(),
	"SplitTrailing": newDefaultPreset(),
	"TabItem":       newDefaultPreset(),
}
`
)
