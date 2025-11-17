package tabs

const (
	FileName = "border.go"

	Template = `{{ $DOT := . -}}
package border

import (
	"fyne.io/fyne/v2/container"
{{- if .UsesRemoteContent }}
	"fyne.io/fyne/v2/widget"
{{- end }}

{{- if .UsesRemoteContent }}

 {{- range $i, $remoteItemNames := .UniqueRemoteItemNames }}
	_{{ call $DOT.Funcs.LowerCase $remoteItemNames }}_ "{{ $DOT.ImportPrefix }}/frontend/screens/{{ $remoteItemNames }}"
	_{{ call $DOT.Funcs.LowerCase $remoteItemNames }}presetting_ "{{ $DOT.ImportPrefix }}/frontend/screens/{{ $remoteItemNames }}/presetting"
 {{- end }}
{{- end }}

{{- if .UsesRemoteContent }}

	_thread_ "{{ .ImportPrefix }}/deps/thread"
{{- else }}

{{- end }}
	_types_ "{{ .ImportPrefix }}/frontend/deps/types"
	_misc_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/misc"
{{- if .UsesLocalContent }}
	_panels_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/panels"
{{- end }}
	_presetting_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/presetting"
)

func Construct(screen *_misc_.Miscellaneous, preset *_presetting_.Preset) (err error) {

	// Areas using local panel content.
{{- range $area := .Areas }}
 {{- if eq (len $area.Area) 0 }}{{ continue }}{{ end }}
 {{- if eq $area.IsLocal false }}{{ continue }}{{ end }}
	var {{ call $DOT.Funcs.LowerCase $area.ItemName }}Panel *_panels_.{{ $area.ItemName }}Panel
{{- end }}

{{- range $area := .Areas }}
 {{- if eq (len $area.Area) 0 }}{{ continue }}{{ end }}
 {{- if eq $area.IsLocal false }}{{ continue }}{{ end }}
	if {{ call $DOT.Funcs.LowerCase $area.ItemName }}Panel, err = _panels_.New{{ $area.ItemName }}Panel(screen, preset.{{ $area.ItemName }}Panel); err != nil {
		return
	}
{{- end }}

	// Construct the border.

	// Start with the areas using a local panel for content.
	screen.Layout.SetBorder(
		container.NewBorder(
{{- if ne (len .Top.Area) 0 }}
 {{- if and (eq .Top.IsLocal true) }}
			{{ call .Funcs.LowerCase .Top.ItemName }}Panel.CanvasObject(),
 {{- else }}
			widget.NewLabel("{{ .Top.Area }}"),
 {{- end }}
{{- else }}
			nil,
{{- end }}

{{- if ne (len .Bottom.Area) 0 }}
 {{- if and (eq .Bottom.IsLocal true) }}
			{{ call .Funcs.LowerCase .Bottom.ItemName }}Panel.CanvasObject(),
 {{- else }}
			widget.NewLabel("{{ .Bottom.Area }}"),
 {{- end }}
{{- else }}
			nil,
{{- end }}

{{- if ne (len .Left.Area) 0 }}
 {{- if and (eq .Left.IsLocal true) }}
			{{ call .Funcs.LowerCase .Left.ItemName }}Panel.CanvasObject(),
 {{- else }}
			widget.NewLabel("{{ .Left.Area }}"),
 {{- end }}
{{- else }}
			nil,
{{- end }}

{{- if ne (len .Right.Area) 0 }}
 {{- if and (eq .Right.IsLocal true) }}
			{{ call .Funcs.LowerCase .Right.ItemName }}Panel.CanvasObject(),
 {{- else }}
			widget.NewLabel("{{ .Right.Area }}"),
 {{- end }}
{{- else }}
			nil,
{{- end }}

{{- if ne (len .Center.Area) 0 }}
 {{- if and (eq .Center.IsLocal true) }}
			{{ call .Funcs.LowerCase .Center.ItemName }}Panel.CanvasObject(),
 {{- else }}
			widget.NewLabel("{{ .Center.Area }}"),
 {{- end }}
{{- else }}
			nil,
{{- end }}
		),
	)

	// Bind the local panels to their area consumers.
{{- range $area := .Areas }}
 {{- if eq (len $area.Area) 0 }}{{ continue }}{{ end }}
 {{- if eq $area.IsLocal false }}{{ continue }}{{ end }}
	// {{ $area.Area }} border area.
	{{ call $DOT.Funcs.LowerCase $area.ItemName }}AreaContentConsumer := _types_.NewBorderAreaContentConsumer(screen.Layout.Border(), screen.Layout.{{ $area.Area }}AreaIndex)
	{{ call $DOT.Funcs.LowerCase $area.ItemName }}Panel.Bind({{ call $DOT.Funcs.LowerCase $area.ItemName }}AreaContentConsumer)
	{{ call $DOT.Funcs.LowerCase $area.ItemName }}Panel.State().Refresh(true)
{{- end }}

{{- if .UsesRemoteContent }}

// Continue by adding the border areas using another screen for content.
	var borderAreaContentConsumer *_types_.BorderAreaContentConsumer
 {{- range $area := .Areas }}
  {{- if eq (len $area.Area) 0 }}{{ continue }}{{ end }}
  {{- if eq $area.IsLocal true }}{{ continue }}{{ end }}
	// {{ $area.Area }} border area.
	{{ call $DOT.Funcs.LowerCase $area.Area }}DefaultPreset := _{{ call $DOT.Funcs.LowerCase $area.ItemName }}presetting_.Presets["Default"]
	if borderAreaContentConsumer, _, err = _{{ call $DOT.Funcs.LowerCase $area.ItemName }}_.NewBorderAreaContentConsumer(
		screen.CTX,
		screen.CTXCancel,
		screen.APP,
		screen.Window,
		screen.Layout.Border(),
		screen.Layout.{{ $area.Area }}AreaIndex,
		{{ call $DOT.Funcs.LowerCase $area.Area }}DefaultPreset,
	); err != nil {
		return
	}
	borderAreaContentConsumer.Refresh(_thread_.IsMainThread())
 {{- end }}
{{- end }}
	return
}
`
)
