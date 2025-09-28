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

{{ if ne (len .RemotePanelNames) 0 }}
	_screenmap_ "{{ .ImportPrefix }}/frontend/screenmap"
{{ end }}
	_misc_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/misc"
	_panels_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/panels"
	_startup_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/startup"
	_types_ "{{ .ImportPrefix }}/frontend/types"
	_ids_ "{{ .ImportPrefix }}/deps/container/{{ .PackageName }}"
{{- range $i, $panelName := .RemotePanelNames }}
 {{- if eq $i 0 }}

	_{{ call $DOT.Funcs.LowerCase $panelName }}startup_ "{{ $DOT.ImportPrefix }}/frontend/screens/{{ $panelName }}/startup"
 {{- else }}
	_{{ call $DOT.Funcs.LowerCase $panelName }}startup_ "{{ $DOT.ImportPrefix }}/frontend/screens/{{ $panelName }}/startup"
 {{- end }}
{{- end }}
)

{{- range $panelName := .LocalPanelNames}}

// Open{{ $panelName }}AccordionItem constructs a {{ $panelName }}AccordionItem.
// The {{ $panelName }}AccordionItem uses the local {{ $panelName }} panel for content.
func Open{{ $panelName }}AccordionItem(screen *_misc_.Miscellaneous, accordionItemTitle string, {{ call $DOT.Funcs.DeCap $panelName }}PanelInitData *_startup_.{{ $panelName }}PanelInitData) (err error) {
	accordionItem := widget.NewAccordionItem(accordionItemTitle, widget.NewLabel("This is the {{ $panelName }} panel."))
	accordionItemContentConsumer := _types_.NewAccordionItemContentConsumer(screen.Layout.Accordion(), accordionItem)
	// accordionItemContentConsumer := _types_.NewAccordionItemContentConsumer(screen.Layout.AccordionConsumer(), screen.Layout.Accordion(), accordionItem, spawned)
	// The {{ $panelName }} panel.
	var panel *_panels_.{{ $panelName }}Panel
	if panel, err = _panels_.New{{ $panelName }}Panel(screen, accordionItemContentConsumer, accordionItem, {{ call $DOT.Funcs.DeCap $panelName }}PanelInitData); err != nil {
		return
	}
	panelProducer := panel.Producer()
	accordionItemContentConsumer.Bind(panelProducer)
	// Add the tab to the layout.
	screen.Layout.AddPanelerAccordionItemConsumer(panel, accordionItem, accordionItemContentConsumer)
	// Update deps.
	_ids_.AddAccordionItem(screen.ScreenID, panel.ID(), _ids_.{{ $panelName }}Item)
	return
}
{{- end }}

{{- range $panelName := .RemotePanelNames}}

// Open{{ $panelName }}AccordionItem constructs a {{ $panelName }}AccordionItem.
// The {{ $panelName }}AccordionItem uses the {{ $panelName }} screen for content.
func Open{{ $panelName }}AccordionItem(screen *_misc_.Miscellaneous, title string, {{ call $DOT.Funcs.DeCap $panelName }}ScreenInitData *_{{ call $DOT.Funcs.LowerCase $panelName }}startup_.ScreenInitData) (err error) {
	accordionItem := widget.NewAccordionItem(title, widget.NewLabel("This is the {{ $panelName }} panel."))
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
		{{ call $DOT.Funcs.DeCap $panelName }}ScreenInitData,
	); err != nil {
		return
	}
	screen.Layout.AddRemoteScreenIDAccordionItemConsumer(accordionItemID, accordionItem, accordionItemContentConsumer)
	// Update deps.
	_ids_.AddAccordionItem(screen.ScreenID, accordionItemID, _ids_.{{ $panelName }}Item)
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
	// Update deps.
	_ids_.RemoveAccordionItem(screen.ScreenID, accordionItemID)
}

`
)
