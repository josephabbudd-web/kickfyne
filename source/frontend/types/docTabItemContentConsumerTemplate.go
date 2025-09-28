package types

type docTabItemContentConsumerTemplateData struct {
	ImportPrefix string
}

const (
	docTabItemContentConsumerFileName = "docTabItemContentConsumer.go"

	docTabItemContentConsumerTemplate = `package types

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

// DocTabsTabItemContentConsumer consumes content from a producer and gives it to a tabItem.
// It is implemented by a tab item.
// DocTabsTabItemContentConsumer implements ContentConsumer.
// DocTabsTabItemContentConsumer implements UnSpawner.
type DocTabsTabItemContentConsumer struct {
	tabbar  *container.DocTabs
	tabItem *container.TabItem

	// producer makes the content for tabItem.
	// 1. Tab icon.
	// 2. Tab label.
	// 3. Tab content.
	producer ContentProducer // A panel's content producer or a screen's content producer.
}

func NewDocTabsTabItemContentConsumer(tabbar *container.DocTabs, tabItem *container.TabItem) (consumer *DocTabsTabItemContentConsumer) {
	consumer = &DocTabsTabItemContentConsumer{
		tabbar:  tabbar,
		tabItem: tabItem,
	}
	return
}

// ContentConsumer implementations.

// Show sets the TabItem's content.
// Show is the implementation of ScreenCanvasWatcher.
func (consumer *DocTabsTabItemContentConsumer) Show(isMainThread bool) {
	if isMainThread {
		consumer.tabItem.Content.Show()
	} else {
		fyne.Do(consumer.tabItem.Content.Show)
	}
}

// IsVisible returns if this content is visible in the window.
// IsVisible is the implementation of ContentConsumer.
func (consumer *DocTabsTabItemContentConsumer) IsVisible() (is bool) {
	is = consumer.tabItem.Disabled() || (consumer.tabbar.Selected() == consumer.tabItem)
	return
}

// Refresh:
// 1. Moves content from the producer to the tabItem.
// 2. Refreshes the tabItem.
// 3. Refreshes the tab-bar.
// Refresh is the implementation of ContentConsumer.
func (consumer *DocTabsTabItemContentConsumer) Refresh(isMainThread bool) {
	if icon := consumer.producer.Icon(consumer); icon != nil {
		consumer.tabItem.Icon = icon
	}
	if label := consumer.producer.Label(consumer); label != nil {
		consumer.tabItem.Text = *label
	}
	if canvasObject := consumer.producer.CanvasObject(consumer); canvasObject != nil {
		consumer.tabItem.Content = canvasObject
	}
	if isMainThread {
		consumer.tabbar.Refresh()
	} else {
		fyne.Do(func() { consumer.tabbar.Refresh() })
	}
}

// Bind binds to the producer and calls the panel or screen's Producer().Bind().
// Bind is the implementation of ContentConsumer.
func (consumer *DocTabsTabItemContentConsumer) Bind(producer ContentProducer) {
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
func (consumer *DocTabsTabItemContentConsumer) UnBind() {
	if consumer.producer == nil {
		// Not bound to a producer.
		return
	}
	// UnBind from the producer.
	producer := consumer.producer
	consumer.producer = nil
	producer.UnBind(consumer)
}

// IsWindowContentConsumer returns false because this is a tabItem consumer.
// IsWindowContentConsumer is the implementation of ContentConsumer.
func (consumer *DocTabsTabItemContentConsumer) IsWindowContentConsumer() (is bool) {
	return
}

func (consumer *DocTabsTabItemContentConsumer) CanUnBind() (canUnBind bool) {
	canUnBind = true
	return
}
`
)
