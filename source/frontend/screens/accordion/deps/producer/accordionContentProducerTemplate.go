package producer

type AccordionContentProducerTemplateData struct {
	ImportPrefix string
}

const (
	AccordionContentProducerFileName = "accordion.go"
	AccordionContentProducerTemplate = `package producer

import (
	"fyne.io/fyne/v2"

	_types_ "{{ .ImportPrefix }}/frontend/types"
)

type accordionContentProducerState struct {
	hasNewCanvasObject bool
	hasNewTitle        bool
	hasNewLabel        bool
}

type AccordionContentProducer struct {
	canvasObject fyne.CanvasObject
	title        string
	label        string
	consumers    map[_types_.ContentConsumer]*accordionContentProducerState
}

func NewAccordionContentProducer(consumer _types_.ContentConsumer) (producer *AccordionContentProducer) {
	producer = &AccordionContentProducer{
		consumers: make(map[_types_.ContentConsumer]*accordionContentProducerState),
	}
	consumer.Bind(producer)
	producer.Bind(consumer)
	return
}

func (producer *AccordionContentProducer) HasWindowConsumer() (has bool) {
	for consumer := range producer.consumers {
		if has = consumer.IsWindowContentConsumer(); has {
			break
		}
	}
	return
}

func (producer *AccordionContentProducer) Refresh(isMainThread bool) {
	for consumer := range producer.consumers {
		consumer.Refresh(isMainThread)
	}
}

func (producer *AccordionContentProducer) SetCanvasObject(canvasObject fyne.CanvasObject) {
	producer.canvasObject = canvasObject
	for _, stat := range producer.consumers {
		stat.hasNewCanvasObject = true
	}
}

func (producer *AccordionContentProducer) SetIcon(icon fyne.Resource) {}

func (producer *AccordionContentProducer) SetTitle(title string) {
	producer.title = title
	for _, stat := range producer.consumers {
		stat.hasNewTitle = true
	}
}

func (producer *AccordionContentProducer) SetLabel(label string) {
	producer.label = label
	for _, stat := range producer.consumers {
		stat.hasNewLabel = true
	}
}

// Implementations of _types_.AccordionContentProducer.

func (producer *AccordionContentProducer) CanvasObjectForce(consumer _types_.ContentConsumer) (canvasObject fyne.CanvasObject) {
	canvasObject = producer.canvasObject
	return
}

func (producer *AccordionContentProducer) CanvasObject(consumer _types_.ContentConsumer) (canvasObject fyne.CanvasObject) {
	if stats, found := producer.consumers[consumer]; found {
		if stats.hasNewCanvasObject {
			stats.hasNewCanvasObject = false
			canvasObject = producer.canvasObject
		}
	}
	return
}

func (producer *AccordionContentProducer) Title(consumer _types_.ContentConsumer) (title *string) {
	if stats, found := producer.consumers[consumer]; found {
		if stats.hasNewTitle {
			stats.hasNewTitle = false
			title = &producer.title
		}
	}
	return
}

func (producer *AccordionContentProducer) Label(consumer _types_.ContentConsumer) (label *string) {
	if stats, found := producer.consumers[consumer]; found {
		if stats.hasNewLabel {
			stats.hasNewLabel = false
			label = &producer.label
		}
	}
	return
}

func (producer *AccordionContentProducer) Icon(consumer _types_.ContentConsumer) (icon fyne.Resource) {
	return
}

func (producer *AccordionContentProducer) Bind(consumer _types_.ContentConsumer) {
	if producer.consumers[consumer] != nil {
		return
	}
	producer.consumers[consumer] = &accordionContentProducerState{}
	consumer.Bind(producer)
}

func (producer *AccordionContentProducer) UnBind(consumer _types_.ContentConsumer) {
	if producer.consumers[consumer] == nil {
		return
	}
	delete(producer.consumers, consumer)
	consumer.UnBind()
}

func (producer *AccordionContentProducer) Die() {
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
// IsVisible is the implementation of AccordionContentProducer.
func (producer *AccordionContentProducer) IsVisible() (is bool) {
	for consumer := range producer.consumers {
		if is = consumer.IsVisible(); is {
			break
		}
	}
	return
}
`
)
