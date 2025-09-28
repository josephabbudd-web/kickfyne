package mainmenu

type mainMenuTemplateData struct {
	ImportPrefix string
}

const (
	mainMenuFileName = "mainmenu.go"

	mainMenuTemplate = `package mainmenu

import (
	"context"
	"log"
	"strings"

	"fyne.io/fyne/v2"

	_screenmap_ "{{ .ImportPrefix }}/frontend/screenmap"
	_types_ "{{ .ImportPrefix }}/frontend/types"
)

var (
	screenNames        []string
	screenNameConsumer = make(map[string]_types_.ContentConsumer)
	menuItems          []*fyne.MenuItem
	window             fyne.Window
	application        fyne.App
)

// Init builds the main menu and adds it to the app.
func Init(ctx context.Context, ctxCancelFunc context.CancelFunc, app fyne.App, w fyne.Window) {
	// Setup.
	application = app
	window = w
	names := strings.Split(app.Metadata().Custom["MainMenu"], " ")
	screenNames = make([]string, 0, len(names))
	for _, name := range names {
		var isValidName bool
		if _, isValidName = _screenmap_.Map[name]; isValidName {
			screenNames = append(screenNames, name)
		} else {
			log.Printf("%q is not a valid screen name in Fyne.toml", name)
		}
	}
	for _, screenName := range screenNames {
		var windowContentConsumer *_types_.WindowContentConsumer
		var screenID string
		var err error
		api := _screenmap_.Map[screenName]
		if windowContentConsumer, screenID, err = api.NewWindowContentConsumer(ctx, ctxCancelFunc, app, w, true, nil); err == nil {
			var screenName string
			if length := strings.Index(screenID, ":"); length > 0 {
				screenName = screenID[0:length]
			} else {
				screenName = screenID
			}
			screenNameConsumer[screenName] = windowContentConsumer
			item := fyne.NewMenuItem(
				screenName,
				func() {
					// Here consumer.Show is being called back in the main thread.
					windowContentConsumer.Show(true)
				},
			)
			menuItems = append(menuItems, item)
		}
	}
	application = app
	window = w
	// Build the sub menu.
	subMenu := fyne.NewMenu(application.Metadata().Name, menuItems...)
	// Build the main menu.
	mainmenu := fyne.NewMainMenu(subMenu)
	fyne.Do(
		func() { window.SetMainMenu(mainmenu) },
	)
	// Show the default screen.
	for _, screenName := range screenNames {
		if consumer := screenNameConsumer[screenName]; consumer != nil {
			consumer.Show(true)
			break
		}
	}
}

// Show the screen named screenName.
func Show(screenName string, isMainThread bool) {
	if consumer := screenNameConsumer[screenName]; consumer != nil {
		consumer.Show(isMainThread)
	} else {
		log.Printf("Unable to show a screen named %q.", screenName)
	}
}
`
)
