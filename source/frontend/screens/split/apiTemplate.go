package split

import (
	_utils_ "github.com/josephabbudd-web/kickfyne/source/utils"
)

const (
	// apiFileName = _utils_.ScreenFileName
	apiFileName = _utils_.APIFileName

	apiNoBETemplate = `{{ $DOT := . -}}
package {{ call .Funcs.LowerCase .PackageName }}

import (
	"context"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	_thread_ "{{ .ImportPrefix }}/deps/thread"
	_types_ "{{ .ImportPrefix }}/frontend/deps/types"
	_layout_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/deps/layout"
	_layoutareas_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/layoutAreas"
	_misc_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/misc"
	_presetting_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/presetting"
)

var screenCount uint = 0

func nextScreenCount() (count uint) {
	count = screenCount
	screenCount++
	return
}

// NewWindowContentConsumer constructs a new screen and returns a window content consumer of the screen's content.
func NewWindowContentConsumer(
	ctx context.Context,
	ctxCancel context.CancelFunc,
	app fyne.App,
	window fyne.Window,
	isInMainMenu bool,
	preset any,
) (
	windowContentConsumer *_types_.WindowContentConsumer,
	screenID string,
	err error,
){
	defer func() {
		if err != nil {
			err = fmt.Errorf("{{ .PackageName }}.NewWindowContentConsumer: %w", err)
		}
	}()

	switch preset := preset.(type) {
	case *_presetting_.Preset:
		// Screen ID.
		screenID = fmt.Sprintf("{{ .PackageName }}:Window:%d", nextScreenCount())
		// Consumer.
		windowContentConsumer = _types_.NewWindowContentConsumer(window, isInMainMenu)
		// Layout.
		_, err = buildLayout(ctx, ctxCancel, app, window, windowContentConsumer, screenID, preset)
	}

	return
}

// NewAppTabsTabItemContentConsumer constructs a new screen and returns a TabItem content consumer of the screen's content.
func NewAppTabsTabItemContentConsumer(
	ctx context.Context,
	ctxCancel context.CancelFunc,
	app fyne.App,
	window fyne.Window,
	appTabs *container.AppTabs,
	tabItem *container.TabItem,
	preset any,
) (
	appTabsTabItemContentConsumer *_types_.AppTabsTabItemContentConsumer,
	screenID string, // id for the caller's appTabsItem that this screen is content for.	
	err error,
) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("{{ .PackageName }}.NewTabItemContentConsumer: %w", err)
		}
	}()

	switch preset := preset.(type) {
	case *_presetting_.Preset:
		// Screen ID.
		screenID = fmt.Sprintf("{{ .PackageName }}:AppTabsTabItem:%d", nextScreenCount())
		// Consumer.
		appTabsTabItemContentConsumer = _types_.NewAppTabsTabItemContentConsumer(appTabs, tabItem)
		// Layout.
		tabItem.Content, err = buildLayout(ctx, ctxCancel, app, window, appTabsTabItemContentConsumer, screenID, preset)
	}

	return
}

// NewDocTabsTabItemContentConsumer constructs a new screen and returns a TabItem content consumer of the screen's content.
func NewDocTabsTabItemContentConsumer(
	ctx context.Context,
	ctxCancel context.CancelFunc,
	app fyne.App,
	window fyne.Window,
	docTabs *container.DocTabs,
	tabItem *container.TabItem,
	preset any,
) (
	docTabsTabItemContentConsumer *_types_.DocTabsTabItemContentConsumer,
	screenID string, // id for the caller's docTabsItem that this screen is content for.	
	err error,
) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("{{ .PackageName }}.NewTabItemContentConsumer: %w", err)
		}
	}()

	switch preset := preset.(type) {
	case *_presetting_.Preset:
		// Screen ID.
		screenID = fmt.Sprintf("{{ .PackageName }}:DocTabsTabItem:%d", nextScreenCount())
		// Consumer.
		docTabsTabItemContentConsumer = _types_.NewDocTabsTabItemContentConsumer(docTabs, tabItem)
		// Layout.
		var content fyne.CanvasObject
		content, err = buildLayout(ctx, ctxCancel, app, window, docTabsTabItemContentConsumer, screenID, preset)
		tabItem.Content = content
	}

	return
}

// NewAccordionItemContentConsumer constructs a new screen and returns a AccordionItem content consumer of the screen's content.
func NewAccordionItemContentConsumer(
	ctx context.Context,
	ctxCancel context.CancelFunc,
	app fyne.App,
	window fyne.Window,
	accordion *widget.Accordion,
	accordionItem *widget.AccordionItem,
	preset any,
) (
	accordionItemContentConsumer *_types_.AccordionItemContentConsumer,
	screenID string, // id for the caller's accordionItem that this screen is content for.	
	err error,
) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("{{ .PackageName }}.NewAccordionItemContentConsumer: %w", err)
		}
	}()

	switch preset := preset.(type) {
	case *_presetting_.Preset:
		// Screen ID.
		screenID = fmt.Sprintf("{{ .PackageName }}:AccordionItem:%d", nextScreenCount())
		// Consumer.
		accordionItemContentConsumer = _types_.NewAccordionItemContentConsumer(accordion, accordionItem)
		// Layout.
		var content fyne.CanvasObject
		content, err = buildLayout(ctx, ctxCancel, app, window, accordionItemContentConsumer, screenID, preset)
		accordionItem.Detail = content
	}

	return
}

// NewBorderCenterAreaContentConsumer constructs a new screen and returns a BorderCenterArea content consumer of the screen's content.
func NewBorderCenterAreaContentConsumer(
	ctx context.Context,
	ctxCancel context.CancelFunc,
	app fyne.App,
	window fyne.Window,
	border *fyne.Container,
	areaIndex int,
	preset any,
) (
	borderAreaContentConsumer *_types_.BorderCenterAreaContentConsumer,
	screenID string, // id for the caller's borderArea that this screen is content for.
	err error,
) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("Split.NewBorderCenterAreaContentConsumer: %w", err)
		}
	}()

	switch preset := preset.(type) {
	case *_presetting_.Preset:
		// Screen ID.
		screenID = fmt.Sprintf("{{ .PackageName }}:BorderCenterArea:%d", nextScreenCount())
		// Consumer.
		borderAreaContentConsumer = _types_.NewBorderCenterAreaContentConsumer(border, areaIndex)
		// Layout the content.
		_, err = buildLayout(ctx, ctxCancel, app, window, borderAreaContentConsumer, screenID, preset)
	}
	return
}

// NewSplitAreaContentConsumer constructs a new screen and returns a SplitArea content consumer of the screen's content.
func NewSplitAreaContentConsumer(
	ctx context.Context,
	ctxCancel context.CancelFunc,
	app fyne.App,
	window fyne.Window,
	preset any,
) (
	splitAreaContentConsumer *_types_.SplitAreaContentConsumer,
	screenID string, // id for the caller's splitArea that this screen is content for.
	err error,
) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("{{ .PackageName }}.NewSplitAreaContentConsumer: %w", err)
		}
	}()

	switch preset := preset.(type) {
	case *_presetting_.Preset:
		// Screen ID.
		screenID = fmt.Sprintf("{{ .PackageName }}:SplitArea:%d", nextScreenCount())
		// Consumer.
		splitAreaContentConsumer = _types_.NewSplitAreaContentConsumer()
		// Layout.
		_, err = buildLayout(ctx, ctxCancel, app, window, splitAreaContentConsumer, screenID, preset)
	}
	return
}

func buildLayout(
	ctx context.Context, ctxCancel context.CancelFunc,
	app fyne.App, window fyne.Window,
	consumer _types_.ContentConsumer,
	screenID string,
	preset *_presetting_.Preset,
) (content fyne.CanvasObject, err error) {
	isMainThread := _thread_.IsMainThread()
	// Build Layout
	layout := _layout_.NewLayout()
	var screen *_misc_.Miscellaneous
	if screen, err = _misc_.NewMiscellaneous(ctx, ctxCancel, app, window, layout, screenID); err != nil {
		return
	}
	// Build the Split content producer.
	var splitContentProducer _types_.ContentProducer
	if splitContentProducer, err = _layoutareas_.Layout(screen, preset, isMainThread); err != nil {
		return
	}
	consumer.Bind(splitContentProducer)
	layout.SetSplitContentProducer(splitContentProducer)
	content = splitContentProducer.CanvasObjectForce(consumer)
	return
}
`
)
