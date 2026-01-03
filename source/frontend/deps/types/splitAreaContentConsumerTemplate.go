package types

const (
	splitAreaContentConsumerFileName = "splitAreaContentConsumer.go"
	splitAreaContentConsumerTemplate = `package types

import (
	"fyne.io/fyne/v2"
)

type SplitAreaPosition int

const (
	SplitPositionLeading SplitAreaPosition = iota
	SplitPositionTrailing
)

// SplitAreaContentConsumer consumes content from a producer and gives it to a area.
// It is implemented by a tab item.
// SplitAreaContentConsumer implements ContentConsumer.
// SplitAreaContentConsumer implements UnSpawner.
type SplitAreaContentConsumer struct {
	position SplitAreaPosition

	// producer makes the content for Split.Leading or Split.Trailing
	producer      ContentProducer       // A panel's content producer or a screen's content producer.
	splitProducer *SplitContentProducer // Treats the two area producers as a single split producer.
	needToRefresh bool
}

func NewSplitAreaContentConsumer() (consumer *SplitAreaContentConsumer) {
	consumer = &SplitAreaContentConsumer{}
	return
}

// SetPosition sets the area to leading or trailing.
func (consumer *SplitAreaContentConsumer) SetPosition(position SplitAreaPosition) {
	consumer.position = position
}

// SetSplitProducer sets the consumer's split.
func (consumer *SplitAreaContentConsumer) SetSplitProducer(splitProducer *SplitContentProducer) {
	consumer.splitProducer = splitProducer
}

// ContentConsumer implementations.

// Show shows the Area's content.
// Show is the implementation of ScreenCanvasWatcher.
func (consumer *SplitAreaContentConsumer) Show(isMainThread bool) {}

// IsVisible returns if this content is visible in the window.
// IsVisible is the implementation of ContentConsumer.
func (consumer *SplitAreaContentConsumer) IsVisible() (is bool) {
	is = consumer.splitProducer.IsVisible()
	return
}

func (consumer *SplitAreaContentConsumer) CanvasObject() (canvasObject fyne.CanvasObject) {
	if consumer.needToRefresh {
		consumer.refreshArea()
	}
	canvasObject = consumer.producer.CanvasObject(consumer)
	return
}

// Refresh: called by the area's content producer.
// Triggers a total refresh.
func (consumer *SplitAreaContentConsumer) Refresh(isMainThread bool) {
	if consumer.producer == nil {
		consumer.needToRefresh = true
		return
	}
	if consumer.splitProducer == nil {
		consumer.needToRefresh = true
		return
	}
	consumer.splitProducer.RefreshSplit(isMainThread)
}

// RefreshSplitArea: called by split's content producer.
// 1. Moves content from the producer to the split area.
func (consumer *SplitAreaContentConsumer) RefreshSplitArea(isMainThread bool) {
	if consumer.producer == nil {
		consumer.needToRefresh = true
		return
	}
	consumer.refreshArea()
}

// refreshArea:
// 1. Moves content from the producer to the split area.
func (consumer *SplitAreaContentConsumer) refreshArea() {
	canvasObject := consumer.producer.CanvasObject(consumer)
	if canvasObject == nil {
		return
	}
	consumer.needToRefresh = false
	switch consumer.position {
	case SplitPositionLeading:
		consumer.splitProducer.SetLeading(canvasObject)
	case SplitPositionTrailing:
		consumer.splitProducer.SetTrailing(canvasObject)
	}
}

// Bind binds to the producer and calls the panel or screen's Producer().Bind().
// Bind is the implementation of ContentConsumer.
// func (consumer *SplitAreaContentConsumer) BindSplitAreaProducer(producer ContentProducer) {
// 	if consumer.producer != nil {
// 		return
// 	}
// 	consumer.producer = producer
// }

// Bind binds to the producer and calls the panel or screen's Producer().Bind().
// Bind is the implementation of ContentConsumer.
func (consumer *SplitAreaContentConsumer) Bind(producer ContentProducer) {
	if consumer.producer != nil {
		return
	}
	consumer.producer = producer
	producer.Bind(consumer)
}

// UnBind calls the producer's UnBind().
// UnBind is the implementation of ContentConsumer.
func (consumer *SplitAreaContentConsumer) UnBind() {
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
func (consumer *SplitAreaContentConsumer) IsWindowContentConsumer() (is bool) {
	return
}

func (consumer *SplitAreaContentConsumer) CanUnBind() (canUnBind bool) {
	canUnBind = true
	return
}

// Producer returns the consumer's content producer.
func (consumer *SplitAreaContentConsumer) Producer() (producer ContentProducer) {
	producer = consumer.producer
	return
}
`
)
