package simple

import (
	_utils_ "github.com/josephabbudd-web/kickfyne/source/utils"
)

type apiPanel struct {
	Name    string
	IsLocal bool
}
type aPITemplateData struct {
	PackageName      string
	AllPanels        []apiPanel
	LocalPanelNames  []string
	AllPanelNames    []string
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

	_layout_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/deps/layout"
	_tabs_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/deps/tabs"
	_producer_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/deps/producer"

	_misc_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/misc"
	_presets_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/presets"

	_types_ "{{ .ImportPrefix }}/frontend/types"
	_ids_ "{{ .ImportPrefix }}/deps/container/{{ .PackageName }}"
)

type InitData  _presets_.ScreenInitData
func NewInitData() (initData *_presets_.ScreenInitData) {
	initData = _presets_.NewScreenInitData()
	return
}

func Presets() (presets map[string]any) {
	presets = make(map[string]any)
	for k, v := range _presets_.Presets {
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
			_ids_.RemoveTabbar(screenID)
		}
	}()

	if startupData == nil {
		startupData = _presets_.DefaultScreenInitData()
	}
	switch startupData := startupData.(type) {
	case *_presets_.ScreenInitData:
		// Screen ID.
		screenID = fmt.Sprintf("{{ .PackageName }}:Window:%d", nextScreenCount())
		// Consumer.
		windowContentConsumer = _types_.NewWindowContentConsumer(window, isInMainMenu)
		// PackageScreen.
		var packageScreen *_misc_.Miscellaneous
		if packageScreen, err = buildLayout(ctx, ctxCancel, app, window, windowContentConsumer, screenID); err != nil {
			return
		}
		// Update deps.
		_ids_.AddTabbar(packageScreen.ScreenID)
		// Get the screen's initializer.
		err = setInitialTabItems(packageScreen, startupData)
	default:
		err = errors.New("startupData is not a *_presets_.ScreenInitData")
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
			err = fmt.Errorf("{{ .PackageName }}.NewTabItemContentConsumer: %w", err)
			_ids_.RemoveTabbar(screenID)
		}
	}()

	if startupData == nil {
		startupData = _presets_.DefaultScreenInitData()
	}
	switch startupData := startupData.(type) {
	case *_presets_.ScreenInitData:
		// Screen ID.
		screenID = fmt.Sprintf("{{ .PackageName }}:TabItem:%d", nextScreenCount())
		// Consumer.
		appTabsTabItemContentConsumer = _types_.NewAppTabsTabItemContentConsumer(appTabs, tabItem)
		// PackageScreen.
		var packageScreen *_misc_.Miscellaneous
		if packageScreen, err = buildLayout(ctx, ctxCancel, app, window, appTabsTabItemContentConsumer, screenID); err != nil {
			return
		}
		// Update deps.
		_ids_.AddTabbar(packageScreen.ScreenID)
		// Get the screen's initializer.
		err = setInitialTabItems(packageScreen, startupData)
	default:
		err = errors.New("startupData is not a *_presets_.ScreenInitData")
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
			err = fmt.Errorf("{{ .PackageName }}.NewTabItemContentConsumer: %w", err)
			_ids_.RemoveTabbar(screenID)
		}
	}()

	if startupData == nil {
		startupData = _presets_.DefaultScreenInitData()
	}
	switch startupData := startupData.(type) {
	case *_presets_.ScreenInitData:
		// Screen ID.
		screenID = fmt.Sprintf("{{ .PackageName }}:TabItem:%d", nextScreenCount())
		// Consumer.
		docTabsTabItemContentConsumer = _types_.NewDocTabsTabItemContentConsumer(docTabs, tabItem)
		// PackageScreen.
		var packageScreen *_misc_.Miscellaneous
		if packageScreen, err = buildLayout(ctx, ctxCancel, app, window, docTabsTabItemContentConsumer, screenID); err != nil {
			return
		}
		// Update deps.
		_ids_.AddTabbar(packageScreen.ScreenID)
		// Get the screen's initializer.
		err = setInitialTabItems(packageScreen, startupData)
	default:
		err = errors.New("startupData is not a *_presets_.ScreenInitData")
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
			_ids_.RemoveTabbar(screenID)
		}
	}()

	if startupData == nil {
		startupData = _presets_.DefaultScreenInitData()
	}
	switch startupData := startupData.(type) {
	case *_presets_.ScreenInitData:
		// Screen ID.
		screenID = fmt.Sprintf("{{ .PackageName }}:AccordionItem:%d", nextScreenCount())
		// Consumer.
		accordionItemContentConsumer = _types_.NewAccordionItemContentConsumer(accordion, accordionItem)
		// PackageScreen.
		var packageScreen *_misc_.Miscellaneous
		if packageScreen, err = buildLayout(ctx, ctxCancel, app, window, accordionItemContentConsumer, screenID); err != nil {
			return
		}
		// Update deps.
		_ids_.AddTabbar(packageScreen.ScreenID)
		// Get the screen's initializer.
		err = setInitialTabItems(packageScreen, startupData)
	default:
		err = errors.New("startupData is not a *_presets_.ScreenInitData")
	}

	return
}

func buildLayout(
	ctx context.Context, ctxCancel context.CancelFunc,
	app fyne.App, window fyne.Window,
	consumer _types_.ContentConsumer,
	screenID string,
) (screen *_misc_.Miscellaneous, err error) {
	// Build the DocTabs content producer.
	docTabsProducer := _producer_.NewDocTabsContentProducer(consumer)
	consumer.Bind(docTabsProducer)

	// Build Layout
	var layout *_layout_.Layout
	if layout, err = _layout_.NewLayout(consumer, docTabsProducer); err != nil {
		return
	}
	if screen, err = _misc_.NewMiscellaneous(ctx, ctxCancel, app, window, layout, screenID); err != nil {
		return
	}
	return
}

func setInitialTabItems(screen *_misc_.Miscellaneous, startupData *_presets_.ScreenInitData) (err error) {
{{ range $panel := .AllPanels }}
 {{ if $panel.IsLocal }}
	if err = _tabs_.Open{{ $panel.Name }}Tab(screen, startupData.{{ $panel.Name }}Panel.TabItemIcon, startupData.{{ $panel.Name }}Panel.TabItemLabel, startupData.{{ $panel.Name }}Panel); err != nil {
		return
	}
 {{- else }}
	if err = _tabs_.Open{{ $panel.Name }}Tab(screen, startupData.{{ $panel.Name }}Screen.TabItemIcon, startupData.{{ $panel.Name }}Screen.TabItemLabel, startupData.{{ $panel.Name }}Screen.ScreenInitData); err != nil {
		return
	}
 {{- end }}
{{- end }}
	return
}
`
)
