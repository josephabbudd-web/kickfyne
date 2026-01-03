package types

const (
	producerStateFileName = "producerState.go"
	producerStateTemplate = `package types

import (
	"fyne.io/fyne/v2"
)

type ProducerState struct {
	canvasObject fyne.CanvasObject
	title        *string
	label        *string
	icon         fyne.Resource
}

func (state *ProducerState) SetCanvasObject(canvasObject fyne.CanvasObject) {
	state.canvasObject = canvasObject
}

func (state *ProducerState) CanvasObject() (canvasObject fyne.CanvasObject) {
	canvasObject = state.canvasObject
	state.canvasObject = nil
	return
}

func (state *ProducerState) SetIcon(icon fyne.Resource) {
	state.icon = icon
}

func (state *ProducerState) Icon() (icon fyne.Resource) {
	icon = state.icon
	state.icon = nil
	return
}

func (state *ProducerState) SetTitle(title *string) {
	state.title = title
}

func (state *ProducerState) Title() (title *string) {
	title = state.title
	state.title = nil
	return
}

func (state *ProducerState) SetLabel(label *string) {
	state.label = label
}

func (state *ProducerState) Label() (label *string) {
	label = state.label
	state.label = nil
	return
}
`
)
