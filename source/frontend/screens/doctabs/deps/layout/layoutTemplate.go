package misc

type LayoutTemplateData struct {
	PackageName  string
	ImportPrefix string
	UseConfigTab bool
}

const (
	LayoutFileName = "layout.go"

	LayoutTemplate = `package layout

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"

	_types_ "{{ .ImportPrefix }}/frontend/deps/types"
	_thread_ "{{ .ImportPrefix }}/deps/thread"
)

// Layout this screen's layout of a container.DocTabs.
type Layout struct {
	tabbarConsumer         _types_.ContentConsumer
	tabItemContentProducer _types_.ContentProducer
	docTabs                *container.DocTabs // This is the canvas object.
	tabItemConsumer        map[*container.TabItem]*_types_.DocTabsTabItemContentConsumer
	tabItemPaneler         map[*container.TabItem]_types_.Paneler
	panelIDPaneler         map[string]_types_.Paneler
	panelIDTabItem         map[string]*container.TabItem
{{- if .UseConfigTab }}
	lastTabItem            *container.TabItem
	configurationTabItem   *container.TabItem
{{- end }}
}

// NewLayout constructs this layout.
func NewLayout(tabbarConsumer _types_.ContentConsumer, tabItemContentProducer _types_.ContentProducer) (layout *Layout, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("{{ .PackageName }}.newLayout: %v", err)
		}
	}()

	docTabs := container.NewDocTabs()
	layout = &Layout{
		tabbarConsumer:         tabbarConsumer,
		tabItemContentProducer: tabItemContentProducer,
		docTabs:                docTabs,
		tabItemConsumer:        make(map[*container.TabItem]*_types_.DocTabsTabItemContentConsumer),
		tabItemPaneler:         make(map[*container.TabItem]_types_.Paneler),
		panelIDPaneler:         make(map[string]_types_.Paneler),
		panelIDTabItem:         make(map[string]*container.TabItem),
	}
	tabItemContentProducer.SetCanvasObject(layout.docTabs)
	layout.SetTabItemHandlers(nil, nil, nil)

return
}

// SetTabItemHandlers sets each item handler.
// Param callBackCloseIntercept returns if the tab must be closed.
func (layout *Layout) SetTabItemHandlers(
	callBackCloseIntercept func(*container.TabItem) (closeTab bool),
	callBackOnSelected func(*container.TabItem),
	callBackOnUnselected func(*container.TabItem),
) {
	// CloseIntercept.
	// Determine if the tab should be closed.
	// Properly close the tab.
	layout.docTabs.CloseIntercept = func(tabItem *container.TabItem) {
{{- if .UseConfigTab }}
		if layout.configurationTabItem == tabItem {
			// Can't close the configuration tab.
			return
		}
{{- end }}
		// func callBackCloseIntercept returns if the tab must be closed.
		if callBackCloseIntercept != nil {
			if !callBackCloseIntercept(tabItem) {
				// Don't close this tab.
				return
			}
		}
		// Close the tab.
		if id, found := layout.TabID(tabItem); found {
			layout.RemoveID(id)
		}
	}
	// OnSelected.
	layout.docTabs.OnSelected = callBackOnSelected
	// OnUnselected.
	layout.docTabs.OnUnselected = func(tabItem *container.TabItem) {
{{- if .UseConfigTab }}
		layout.lastTabItem = tabItem
{{- end }}
		if callBackOnUnselected != nil {
			callBackOnUnselected(tabItem)
		}
	}

}

{{- if .UseConfigTab }}

func (layout *Layout) Back() {
	if layout.lastTabItem != nil {
		layout.docTabs.Select(layout.lastTabItem)
		layout.lastTabItem = nil
		return
	}
	for _, tabItem := range layout.docTabs.Items {
		if tabItem == layout.configurationTabItem {
			continue
		}
		if tabItem.Disabled() {
			continue
		}
		layout.docTabs.Select(tabItem)
		return
	}
}

func (layout *Layout) AddConfigurationTabItemConsumer(paneler _types_.Paneler, tabItem *container.TabItem, consumer *_types_.DocTabsTabItemContentConsumer) {
	layout.AddPanelerTabItemConsumer(paneler, tabItem, consumer)
	layout.configurationTabItem = tabItem
}
{{- end }}

func (layout *Layout) TabbarConsumer() (tabbarConsumer _types_.ContentConsumer) {
	tabbarConsumer = layout.tabbarConsumer
	return
}

func (layout *Layout) AddPanelerTabItemConsumer(paneler _types_.Paneler, tabItem *container.TabItem, consumer *_types_.DocTabsTabItemContentConsumer) {
	isMainThread := _thread_.IsMainThread()
	if isMainThread {
		layout.docTabs.Append(tabItem)
	} else {
		fyne.DoAndWait(func() { layout.docTabs.Append(tabItem) })
	}
	layout.tabItemConsumer[tabItem] = consumer
	layout.tabItemPaneler[tabItem] = paneler
	panelID := paneler.ID()
	layout.panelIDPaneler[panelID] = paneler
	layout.panelIDTabItem[panelID] = tabItem
	paneler.State().Refresh(isMainThread)
}

func (layout *Layout) AddRemoteScreenIDTabItemConsumer(remoteScreenID string, tabItem *container.TabItem, consumer *_types_.DocTabsTabItemContentConsumer) {
	if _thread_.IsMainThread() {
		layout.docTabs.Append(tabItem)
	} else {
		fyne.DoAndWait(func() { layout.docTabs.Append(tabItem) })
	}
	layout.tabItemConsumer[tabItem] = consumer
	layout.tabItemPaneler[tabItem] = nil
	layout.panelIDPaneler[remoteScreenID] = nil
	layout.panelIDTabItem[remoteScreenID] = tabItem
}

func (layout *Layout) TabID(tabItem *container.TabItem) (tabID string, found bool) {
	if paneler := layout.tabItemPaneler[tabItem]; paneler != nil {
		tabID = paneler.ID()
		found = true
	} else {
		for id, item := range layout.panelIDTabItem {
			if item == tabItem {
				tabID = id
				found = true
				break
			}
		}
	}
	return
}

func (layout *Layout) RemoveID(removeID string) {
	var paneler _types_.Paneler
	var tabItem *container.TabItem
	if paneler = layout.panelIDPaneler[removeID]; paneler == nil {
		return
	}
	tabItem = layout.panelIDTabItem[removeID]
{{- if .UseConfigTab }}
	if tabItem == layout.configurationTabItem {
		// Can't remove the configuration tab.
		return
	}
	if tabItem == layout.lastTabItem {
		layout.lastTabItem = nil
	}
{{- end }}
	layout.docTabs.Remove(tabItem)
	paneler.UnBindCleanUP()
	delete(layout.tabItemPaneler, tabItem)
	delete(layout.tabItemConsumer, tabItem)
	panelID := paneler.ID()
	delete(layout.panelIDPaneler, panelID)
	delete(layout.panelIDTabItem, panelID)
}

func (layout *Layout) Producer() (producer _types_.ContentProducer) {
	producer = layout.tabItemContentProducer
	return
}

// Refresh refreshes the tabbar tabs only. Not their content.
func (layout *Layout) Refresh() {
	layout.docTabs.Refresh()
}

func (layout *Layout) Tabbar() (tabbar *container.DocTabs) {
	tabbar = layout.docTabs
	return
}
`
)
