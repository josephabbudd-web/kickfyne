package misc

const (
	MiscellaneousFileName = "miscellaneous.go"

	MiscellaneousTemplate = `{{ $DOT := . -}}
package misc

import (
	"context"

	"fyne.io/fyne/v2"

	_layout_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/deps/layout"
)

// Miscellaneous is a variety of components for this layout and panels.
type Miscellaneous struct {
	CTX       context.Context
	CTXCancel context.CancelFunc
	APP       fyne.App
	Window    fyne.Window

	Layout   *_layout_.Layout
	ScreenID string
}

// NewMiscellaneous constrtucts a Miscellaneous.
// Its parts of the screen that can be deps in one struct.
func NewMiscellaneous(ctx context.Context, ctxCancel context.CancelFunc, app fyne.App, w fyne.Window, layout *_layout_.Layout, screenID string) (components *Miscellaneous, err error) {
	components = &Miscellaneous{
		CTX:       ctx,
		CTXCancel: ctxCancel,
		APP:       app,
		Window:    w,
		Layout:    layout,
		ScreenID:  screenID,
	}
	return
}
`
)
