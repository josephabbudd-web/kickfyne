package frontend

type frontendTemplateData struct {
	ImportPrefix string
	ScreenNames  []string
}

const (
	frontendFileName = "frontend.go"

	frontendNoBETemplate = `{{ $DOT := . -}}
package frontend

import (
	"context"
	"fmt"

	"fyne.io/fyne/v2"

	_mainmenu_ "{{ .ImportPrefix }}/frontend/mainmenu"
	_screenmap_ "{{ .ImportPrefix }}/frontend/screenmap"

{{ range $screenName := .ScreenNames }}
	_{{ $screenName }}_ "{{ $DOT.ImportPrefix }}/frontend/screens/{{ $screenName }}"
{{- end }}
)

func Start(ctx context.Context, ctxCancelFunc context.CancelFunc, app fyne.App, window fyne.Window) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("frontend.Start: %w", err)
		}
	}()

	// Set the screen map.
{{ range $screenName := .ScreenNames }}
	_screenmap_.Map["{{ $screenName }}"] = &_screenmap_.API{
		NewWindowContentConsumer:         _{{ $screenName }}_.NewWindowContentConsumer,
		NewAppTabsTabItemContentConsumer: _{{ $screenName }}_.NewAppTabsTabItemContentConsumer,
		NewDocTabsTabItemContentConsumer: _{{ $screenName }}_.NewDocTabsTabItemContentConsumer,
		NewAccordionItemContentConsumer:  _{{ $screenName }}_.NewAccordionItemContentConsumer,
	}
{{- end }}

	// Set the screen presets map.
{{ range $screenName := .ScreenNames }}
	_screenmap_.PresetsMap["{{ $screenName }}"] = _{{ $screenName }}_.Presets()
{{- end }}

	// Initialize main menu.
	// The developer must ensure that all panel groups should get initialized from main menu.
	_mainmenu_.Init(ctx, ctxCancelFunc, app, window)

	return
}
`
)
