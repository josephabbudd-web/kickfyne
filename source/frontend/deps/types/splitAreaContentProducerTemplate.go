package types

const (
	splitAreaContentProducerFileName = "splitAreaContentProducer.go"
	splitAreaContentProducerTemplate = `package types

import (
	"fyne.io/fyne/v2"
)

type SplitAreaContentProducer struct {
	canvasObject fyne.CanvasObject
	title        string
	label        string
	icon         fyne.Resource
	consumer     *SplitAreaContentConsumer
	state        *ProducerState
}

func NewSplitAreaContentProducer(canvasObject fyne.CanvasObject, consumer *SplitAreaContentConsumer) (producer *SplitAreaContentProducer) {
	producer = &SplitAreaContentProducer{
		canvasObject: canvasObject,
		consumer:     consumer,
		state:        &ProducerState{},
	}
	producer.state.SetCanvasObject(producer.canvasObject)
	consumer.Bind(producer)
	producer.Bind(consumer)
	return
}

func (producer *SplitAreaContentProducer) SplitAreaConsumer() (consumer *SplitAreaContentConsumer) {
	consumer = producer.consumer
	return
}

func (producer *SplitAreaContentProducer) Consumer() (consumer ContentConsumer) {
	consumer = producer.consumer
	return
}

func (producer *SplitAreaContentProducer) HasWindowConsumer() (has bool) {
	return
}

func (producer *SplitAreaContentProducer) Refresh(isMainThread bool) {
	producer.consumer.Refresh(isMainThread)
}

// Implementations of SplitAreaContentProducer.

func (producer *SplitAreaContentProducer) CanvasObjectForce(consumer ContentConsumer) (canvasObject fyne.CanvasObject) {
	canvasObject = producer.canvasObject
	return
}

func (producer *SplitAreaContentProducer) Bind(_ ContentConsumer) {}

func (producer *SplitAreaContentProducer) UnBind(consumer ContentConsumer) {
	if consumer != producer.consumer {
		return
	}
	producer.consumer = nil
	consumer.UnBind()
}

func (producer *SplitAreaContentProducer) Die() {
	// Get a list of consumers from the map.
	if producer.consumer != nil {
		producer.UnBind(producer.consumer)
	}
}

// IsVisible returns if this content is visible in the window.
// IsVisible is the implementation of SplitAreaContentProducer.
func (producer *SplitAreaContentProducer) IsVisible() (is bool) {
	is = producer.consumer.IsVisible()
	return
}

// State

func (producer *SplitAreaContentProducer) SetCanvasObject(canvasObject fyne.CanvasObject) {
	producer.canvasObject = canvasObject
	producer.state.SetCanvasObject(producer.canvasObject)
}

func (producer *SplitAreaContentProducer) CanvasObject(consumer ContentConsumer) (canvasObject fyne.CanvasObject) {
	if consumer == producer.consumer {
		canvasObject = producer.state.CanvasObject()
	}
	return
}

func (producer *SplitAreaContentProducer) SetIcon(icon fyne.Resource) {
	producer.icon = icon
	producer.state.SetIcon(producer.icon)
}

func (producer *SplitAreaContentProducer) Icon(consumer ContentConsumer) (icon fyne.Resource) {
	if consumer == producer.consumer {
		icon = producer.state.Icon()
	}
	return
}

func (producer *SplitAreaContentProducer) SetLabel(label string) {
	producer.label = label
	producer.state.SetLabel(&producer.label)
}

func (producer *SplitAreaContentProducer) Label(consumer ContentConsumer) (label *string) {
	if consumer == producer.consumer {
		label = producer.state.Label()
	}
	return
}

func (producer *SplitAreaContentProducer) SetTitle(title string) {
	producer.title = title
	producer.state.SetTitle(&producer.title)
}

func (producer *SplitAreaContentProducer) Title(consumer ContentConsumer) (title *string) {
	if consumer == producer.consumer {
		title = producer.state.Title()
	}
	return
}
`
)
