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

	_producer_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/deps/producer"
	_misc_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/misc"
	_types_ "{{ .ImportPrefix }}/frontend/deps/types"
)

// Content is the content for the {{ .PanelName }} panel.
// KICKFYNE TODO: Correct this panel's doc comment.
type Content struct {
	state    *State
	producer *_producer_.TabItemContentProducer
	content  fyne.CanvasObject
	screen   *_misc_.Miscellaneous
	tabItem  *container.TabItem

	// Widgets with variable state. See state.go.
	heading     *widget.Label
	description *widget.Label
}

// NewContent initializes this panel's content.
// Returns the panel's content and the error.
func NewContent(tabItemContentConsumer *_types_.DocTabsTabItemContentConsumer, screen *_misc_.Miscellaneous, tabItem *container.TabItem, paneler _types_.Paneler) (panelContent *Content, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("{{ .PanelName }}Panel.NewContent: %w", err)
		}
	}()

	_ = paneler

	// Create the components of this panel's content.
	panelContent = &Content{
		producer: _producer_.NewTabItemContentProducer(tabItemContentConsumer),
		screen:   screen,
		tabItem:  tabItem,

		// Widgets with variable state. See state.go.
		heading:     widget.NewLabelWithStyle("", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		description: widget.NewLabelWithStyle("", fyne.TextAlignLeading, fyne.TextStyle{Monospace: true}),
	}

	// Layout the components.
	panelContent.content = container.NewVBox(
		panelContent.heading,
		panelContent.description,
	)
	return
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

	ConfigContentTemplate = `{{ $DOT := . -}}
package {{ .PanelName }}Panel

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	_producer_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/deps/producer"
	_misc_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/misc"
	_types_ "{{ .ImportPrefix }}/frontend/deps/types"
)

// Content is the content for the Settings panel.
// KICKFYNE TODO: Correct this panel's doc comment.
type Content struct {
	state    *State
	producer *_producer_.TabItemContentProducer
	content  fyne.CanvasObject
	screen   *_misc_.Miscellaneous
	tabItem  *container.TabItem

	// Widgets with variable state. See state.go.
	heading     *widget.Label
	description *widget.Label
}

// NewContent initializes this panel's content.
// Returns the panel's content and the error.
func NewContent(tabItemContentConsumer *_types_.DocTabsTabItemContentConsumer, screen *_misc_.Miscellaneous, tabItem *container.TabItem, paneler _types_.Paneler) (panelContent *Content, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("SettingsPanel.NewContent: %w", err)
		}
	}()

	_ = paneler

	// Create the components of this panel's content.
	panelContent = &Content{
		producer: _producer_.NewTabItemContentProducer(tabItemContentConsumer),
		screen:   screen,
		tabItem:  tabItem,

		// Widgets with variable state. See state.go.
		heading:     widget.NewLabelWithStyle("", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		description: widget.NewLabelWithStyle("", fyne.TextAlignLeading, fyne.TextStyle{Monospace: true}),
	}

	// Layout the components.
	topCentered := container.New(
		layout.NewCenterLayout(),
		widget.NewButtonWithIcon(
			"", // "Top",
			theme.MoveUpIcon(),
			func() {
				panelContent.screen.Layout.Tabbar().SetTabLocation(container.TabLocationTop)
			},
		),
	)
	leftRightCentered := container.New(
		layout.NewCenterLayout(),
		container.NewHBox(
			container.NewHBox(
				widget.NewButtonWithIcon(
					"", // "Left",
					theme.NavigateBackIcon(),
					func() {
						panelContent.screen.Layout.Tabbar().SetTabLocation(container.TabLocationLeading)
					},
				),
				widget.NewButtonWithIcon(
					"", // "Right",
					theme.NavigateNextIcon(),
					func() {
						panelContent.screen.Layout.Tabbar().SetTabLocation(container.TabLocationTrailing)
					},
				),
			),
		),
	)
	bottomCentered := container.New(
		layout.NewCenterLayout(),
		widget.NewButtonWithIcon(
			"", // "Bottom",
			theme.MoveDownIcon(),
			func() {
				panelContent.screen.Layout.Tabbar().SetTabLocation(container.TabLocationBottom)
			},
		),
	)
	innerVBox := container.NewVBox(
		panelContent.heading,
		panelContent.description,
		topCentered,
		leftRightCentered,
		bottomCentered,
	)
	// SuccessImportance
	dismissButton := widget.NewButton(
		"Dismiss",
		func() {
			panelContent.screen.Layout.Back()
		},
	)
	dismissButton.Importance = widget.LowImportance
	panelContent.content = container.New(
		layout.NewCenterLayout(),
		container.NewVBox(
			innerVBox,
			dismissButton,
		),
	)
	return
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
