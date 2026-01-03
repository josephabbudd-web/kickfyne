package panel

import (
	_utils_ "github.com/josephabbudd-web/kickfyne/source/utils"
)

const (
	ContentFileName = _utils_.ContentFileName

	ContentTemplate = `{{ $DOT := . -}}
package {{ .PanelName }}Panel

import (
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	_types_ "{{ .ImportPrefix }}/frontend/deps/types"
	_misc_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/misc"
)

// Content is the content for the {{ .PanelName }} panel.
// KICKFYNE TODO: Correct this panel's doc comment.
type Content struct {
	state    *State
	producer *_types_.SplitAreaContentProducer
	content  fyne.CanvasObject
	screen   *_misc_.Miscellaneous

	// Widgets with variable state. See state.go.
	heading     *widget.Label
	description *widget.Label
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
		screen: screen,

		// Widgets with variable state. See state.go.
		heading:     widget.NewLabelWithStyle("", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		description: widget.NewLabelWithStyle("", fyne.TextAlignLeading, fyne.TextStyle{Monospace: true}),
	}

	// Layout the components.
	panelContent.content = container.NewVBox(
		panelContent.heading,
		panelContent.description,
		widget.NewButton("{{ .PanelName }} 1", func() { log.Println("{{ .PanelName }} 1") }),
		widget.NewButton("{{ .PanelName }} 2", func() { log.Println("{{ .PanelName }} 2") }),
		widget.NewButton("{{ .PanelName }} 3", func() { log.Println("{{ .PanelName }} 3") }),
		widget.NewButton("{{ .PanelName }} 4", func() { log.Println("{{ .PanelName }} 4") }),
		widget.NewButton("{{ .PanelName }} 5", func() { log.Println("{{ .PanelName }} 5") }),
	)

	splitAreaContentConsumer := _types_.NewSplitAreaContentConsumer()
	panelContent.producer = _types_.NewSplitAreaContentProducer(panelContent.content, splitAreaContentConsumer)
	panelContent.producer.Bind(splitAreaContentConsumer)

	return
}

// CanvasObject returns the panel's content as a fyne.CanvasObject.
func (panelContent *Content) CanvasObject() (canvasObject fyne.CanvasObject) {
	canvasObject = panelContent.content
	return
}

func (panelContent *Content) Producer() (producer *_types_.SplitAreaContentProducer) {
	producer = panelContent.producer
	return
}`
)
