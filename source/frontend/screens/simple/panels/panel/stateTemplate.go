package panel

import (
	_utils_ "github.com/josephabbudd-web/kickfyne/source/utils"
)

const (
	StateFileName = _utils_.StateFileName

	StateTemplate = `package {{ .PanelName }}Panel

import(
	"fyne.io/fyne/v2"

	_thread_ "{{ .ImportPrefix }}/deps/thread"
	_types_ "{{ .ImportPrefix }}/frontend/deps/types"
)

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

// LoadPreset is called by the panel's constructor.
func (state *State) LoadPreset(preset *Preset) {
	state.Set(
		state.SetHeading(preset.Heading),
		state.SetDescription(preset.Description),
	)
}

// Refresh is not used in this package.
// It is part of the frontend/deps/types/Stater implementation.
func (state *State) Refresh(isMainThread bool) {}

// Set sets the state.
// It is part of the frontend/deps/types/Stater implementation.
func (state *State) Set(setters ..._types_.StateSetter) {
	isMainThread := _thread_.IsMainThread()
	var refreshCanvasObject bool
	for _, setter := range setters {
		if setter(isMainThread) {
			refreshCanvasObject = true
		}
	}
	if refreshCanvasObject {
		state.content.screen.Layout.RefreshIfCurrent(state.id, state.content.content, isMainThread)
	}
}

// ID is this panel's id.
func (state *State) ID() (id string) {
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

/* An example setter.
	func SetMyContentMemberAndFyneWidget() (setter _types_.StateSetter) {
		setWithLock := func() {
			state.mutex.Lock()
			defer state.mutex.Unlock()

			// anyBoolMember is not a widget.
			state.content.anyBoolMember = true
		}
		setFyne := func() {
			// Fyne widgets.
			state.content.heading.SetText(headingText)
		}
		setter = func(isMainThread bool) (refreshCanvasObject bool) {
			refreshCanvasObject = true // The view is updated so refresh.
			setWithLock()
			if isMainThread {
				setFyne()
			} else {
				fyne.DoAndWait(setFyne)
			}
			return
		}
		return
	}
*/
`
)
