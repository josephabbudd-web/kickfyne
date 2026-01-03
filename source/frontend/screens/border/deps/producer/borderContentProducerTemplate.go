package producer

const (
	BorderContentProducerFileName = "border.go"
	BorderContentProducerTemplate = `package producer

import (
	"fyne.io/fyne/v2"

	_types_ "{{ .ImportPrefix }}/frontend/deps/types"
)

type borderContentProducerState struct {
	hasNewCanvasObject bool
	hasNewTitle        bool
}

type BorderContentProducer struct {
	canvasObject fyne.CanvasObject
	title        string
	consumers    map[_types_.ContentConsumer]*borderContentProducerState
}

func NewBorderContentProducer(consumer _types_.ContentConsumer) (producer *BorderContentProducer) {
	producer = &BorderContentProducer{
		consumers: make(map[_types_.ContentConsumer]*borderContentProducerState),
	}
	consumer.Bind(producer)
	producer.Bind(consumer)
	return
}

func (producer *BorderContentProducer) HasWindowConsumer() (has bool) {
	for consumer := range producer.consumers {
		if has = consumer.IsWindowContentConsumer(); has {
			break
		}
	}
	return
}

func (producer *BorderContentProducer) Refresh(isMainThread bool) {
	for consumer := range producer.consumers {
		consumer.Refresh(isMainThread)
	}
}

func (producer *BorderContentProducer) SetCanvasObject(canvasObject fyne.CanvasObject) {
	producer.canvasObject = canvasObject
	for _, stat := range producer.consumers {
		stat.hasNewCanvasObject = true
	}
}

func (producer *BorderContentProducer) SetIcon(icon fyne.Resource) {}

func (producer *BorderContentProducer) SetTitle(title string) {
	producer.title = title
	for _, stat := range producer.consumers {
		stat.hasNewTitle = true
	}
}

func (producer *BorderContentProducer) SetLabel(label string) {}

// Implementations of _types_.BorderContentProducer.

func (producer *BorderContentProducer) CanvasObjectForce(consumer _types_.ContentConsumer) (canvasObject fyne.CanvasObject) {
	canvasObject = producer.canvasObject
	return
}

func (producer *BorderContentProducer) CanvasObject(consumer _types_.ContentConsumer) (canvasObject fyne.CanvasObject) {
	if stats, found := producer.consumers[consumer]; found {
		if stats.hasNewCanvasObject {
			stats.hasNewCanvasObject = false
			canvasObject = producer.canvasObject
		}
	}
	return
}

func (producer *BorderContentProducer) Title(consumer _types_.ContentConsumer) (title *string) {
	if stats, found := producer.consumers[consumer]; found {
		if stats.hasNewTitle {
			stats.hasNewTitle = false
			title = &producer.title
		}
	}
	return
}

func (producer *BorderContentProducer) Label(consumer _types_.ContentConsumer) (label *string) {
	return
}

func (producer *BorderContentProducer) Icon(consumer _types_.ContentConsumer) (icon fyne.Resource) {
	return
}

func (producer *BorderContentProducer) Bind(consumer _types_.ContentConsumer) {
	if producer.consumers[consumer] != nil {
		return
	}
	producer.consumers[consumer] = &borderContentProducerState{}
	consumer.Bind(producer)
}

func (producer *BorderContentProducer) UnBind(consumer _types_.ContentConsumer) {
	if producer.consumers[consumer] == nil {
		return
	}
	delete(producer.consumers, consumer)
	consumer.UnBind()
}

func (producer *BorderContentProducer) Die() {
	// Get a list of consumers from the map.
	consumers := make([]_types_.ContentConsumer, len(producer.consumers))
	var i int = 0
	for consumer := range producer.consumers {
		consumers[i] = consumer
		i++
	}
	for _, consumer := range consumers {
		producer.UnBind(consumer)
	}
}

// IsVisible returns if this content is visible in the window.
// IsVisible is the implementation of BorderContentProducer.
func (producer *BorderContentProducer) IsVisible() (is bool) {
	for consumer := range producer.consumers {
		if is = consumer.IsVisible(); is {
			break
		}
	}
	return
}
`
)
