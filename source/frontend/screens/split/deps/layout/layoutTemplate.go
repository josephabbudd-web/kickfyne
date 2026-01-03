package misc

const (
	LayoutFileName = "layout.go"

	LayoutTemplate = `package layout

import (
	_types_ "{{ .ImportPrefix }}/frontend/deps/types"
)

// Layout this screen's layout of a container.AppTabs.
type Layout struct {
	splitContentProducer _types_.ContentProducer
}

// NewLayout constructs this layout.
func NewLayout() (layout *Layout) {
	layout = &Layout{}
	return
}

func (layout *Layout) SetSplitContentProducer(splitContentProducer _types_.ContentProducer) {
	layout.splitContentProducer = splitContentProducer
}

func (layout *Layout) Producer() (producer _types_.ContentProducer) {
	producer = layout.splitContentProducer
	return
}
`
)
