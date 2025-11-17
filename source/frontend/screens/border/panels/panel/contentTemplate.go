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

	_producer_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/deps/producer"
	_misc_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/misc"
	_types_ "{{ .ImportPrefix }}/frontend/deps/types"
)

// Content is the content for the {{ .PanelName }} panel.
// KICKFYNE TODO: Correct this panel's doc comment.
type Content struct {
	state    *State
	producer *_producer_.AreaContentProducer
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
		// producer: _producer_.NewAreaContentProducer(tabItemContentConsumer),
		screen:   screen,

		// Widgets with variable state. See state.go.
		heading:     widget.NewLabelWithStyle("", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		description: widget.NewLabelWithStyle("", fyne.TextAlignLeading, fyne.TextStyle{Monospace: true}),
	}

	// Layout the components.
	panelContent.content = container.NewVBox(
		panelContent.heading,
		panelContent.description,
{{- if eq .Area "Top" "Bottom" }}
		container.NewHBox(
			widget.NewButton("{{ .Area }} 1", func() { log.Println("{{ .Area }} 1") }),
			widget.NewButton("{{ .Area }} 2", func() { log.Println("{{ .Area }} 2") }),
			widget.NewButton("{{ .Area }} 3", func() { log.Println("{{ .Area }} 3") }),
			widget.NewButton("{{ .Area }} 4", func() { log.Println("{{ .Area }} 4") }),
			widget.NewButton("{{ .Area }} 5", func() { log.Println("{{ .Area }} 5") }),
		),
{{- else }}
			widget.NewButton("{{ .Area }} 1", func() { log.Println("{{ .Area }} 1") }),
			widget.NewButton("{{ .Area }} 2", func() { log.Println("{{ .Area }} 2") }),
			widget.NewButton("{{ .Area }} 3", func() { log.Println("{{ .Area }} 3") }),
			widget.NewButton("{{ .Area }} 4", func() { log.Println("{{ .Area }} 4") }),
			widget.NewButton("{{ .Area }} 5", func() { log.Println("{{ .Area }} 5") }),
{{- end }}
	)
	return
}
func (panelContent *Content) Bind(borderAreaContentConsumer *_types_.BorderAreaContentConsumer) {
	panelContent.producer = _producer_.NewAreaContentProducer(borderAreaContentConsumer)
}

// CanvasObject returns the panel's content as a fyne.CanvasObject.
func (panelContent *Content) CanvasObject() (canvasObject fyne.CanvasObject) {
	canvasObject = panelContent.content
	return
}

func (panelContent *Content) Producer() (producer _types_.ContentProducer) {
	producer = panelContent.producer
	return
}
`
)
