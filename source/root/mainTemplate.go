package root

import (
	_utils_ "github.com/josephabbudd-web/kickfyne/source/utils"
)

const (
	MainFileName = "main.go"
)

type mainTemplateData struct {
	ImportPrefix string
	AppName      string
	Funcs        _utils_.Funcs
}

var mainTemplate = `package main

import (
	"context"
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/driver/desktop"

	_deps_ "{{ .ImportPrefix }}/deps"
	_paths_ "{{ .ImportPrefix }}/deps/paths"
 	_thread_ "{{ .ImportPrefix }}/deps/thread"
	_frontend_ "{{ .ImportPrefix }}/frontend"
)

const (
	envTrue  = "1"
	envFalse = ""
)

var exitError error

func main() {

	defer func() {
		if exitError != nil {
			log.Printf("main: err is %s", exitError.Error())
			os.Exit(1)
		}
	}()

	if len(os.Getenv("FYNE_SCALE")) == 0 {
		os.Setenv("FYNE_SCALE", "1")
	}
	if len(os.Getenv("FYNE_THEME")) == 0 {
		os.Setenv("FYNE_THEME", "dark")
	}
	os.Setenv("USETESTPATH", envFalse)
	os.Setenv("CWT_TESTING", envFalse)

	application := app.New()
	appName := application.Metadata().Name
	window := application.NewWindow(appName)

	// Cancel.
	ctx, ctxCancel := context.WithCancel(context.Background())
	defer ctxCancel()

	// The Error Channel.
	errCh := make(chan error, 2)
	go monitor(window, ctx, errCh)

	// Set the main thread ID.
	if exitError = _thread_.SetMainThreadID(); exitError != nil {
		return
	}

	// Start deps.
	if exitError = _deps_.Start(ctx, ctxCancel); exitError != nil {
		return
	}

	// The app icon.
	iconPath := _paths_.ImagePath("app-icon.jpeg")
	var iconResource fyne.Resource
	if iconResource, exitError = fyne.LoadResourceFromPath(iconPath); exitError != nil {
		return
	} else {
		application.SetIcon(iconResource)
	}
	// KICKFYNE TODO:
	// If you want application system tray:
	// 1. Uncomment the next line.
	// 2. Customize func systemTray if required.
	systemTray(application, window)

	// Start the front end.
	if exitError = _frontend_.Start(ctx, ctxCancel, application, window); exitError != nil {
		return
	}

	size := size16x9(1000, 0)
	window.Resize(size)
	window.CenterOnScreen()
	window.Show()

	// Start Fyne's event cycle.
	application.Run()
}

func monitor(window fyne.Window, ctx context.Context, errCh chan error) {
	select {
	case <-ctx.Done():
		fyne.Do(func() { window.Close() })
		return
	case exitError = <-errCh:
		fyne.Do(func() { window.Close() })
		return
	}
}

func size16x9(width, height int) (size fyne.Size) {
	var newWidth float32
	var newHeight float32
	switch {
	case width != 0:
		if width < 0 {
			width = 0 - width
		}
		r := width / 16
		newWidth = float32(r * 16)
		newHeight = float32(r * 9)
	case height != 0:
		if height < 0 {
			height = 0 - height
		}
		r := height / 9
		newWidth = float32(r * 16)
		newHeight = float32(r * 9)
	default:
		// default to 720 width.
		r := 720 / 16
		newWidth = float32(r * 16)
		newHeight = float32(r * 9)
	}
	size = fyne.Size{Width: newWidth, Height: newHeight}
	return
}

// systemTray will create and set the system tray.
// app.Icon must be set for there to be a system tray.
// KICKFYNE TODO: Do something useful with the system tray if you want to use it.
func systemTray(application fyne.App, window fyne.Window) {

	var desktopApp desktop.App
	var ok bool
	if desktopApp, ok = application.(desktop.App); !ok {
		// Not running on a desktop.
		return
	}
	var icon fyne.Resource
	if icon = application.Icon(); icon == nil {
		log.Println("The application's Icon was not set.")
		// The application's Icon was not set.
		return
	}

	item := fyne.NewMenuItem(
		"Show",
		func() { window.Show() },
	)
	systemTrayMenu := fyne.NewMenu(
		"Crud",
		item,
	)
	desktopApp.SetSystemTrayMenu(systemTrayMenu)
	desktopApp.SetSystemTrayIcon(icon)
	window.SetCloseIntercept(
		func() {
			window.Hide()
		},
	)
}
`
