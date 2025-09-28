package panel

import (
	_utils_ "github.com/josephabbudd-web/kickfyne/source/utils"
)

type StateTemplateData struct {
	PackageName  string
	PanelName    string
	ImportPrefix string
}

const (
	StateFileName = _utils_.StateFileName

	StateTemplate = `package {{ .PanelName }}Panel

import (
	"fyne.io/fyne/v2"

	_thread_ "{{ .ImportPrefix }}/deps/thread"
	_startup_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/startup"
	_types_ "{{ .ImportPrefix }}/frontend/types"
)

type Getters struct {
	ID          func() string
	TabLabel    func() string
	TabIcon     func() fyne.Resource
	Heading     func() string
	Description func() string
}

// State is the state for the {{ .PanelName }} panel.
// Panel and Tab have the state id.
type State struct {
	id       string
	tabLabel string
	tabIcon  fyne.Resource
	content  *Content
}

// NewState constructs a new content state.
// It may or may not make some initial settings.
func NewState(
	content  *Content,
	screenID string,
) (state *State, err error) {
	state = &State{
		id:      screenID + ".{{ .PanelName }}",
		content: content,
	}
	content.state = state
	return
}

// LoadStartupData is called by the panel's constructor.
func (state *State) LoadStartupData(startupData *_startup_.{{ .PanelName }}PanelInitData) {
	state.Set(
		state.SetTabIcon(startupData.TabItemIcon),
		state.SetTabLabel(startupData.TabItemLabel),
		state.SetHeading(startupData.Heading),
		state.SetDescription(startupData.Description),
	)
}

// Refresh moves the state's new content to the producer.
// This new content was created int func NewContent before the tab item consumer and screen producer were added.
// This is only called in func layout.AddPanelerTabItemConsumer.
// It is part of the frontend/types/Stater implementation.
func (state *State) Refresh(isMainThread bool) {
	producer := state.content.producer
	producer.SetCanvasObject(state.content.content)
	producer.Refresh(isMainThread)
}

// Set sets the state.
// It is part of the frontend/types/Stater implementation.
func (state *State) Set(setters ..._types_.StateSetter) {
	isMainThread := _thread_.IsMainThread()
	var refreshCanvasObject bool
	for _, setter := range setters {
		refreshCanvasObject = refreshCanvasObject || setter(isMainThread)
	}
	if refreshCanvasObject {
		state.content.producer.Refresh(isMainThread)
	}
}

// Get returns each getter.
// It is part of the frontend/types/Stater implementation.
func (state *State) Get() (getters any) {
	getters = Getters{
		ID:          state.getID,
		TabLabel:    state.getTabLabel,
		TabIcon:     state.getTabIcon,
		Heading:     state.getHeading,
		Description: state.getDescription,
	}
	return
}

// The Panel and Tab use this for an ID.
func (state *State) getID() (id string) {
	id = state.id
	return
}

// Tab label.
func (state *State) SetTabLabel(label string) (setter _types_.StateSetter) {
	state.tabLabel = label
	setter = func(isMainThread bool) (refreshCanvasObject bool) {
		if isMainThread {
			state.content.producer.SetLabel(state.tabLabel)
		} else {
			fyne.Do(
				func() {
					state.content.producer.SetLabel(state.tabLabel);
				},
			)
		}
		return
	}
	return
}

func (state *State) getTabLabel() (label string) {
	label = state.tabLabel
	return
}

// Tab icon.
func (state *State) SetTabIcon(icon fyne.Resource) (setter _types_.StateSetter) {
	state.tabIcon = icon
	setter = func(isMainThread bool) (refreshCanvasObject bool) {
		if isMainThread {
			state.content.producer.SetIcon(state.tabIcon)
		} else {
			fyne.Do(
				func() {
					state.content.producer.SetIcon(state.tabIcon);
				},
			)
		}
		return
	}
	return
}

func (state *State) getTabIcon() (icon fyne.Resource) {
	icon = state.tabIcon
	return
}

// Heading is a widget with variable state.

// SetHeading returns a _types_.Setter that sets the content's heading widget's text.
func (state *State) SetHeading(heading string) (setter _types_.StateSetter) {
	setter = func(isMainThread bool) (refreshCanvasObject bool) {
		if isMainThread {
			state.content.heading.Text = heading
		} else {
			fyne.Do(
				func() {
					state.content.heading.Text = heading;
				},
			)
		}
		refreshCanvasObject = true
		return
	}
	return
}

func (state *State) getHeading() (heading string) {
	heading = state.content.heading.Text
	return
}

// Description is a widget with variable state.

// SetDescription returns a _types_.Setter that sets the content's description widget's text.
func (state *State) SetDescription(description string) (setter _types_.StateSetter) {
	setter = func(isMainThread bool) (refreshCanvasObject bool) {
		if isMainThread {
			state.content.description.Text = description
		} else {
			fyne.Do(
				func() {
					state.content.description.Text = description;
				},
			)
		}
		refreshCanvasObject = true
		return
	}
	return
}

func (state *State) getDescription() (description string) {
	description = state.content.description.Text
	return
}
`
)
