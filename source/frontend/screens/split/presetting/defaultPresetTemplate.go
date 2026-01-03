package presetting

const (
	DefaultPresetTemplate = `package presetting

import(
{{- if .Leading.IsLocal }}
	_leadingpanel_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/panels/LeadingPanel"
{{- else }}
	_{{ call .Funcs.LowerCase .Leading.ScreenName }}screenpresetting_ "{{ .ImportPrefix }}/frontend/screens/{{ .Leading.ScreenName }}/presetting"
{{- end }}
{{- if .Trailing.IsLocal }}
	_trailingpanel_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/panels/TrailingPanel"
{{- else }}
	_{{ call .Funcs.LowerCase .Trailing.ScreenName }}screenpresetting_ "{{ .ImportPrefix }}/frontend/screens/{{ .Trailing.ScreenName }}/presetting"
{{- end }}
)

/* KICKFYNE TODO:
To add a custom preset
- Add a custom preset in each panels/«panel-name»Panel/preset.go.
- You could also add a custom preset in screens that are used for accordion item content.
- Copy this file and use it as a template for the new preset constructor.
- Update api.go var Presets.
*/

// newDefaultPreset() constructs the default Preset.
func newDefaultPreset(direction Direction) (preset *Preset) {
	preset = new(
		direction,
{{- if .Leading.IsLocal }}
		_leadingpanel_.NewDefaultPreset(),
{{- else }}
		new{{ .Leading.ScreenName }}ScreenPreset(_{{ call .Funcs.LowerCase .Leading.ScreenName }}screenpresetting_.Presets["SplitLeading"]),
{{- end }}
{{- if .Trailing.IsLocal }}
		_trailingpanel_.NewDefaultPreset(),
{{- else }}
		new{{ .Trailing.ScreenName }}ScreenPreset(_{{ call .Funcs.LowerCase .Trailing.ScreenName }}screenpresetting_.Presets["SplitTrailing"]),
{{- end }}
	)
	return
}
`
)
