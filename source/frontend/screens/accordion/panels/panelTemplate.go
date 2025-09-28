package panels

import (
	"github.com/josephabbudd-web/kickfyne/source/utils"
)

type PanelTemplateData struct {
	PanelName    string
	PackageName  string
	ImportPrefix string
	Funcs        utils.Funcs
}

const (
	PanelFileNameSuffix = "Panel.go"

	PanelNoBETemplate = `package panels

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"

	_misc_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/misc"
	_content_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/panels/{{ .PanelName }}Panel"
	_startup_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/startup"
	_types_ "{{ .ImportPrefix }}/frontend/types"
)

// {{ .PanelName }}Panel is a {{ .PanelName }} panel.
// It's content is at {{ .PanelName }}Panel/content.go.
// It's content's state is at {{ .PanelName }}Panel/state.go.
type {{ .PanelName }}Panel struct {
	content       *_content_.Content
	state         *_content_.State
	accordionItem *widget.AccordionItem
	screen        *_misc_.Miscellaneous
}

// New{{ .PanelName }}Panel initializes this panel.
// Returns the panel and the error.
func New{{ .PanelName }}Panel(screen *_misc_.Miscellaneous, accordionItemContentConsumer *_types_.AccordionItemContentConsumer, accordionItem *widget.AccordionItem, {{ call .Funcs.DeCap .PanelName }}PanelInitData *_startup_.{{ .PanelName }}PanelInitData) (panel *{{ .PanelName }}Panel, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("{{ .PanelName }}Panel.New{{ .PanelName }}Panel: %w", err)
		}
	}()

	panel = &{{ .PanelName }}Panel{
		screen:        screen,
		accordionItem: accordionItem,
	}
	if panel.content, err = _content_.NewContent(accordionItemContentConsumer, screen, accordionItem, panel); err != nil {
		return
	}
	if panel.state, err = _content_.NewState(
		panel.content,
		screen.ScreenID,
	); err != nil {
		return
	}
	panel.state.LoadStartupData({{ call .Funcs.DeCap .PanelName }}PanelInitData)

	return
}

// ID returns the panel's id.
func (panel *{{ .PanelName }}Panel) ID() (id string) {
	getters := panel.state.Get().(_content_.Getters)
	id = getters.ID()
	return
}

// Show doesn't do anything becuase this panel is always shown by it's accordionItem.
func (panel *{{ .PanelName }}Panel) Show(isMainThread bool) {}

// Producer returns the panel's producer.
// func (panel *{{ .PanelName }}Panel) Producer() (producer *_producer_.ContentProducer) {
func (panel *{{ .PanelName }}Panel) Producer() (producer _types_.ContentProducer) {
	producer = panel.content.Producer()
	return
}

// CanvasObject returns the panel's content.
func (panel *{{ .PanelName }}Panel) CanvasObject() (canvasObject fyne.CanvasObject) {
	canvasObject = panel.content.CanvasObject()
	return
}

// Required cleanup after unbinding.
func (panel *{{ .PanelName }}Panel) UnBindCleanUP() {}

// State returns the panel's state.
func (panel *{{ .PanelName }}Panel) State() (state _types_.Stater) {
	state = panel.state
	return
}
`
)
