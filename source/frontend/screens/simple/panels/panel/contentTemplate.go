package panel

import (
	_utils_ "github.com/josephabbudd-web/kickfyne/source/utils"
)

type TemplateData struct {
	PackageName     string
	PanelName       string
	LocalPanelNames []string
	ImportPrefix    string
	Funcs           _utils_.Funcs
}

const (
	ContentFileName = _utils_.ContentFileName

	ContentTemplate = `{{ $DOT := . -}}
package {{ .PanelName }}Panel

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	_thread_ "{{ .ImportPrefix }}/deps/thread"
	_misc_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/misc"
)

// Content is the content for the {{ .PanelName }} panel.
// KICKFYNE TODO: Correct this panel's doc comment.
type Content struct {
	state   *State
	content fyne.CanvasObject
	screen  *_misc_.Miscellaneous

	// Widgets with variable state. See state.go.
	heading     *widget.Label
	description *widget.Label
	switcher    *widget.Select
}

// NewContent initializes this panel's content.
// Returns the panel's content and the error.
func NewContent(screen *_misc_.Miscellaneous) (panelContent *Content, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("{{ .PanelName }}Panel.NewContent: %w", err)
		}
	}()

	// Create the components of this panel's content.
	panelContent = &Content{
		screen:      screen,
		heading:     widget.NewLabelWithStyle("", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		description: widget.NewLabelWithStyle("", fyne.TextAlignLeading, fyne.TextStyle{Monospace: true}),
	}

	// Layout the components.
	objects := make([]fyne.CanvasObject, 0, 3)
	objects = append(objects, panelContent.heading, panelContent.description)
	panelContent.setSwitcher()
	if panelContent.switcher != nil {
		hbox := container.NewHBox(
			widget.NewLabel("Switch to another panel."),
			panelContent.switcher,
		)
		objects = append(objects, hbox)
	}
	panelContent.content = container.NewVBox(objects...)
	return
}

// CanvasObject returns the panel's content as a fyne.CanvasObject.
func (panelContent *Content) CanvasObject() (canvasObject fyne.CanvasObject) {
	canvasObject = panelContent.content
	return
}

// setSwitcher demonstrates how to switch panels in a Simple Screen.
// KICKFYNE TODO: You will probably want to:
// - Remove the line in NewContent that calls setSwitcher().
// - Remove panelContent.switcher.
// - Remove this func.
func (panelContent *Content) setSwitcher() {
	panelMap := panelContent.screen.Panelers.Map()
	var length int
	if length = len(panelMap); length == 1 {
		return
	}
	options := make([]string, 0, length)
	for name := range panelMap {
		if name == "{{ .PanelName }}" {
			continue
		}
		options = append(options, name)
	}
	panelContent.switcher = widget.NewSelect(
		options,
		func(value string) {
			if len(value) == 0 {
				return
			}
			panelerMap := panelContent.screen.Panelers.Map()
			paneler := panelerMap[value]
			isMainThread := _thread_.IsMainThread()
			paneler.Show(isMainThread)			
			panelContent.switcher.ClearSelected()
		},
	)
}

`
)
