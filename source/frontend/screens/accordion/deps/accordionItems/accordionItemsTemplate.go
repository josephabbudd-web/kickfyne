package accordionitems

import (
	_utils_ "github.com/josephabbudd-web/kickfyne/source/utils"
)

type TemplateData struct {
	PackageName      string
	LocalPanelNames  []string
	RemotePanelNames []string
	ImportPrefix     string
	Funcs            _utils_.Funcs
}

const (
	FileName = "accordionItems.go"

	NoBETemplate = `{{ $DOT := . -}}
package accordionitems

import (
	"fyne.io/fyne/v2/widget"

{{- range $i, $panelName := .LocalPanelNames }}
 {{- if eq $i 0 }}

	_{{ call $DOT.Funcs.LowerCase $panelName }}panel_ "{{ $DOT.ImportPrefix }}/frontend/screens/{{ $DOT.PackageName }}/panels/{{ $panelName }}Panel"
 {{- else }}
	_{{ call $DOT.Funcs.LowerCase $panelName }}panel_ "{{ $DOT.ImportPrefix }}/frontend/screens/{{ $DOT.PackageName }}/panels/{{ $panelName }}Panel"
 {{- end }}
{{- end }}

{{ if ne (len .RemotePanelNames) 0 }}
	_screenmap_ "{{ .ImportPrefix }}/frontend/deps/screenmap"
{{ end }}
	_misc_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/misc"
	_panels_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/panels"
	_presetting_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/presetting"
	_types_ "{{ .ImportPrefix }}/frontend/deps/types"
)

func SetInitialAccordionItems(screen *_misc_.Miscellaneous, screenPreset *_presetting_.Preset) (err error) {
{{- range $panelName := .LocalPanelNames}}
	if err = Open{{ $panelName }}AccordionItem(screen, screenPreset.{{ $panelName }}Panel); err != nil {
		return
	}
{{- end }}

{{ range $panelName := .RemotePanelNames}}
	if err = Open{{ $panelName }}AccordionItem(screen, screenPreset.{{ $panelName }}Screen); err != nil {
		return
	}
{{- end }}
	return
}

{{- range $panelName := .LocalPanelNames}}

// Open{{ $panelName }}AccordionItem constructs a {{ $panelName }}AccordionItem.
// The {{ $panelName }}AccordionItem uses the local {{ $panelName }} panel for content.
func Open{{ $panelName }}AccordionItem(screen *_misc_.Miscellaneous, preset *_{{ call $DOT.Funcs.LowerCase $panelName }}panel_.Preset) (err error) {
	accordionItem := widget.NewAccordionItem(preset.AccordionItemTitle, widget.NewLabel("This is the {{ $panelName }} panel."))
	accordionItemContentConsumer := _types_.NewAccordionItemContentConsumer(screen.Layout.Accordion(), accordionItem)
	// accordionItemContentConsumer := _types_.NewAccordionItemContentConsumer(screen.Layout.AccordionConsumer(), screen.Layout.Accordion(), accordionItem, spawned)
	// The {{ $panelName }} panel.
	var panel *_panels_.{{ $panelName }}Panel
	if panel, err = _panels_.New{{ $panelName }}Panel(screen, accordionItemContentConsumer, accordionItem, preset); err != nil {
		return
	}
	panelProducer := panel.Producer()
	accordionItemContentConsumer.Bind(panelProducer)
	// Add the tab to the layout.
	screen.Layout.AddPanelerAccordionItemConsumer(panel, accordionItem, accordionItemContentConsumer)
	return
}
{{- end }}

{{- range $panelName := .RemotePanelNames}}

// Open{{ $panelName }}AccordionItem constructs a {{ $panelName }}AccordionItem.
// The {{ $panelName }}AccordionItem uses the {{ $panelName }} screen for content.
func Open{{ $panelName }}AccordionItem(screen *_misc_.Miscellaneous, preset *_presetting_.{{ $panelName }}ScreenPreset) (err error) {
	accordionItem := widget.NewAccordionItem(preset.AccordionItemTitle, widget.NewLabel("This is the {{ $panelName }} panel."))
	api := _screenmap_.Map["{{ $panelName }}"]
	var accordionItemContentConsumer *_types_.AccordionItemContentConsumer
	var accordionItemID string
	if accordionItemContentConsumer, accordionItemID, err = api.NewAccordionItemContentConsumer(
		screen.CTX,
		screen.CTXCancel,
		screen.APP,
		screen.Window,
		screen.Layout.Accordion(),
		accordionItem,
		preset.Preset,
	); err != nil {
		return
	}
	screen.Layout.AddRemoteScreenIDAccordionItemConsumer(accordionItemID, accordionItem, accordionItemContentConsumer)
	return
}
{{- end }}

func CloseAccordionItem(screen *_misc_.Miscellaneous, accordionItem *widget.AccordionItem) {
	if accordionItemID, found := screen.Layout.AccordionItemID(accordionItem); found {
		CloseAccordionItemID(screen, accordionItemID)
	}
}

func CloseAccordionItemID(screen *_misc_.Miscellaneous, accordionItemID string) {
	screen.Layout.RemoveID(accordionItemID)
}

`
)
