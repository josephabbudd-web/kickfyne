package presetting

const (
	RemotePresetTemplate = `{{ $DOT := . -}}
package presetting

/* KICKFYNE TODO:
In each panels/«panel-name»Panel/preset.go.
 - Customize the func NewDefaultPreset().
 - Add any new customized preset funcs.
In this file.
 - Add a preset funcs corresponding to the custom panel presets and the custom screen presets.
 - Update api.go var Presets.
*/

import(
{{ range $panelName := .RemotePanelNames }}
	_{{ call $DOT.Funcs.LowerCase $panelName }}screenpresetting_ "{{ $DOT.ImportPrefix }}/frontend/screens/{{ $panelName }}/presetting"
{{- end }}
)

{{- range $i, $panelName := .RemotePanelNames }}

type {{ $panelName }}ScreenPreset struct {
 	// Remote {{ $panelName }} screen preset data.
	AccordionItemTitle string
	Preset *_{{ call $DOT.Funcs.LowerCase $panelName }}screenpresetting_.Preset
}

func new{{ $panelName }}ScreenPreset(accordionItemTitle string, {{ call $DOT.Funcs.LowerCase $panelName }}Preset *_{{ call $DOT.Funcs.LowerCase $panelName }}screenpresetting_.Preset) (preset *{{ $panelName }}ScreenPreset) {
	preset = &{{ $panelName }}ScreenPreset{
		AccordionItemTitle: accordionItemTitle,
		Preset:             {{ call $DOT.Funcs.LowerCase $panelName }}Preset,
	}
	return
}
{{- end }}
`

	NoRemotePresetTemplate = `{{ $DOT := . -}}
package presetting

// This package has no remote DocTabs so there are no remote presets here.
`
)
