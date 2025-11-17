package mainmenu

type initTemplateData struct {
	ImportPrefix string
}

const (
	initFileName = "init.go"

	initTemplate = `package mainmenu

import (
	"context"
	"log"

	"fyne.io/fyne/v2"

	_screenmap_ "{{ .ImportPrefix }}/frontend/deps/screenmap"
	_types_ "{{ .ImportPrefix }}/frontend/deps/types"
)

type validMainMenuItem struct {
	label  string
	screen string
	preset any
}

var (
	application         fyne.App
	window              fyne.Window
	screenLabelConsumer = make(map[string]_types_.ContentConsumer)
)

// Init builds the main menu and adds it to the app.
func Init(ctx context.Context, ctxCancelFunc context.CancelFunc, app fyne.App, w fyne.Window) {
	// Setup.
	application = app
	window = w
	valids := validateMainMenuItems()
	menuItems := make([]*fyne.MenuItem, 0, len(valids))
	var first *_types_.WindowContentConsumer
	var opener *_types_.WindowContentConsumer
	for i, valid := range valids {
		var windowContentConsumer *_types_.WindowContentConsumer
		var err error
		api := _screenmap_.Map[valid.screen]
		if windowContentConsumer, _, err = api.NewWindowContentConsumer(ctx, ctxCancelFunc, app, w, true, valid.preset); err == nil {
			screenLabelConsumer[valid.label] = windowContentConsumer
			if i == 0 {
				first = windowContentConsumer
			}
			if valid.screen == openingScreen.screen && valid.preset == openingScreen.preset {
				// This is also the opener.
				opener = windowContentConsumer
			}
			item := fyne.NewMenuItem(
				valid.label,
				func() {
					// Here consumer.Show is being called back in the main thread.
					windowContentConsumer.Show(true)
				},
			)
			menuItems = append(menuItems, item)
		}
	}
	// Build the sub menu.
	subMenu := fyne.NewMenu(application.Metadata().Name, menuItems...)
	// Build the main menu.
	mainmenu := fyne.NewMainMenu(subMenu)
	fyne.Do(
		func() { window.SetMainMenu(mainmenu) },
	)
	// Show the default screen.
	if opener == nil {
		if opener = openingScreenWindowContentConsumer(ctx, ctxCancelFunc, app, w); opener == nil {
			opener = first
		}
	}
	opener.Show(true)
}

func openingScreenWindowContentConsumer(ctx context.Context, ctxCancelFunc context.CancelFunc, app fyne.App, w fyne.Window) (windowContentConsumer *_types_.WindowContentConsumer) {
	var api *_screenmap_.API
	if api = _screenmap_.Map[openingScreen.screen]; api == nil {
		log.Printf("%q is not a valid openingScreen.screen in frontend/deps/mainmenu/mainmenu.go.", openingScreen.screen)
		return
	}
	var presets map[string]any
	var preset any
	presets = _screenmap_.PresetsMap[openingScreen.screen]
	if preset = presets[openingScreen.preset]; preset == nil {
		log.Printf("%q is not a valid openingScreen.preset in frontend/deps/mainmenu/mainmenu.go.", openingScreen.preset)
		return
	}
	var err error
	if windowContentConsumer, _, err = api.NewWindowContentConsumer(ctx, ctxCancelFunc, app, w, true, preset); err != nil {
		windowContentConsumer = nil
	}
	return
}

func validateMainMenuItems() (valids []validMainMenuItem) {
	valids = make([]validMainMenuItem, 0, len(mainMenuItems))
	for _, item := range mainMenuItems {
		var isValid bool
		if _, isValid = _screenmap_.Map[item.screen]; !isValid {
			log.Printf("%q is not a valid screen name in frontend/deps/mainmenu/mainmenu.go.", item.screen)
			continue
		}
		presets := _screenmap_.PresetsMap[item.screen]
		var preset any
		if preset, isValid = presets[item.preset]; !isValid {
			log.Printf("%q is not a valid preset of the screen %q in frontend/deps/mainmenu/mainmenu.go.", item.preset, item.screen)
			continue
		}
		for _, valid := range valids {
			if valid.label == item.label {
				log.Printf("The menu label %q is used more than once in frontend/deps/mainmenu/mainmenu.go.", item.label)
				continue
			}
		}
		newValid := validMainMenuItem{
			label:  item.label,
			screen: item.screen,
			preset: preset,
		}
		valids = append(valids, newValid)
	}
	return
}
`
)
