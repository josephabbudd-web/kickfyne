package screenmap

type screenMapTemplateData struct {
	ImportPrefix string
}

const (
	screenMapFileName = "screenmap.go"

	screenMapNoBETemplate = `{{ $DOT := . -}}
package screenmap

import (
	"context"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	_types_ "{{ .ImportPrefix }}/frontend/deps/types"
)

type API struct {
	NewBorderAreaContentConsumer func(
		ctx context.Context,
		ctxCancel context.CancelFunc,
		app fyne.App,
		window fyne.Window,
		border *fyne.Container,
		areaIndex int,
		preset any,
	) (
		borderAreaContentConsumer *_types_.BorderAreaContentConsumer,
		screenID string, // id for the caller's borderArea that this screen is content for.
		err error,
	)

	NewWindowContentConsumer func(
		ctx context.Context,
		ctxCancel context.CancelFunc,
		app fyne.App,
		w fyne.Window,
		isInMainMenu bool,
		screenPreset any,
	) (
		windowContentConsumer *_types_.WindowContentConsumer,
		screenID string,
		err error,
	)

	NewAppTabsTabItemContentConsumer func(
		ctx context.Context,
		ctxCancel context.CancelFunc,
		app fyne.App,
		w fyne.Window,
		tabbar *container.AppTabs,
		tabItem *container.TabItem,
		screenPreset any,
	) (
		appTabsTabItemContentConsumer *_types_.AppTabsTabItemContentConsumer,
		screenID string, // id for the caller's appTabsItem that this screen is content for.
		err error,
	)

	NewDocTabsTabItemContentConsumer func(
		ctx context.Context,
		ctxCancel context.CancelFunc,
		app fyne.App,
		w fyne.Window,
		tabbar *container.DocTabs,
		tabItem *container.TabItem,
		screenPreset any,
	) (
		docTabsTabItemContentConsumer *_types_.DocTabsTabItemContentConsumer,
		screenID string, // id for the caller's docTabsItem that this screen is content for.	
		err error,
	)

	NewAccordionItemContentConsumer func(
		ctx context.Context,
		ctxCancel context.CancelFunc,
		app fyne.App,
		w fyne.Window,
		accordion *widget.Accordion,
		accordionItem *widget.AccordionItem,
		screenPreset any,
	) (
		accordionItemContentConsumer *_types_.AccordionItemContentConsumer,
		screenID string, // id for the caller's accordionItem that this screen is content for.	
		err error,
	)

}

var Map = make(map[string]*API)
var PresetsMap = make(map[string]map[string]any)
`
)
