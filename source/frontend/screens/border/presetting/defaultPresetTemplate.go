package presetting

const (
	DefaultPresetTemplate = `{{ $DOT := . -}}
package presetting

import(
{{- range $area := .Areas }}
 {{- if eq (len $area.Area) 0 }}{{ continue }}{{ end }}
 {{- if eq $area.IsLocal false }}{{ continue }}{{ end }}
	_{{ call $DOT.Funcs.LowerCase $area.ItemName }}panel_ "{{ $DOT.ImportPrefix }}/frontend/screens/{{ $DOT.PackageName }}/panels/{{ $area.ItemName }}Panel"
{{- end }}
{{- range $remoteItemNames := .UniqueRemoteItemNames }}
	_{{ call $DOT.Funcs.LowerCase $remoteItemNames }}screenpresetting_ "{{ $DOT.ImportPrefix }}/frontend/screens/{{ $remoteItemNames }}/presetting"
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
func newDefaultPreset() (preset *Preset) {
	preset = new(
{{- range $area := .Areas }}
 {{- if eq (len $area.Area) 0 }}{{ continue }}{{ end }}
 {{- if eq $area.IsLocal false }}{{ continue }}{{ end }}
		_{{ call $DOT.Funcs.LowerCase $area.ItemName }}panel_.NewDefaultPreset(),
{{- end }}
{{- range $remoteItemNames := .UniqueRemoteItemNames }}
		new{{ $remoteItemNames }}ScreenPreset(_{{ call $DOT.Funcs.LowerCase $remoteItemNames }}screenpresetting_.Presets["BorderCenter"]),
{{- end }}
	)
	return
}
`
)
