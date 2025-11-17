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

	_misc_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/misc"
	_{{ call .Funcs.LowerCase .PanelName }}panel_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/panels/{{ .PanelName }}Panel"
	_types_ "{{ .ImportPrefix }}/frontend/deps/types"
)


// {{ .PanelName }}Panel is a {{ .PanelName }} panel.
// The panel's content is at {{ .PanelName }}Panel/content.go.
// The panel's content-state is at {{ .PanelName }}Panel/state.go.
// The panel's state-presets are at {{ .PanelName }}Panel/preset.go.
type {{ .PanelName }}Panel struct {
	content *_{{ call .Funcs.LowerCase .PanelName }}panel_.Content
	state   *_{{ call .Funcs.LowerCase .PanelName }}panel_.State
	screen  *_misc_.Miscellaneous
}

// New{{ .PanelName }}Panel initializes this panel.
// Returns the panel and the error.
func New{{ .PanelName }}Panel(screen *_misc_.Miscellaneous, preset *_{{ call .Funcs.LowerCase .PanelName }}panel_.Preset) (panel *{{ .PanelName }}Panel, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("{{ .PanelName }}Panel.New{{ .PanelName }}Panel: %w", err)
		}
	}()

	panel = &{{ .PanelName }}Panel{
		screen: screen,
	}
	if panel.content, err = _{{ call .Funcs.LowerCase .PanelName }}panel_.NewContent(screen); err != nil {
		return
	}
	if panel.state, err = _{{ call .Funcs.LowerCase .PanelName }}panel_.NewState(panel.content, screen.ScreenID); err != nil {
		return
	}
	panel.state.LoadPreset(preset)

	return
}

// ID returns this panel's id. Same as this panel's id.
func (panel *{{ .PanelName }}Panel) ID() (id string) {
	id = panel.state.ID()
	return
}

// Show shows this panel and hides the others.
func (panel *{{ .PanelName }}Panel) Show(isMainThread bool) {
	panel.screen.Layout.Set{{ .PanelName }}PanelCanvasObject(panel.ID(), panel.content.CanvasObject())
	panel.screen.Layout.Producer().Refresh(isMainThread)
}

// Producer returns the panel's producer.
// func (panel *{{ .PanelName }}Panel) Producer() (producer *_producer_.ContentProducer) {
func (panel *{{ .PanelName }}Panel) Producer() (producer _types_.ContentProducer) {
	producer = panel.screen.Layout.Producer()
	return
}

// Returns the panel's state.
func (panel *{{ .PanelName }}Panel) State() (state _types_.Stater) {
	state = panel.state
	return
}

// CanvasObject returns the panel's content.
func (panel *{{ .PanelName }}Panel) CanvasObject() (canvasObject fyne.CanvasObject) {
	canvasObject = panel.content.CanvasObject()
	return
}

// Required cleanup after unbinding.
func (panel *{{ .PanelName }}Panel) UnBindCleanUP() {
}
`
)
