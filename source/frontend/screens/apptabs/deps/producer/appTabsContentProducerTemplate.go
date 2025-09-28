package producer

type AppTabsContentProducerTemplateData struct {
	ImportPrefix string
}

const (
	AppTabsContentProducerFileName = "appTabs.go"
	AppTabsContentProducerTemplate = `package producer

import (
	"fyne.io/fyne/v2"

	_types_ "{{ .ImportPrefix }}/frontend/types"
)

type appTabsContentProducerState struct {
	hasNewCanvasObject bool
	hasNewTitle        bool
}

type AppTabsContentProducer struct {
	canvasObject fyne.CanvasObject
	title        string
	consumers    map[_types_.ContentConsumer]*appTabsContentProducerState
}

func NewAppTabsContentProducer(consumer _types_.ContentConsumer) (producer *AppTabsContentProducer) {
	producer = &AppTabsContentProducer{
		consumers: make(map[_types_.ContentConsumer]*appTabsContentProducerState),
	}
	consumer.Bind(producer)
	producer.Bind(consumer)
	return
}

func (producer *AppTabsContentProducer) HasWindowConsumer() (has bool) {
	for consumer := range producer.consumers {
		if has = consumer.IsWindowContentConsumer(); has {
			break
		}
	}
	return
}

func (producer *AppTabsContentProducer) Refresh(isMainThread bool) {
	for consumer := range producer.consumers {
		consumer.Refresh(isMainThread)
	}
}

func (producer *AppTabsContentProducer) SetCanvasObject(canvasObject fyne.CanvasObject) {
	producer.canvasObject = canvasObject
	for _, stat := range producer.consumers {
		stat.hasNewCanvasObject = true
	}
}

func (producer *AppTabsContentProducer) SetIcon(icon fyne.Resource) {}

func (producer *AppTabsContentProducer) SetTitle(title string) {
	producer.title = title
	for _, stat := range producer.consumers {
		stat.hasNewTitle = true
	}
}

func (producer *AppTabsContentProducer) SetLabel(label string) {}

// Implementations of _types_.AppTabsContentProducer.

func (producer *AppTabsContentProducer) CanvasObjectForce(consumer _types_.ContentConsumer) (canvasObject fyne.CanvasObject) {
	canvasObject = producer.canvasObject
	return
}

func (producer *AppTabsContentProducer) CanvasObject(consumer _types_.ContentConsumer) (canvasObject fyne.CanvasObject) {
	if stats, found := producer.consumers[consumer]; found {
		if stats.hasNewCanvasObject {
			stats.hasNewCanvasObject = false
			canvasObject = producer.canvasObject
		}
	}
	return
}

func (producer *AppTabsContentProducer) Title(consumer _types_.ContentConsumer) (title *string) {
	if stats, found := producer.consumers[consumer]; found {
		if stats.hasNewTitle {
			stats.hasNewTitle = false
			title = &producer.title
		}
	}
	return
}

func (producer *AppTabsContentProducer) Label(consumer _types_.ContentConsumer) (label *string) {
	return
}

func (producer *AppTabsContentProducer) Icon(consumer _types_.ContentConsumer) (icon fyne.Resource) {
	return
}

func (producer *AppTabsContentProducer) Bind(consumer _types_.ContentConsumer) {
	if producer.consumers[consumer] != nil {
		return
	}
	producer.consumers[consumer] = &appTabsContentProducerState{}
	consumer.Bind(producer)
}

func (producer *AppTabsContentProducer) UnBind(consumer _types_.ContentConsumer) {
	if producer.consumers[consumer] == nil {
		return
	}
	delete(producer.consumers, consumer)
	consumer.UnBind()
}

func (producer *AppTabsContentProducer) Die() {
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
// IsVisible is the implementation of AppTabsContentProducer.
func (producer *AppTabsContentProducer) IsVisible() (is bool) {
	for consumer := range producer.consumers {
		if is = consumer.IsVisible(); is {
			break
		}
	}
	return
}
`
)
