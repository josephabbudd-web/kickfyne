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

	_frontend_ "{{ .ImportPrefix }}/frontend"
	_deps_ "{{ .ImportPrefix }}/deps"
 	_thread_ "{{ .ImportPrefix }}/deps/thread"
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

	a := app.New()
	appName := a.Metadata().Name
	w := a.NewWindow(appName)

	// Cancel.
	ctx, ctxCancel := context.WithCancel(context.Background())
	w.SetCloseIntercept(
		ctxCancel,
	)
	errCh := make(chan error, 2)
	go monitor(w, ctx, errCh)

	// Set the main thread ID.
	if exitError = _thread_.SetMainThreadID(); exitError != nil {
		return
	}

	// Start deps.
	if exitError = _deps_.Start(ctx, ctxCancel); exitError != nil {
		return
	}

	// Start the front end.
	if exitError = _frontend_.Start(ctx, ctxCancel, a, w); exitError != nil {
		return
	}

	size := size16x9(1000, 0)
	w.Resize(size)
	w.CenterOnScreen()
	w.Show()

	// Start Fyne's event cycle.
	a.Run()
}

func monitor(w fyne.Window, ctx context.Context, errCh chan error) {
	select {
	case <-ctx.Done():
		fyne.Do(func() { w.Close() })
		return
	case exitError = <-errCh:
		fyne.Do(func() { w.Close() })
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
`
