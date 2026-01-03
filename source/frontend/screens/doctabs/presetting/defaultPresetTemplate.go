package presetting

const (
	DefaultPresetTemplate = `{{ $DOT := . -}}
package presetting

import(
{{- range $panelName := .LocalPanelNames }}
	_{{ call $DOT.Funcs.LowerCase $panelName }}panel_ "{{ $DOT.ImportPrefix }}/frontend/screens/{{ $DOT.PackageName }}/panels/{{ $panelName }}Panel"
{{- end }}
{{- if ne 0 (len .RemotePanelNames) }}

{{ end }}
{{- range $panelName := .RemotePanelNames }}
	_{{ call $DOT.Funcs.LowerCase $panelName }}screenpresetting_ "{{ $DOT.ImportPrefix }}/frontend/screens/{{ $panelName }}/presetting"
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
{{- range $panelName := .LocalPanelNames }}
		_{{ call $DOT.Funcs.LowerCase $panelName }}panel_.NewDefaultPreset(),
{{- end }}
{{- range $panelName := .RemotePanelNames }}
		new{{ $panelName }}ScreenPreset(
			"", // tab icon
			"{{ $panelName }}", // tab label
			_{{ call $DOT.Funcs.LowerCase $panelName }}screenpresetting_.Presets["TabItem"],
		),
{{- end }}
	)
	return
}
`
)
