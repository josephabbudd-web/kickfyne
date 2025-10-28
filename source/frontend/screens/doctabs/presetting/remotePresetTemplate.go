package presetting

const (
	RemotePresetTemplate = `{{ $DOT := . -}}
package presetting

import(
	"fyne.io/fyne/v2"

{{ range $panelName := .RemotePanelNames }}
	_{{ call $DOT.Funcs.LowerCase $panelName }}screenpresetting_ "{{ $DOT.ImportPrefix }}/frontend/screens/{{ $panelName }}/presetting"
{{- end }}
)

{{- range $i, $panelName := .RemotePanelNames }}

type {{ $panelName }}ScreenPreset struct {
 	// Remote {{ $panelName }} screen preset data.
	TabItemIconName fyne.ThemeIconName
	TabItemLabel    string
	Preset          *_{{ call $DOT.Funcs.LowerCase $panelName }}screenpresetting_.Preset
}

func new{{ $panelName }}ScreenPreset(tabItemIconName fyne.ThemeIconName, tabItemLabel string, {{ call $DOT.Funcs.LowerCase $panelName }}ScreenPreset *_{{ call $DOT.Funcs.LowerCase $panelName }}screenpresetting_.Preset) (preset *{{ $panelName }}ScreenPreset) {
	preset = &{{ $panelName }}ScreenPreset{
		TabItemIconName: tabItemIconName,
		TabItemLabel:    tabItemLabel,
		Preset:          {{ call $DOT.Funcs.LowerCase $panelName }}ScreenPreset,
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
