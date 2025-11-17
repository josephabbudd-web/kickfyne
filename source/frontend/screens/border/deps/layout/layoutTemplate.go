package misc

const (
	LayoutFileName = "layout.go"

	LayoutTemplate = `package layout

import (
	"fyne.io/fyne/v2"

	_types_ "{{ .ImportPrefix }}/frontend/deps/types"
)

// Layout this screen's layout of a container.AppTabs.
type Layout struct {
	borderConsumer        _types_.ContentConsumer
	borderContentProducer _types_.ContentProducer
	border                *fyne.Container
	TopAreaIndex          int
	BottomAreaIndex       int
	LeftAreaIndex         int
	RightAreaIndex        int
	CenterAreaIndex       int
}

// NewLayout constructs this layout.
func NewLayout(borderConsumer _types_.ContentConsumer, borderContentProducer _types_.ContentProducer) (layout *Layout) {
	layout = &Layout{
		borderConsumer:        borderConsumer,
		borderContentProducer: borderContentProducer,
		TopAreaIndex:          {{ .TopIndex }},
		BottomAreaIndex:       {{ .BottomIndex }},
		LeftAreaIndex:         {{ .LeftIndex }},
		RightAreaIndex:        {{ .RightIndex }},
		CenterAreaIndex:       {{ .CenterIndex }},
	}
	return
}

func (layout *Layout) SetBorder(border *fyne.Container) {
	layout.border = border
	layout.borderContentProducer.SetCanvasObject(border)
}

func (layout *Layout) BorderConsumer() (borderConsumer _types_.ContentConsumer) {
	borderConsumer = layout.borderConsumer
	return
}

func (layout *Layout) Producer() (producer _types_.ContentProducer) {
	producer = layout.borderContentProducer
	return
}

// Refresh refreshes the tabbar tabs only. Not their content.
func (layout *Layout) Refresh() {
	layout.border.Refresh()
}

func (layout *Layout) Border() (tabbar *fyne.Container) {
	tabbar = layout.border
	return
}
`
)
