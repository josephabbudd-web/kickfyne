package types

const (
	interfacesFileName = "interfaces.go"

	interfacesNoBETemplate = `package types

import (
	"fyne.io/fyne/v2"
)

type StateSetter func(isMainThread bool) (refreshCanvasObject bool)

type Stater interface {
	Set(setters ...StateSetter)
	Refresh(isMainThread bool)
}

// Panel
type Paneler interface {
	ID() (id string)
	Show(isMainThread bool)
	State() (state Stater)
	CanvasObject() (content fyne.CanvasObject)
	Producer() (producer ContentProducer)
	UnBindCleanUP()
}

type ContentConsumer interface {
	Show(isMainThread bool)
	IsVisible() (is bool)
	Refresh(isMainThread bool)
	Bind(producer ContentProducer)
	UnBind() // Call producer.UnBind(self). Delete self. WindowContentProducer does nothing.
	IsWindowContentConsumer() (is bool)
}

// ContentProducer produces the content for a ContentConsumer.
type ContentProducer interface {
	// Window, TabItem, AccordionItem funcs.
	CanvasObjectForce(consumer ContentConsumer) (canvasObject fyne.CanvasObject)
	Bind(consumer ContentConsumer)
	UnBind(consumer ContentConsumer) //Stop using this consumer. Delete the package if no other consumers.
	IsVisible() (is bool)
	Refresh(isMainThread bool)

	// State canvas object.
	CanvasObject(consumer ContentConsumer) (canvasObject fyne.CanvasObject)
	SetCanvasObject(canvasObject fyne.CanvasObject)
	// State window title.
	Title(consumer ContentConsumer) (title *string)
	SetTitle(title string)
	// State TabItem Icon.
	Icon(consumer ContentConsumer) (icon fyne.Resource)
	SetIcon(icon fyne.Resource)
	// State TabItem Label and AccordionItem Label.
	Label(consumer ContentConsumer) (label *string)
	SetLabel(label string)
}
`
)
