package types

const (
	borderAreaContentConsumerFileName = "borderCenterAreaContentConsumer.go"
	borderAreaContentConsumerTemplate = `package types

import (
	"fyne.io/fyne/v2"
)

// BorderCenterAreaContentConsumer consumes content from a producer and gives it to a area.
// It is implemented by a tab item.
// BorderCenterAreaContentConsumer implements ContentConsumer.
// BorderCenterAreaContentConsumer implements UnSpawner.
type BorderCenterAreaContentConsumer struct {
	border    *fyne.Container
	areaIndex int //   fyne.CanvasObject

	// producer makes the content for area.
	// 1. Tab icon.
	// 2. Tab label.
	// 3. Tab content.
	producer ContentProducer // A panel's content producer or a screen's content producer.
}

func NewBorderCenterAreaContentConsumer(border *fyne.Container, areaIndex int) (consumer *BorderCenterAreaContentConsumer) {
	consumer = &BorderCenterAreaContentConsumer{
		border:    border,
		areaIndex: areaIndex,
	}
	return
}

// SetBorder sets the consumer's border.
func (consumer *BorderCenterAreaContentConsumer) SetBorder(border *fyne.Container) {
	consumer.border = border
}

// ContentConsumer implementations.

// Show shows the Area's content.
// Show is the implementation of ScreenCanvasWatcher.
func (consumer *BorderCenterAreaContentConsumer) Show(isMainThread bool) {}

// IsVisible returns if this content is visible in the window.
// IsVisible is the implementation of ContentConsumer.
func (consumer *BorderCenterAreaContentConsumer) IsVisible() (is bool) {
	is = consumer.border.Visible()
	return
}

// Refresh:
// 1. Moves content from the producer to the area.
// 2. Refreshes the area.
// 3. Refreshes the tab-bar.
// Refresh is the implementation of ContentConsumer.
func (consumer *BorderCenterAreaContentConsumer) Refresh(isMainThread bool) {
	if canvasObject := consumer.producer.CanvasObject(consumer); canvasObject != nil {
		consumer.border.Objects[consumer.areaIndex] = canvasObject
	}
	if isMainThread {
		consumer.border.Refresh()
	} else {
		fyne.Do(func() { consumer.border.Refresh() })
	}
}

// Bind binds to the producer and calls the panel or screen's Producer().Bind().
// Bind is the implementation of ContentConsumer.
func (consumer *BorderCenterAreaContentConsumer) Bind(producer ContentProducer) {
	if consumer.producer != nil {
		// Already bound to a producer.
		return
	}
	// Bind to the producer.
	consumer.producer = producer
	producer.Bind(consumer)
}

// UnBind calls the producer's UnBind() and then unspawns.
// UnBind is the implementation of ContentConsumer.
func (consumer *BorderCenterAreaContentConsumer) UnBind() {
	if consumer.producer == nil {
		// Not bound to a producer.
		return
	}
	// UnBind from the producer.
	producer := consumer.producer
	consumer.producer = nil
	producer.UnBind(consumer)
}

// IsWindowContentConsumer returns false because this is a area consumer.
// IsWindowContentConsumer is the implementation of ContentConsumer.
func (consumer *BorderCenterAreaContentConsumer) IsWindowContentConsumer() (is bool) {
	return
}

func (consumer *BorderCenterAreaContentConsumer) CanUnBind() (canUnBind bool) {
	canUnBind = true
	return
}
`
)
