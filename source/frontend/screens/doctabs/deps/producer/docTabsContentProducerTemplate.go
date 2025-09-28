package producer

type DocTabsContentProducerTemplateData struct {
	ImportPrefix string
}

const (
	DocTabsContentProducerFileName = "docTabs.go"
	DocTabsContentProducerTemplate = `package producer

import (
	"fyne.io/fyne/v2"

	_types_ "{{ .ImportPrefix }}/frontend/types"
)

type docTabsContentProducerState struct {
	hasNewCanvasObject bool
	hasNewTitle        bool
}

type DocTabsContentProducer struct {
	canvasObject fyne.CanvasObject
	title        string
	consumers    map[_types_.ContentConsumer]*docTabsContentProducerState
}

func NewDocTabsContentProducer(consumer _types_.ContentConsumer) (producer *DocTabsContentProducer) {
	producer = &DocTabsContentProducer{
		consumers: make(map[_types_.ContentConsumer]*docTabsContentProducerState),
	}
	consumer.Bind(producer)
	producer.Bind(consumer)
	return
}

func (producer *DocTabsContentProducer) HasWindowConsumer() (has bool) {
	for consumer := range producer.consumers {
		if has = consumer.IsWindowContentConsumer(); has {
			break
		}
	}
	return
}

func (producer *DocTabsContentProducer) Refresh(isMainThread bool) {
	for consumer := range producer.consumers {
		consumer.Refresh(isMainThread)
	}
}

func (producer *DocTabsContentProducer) SetCanvasObject(canvasObject fyne.CanvasObject) {
	producer.canvasObject = canvasObject
	for _, stat := range producer.consumers {
		stat.hasNewCanvasObject = true
	}
}

func (producer *DocTabsContentProducer) SetIcon(icon fyne.Resource) {}

func (producer *DocTabsContentProducer) SetTitle(title string) {
	producer.title = title
	for _, stat := range producer.consumers {
		stat.hasNewTitle = true
	}
}

func (producer *DocTabsContentProducer) SetLabel(label string) {}

// Implementations of _types_.DocTabsContentProducer.

func (producer *DocTabsContentProducer) CanvasObjectForce(consumer _types_.ContentConsumer) (canvasObject fyne.CanvasObject) {
	canvasObject = producer.canvasObject
	return
}

func (producer *DocTabsContentProducer) CanvasObject(consumer _types_.ContentConsumer) (canvasObject fyne.CanvasObject) {
	if stats, found := producer.consumers[consumer]; found {
		if stats.hasNewCanvasObject {
			stats.hasNewCanvasObject = false
			canvasObject = producer.canvasObject
		}
	}
	return
}

func (producer *DocTabsContentProducer) Title(consumer _types_.ContentConsumer) (title *string) {
	if stats, found := producer.consumers[consumer]; found {
		if stats.hasNewTitle {
			stats.hasNewTitle = false
			title = &producer.title
		}
	}
	return
}

func (producer *DocTabsContentProducer) Label(consumer _types_.ContentConsumer) (label *string) {
	return
}

func (producer *DocTabsContentProducer) Icon(consumer _types_.ContentConsumer) (icon fyne.Resource) {
	return
}

func (producer *DocTabsContentProducer) Bind(consumer _types_.ContentConsumer) {
	if producer.consumers[consumer] != nil {
		return
	}
	producer.consumers[consumer] = &docTabsContentProducerState{}
	consumer.Bind(producer)
}

func (producer *DocTabsContentProducer) UnBind(consumer _types_.ContentConsumer) {
	if producer.consumers[consumer] == nil {
		return
	}
	delete(producer.consumers, consumer)
	consumer.UnBind()
}

func (producer *DocTabsContentProducer) Die() {
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
// IsVisible is the implementation of DocTabsContentProducer.
func (producer *DocTabsContentProducer) IsVisible() (is bool) {
	for consumer := range producer.consumers {
		if is = consumer.IsVisible(); is {
			break
		}
	}
	return
}
`
)
