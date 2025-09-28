package producer

type ProducerTemplateData struct {
	ImportPrefix string
}

const (
	ProducerFileName = "producer.go"
	ProducerTemplate = `package producer

import (
	"fyne.io/fyne/v2"

	_types_ "{{ .ImportPrefix }}/frontend/types"
)

type consumerStat struct {
	hasNewCanvasObject bool
	hasNewTitle        bool
	hasNewLabel        bool
	hasNewIcon         bool
}

type ContentProducer struct {
	canvasObject fyne.CanvasObject
	title        string
	label        string
	icon         fyne.Resource
	consumers    map[_types_.ContentConsumer]*consumerStat
}

func NewContentProducer(consumer _types_.ContentConsumer) (producer *ContentProducer) {
	producer = &ContentProducer{
		consumers: make(map[_types_.ContentConsumer]*consumerStat),
	}
	consumer.Bind(producer)
	producer.Bind(consumer)
	return
}

func (producer *ContentProducer) HasWindowConsumer() (has bool) {
	for consumer := range producer.consumers {
		if has = consumer.IsWindowContentConsumer(); has {
			break
		}
	}
	return
}

func (producer *ContentProducer) Refresh(isMainThread bool) {
	for consumer := range producer.consumers {
		consumer.Refresh(isMainThread)
	}
}

func (producer *ContentProducer) SetCanvasObject(canvasObject fyne.CanvasObject) {
	producer.canvasObject = canvasObject
	for _, stat := range producer.consumers {
		stat.hasNewCanvasObject = true
	}
}

func (producer *ContentProducer) SetIcon(icon fyne.Resource) {
	producer.icon = icon
	for _, stat := range producer.consumers {
		stat.hasNewIcon = true
	}
}

func (producer *ContentProducer) SetTitle(title string) {
	producer.title = title
	for _, stat := range producer.consumers {
		stat.hasNewTitle = true
	}
}

func (producer *ContentProducer) SetLabel(label string) {
	producer.label = label
	for _, stat := range producer.consumers {
		stat.hasNewLabel = true
	}
}

// Implementations of _types_.ContentProducer.

func (producer *ContentProducer) CanvasObjectForce(consumer _types_.ContentConsumer) (canvasObject fyne.CanvasObject) {
	canvasObject = producer.canvasObject
	return
}

func (producer *ContentProducer) CanvasObject(consumer _types_.ContentConsumer) (canvasObject fyne.CanvasObject) {
	if stats, found := producer.consumers[consumer]; found {
		if stats.hasNewCanvasObject {
			stats.hasNewCanvasObject = false
			canvasObject = producer.canvasObject
		} else if consumer.IsWindowContentConsumer() {
			canvasObject = producer.canvasObject
		}
	}
	return
}

func (producer *ContentProducer) Title(consumer _types_.ContentConsumer) (title *string) {
	if stats, found := producer.consumers[consumer]; found {
		if stats.hasNewTitle {
			stats.hasNewTitle = false
			title = &producer.title
		} else if consumer.IsWindowContentConsumer() {
			title = &producer.title
		}
	}
	return
}

func (producer *ContentProducer) Label(consumer _types_.ContentConsumer) (label *string) {
	if stats, found := producer.consumers[consumer]; found {
		if stats.hasNewLabel {
			stats.hasNewLabel = false
			label = &producer.label
		} else if consumer.IsWindowContentConsumer() {
			label = &producer.label
		}
	}
	return
}

func (producer *ContentProducer) Icon(consumer _types_.ContentConsumer) (icon fyne.Resource) {
	if stats, found := producer.consumers[consumer]; found {
		if stats.hasNewIcon {
			stats.hasNewIcon = false
			icon = producer.icon
		} else if consumer.IsWindowContentConsumer() {
			icon = producer.icon
		}
	}
	return
}

func (producer *ContentProducer) Bind(consumer _types_.ContentConsumer) {
	producer.consumers[consumer] = &consumerStat{}
}

func (producer *ContentProducer) UnBind(consumer _types_.ContentConsumer) {
	delete(producer.consumers, consumer)
}

// IsVisible returns if this content is visible in the window.
// IsVisible is the implementation of ContentProducer.
func (producer *ContentProducer) IsVisible() (is bool) {
	for consumer := range producer.consumers {
		if is = consumer.IsVisible(); is {
			break
		}
	}
	return
}
`
)
