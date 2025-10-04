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
	AllPanelNames    []string
	LocalPanelNames  []string
	RemotePanelNames []string
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

	_accordionitems_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/deps/accordionitems"
	_layout_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/deps/layout"
	_misc_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/misc"
	_producer_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/deps/producer"
	_startup_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/startup"
	_types_ "{{ .ImportPrefix }}/frontend/types"
	_ids_ "{{ .ImportPrefix }}/deps/container/{{ .PackageName }}"
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

type InitializeScreenData struct {
{{- range $panelName := .LocalPanelNames }}

	// Initialize data for the {{ $panelName }} tab which uses the {{ $panelName }}Panel for content.
	{{ $panelName }}TabLabel string
	{{ $panelName }}TabIcon fyne.Resource
	{{ $panelName }}Heading string
	{{ $panelName }}Description string
{{- end }}

{{- range $panelName := .RemotePanelNames }}

	// Initialize data for the {{ $panelName }} tab which uses the {{ $panelName }} screen for content.
	// This datum should be the key or id that tells the {{ $panelName }} screen what to load into it's panel's.
	// Example: IRC Server & Channel.
	// Example: IRC UserName
{{- end }}
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
			_ids_.RemoveAccordion(screenID)
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
		// PackageScreen.
		var packageScreen *_misc_.Miscellaneous
		if packageScreen, err = buildLayout(ctx, ctxCancel, app, window, windowContentConsumer, screenID); err != nil {
			return
		}
		// Update deps.
		_ids_.AddAccordion(packageScreen.ScreenID)
		// Get the screen's initializer.
		err = setInitialAccordionItems(packageScreen, startupData)
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
			_ids_.RemoveAccordion(screenID)
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
		// PackageScreen.
		var packageScreen *_misc_.Miscellaneous
		if packageScreen, err = buildLayout(ctx, ctxCancel, app, window, accordionItemContentConsumer, screenID); err != nil {
			return
		}
		// Update deps.
		_ids_.AddAccordion(packageScreen.ScreenID)
		// Get the screen's initializer.
		err = setInitialAccordionItems(packageScreen, startupData)
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
			err = fmt.Errorf("{{ .PackageName }}.NewTabItemContentConsumer: %w", err)
			_ids_.RemoveAccordion(screenID)
		}
	}()

	if startupData == nil {
		startupData = _startup_.DefaultScreenInitData()
	}
	switch startupData := startupData.(type) {
		case *_startup_.ScreenInitData:
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
		_ids_.AddAccordion(packageScreen.ScreenID)
		// Get the screen's initializer.
		err = setInitialAccordionItems(packageScreen, startupData)
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
			err = fmt.Errorf("{{ .PackageName }}.NewTabItemContentConsumer: %w", err)
			_ids_.RemoveAccordion(screenID)
		}
	}()

	if startupData == nil {
		startupData = _startup_.DefaultScreenInitData()
	}
	switch startupData := startupData.(type) {
	case *_startup_.ScreenInitData:
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
		_ids_.AddAccordion(packageScreen.ScreenID)
		// Get the screen's initializer.
		err = setInitialAccordionItems(packageScreen, startupData)
	default:
		err = errors.New("startupData is not a *_startup_.ScreenInitData")
	}

	return
}

func buildLayout(
	ctx context.Context, ctxCancel context.CancelFunc,
	app fyne.App, window fyne.Window,
	consumer _types_.ContentConsumer,
	screenID string,
) (screen *_misc_.Miscellaneous, err error) {
	// Build the Accordion content producer.
	accordionContentProducer := _producer_.NewAccordionContentProducer(consumer)
	consumer.Bind(accordionContentProducer)

	// Build Layout
	var layout *_layout_.Layout
	if layout, err = _layout_.NewLayout(consumer, accordionContentProducer); err != nil {
		return
	}
	if screen, err = _misc_.NewMiscellaneous(ctx, ctxCancel, app, window, layout, screenID); err != nil {
		return
	}
	return
}

func setInitialAccordionItems(screen *_misc_.Miscellaneous, startupData *_startup_.ScreenInitData) (err error) {
{{- range $panel := .AllPanels}}
 {{- if $panel.IsLocal }}
	if err = _accordionitems_.Open{{ $panel.Name }}AccordionItem(screen, startupData.{{ $panel.Name }}Panel.AccordionItemTitle, startupData.{{ $panel.Name }}Panel); err != nil {
		return
	}
 {{- else }}
	if err = _accordionitems_.Open{{ $panel.Name }}AccordionItem(screen, startupData.{{ $panel.Name }}Screen.AccordionItemTitle, startupData.{{ $panel.Name }}Screen.ScreenInitData); err != nil {
		return
	}
 {{- end }}
{{- end }}
	return
}
`
)
