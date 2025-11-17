package mainmenu

type mainMenuTemplateData struct {
	ImportPrefix string
}

const (
	mainMenuTemplate = `package mainmenu

import (
	"context"

	"fyne.io/fyne/v2"

	_screenmap_ "{{ .ImportPrefix }}/frontend/deps/screenmap"
	_types_ "{{ .ImportPrefix }}/frontend/deps/types"
)

// label is the text displayed in the application's main menu.
// screen is the name of the screen package folder in frontend/screens/.
// preset is the name of the screen package's preset setting in the package's startup/ folder.
type MainMenuItem struct {
	Label      string
	ScreenName string
	PresetName string
}

var (
	application         fyne.App
	window              fyne.Window
	screenLabelConsumer = make(map[string]_types_.ContentConsumer)
)

// Start builds the main menu and adds it to the app.
func Start(ctx context.Context, ctxCancelFunc context.CancelFunc, app fyne.App, w fyne.Window, mainMenuItems []MainMenuItem) (err error) {
	// Setup.
	application = app
	window = w
	menuItems := make([]*fyne.MenuItem, 0, len(mainMenuItems))
	for _, valid := range mainMenuItems {
		var windowContentConsumer *_types_.WindowContentConsumer
		presets := _screenmap_.PresetsMap[valid.ScreenName]
		preset := presets[valid.PresetName]
		api := _screenmap_.Map[valid.ScreenName]
		if windowContentConsumer, _, err = api.NewWindowContentConsumer(ctx, ctxCancelFunc, app, w, true, preset); err != nil {
			return
		}
		screenLabelConsumer[valid.Label] = windowContentConsumer
		item := fyne.NewMenuItem(
			valid.Label,
			func() {
				// Here consumer.Show is being called back in the main thread.
				windowContentConsumer.Show(true)
			},
		)
		menuItems = append(menuItems, item)
	}
	// Build the sub menu.
	subMenu := fyne.NewMenu(application.Metadata().Name, menuItems...)
	// Build the main menu.
	mainmenu := fyne.NewMainMenu(subMenu)
	fyne.Do(
		func() { window.SetMainMenu(mainmenu) },
	)
	return
}
`
)
