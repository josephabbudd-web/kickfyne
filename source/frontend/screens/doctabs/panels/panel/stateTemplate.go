package panel

import (
	_utils_ "github.com/josephabbudd-web/kickfyne/source/utils"
)

const (
	StateFileName = _utils_.StateFileName

	StateTemplate = `package {{ .PanelName }}Panel

import (
	"fyne.io/fyne/v2"

	_thread_ "{{ .ImportPrefix }}/deps/thread"
	_types_ "{{ .ImportPrefix }}/frontend/types"
)

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

// LoadPreset is called by the panel's constructor.
func (state *State) LoadPreset(preset *Preset) {
	var icon fyne.Resource
	if len(preset.TabItemIconName) > 0 {
		icon = state.content.screen.APP.Settings().Theme().Icon(preset.TabItemIconName)
	}
	state.Set(
		state.SetTabIcon(icon),
		state.SetTabLabel(preset.TabItemLabel),
		state.SetHeading(preset.Heading),
		state.SetDescription(preset.Description),
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
		if setter(isMainThread) {
			refreshCanvasObject = true
		}
	}
	if refreshCanvasObject {
		state.content.producer.Refresh(isMainThread)
	}
}

// ID is this panel's id and the tab item's id.
func (state *State) ID() (id string) {
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
