package presetting

const (
	APITemplate = `package presetting
{{- if .UsesLocalContent }}

import(
{{- if .Leading.IsLocal }}
	_leadingpanel_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/panels/LeadingPanel"
{{- end }}
{{- if .Trailing.IsLocal }}
	_trailingpanel_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/panels/TrailingPanel"
{{- end }}
)
{{- end }}

type Direction int

const (
	Horizontal Direction = iota
	Vertical
)

type Preset struct {
	Direction Direction
{{- if .Leading.IsLocal }}
	LeadingPanel *_leadingpanel_.Preset
{{- else }}
	{{ .Leading.ScreenName }}Screen *{{ .Leading.ScreenName }}ScreenPreset
{{- end }}
{{- if .Trailing.IsLocal }}
	TrailingPanel *_trailingpanel_.Preset
{{- else }}
	{{ .Trailing.ScreenName }}Screen *{{ .Trailing.ScreenName }}ScreenPreset
{{- end }}
}

// new constructs a new Preset.
func new(direction Direction, presets ...any) (screenPreset *Preset) {
	screenPreset = &Preset{
		Direction: direction,
	}
	for _, preset := range presets {
		switch preset := preset.(type) {
{{- if .Leading.IsLocal }}
		case *_leadingpanel_.Preset:
			screenPreset.LeadingPanel = preset
{{- else }}
		case *{{ .Leading.ScreenName }}ScreenPreset:
			screenPreset.{{ .Leading.ScreenName }}Screen = preset
{{- end }}
{{- if .Trailing.IsLocal }}
		case *_trailingpanel_.Preset:
			screenPreset.TrailingPanel = preset
{{- else }}
		case *{{ .Trailing.ScreenName }}ScreenPreset:
			screenPreset.{{ .Trailing.ScreenName }}Screen = preset
{{- end }}
		}
	}
	return
}

// See func Presets() (screenPreset map[string]any) in this package's api.go.
var Presets = map[string]*Preset {
	"Window":        newDefaultPreset(Vertical),
	"AccordionItem": newDefaultPreset(Vertical),
	"BorderCenter":  newDefaultPreset(Vertical),
	"SplitLeading":  newDefaultPreset(Horizontal),
	"SplitTrailing": newDefaultPreset(Horizontal),
	"TabItem":       newDefaultPreset(Vertical),
}
`
)
