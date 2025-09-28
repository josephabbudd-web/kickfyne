package misc

import (
	_utils_ "github.com/josephabbudd-web/kickfyne/source/utils"
)

type LayoutTemplateData struct {
	PackageName     string
	ImportPrefix    string
	Funcs           _utils_.Funcs
	LocalPanelNames []string
}

const (
	LayoutFileName = "layout.go"

	LayoutTemplate = `package layout

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	fynelayout "fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	_producer_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/deps/producer"
)
{{- if eq (len .LocalPanelNames) 1 }}

// Layout is a Simple layout of it's panel.
 {{- else }}

// Layout is a Simple layout.
// It lays out 1 of it's {{ len .LocalPanelNames }} panels.
{{- end }}
// It is this screen's layout.
type Layout struct {
	panelID            string
	panelCanvasObject  fyne.CanvasObject
	screenCanvasObject fyne.CanvasObject
	contentProducer    *_producer_.ContentProducer
}

// NewLayout constructs this layout.
func NewLayout(producer *_producer_.ContentProducer) (layout *Layout, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("{{ .PackageName }}.newLayout: %v", err)
		}
	}()

	layout = &Layout{
		contentProducer: producer,
		panelCanvasObject:    widget.NewLabel("{{ .PackageName }} screen."), // Temporary content replaced by actual panel content.
	}
	err = layout.layoutScreenCanvasObject()

	return
}

func (layout *Layout) Producer() (producer *_producer_.ContentProducer) {
	producer = layout.contentProducer
	return
}

func (layout *Layout) RefreshIfCurrent(panelID string, panelCanvasObject fyne.CanvasObject) (err error) {
	if layout.panelID == panelID {
		layout.panelCanvasObject = panelCanvasObject
		if err = layout.layoutScreenCanvasObject(); err != nil {
			return
		}
		layout.contentProducer.SetCanvasObject(layout.screenCanvasObject)
	}
	return
}

{{- range $panelName := .LocalPanelNames }}

// Set{{ $panelName }}PanelCanvasObject sets the {{ $panelName }}Panel's canvas object in the layout.
func (layout *Layout) Set{{ $panelName }}PanelCanvasObject(panelID string, panelCanvasObject fyne.CanvasObject) (err error) {
	layout.panelCanvasObject = panelCanvasObject
	layout.panelID = panelID
	if err = layout.layoutScreenCanvasObject(); err != nil {
		return
	}
	layout.contentProducer.SetCanvasObject(layout.screenCanvasObject)
	return
}
{{- end }}

// layoutScreenCanvasObject lays out the screen's panel's.
// A stack layout of a single panel.
func (layout *Layout) layoutScreenCanvasObject() (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("error {{ .PackageName }}:layoutScreenCanvasObject: %q", err.Error())
		}
	}()

	layedOut := container.New(
		fynelayout.NewStackLayout(),
		layout.panelCanvasObject,
	)

	if err == nil {
		// No error so use the new canvas object.
		layout.screenCanvasObject = layedOut
	}
	return
}
`
)
