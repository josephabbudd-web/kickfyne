package types

const (
	accordionItemContentConsumerFileName = "accordionItemContentConsumer.go"
	accordionItemContentConsumerTemplate = `package types

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// AccordionItemContentConsumer consumes content from a producer and gives it to an accordionItem.
// AccordionItemContentConsumer implements ContentConsumer.
// AccordionItemContentConsumer implements UnSpawner.
type AccordionItemContentConsumer struct {
	accordion     *widget.Accordion
	accordionItem *widget.AccordionItem

	// producer makes the content for accordionItem.
	// 1. accordionItem.Title (label).
	// 2. accordionItem.Detail (content).
	// producer can be from
	// 1. a separate screen.
	// 2. a panel in this screen.
	producer ContentProducer // A panel's content producer or a screen's content producer.
}

func NewAccordionItemContentConsumer(accordion *widget.Accordion, accordionItem *widget.AccordionItem) (consumer *AccordionItemContentConsumer) {
	consumer = &AccordionItemContentConsumer{
		accordion:     accordion,
		accordionItem: accordionItem,
	}
	return
}

// ContentConsumer implementations.

// Show sets the AccordionItem's content.
// Show is the implementation of ContentConsumer.
func (consumer *AccordionItemContentConsumer) Show(isMainThread bool) {
	if isMainThread {
		consumer.accordionItem.Detail.Show()
	} else {
		fyne.Do(consumer.accordionItem.Detail.Show)
	}
}

// IsVisible returns if this content is visible in the window.
// IsVisible is the implementation of ContentConsumer.
func (consumer *AccordionItemContentConsumer) IsVisible() (is bool) {
	is = consumer.accordionItem.Open
	return
}

// Refresh:
// 1. Moves content from the producer to the accordionIItem.
// 2. Has the accordionIItem refresh.
// Refresh is the implementation of ContentConsumer.
func (consumer *AccordionItemContentConsumer) Refresh(isMainThread bool) {
	if label := consumer.producer.Label(consumer); label != nil {
		consumer.accordionItem.Title = *label
	}
	if canvasObject := consumer.producer.CanvasObject(consumer); canvasObject != nil {
		consumer.accordionItem.Detail = canvasObject
	}
	if isMainThread {
		consumer.accordion.Refresh()
	} else {
		fyne.Do(func() { consumer.accordion.Refresh() })
	}
}

// Bind binds to the producer and calls the panel or screen's Producer().Bind().
// Bind is the implementation of ContentConsumer.
func (consumer *AccordionItemContentConsumer) Bind(producer ContentProducer) {
	if consumer.producer != nil {
		// Already bound to a producer.
		return
	}
	// Bind to the producer.
	consumer.producer = producer
	producer.Bind(consumer)
}

// UnBind calls the producer's UnBind() and then unspawns.
// UnBind is the implementation of ContentConsumer.
func (consumer *AccordionItemContentConsumer) UnBind() {
	if consumer.producer == nil {
		// Not bound to a producer.
		return
	}
	producer := consumer.producer
	consumer.producer = nil
	producer.UnBind(consumer)
	// Remove the accordion item from the accordion.
	consumer.accordion.Remove(consumer.accordionItem)
}

// IsWindowContentConsumer returns false because this is a accordionItem consumer.
// IsWindowContentConsumer is the implementation of ContentConsumer.
func (consumer *AccordionItemContentConsumer) IsWindowContentConsumer() (is bool) {
	return
}
`
)
