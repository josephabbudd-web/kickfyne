package simple

import (
	_utils_ "github.com/josephabbudd-web/kickfyne/source/utils"
)

type aPITemplateData struct {
	PackageName      string
	LocalPanelNames  []string
	DefaultPanelName string
	ImportPrefix     string
	Funcs            _utils_.Funcs
}

const (
	// aPIFileName = _utils_.ScreenFileName
	aPIFileName = _utils_.APIFileName

	aPINoBETemplate = `{{ $DOT := . -}}
package {{ call .Funcs.LowerCase .PackageName }}

import (
	"context"
	"errors"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	_misc_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/misc"
	_panels_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/panels"
	_startup_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/startup"

	_layout_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/deps/layout"
	_panelers_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/deps/panelers"
	_producer_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/deps/producer"

	_types_ "{{ .ImportPrefix }}/frontend/types"
	_thread_ "{{ .ImportPrefix }}/deps/thread"
)

type InitData  _startup_.ScreenInitData
func NewInitData() (initData *_startup_.ScreenInitData) {
	initData = _startup_.NewScreenInitData()
	return
}

func Presets() (presets map[string]any) {
	presets = make(map[string]any)
	for k, v := range _startup_.Presets {
		presets[k] = v
	}
	return
}

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
	startupData any,
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

	if startupData == nil {
		startupData = _startup_.DefaultScreenInitData()
	}
	switch startupData := startupData.(type) {
	case *_startup_.ScreenInitData:
		// Screen ID.
		screenID = fmt.Sprintf("{{ .PackageName }}:Window:%d", nextScreenCount())
		// Consumer.
		windowContentConsumer = _types_.NewWindowContentConsumer(window, isInMainMenu)
		// Screen.
		var packageScreen *_misc_.Miscellaneous
		if packageScreen, err = newScreen(ctx, ctxCancel, app, window, windowContentConsumer, screenID, startupData); err != nil {
			return
		}

		// This screen only show 1 of it's panels at a time.
		// Show the default panel.
		isMainThread := _thread_.IsMainThread()
		packageScreen.Panelers.DefaultPanel.Show(isMainThread)
		packageScreen.Layout.Producer().Refresh(isMainThread)
	default:
		err = errors.New("startupData is not a *_startup_.ScreenInitData")
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
	startupData any,
) (
	appTabsTabItemContentConsumer *_types_.AppTabsTabItemContentConsumer,
	screenID string, // id for the caller's appTabsItem that this screen is content for.	
	err error,
) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("{{ .PackageName }}.NewAppTabsTabItemContentConsumer: %w", err)
		}
	}()

	if startupData == nil {
		startupData = _startup_.DefaultScreenInitData()
	}
	switch startupData := startupData.(type) {
	case *_startup_.ScreenInitData:
		// Screen ID.
		screenID = fmt.Sprintf("{{ .PackageName }}:AppTabsTabItem:%d", nextScreenCount())
		// Consumer.
		appTabsTabItemContentConsumer = _types_.NewAppTabsTabItemContentConsumer(appTabs, tabItem)
		// Screen.
		var packageScreen *_misc_.Miscellaneous
		if packageScreen, err = newScreen(ctx, ctxCancel, app, window, appTabsTabItemContentConsumer, screenID, startupData); err != nil {
			return
		}

		// This screen only show 1 of it's panels at a time.
		// Show the default panel.
		isMainThread := _thread_.IsMainThread()
		packageScreen.Panelers.DefaultPanel.Show(isMainThread)
		packageScreen.Layout.Producer().Refresh(isMainThread)
	default:
		err = errors.New("startupData is not a *_startup_.ScreenInitData")
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
	startupData any,
) (
	docTabsTabItemContentConsumer *_types_.DocTabsTabItemContentConsumer,
	screenID string, // id for the caller's docTabsItem that this screen is content for.	
	err error,
) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("{{ .PackageName }}.NewDocTabsTabItemContentConsumer: %w", err)
		}
	}()

	if startupData == nil {
		startupData = _startup_.DefaultScreenInitData()
	}
	switch startupData := startupData.(type) {
	case *_startup_.ScreenInitData:
		// Screen ID.
		screenID = fmt.Sprintf("{{ .PackageName }}:DocTabsTabItem:%d", nextScreenCount())
		// Consumer.
		docTabsTabItemContentConsumer = _types_.NewDocTabsTabItemContentConsumer(docTabs, tabItem)
		// Screen.
		var packageScreen *_misc_.Miscellaneous
		if packageScreen, err = newScreen(ctx, ctxCancel, app, window, docTabsTabItemContentConsumer, screenID, startupData); err != nil {
			return
		}

		// This screen only show 1 of it's panels at a time.
		// Show the default panel.
		isMainThread := _thread_.IsMainThread()
		packageScreen.Panelers.DefaultPanel.Show(isMainThread)
		packageScreen.Layout.Producer().Refresh(isMainThread)
	default:
		err = errors.New("startupData is not a *_startup_.ScreenInitData")
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
	startupData any,
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

	if startupData == nil {
		startupData = _startup_.DefaultScreenInitData()
	}
	switch startupData := startupData.(type) {
	case *_startup_.ScreenInitData:
		// Screen ID.
		screenID = fmt.Sprintf("{{ .PackageName }}:AccordionItem:%d", nextScreenCount())
		// Consumer.
		accordionItemContentConsumer = _types_.NewAccordionItemContentConsumer(accordion, accordionItem)
		// Screen.
		var packageScreen *_misc_.Miscellaneous
		if packageScreen, err = newScreen(ctx, ctxCancel, app, window, accordionItemContentConsumer, screenID, startupData); err != nil {
			return
		}

		// This screen only show 1 of it's panels at a time.
		// Show the default panel.
		isMainThread := _thread_.IsMainThread()
		packageScreen.Panelers.DefaultPanel.Show(isMainThread)
		packageScreen.Layout.Producer().Refresh(isMainThread)
	default:
		err = errors.New("startupData is not a *_startup_.ScreenInitData")
	}
	return
}

func newScreen(
	ctx context.Context, ctxCancel context.CancelFunc,
	app fyne.App, window fyne.Window,
	consumer _types_.ContentConsumer,
	id string,
	startupData *_startup_.ScreenInitData,
) (screen *_misc_.Miscellaneous, err error) {
	// Build the content & producer.
	producer := _producer_.NewContentProducer(consumer)
	producer.Bind(consumer)

	// Build Layout
	var layout *_layout_.Layout
	if layout, err = _layout_.NewLayout(producer); err != nil {
		return
	}
	if screen, err = _misc_.NewMiscellaneous(ctx, ctxCancel, app, window, layout, id); err != nil {
		return
	}

	// Build each panel.
{{- range $panelName := .LocalPanelNames }}
	// {{ $panelName }} panel.
	var {{ call $DOT.Funcs.DeCap $panelName }}Panel *_panels_.{{ $panelName }}Panel
	if {{ call $DOT.Funcs.DeCap $panelName }}Panel, err = _panels_.New{{ $panelName }}Panel(screen, startupData.{{ $panelName }}Panel); err != nil {
		return
	}
	screen.Panelers.{{ $panelName }} = {{ call $DOT.Funcs.DeCap $panelName }}Panel
{{- end }}
	// Build Panelers.
	screen.Panelers = &_panelers_.Panelers{}
{{- range $panelName := .LocalPanelNames }}
	screen.Panelers.{{ $panelName }} = {{ call $DOT.Funcs.DeCap $panelName }}Panel
{{- end }}
	screen.Panelers.DefaultPanel = {{ call .Funcs.DeCap .DefaultPanelName }}Panel
return
}
`
)
