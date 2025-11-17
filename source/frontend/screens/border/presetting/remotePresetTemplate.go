package presetting

const (
	RemotePresetTemplate = `{{ $DOT := . -}}
package presetting

import(
{{- range $remoteItemName := .UniqueRemoteItemNames }}
	_{{ call $DOT.Funcs.LowerCase $remoteItemName }}screenpresetting_ "{{ $DOT.ImportPrefix }}/frontend/screens/{{ $remoteItemName }}/presetting"
{{- end }}
)

{{ range $remoteItemName := .UniqueRemoteItemNames }}
 {{- range $area := $DOT.Areas }}
  {{- if eq $area.IsLocal true }}{{ continue }}{{ end }}
  {{- if eq $area.ItemName $remoteItemName }}
// The {{ $area.Area }} border area uses the {{ $remoteItemName }} screen for content.
  {{- end }}
 {{- end }}

type {{ $remoteItemName }}ScreenPreset struct {
	Preset *_{{ call $DOT.Funcs.LowerCase $remoteItemName }}screenpresetting_.Preset
}
func new{{ $remoteItemName }}ScreenPreset({{ call $DOT.Funcs.LowerCase $remoteItemName }}ScreenPreset *_{{ call $DOT.Funcs.LowerCase $remoteItemName }}screenpresetting_.Preset) (preset *{{ $remoteItemName }}ScreenPreset) {
	preset = &{{ $remoteItemName }}ScreenPreset{
		Preset: {{ call $DOT.Funcs.LowerCase $remoteItemName }}ScreenPreset,
	}
	return
}
{{- end }}
`

	NoRemotePresetTemplate = `package presetting

// There are no border areas using content from other screens so there are no remote presets here.
`
)
