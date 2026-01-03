package producer

const (
	BorderCenterAreaContentProducerFileName = "borderCenterArea.go"
	BorderCenterAreaContentProducerTemplate = `package producer

import (
	"fyne.io/fyne/v2"

	_types_ "{{ .ImportPrefix }}/frontend/deps/types"
)

type BorderCenterAreaContentProducerState struct {
	hasNewCanvasObject bool
	hasNewTitle        bool
	hasNewLabel        bool
	hasNewIcon         bool
}

type BorderCenterAreaContentProducer struct {
	canvasObject fyne.CanvasObject
	title        string
	label        string
	icon         fyne.Resource
	consumers    map[_types_.ContentConsumer]*BorderCenterAreaContentProducerState
}

func NewBorderCenterAreaContentProducer(consumer _types_.ContentConsumer) (producer *BorderCenterAreaContentProducer) {
	producer = &BorderCenterAreaContentProducer{
		consumers: make(map[_types_.ContentConsumer]*BorderCenterAreaContentProducerState),
	}
	consumer.Bind(producer)
	producer.Bind(consumer)
	return
}

func (producer *BorderCenterAreaContentProducer) HasWindowConsumer() (has bool) {
	for consumer := range producer.consumers {
		if has = consumer.IsWindowContentConsumer(); has {
			break
		}
	}
	return
}

func (producer *BorderCenterAreaContentProducer) Refresh(isMainThread bool) {
	for consumer := range producer.consumers {
		consumer.Refresh(isMainThread)
	}
}

func (producer *BorderCenterAreaContentProducer) SetCanvasObject(canvasObject fyne.CanvasObject) {
	producer.canvasObject = canvasObject
	for _, stat := range producer.consumers {
		stat.hasNewCanvasObject = true
	}
}

func (producer *BorderCenterAreaContentProducer) SetIcon(icon fyne.Resource) {
	producer.icon = icon
	for _, stat := range producer.consumers {
		stat.hasNewIcon = true
	}
}

func (producer *BorderCenterAreaContentProducer) SetTitle(title string) {
	producer.title = title
	for _, stat := range producer.consumers {
		stat.hasNewTitle = true
	}
}

func (producer *BorderCenterAreaContentProducer) SetLabel(label string) {
	producer.label = label
	for _, stat := range producer.consumers {
		stat.hasNewLabel = true
	}
}

// Implementations of _types_.BorderCenterAreaContentProducer.

func (producer *BorderCenterAreaContentProducer) CanvasObjectForce(consumer _types_.ContentConsumer) (canvasObject fyne.CanvasObject) {
	canvasObject = producer.canvasObject
	return
}

func (producer *BorderCenterAreaContentProducer) CanvasObject(consumer _types_.ContentConsumer) (canvasObject fyne.CanvasObject) {
	if stats, found := producer.consumers[consumer]; found {
		if stats.hasNewCanvasObject {
			stats.hasNewCanvasObject = false
			canvasObject = producer.canvasObject
		}
	}
	return
}

func (producer *BorderCenterAreaContentProducer) Title(consumer _types_.ContentConsumer) (title *string) {
	if stats, found := producer.consumers[consumer]; found {
		if stats.hasNewTitle {
			stats.hasNewTitle = false
			title = &producer.title
		}
	}
	return
}

func (producer *BorderCenterAreaContentProducer) Label(consumer _types_.ContentConsumer) (label *string) {
	if stats, found := producer.consumers[consumer]; found {
		if stats.hasNewLabel {
			stats.hasNewLabel = false
			label = &producer.label
		}
	}
	return
}

func (producer *BorderCenterAreaContentProducer) Icon(consumer _types_.ContentConsumer) (icon fyne.Resource) {
	if stats, found := producer.consumers[consumer]; found {
		if stats.hasNewIcon {
			stats.hasNewIcon = false
			icon = producer.icon
		}
	}
	return
}

func (producer *BorderCenterAreaContentProducer) Bind(consumer _types_.ContentConsumer) {
	if producer.consumers[consumer] != nil {
		return
	}
	producer.consumers[consumer] = &BorderCenterAreaContentProducerState{}
	consumer.Bind(producer)
}

func (producer *BorderCenterAreaContentProducer) UnBind(consumer _types_.ContentConsumer) {
	if producer.consumers[consumer] == nil {
		return
	}
	delete(producer.consumers, consumer)
	consumer.UnBind()
}

func (producer *BorderCenterAreaContentProducer) Die() {
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
// IsVisible is the implementation of BorderCenterAreaContentProducer.
func (producer *BorderCenterAreaContentProducer) IsVisible() (is bool) {
	for consumer := range producer.consumers {
		if is = consumer.IsVisible(); is {
			break
		}
	}
	return
}
`
)
