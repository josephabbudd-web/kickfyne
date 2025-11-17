package producer

const (
	BorderAreaContentProducerFileName = "borderarea.go"
	BorderAreaContentProducerTemplate = `package producer

import (
	"fyne.io/fyne/v2"

	_types_ "{{ .ImportPrefix }}/frontend/deps/types"
)

type AreaContentProducerState struct {
	hasNewCanvasObject bool
	hasNewTitle        bool
	hasNewLabel        bool
	hasNewIcon         bool
}

type AreaContentProducer struct {
	canvasObject fyne.CanvasObject
	title        string
	label        string
	icon         fyne.Resource
	consumers    map[_types_.ContentConsumer]*AreaContentProducerState
}

func NewAreaContentProducer(consumer _types_.ContentConsumer) (producer *AreaContentProducer) {
	producer = &AreaContentProducer{
		consumers: make(map[_types_.ContentConsumer]*AreaContentProducerState),
	}
	consumer.Bind(producer)
	producer.Bind(consumer)
	return
}

func (producer *AreaContentProducer) HasWindowConsumer() (has bool) {
	for consumer := range producer.consumers {
		if has = consumer.IsWindowContentConsumer(); has {
			break
		}
	}
	return
}

func (producer *AreaContentProducer) Refresh(isMainThread bool) {
	for consumer := range producer.consumers {
		consumer.Refresh(isMainThread)
	}
}

func (producer *AreaContentProducer) SetCanvasObject(canvasObject fyne.CanvasObject) {
	producer.canvasObject = canvasObject
	for _, stat := range producer.consumers {
		stat.hasNewCanvasObject = true
	}
}

func (producer *AreaContentProducer) SetIcon(icon fyne.Resource) {
	producer.icon = icon
	for _, stat := range producer.consumers {
		stat.hasNewIcon = true
	}
}

func (producer *AreaContentProducer) SetTitle(title string) {
	producer.title = title
	for _, stat := range producer.consumers {
		stat.hasNewTitle = true
	}
}

func (producer *AreaContentProducer) SetLabel(label string) {
	producer.label = label
	for _, stat := range producer.consumers {
		stat.hasNewLabel = true
	}
}

// Implementations of _types_.AreaContentProducer.

func (producer *AreaContentProducer) CanvasObjectForce(consumer _types_.ContentConsumer) (canvasObject fyne.CanvasObject) {
	canvasObject = producer.canvasObject
	return
}

func (producer *AreaContentProducer) CanvasObject(consumer _types_.ContentConsumer) (canvasObject fyne.CanvasObject) {
	if stats, found := producer.consumers[consumer]; found {
		if stats.hasNewCanvasObject {
			stats.hasNewCanvasObject = false
			canvasObject = producer.canvasObject
		}
	}
	return
}

func (producer *AreaContentProducer) Title(consumer _types_.ContentConsumer) (title *string) {
	if stats, found := producer.consumers[consumer]; found {
		if stats.hasNewTitle {
			stats.hasNewTitle = false
			title = &producer.title
		}
	}
	return
}

func (producer *AreaContentProducer) Label(consumer _types_.ContentConsumer) (label *string) {
	if stats, found := producer.consumers[consumer]; found {
		if stats.hasNewLabel {
			stats.hasNewLabel = false
			label = &producer.label
		}
	}
	return
}

func (producer *AreaContentProducer) Icon(consumer _types_.ContentConsumer) (icon fyne.Resource) {
	if stats, found := producer.consumers[consumer]; found {
		if stats.hasNewIcon {
			stats.hasNewIcon = false
			icon = producer.icon
		}
	}
	return
}

func (producer *AreaContentProducer) Bind(consumer _types_.ContentConsumer) {
	if producer.consumers[consumer] != nil {
		return
	}
	producer.consumers[consumer] = &AreaContentProducerState{}
	consumer.Bind(producer)
}

func (producer *AreaContentProducer) UnBind(consumer _types_.ContentConsumer) {
	if producer.consumers[consumer] == nil {
		return
	}
	delete(producer.consumers, consumer)
	consumer.UnBind()
}

func (producer *AreaContentProducer) Die() {
	// Get a list of consumers from the map.
	consumers := make([]_types_.ContentConsumer, len(producer.consumers))
	var i int = 0
	for consumer, _ := range producer.consumers {
		consumers[i] = consumer
		i++
	}
	for _, consumer := range consumers {
		producer.UnBind(consumer)
	}
}

// IsVisible returns if this content is visible in the window.
// IsVisible is the implementation of AreaContentProducer.
func (producer *AreaContentProducer) IsVisible() (is bool) {
	for consumer := range producer.consumers {
		if is = consumer.IsVisible(); is {
			break
		}
	}
	return
}
`
)
