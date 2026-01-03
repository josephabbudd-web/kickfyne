package types

const (
	splitContentProducerFileName = "splitContentProducer.go"
	splitContentProducerTemplate = `package types

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"

	_thread_ "{{ .ImportPrefix }}/deps/thread"
)

// SplitContentProducer consumes content from it's 2 SplitAreaContentConsumers and gives it to it's widget.Split.
// SplitContentProducer gives it's widget.Split.CanvasObject to the Conumer it's bound to.
// SplitContentProducer implements ContentConsumer.
// SplitContentProducer implements UnSpawner.
type SplitContentProducer struct {
	split    *container.Split
	leading  *SplitAreaContentConsumer
	trailing *SplitAreaContentConsumer

	title     string
	label     string
	icon      fyne.Resource
	consumers map[ContentConsumer]*ProducerState
}

func NewSplitContentProducer(split *container.Split, leading, trailing *SplitAreaContentConsumer) (producer *SplitContentProducer) {
	producer = &SplitContentProducer{
		split:     split,
		leading:   leading,
		trailing:  trailing,
		consumers: make(map[ContentConsumer]*ProducerState),
	}
	leading.SetSplitProducer(producer)
	trailing.SetSplitProducer(producer)
	return
}

// // SetLeading sets the consumer's split's trailing canvasObject.
func (producer *SplitContentProducer) SetLeading(canvasObject fyne.CanvasObject) {
	producer.split.Leading = nil
	producer.split.Leading = canvasObject
}

// // SetTrailing sets the consumer's split's trailing canvasObject.
func (producer *SplitContentProducer) SetTrailing(canvasObject fyne.CanvasObject) {
	producer.split.Trailing = nil
	producer.split.Trailing = canvasObject
}

func (producer *SplitContentProducer) HasWindowConsumer() (has bool) {
	for consumer := range producer.consumers {
		if has = consumer.IsWindowContentConsumer(); has {
			break
		}
	}
	return
}

// RefreshSplit:
// 1. Moves content from the producer to the area.
// 2. Refreshes the area.
// 3. Refreshes the split.
func (producer *SplitContentProducer) RefreshSplit(isMainThread bool) {
	if producer.leading == nil || producer.trailing == nil {
		return
	}
	producer.leading.RefreshSplitArea(isMainThread)
	producer.trailing.RefreshSplitArea(isMainThread)
	if producer.split.Leading == nil || producer.split.Trailing == nil {
		return
	}
	producer.split.Refresh()

	// Update the consumers.
	for consumer, state := range producer.consumers {
		state.SetCanvasObject(producer.split)
		consumer.Refresh(isMainThread)
	}
}

// Refresh:
// 1. Moves content from the producer to the area.
// 2. Refreshes the area.
// 3. Refreshes the split.
// Refresh is the implementation of ContentConsumer.
func (producer *SplitContentProducer) Refresh(_ bool) {}

// ContentConsumer implementations.

// Show shows the Area's content.
// Show is the implementation of ScreenCanvasWatcher.
func (producer *SplitContentProducer) Show(isMainThread bool) {}

// IsVisible returns if this content is visible in the window.
// IsVisible is the implementation of ContentConsumer.
func (producer *SplitContentProducer) IsVisible() (is bool) {
	is = producer.leading.IsVisible()
	return
}

func (producer *SplitContentProducer) CanvasObjectForce(consumer ContentConsumer) (canvasObject fyne.CanvasObject) {
	if producer.split.Leading == nil || producer.split.Trailing == nil {
		producer.RefreshSplit(_thread_.IsMainThread())
	}
	if producer.split.Leading == nil || producer.split.Trailing == nil {
		return
	}
	canvasObject = producer.split
	return
}

// Bind binds to the producer and calls the panel or screen's Producer().Bind().
// Bind is the implementation of SplitContentProducer.
func (producer *SplitContentProducer) Bind(consumer ContentConsumer) {
	state := &ProducerState{}
	state.SetCanvasObject(producer.split)
	state.SetIcon(producer.icon)
	state.SetLabel(&producer.label)
	state.SetTitle(&producer.title)
	producer.consumers[consumer] = state
	consumer.Bind(producer)
}

// UnBind calls the producer's UnBind().
// UnBind is the implementation of SplitContentProducer.
func (producer *SplitContentProducer) UnBind(consumer ContentConsumer) {
	if producer.consumers[consumer] == nil {
		return
	}
	delete(producer.consumers, consumer)
	consumer.UnBind()
}

func (producer *SplitContentProducer) CanUnBind() (canUnBind bool) {
	canUnBind = true
	return
}

// State

func (producer *SplitContentProducer) SetCanvasObject(canvasObject fyne.CanvasObject) {
	_ = canvasObject
}

func (producer *SplitContentProducer) CanvasObject(consumer ContentConsumer) (canvasObject fyne.CanvasObject) {
	if state, found := producer.consumers[consumer]; found {
		canvasObject = state.CanvasObject()
	}
	return
}

func (producer *SplitContentProducer) SetIcon(icon fyne.Resource) {
	producer.icon = icon
	for _, state := range producer.consumers {
		state.SetIcon(producer.icon)
	}
}

func (producer *SplitContentProducer) Icon(consumer ContentConsumer) (icon fyne.Resource) {
	if state, found := producer.consumers[consumer]; found {
		icon = state.Icon()
	}
	return
}

func (producer *SplitContentProducer) SetLabel(label string) {
	producer.label = label
	for _, state := range producer.consumers {
		state.SetLabel(&producer.label)
	}
}

func (producer *SplitContentProducer) Label(consumer ContentConsumer) (label *string) {
	if state, found := producer.consumers[consumer]; found {
		label = state.Label()
	}
	return
}

func (producer *SplitContentProducer) SetTitle(title string) {
	producer.title = title
	for _, state := range producer.consumers {
		state.SetTitle(&producer.title)
	}
}

func (producer *SplitContentProducer) Title(consumer ContentConsumer) (title *string) {
	if state, found := producer.consumers[consumer]; found {
		title = state.Title()
	}
	return
}
`
)
