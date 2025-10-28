package presetting

const (
	DefaultPresetTemplate = `{{ $DOT := . -}}
package presetting
{{- if ne (len .LocalPanelNames) 0 }}

import(
{{- range $panelName := .LocalPanelNames }}
	_{{ call $DOT.Funcs.LowerCase $panelName }}panel_ "{{ $DOT.ImportPrefix }}/frontend/screens/{{ $DOT.PackageName }}/panels/{{ $panelName }}Panel"
{{- end }}
)
{{- end }}

/* KICKFYNE TODO:
In each panels/«panel-name»Panel/preset.go.
 - Customize the func NewDefaultPreset().
 - Add any new customized preset funcs.
In a new file.
 - Name the new file like this: "imagePreset.go"
 - Use this file as a template for the new file.
 - Add a preset funcs corresponding to the custom panel presets and the custom screen presets.
 - Update api.go var Presets.
*/

// newDefaultPreset() constructs the default Preset.
func newDefaultPreset() (screenPreset *Preset) {
	screenPreset = new(
{{- range $panelName := .LocalPanelNames }}
		_{{ call $DOT.Funcs.LowerCase $panelName }}panel_.NewDefaultPreset(),
{{- end }}
	)
	return
}
`
)
