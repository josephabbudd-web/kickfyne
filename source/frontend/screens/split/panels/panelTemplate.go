package panels

import (
	_utils_ "github.com/josephabbudd-web/kickfyne/source/utils"
)

type PanelTemplateData struct {
	PackageName  string
	PanelName    string
	Area         string
	ImportPrefix string
	Funcs        _utils_.Funcs
}

const (
	PanelFileNameSuffix = "Panel.go"

	PanelTemplate = `package panels

import (
	"fyne.io/fyne/v2"

	_types_ "{{ .ImportPrefix }}/frontend/deps/types"
	_misc_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/misc"
	_{{ call .Funcs.LowerCase .PanelName }}panel_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/panels/{{ .PanelName }}Panel"
)

// {{ .PanelName }}Panel is a {{ .PanelName }} panel.
// It's content is at {{ .PanelName }}Panel/content.go.
// It's content's state is at {{ .PanelName }}Panel/state.go.
type {{ .PanelName }}Panel struct {
	content   *_{{ call .Funcs.LowerCase .PanelName }}panel_.Content
	state     *_{{ call .Funcs.LowerCase .PanelName }}panel_.State
	screen    *_misc_.Miscellaneous
}

// New{{ .PanelName }}Panel initializes this panel.
// Returns the panel and the error.
func New{{ .PanelName }}Panel(screen *_misc_.Miscellaneous, preset *_{{ call .Funcs.LowerCase .PanelName }}panel_.Preset) (panel *{{ .PanelName }}Panel, err error) {
	panel = &{{ .PanelName }}Panel{
		screen: screen,
	}
	if panel.content, err = _{{ call .Funcs.LowerCase .PanelName }}panel_.NewContent(screen); err != nil {
		return
	}
	if panel.state, err = _{{ call .Funcs.LowerCase .PanelName }}panel_.NewState(panel.content, panel.screen.ScreenID); err != nil {
		return
	}
	panel.state.LoadInitialPreset(preset)
	return
}

// ID returns the panel's id.
func (panel *{{ .PanelName }}Panel) ID() (id string) {
	id = panel.state.ID()
	return
}

// Show doesn't do anything becuase this panel is always shown by it's tabItem.
func (panel *{{ .PanelName }}Panel) Show(isMainThread bool) {}

// Producer returns the panel's producer.
func (panel *{{ .PanelName }}Panel) Producer() (producer *_types_.SplitAreaContentProducer) {
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
