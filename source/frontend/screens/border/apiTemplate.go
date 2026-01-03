package border

import (
	_utils_ "github.com/josephabbudd-web/kickfyne/source/utils"
)

type apiTemplateData struct {
	PackageName  string
	ImportPrefix string
	Funcs        _utils_.Funcs
}

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

	_layout_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/deps/layout"
	_producer_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/deps/producer"
	_border_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/deps/border"
	_misc_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/misc"
	_presetting_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/presetting"
	_types_ "{{ .ImportPrefix }}/frontend/deps/types"
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
		// PackageScreen.
		var packageScreen *_misc_.Miscellaneous
		if _, packageScreen, err = buildLayout(ctx, ctxCancel, app, window, windowContentConsumer, screenID); err != nil {
			return
		}
		err = _border_.Construct(packageScreen, preset)
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
		// PackageScreen.
		var packageScreen *_misc_.Miscellaneous
		if tabItem.Content, packageScreen, err = buildLayout(ctx, ctxCancel, app, window, appTabsTabItemContentConsumer, screenID); err != nil {
			return
		}
		// Layout the border areas.
		if err = _border_.Construct(packageScreen, preset); err != nil {
			return
		}
		tabItem.Content = packageScreen.Layout.Border()
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
		// PackageScreen.
		var packageScreen *_misc_.Miscellaneous
		if tabItem.Content, packageScreen, err = buildLayout(ctx, ctxCancel, app, window, docTabsTabItemContentConsumer, screenID); err != nil {
			return
		}
		// Layout the border areas.
		if err = _border_.Construct(packageScreen, preset); err != nil {
			return
		}
		tabItem.Content = packageScreen.Layout.Border()
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
		// PackageScreen.
		var packageScreen *_misc_.Miscellaneous
		if _, packageScreen, err = buildLayout(ctx, ctxCancel, app, window, accordionItemContentConsumer, screenID); err != nil {
			return
		}
		// Layout the border areas.
		err = _border_.Construct(packageScreen, preset)
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
			err = fmt.Errorf("{{ .PackageName }}.NewBorderCenterAreaContentConsumer: %w", err)
		}
	}()

	switch preset := preset.(type) {
	case *_presetting_.Preset:
		// Screen ID.
		screenID = fmt.Sprintf("{{ .PackageName }}:BorderCenterArea:%d", nextScreenCount())
		// Consumer.
		borderAreaContentConsumer = _types_.NewBorderCenterAreaContentConsumer(border, areaIndex)
		// PackageScreen.
		var packageScreen *_misc_.Miscellaneous
		if _, packageScreen, err = buildLayout(ctx, ctxCancel, app, window, borderAreaContentConsumer, screenID); err != nil {
			return
		}
		// Layout the border areas.
		err = _border_.Construct(packageScreen, preset)
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
		// PackageScreen.
		var packageScreen *_misc_.Miscellaneous
		if _, packageScreen, err = buildLayout(ctx, ctxCancel, app, window, splitAreaContentConsumer, screenID); err != nil {
			return
		}
		// Layout the border areas.
		err = _border_.Construct(packageScreen, preset)
	}
	return
}

func buildLayout(
	ctx context.Context, ctxCancel context.CancelFunc,
	app fyne.App, window fyne.Window,
	consumer _types_.ContentConsumer,
	screenID string,
) (content fyne.CanvasObject, screen *_misc_.Miscellaneous, err error) {
	// Build the AppTabs content producer.
	borderProducer := _producer_.NewBorderContentProducer(consumer)
	consumer.Bind(borderProducer)
	// Build Layout
	layout := _layout_.NewLayout(consumer, borderProducer)
	screen, err = _misc_.NewMiscellaneous(ctx, ctxCancel, app, window, layout, screenID)
	return
}
`
)
