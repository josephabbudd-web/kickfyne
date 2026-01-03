package presetting

const (
	RemotePresetTemplate = `package presetting

import(
{{- if eq false .Leading.IsLocal }}
	_{{ call .Funcs.LowerCase .Leading.ScreenName }}screenpresetting_ "{{ .ImportPrefix }}/frontend/screens/{{ .Leading.ScreenName }}/presetting"
{{- end }}
{{- if eq false .Trailing.IsLocal }}
	_{{ call .Funcs.LowerCase .Trailing.ScreenName }}screenpresetting_ "{{ .ImportPrefix }}/frontend/screens/{{ .Trailing.ScreenName }}/presetting"
{{- end }}
)

{{- if eq false .Leading.IsLocal }}

// The Leading Split area uses the {{ .Leading.ScreenName }} screen for content.
type {{ .Leading.ScreenName }}ScreenPreset struct {
	Preset *_{{ call .Funcs.LowerCase .Leading.ScreenName }}screenpresetting_.Preset
}
func new{{ .Leading.ScreenName }}ScreenPreset({{ call .Funcs.LowerCase .Leading.ScreenName }}ScreenPreset *_{{ call .Funcs.LowerCase .Leading.ScreenName }}screenpresetting_.Preset) (preset *{{ .Leading.ScreenName }}ScreenPreset) {
	preset = &{{ .Leading.ScreenName }}ScreenPreset{
		Preset: {{ call .Funcs.LowerCase .Leading.ScreenName }}ScreenPreset,
	}
	return
}
{{- end }}

{{- if eq false .Trailing.IsLocal }}{{ if eq true .Leading.IsLocal}}

{{ end -}}
// The Trailing Split area uses the {{ .Trailing.ScreenName }} screen for content.
type {{ .Trailing.ScreenName }}ScreenPreset struct {
	Preset *_{{ call .Funcs.LowerCase .Trailing.ScreenName }}screenpresetting_.Preset
}
func new{{ .Trailing.ScreenName }}ScreenPreset({{ call .Funcs.LowerCase .Trailing.ScreenName }}ScreenPreset *_{{ call .Funcs.LowerCase .Trailing.ScreenName }}screenpresetting_.Preset) (preset *{{ .Trailing.ScreenName }}ScreenPreset) {
	preset = &{{ .Trailing.ScreenName }}ScreenPreset{
		Preset: {{ call .Funcs.LowerCase .Trailing.ScreenName }}ScreenPreset,
	}
	return
}
{{- end }}
`

	NoRemotePresetTemplate = `package presetting

// There are no border areas using content from other screens so there are no remote presets here.
`
)
