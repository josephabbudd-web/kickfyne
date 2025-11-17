package apptabs

import _utils_ "github.com/josephabbudd-web/kickfyne/source/utils"

type handlersTemplateData struct {
	PackageName  string
	ImportPrefix string
	Funcs        _utils_.Funcs
}

const (
	handlersFileName = "tabItemHandlers.go"
	handlersTemplate = `
package {{ call .Funcs.LowerCase .PackageName }}

import (
	"fyne.io/fyne/v2/container"

	_layout_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/deps/layout"
)

// KICKFYNE TODO:
// If you want a handler:
// - Add the functionality to func closeIntercept, onSelected or onUnselected.
// - Pass the handler(s) in func setTabItemHandlers.
func setTabItemHandlers(layout *_layout_.Layout) {
	layout.SetTabItemHandlers(
		nil, // onSelected
		nil, // onUnselected
	)
}

// onSelected is called after the tab is selected and it's content is displayed.
// Param selectedTab is the tab the user selected.
func onSelected(selectedTab *container.TabItem) {
}

// onUnselected is called before the tab is unselected and while it's content is still displayed.
// Param unselectedTab is the tab that will no longer be selected because the user selected another tab.
func onUnselected(unselectedTab *container.TabItem) {
}
`
)
