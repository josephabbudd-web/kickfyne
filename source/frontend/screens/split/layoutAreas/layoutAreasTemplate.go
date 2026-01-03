package layoutitems

const (
	FileName = "layoutAreas.go"
	Template = `package layoutareas

import (
	"fyne.io/fyne/v2/container"

	_types_ "{{ .ImportPrefix }}/frontend/deps/types"
	_misc_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/misc"
	_presetting_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/presetting"
{{- if .UsesLocalContent}}
	_panels_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/panels"
{{- end }}
{{- if eq false .Leading.IsLocal }}
	_{{ call .Funcs.LowerCase .Leading.ScreenName }}_ "{{ .ImportPrefix }}/frontend/screens/{{ .Leading.ScreenName }}"
{{- end }}
{{- if eq false .Trailing.IsLocal }}
	_{{ call .Funcs.LowerCase .Trailing.ScreenName }}_ "{{ .ImportPrefix }}/frontend/screens/{{ .Trailing.ScreenName }}"
{{- end }}
)

// Layout lays out the split vercially or horizontally.
// KICKFYNE TODO: Edit func Layout to layout vertically or horizontally.
func Layout(screen *_misc_.Miscellaneous, preset *_presetting_.Preset, isMainThread bool) (splitContentProducer *_types_.SplitContentProducer, err error) {
	if preset.Direction == _presetting_.Horizontal {
		splitContentProducer, err = singleProducer(
			container.NewHSplit(nil, nil),
			screen,
			preset,
			isMainThread,
		)
	} else {
		splitContentProducer, err = singleProducer(
			container.NewVSplit(nil, nil),
			screen,
			preset,
			isMainThread,
		)
	}
	return
}

func singleProducer(split *container.Split, screen *_misc_.Miscellaneous, preset *_presetting_.Preset, isMainThread bool) (splitContentProducer *_types_.SplitContentProducer, err error) {
	{{ if .Leading.IsLocal }}
	// Leading.
	var leadingPanel *_panels_.LeadingPanel
	if leadingPanel, err = _panels_.NewLeadingPanel(screen, preset.LeadingPanel); err != nil {
		return
	}
	leadingSplitAreaContentConsumer := leadingPanel.Producer().SplitAreaConsumer()
{{- else }}
	// Leading.
	var leadingSplitAreaContentConsumer *_types_.SplitAreaContentConsumer
	if leadingSplitAreaContentConsumer, _, err = _{{ call .Funcs.LowerCase .Leading.ScreenName }}_.NewSplitAreaContentConsumer(
		screen.CTX,
		screen.CTXCancel,
		screen.APP,
		screen.Window,
		preset.{{ .Leading.ScreenName }}Screen.Preset,
	); err != nil {
		return
	}
{{- end }}
	leadingSplitAreaContentConsumer.SetPosition(_types_.SplitPositionLeading)

{{- if .Trailing.IsLocal }}
	// Trailing.
	var trailingPanel *_panels_.TrailingPanel
	if trailingPanel, err = _panels_.NewTrailingPanel(screen, preset.TrailingPanel); err != nil {
		return
	}
	trailingSplitAreaContentConsumer := trailingPanel.Producer().SplitAreaConsumer()
{{- else }}
	// Trailing.
	var trailingSplitAreaContentConsumer *_types_.SplitAreaContentConsumer
	if trailingSplitAreaContentConsumer, _, err = _{{ call .Funcs.LowerCase .Trailing.ScreenName }}_.NewSplitAreaContentConsumer(
		screen.CTX,
		screen.CTXCancel,
		screen.APP,
		screen.Window,
		preset.{{ .Trailing.ScreenName }}Screen.Preset,
	); err != nil {
		return
	}
{{- end }}
	trailingSplitAreaContentConsumer.SetPosition(_types_.SplitPositionTrailing)

	// The SplitContentProducer.
	splitContentProducer = _types_.NewSplitContentProducer(
		split,
		leadingSplitAreaContentConsumer,
		trailingSplitAreaContentConsumer,
	)
	splitContentProducer.RefreshSplit(isMainThread)

	return
}
`
)
