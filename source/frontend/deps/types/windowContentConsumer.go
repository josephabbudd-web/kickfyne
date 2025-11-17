package types

const (
	windowContentConsumerFileName = "windowContentConsumer.go"
	windowContentConsumerTemplate = `package types

import (
	"fyne.io/fyne/v2"
)

// currentWindowContentConsumer is the current consumer for the app window content.
var currentWindowContentConsumer *WindowContentConsumer

// WindowContentConsumer consumes content from a producer and gives it to a window.
// WindowContentConsumer implements ContentConsumer
// Only 1 WindowContentConsumer is displayed in the app at a time.
type WindowContentConsumer struct {
	window       fyne.Window
	producer     ContentProducer
	isInMainMenu bool
}

func NewWindowContentConsumer(window fyne.Window, isInMainMenu bool) (consumer *WindowContentConsumer) {
	consumer = &WindowContentConsumer{
		window:       window,
		isInMainMenu: isInMainMenu,
	}
	return
}

// ContentConsumer implementations.

func (consumer *WindowContentConsumer) IsInMainMenu() (isInMainMenu bool) {
	isInMainMenu = consumer.isInMainMenu
	return
}

// Show sets consumer as the window's content.
// Show is the implementation of ContentConsumer.
func (consumer *WindowContentConsumer) Show(isMainThread bool) {
	if isMainThread {
		consumer.show()
	} else {
		fyne.Do(
			func() { consumer.show() },
		)
	}
}

// show this window consumer.
func (consumer *WindowContentConsumer) show() {
	currentWindowContentConsumer = consumer
	title := consumer.producer.Title(consumer)
	if title != nil && len(*title) > 0 {
		consumer.window.SetTitle(*title)
	}
	canvasObject := consumer.producer.CanvasObjectForce(consumer)
	consumer.window.SetContent(canvasObject)
}

// IsVisible returns if this content is visible in the window.
// Show is the implementation of ContentConsumer.
func (consumer *WindowContentConsumer) IsVisible() (is bool) {
	is = currentWindowContentConsumer == consumer
	return
}

// Bind binds to the producer and calls the screen's Producer().Bind().
// Bind is the implementation of ContentConsumer.
func (consumer *WindowContentConsumer) Bind(producer ContentProducer) {
	if consumer.producer != nil {
		// Already bound to a producer.
		return
	}
	// Bind to the producer.
	consumer.producer = producer
	producer.Bind(consumer)
}

func (consumer *WindowContentConsumer) CanUnBind() (canUnBind bool) {
	canUnBind = !consumer.isInMainMenu
	return
}

// UnBind unbinds the consumer from it's producer.
// Bind is the implementation of ContentConsumer.
func (consumer *WindowContentConsumer) UnBind() {
	if consumer.producer == nil {
		// Not bound to a producer.
		return
	}
	// UnBind from the producer.
	producer := consumer.producer
	consumer.producer = nil
	producer.UnBind(consumer)
}

// Refresh refreshes the window if there is something to refresh.
// Refresh is thread safe.
// Refresh is the implementation of ContentConsumer.
func (consumer *WindowContentConsumer) Refresh(isMainThread bool) {
	if isMainThread {
		consumer.show()
	} else {
		fyne.Do(
			func() { consumer.show() },
		)
	}
}

// IsWindowContentConsumer returns true because this is a window consumer.
// IsWindowContentConsumer is the implementation of ContentConsumer.
func (consumer *WindowContentConsumer) IsWindowContentConsumer() (is bool) {
	is = true
	return
}
`
)
