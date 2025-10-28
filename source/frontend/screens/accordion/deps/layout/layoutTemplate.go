package misc

type LayoutTemplateData struct {
	PackageName  string
	ImportPrefix string
}

const (
	LayoutFileName = "layout.go"

	LayoutTemplate = `package layout

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"

	_types_ "{{ .ImportPrefix }}/frontend/types"
	_thread_ "{{ .ImportPrefix }}/deps/thread"
)

// Layout this screen's layout of a widget.AccordionItem.
type Layout struct {
	accordionConsumer             _types_.ContentConsumer
	accordionItemContentProducer  _types_.ContentProducer
	accordion                     *widget.Accordion // This is the canvas object.
	accordionItemConsumer         map[*widget.AccordionItem]*_types_.AccordionItemContentConsumer
	accordionItemPaneler          map[*widget.AccordionItem]_types_.Paneler
	panelIDPaneler                map[string]_types_.Paneler
	panelIDAccordionItem          map[string]*widget.AccordionItem
}

// NewLayout constructs this layout.
func NewLayout(accordionConsumer _types_.ContentConsumer, accordionItemContentProducer _types_.ContentProducer) (layout *Layout, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("{{ .PackageName }}.newLayout: %v", err)
		}
	}()

	layout = &Layout{
		accordionConsumer:            accordionConsumer,
		accordionItemContentProducer: accordionItemContentProducer,
		accordion:                    widget.NewAccordion(),

		accordionItemConsumer:        make(map[*widget.AccordionItem]*_types_.AccordionItemContentConsumer),
		accordionItemPaneler:         make(map[*widget.AccordionItem]_types_.Paneler),
		panelIDPaneler:               make(map[string]_types_.Paneler),
		panelIDAccordionItem:         make(map[string]*widget.AccordionItem),
	}
	accordionItemContentProducer.SetCanvasObject(layout.accordion)

	return
}

func (layout *Layout) AccordionConsumer() (accordionConsumer _types_.ContentConsumer) {
	accordionConsumer = layout.accordionConsumer
	return
}

func (layout *Layout) AddPanelerAccordionItemConsumer(paneler _types_.Paneler, accordionItem *widget.AccordionItem, consumer *_types_.AccordionItemContentConsumer) {
	isMainThread := _thread_.IsMainThread()
	if isMainThread {
		layout.accordion.Append(accordionItem)
	} else {
		fyne.DoAndWait(func() { layout.accordion.Append(accordionItem) })
	}
	layout.accordionItemConsumer[accordionItem] = consumer
	layout.accordionItemPaneler[accordionItem] = paneler
	panelID := paneler.ID()
	layout.panelIDPaneler[panelID] = paneler
	layout.panelIDAccordionItem[panelID] = accordionItem
	paneler.State().Refresh(isMainThread)
}

func (layout *Layout) AddRemoteScreenIDAccordionItemConsumer(remoteScreenID string, accordionItem *widget.AccordionItem, consumer *_types_.AccordionItemContentConsumer) {
	if _thread_.IsMainThread() {
		layout.accordion.Append(accordionItem)
	} else {
		fyne.DoAndWait(func() { layout.accordion.Append(accordionItem) })
	}
	layout.accordionItemConsumer[accordionItem] = consumer
	layout.accordionItemPaneler[accordionItem] = nil
	layout.panelIDPaneler[remoteScreenID] = nil
	layout.panelIDAccordionItem[remoteScreenID] = accordionItem
}

func (layout *Layout) AccordionItemID(accordionItem *widget.AccordionItem) (accordionItemID string, found bool) {
	if paneler := layout.accordionItemPaneler[accordionItem]; paneler != nil {
		accordionItemID = paneler.ID()
		found = true
	} else {
		for id, item := range layout.panelIDAccordionItem {
			if item == accordionItem {
				accordionItemID = id
				found = true
				break
			}
		}
	}
	return
}

func (layout *Layout) RemoveID(removeID string) {
	var paneler _types_.Paneler
	var accordionItem *widget.AccordionItem
	if paneler = layout.panelIDPaneler[removeID]; paneler == nil {
		return
	}
	accordionItem = layout.panelIDAccordionItem[removeID]
	layout.accordion.Remove(accordionItem)
	paneler.UnBindCleanUP()
	delete(layout.accordionItemPaneler, accordionItem)
	delete(layout.accordionItemConsumer, accordionItem)
	panelID := paneler.ID()
	delete(layout.panelIDPaneler, panelID)
	delete(layout.panelIDAccordionItem, panelID)
}

func (layout *Layout) Producer() (producer _types_.ContentProducer) {
	producer = layout.accordionItemContentProducer
	return
}

// Refresh refreshes the accordion tabs only. Not their content.
func (layout *Layout) Refresh() {
	layout.accordion.Refresh()
}

func (layout *Layout) Accordion() (accordion *widget.Accordion) {
	accordion = layout.accordion
	return
}
`
)
