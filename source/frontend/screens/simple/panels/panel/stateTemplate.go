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

import(
	"fyne.io/fyne/v2"

	_thread_ "{{ .ImportPrefix }}/deps/thread"
	_startup_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/startup"
	_types_ "{{ .ImportPrefix }}/frontend/types"
)

type Getters struct {
	ID          func() string
	Heading     func() string
	Description func() string
}

// State is the state for the {{ .PanelName }} panel.
type State struct {
	id      string
	content *Content

}

// NewState constructs a new content state.
// It may or may not make some initial settings.
func NewState(
	content *Content,
	screenID string,
) (state *State, err error) {
	state = &State{
		content: content,
		id:      screenID + ".{{ .PanelName }}",
	}
	content.state = state
	return
}

// LoadStartupData is called by the panel's constructor.
func (state *State) LoadStartupData(startupData *_startup_.{{ .PanelName }}PanelInitData) {
	state.Set(
		state.SetHeading(startupData.Heading),
		state.SetDescription(startupData.Description),
	)
}

// Refresh is not used in this package.
// It is part of the frontend/types/Stater implementation.
func (state *State) Refresh(isMainThread bool) {}

// Set sets the state.
// It is part of the frontend/types/Stater implementation.
func (state *State) Set(setters ..._types_.StateSetter) {
	isMainThread := _thread_.IsMainThread()
	var refreshCanvasObject bool
	for _, setter := range setters {
		refreshCanvasObject = refreshCanvasObject || setter(isMainThread)
	}
	if refreshCanvasObject {
		state.content.screen.Layout.RefreshIfCurrent(state.id, state.content.content)
	}
}

// Get returns each getter.
// It is part of the frontend/types/Stater implementation.
func (state *State) Get() (getters any) {
	getters = Getters{
		ID:          state.getID,
		Heading:     state.getHeading,
		Description: state.getDescription,
	}
	return
}

// ID is this panel's id.
func (state *State) getID() (id string) {
	id = state.id
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
