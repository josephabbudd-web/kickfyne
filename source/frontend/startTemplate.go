package frontend

import (
	_utils_ "github.com/josephabbudd-web/kickfyne/source/utils"
)

type frontendTemplateData struct {
	ImportPrefix string
	ScreenNames  []string
	Funcs        _utils_.Funcs
}

const (
	frontendFileName = "start.go"

	frontendNoBETemplate = `{{ $DOT := . -}}
package frontend

import (
	"context"
	"fmt"

	"fyne.io/fyne/v2"

	_mainmenu_ "{{ .ImportPrefix }}/frontend/deps/mainmenu"
	_screenmap_ "{{ .ImportPrefix }}/frontend/deps/screenmap"
	_types_ "{{ .ImportPrefix }}/frontend/deps/types"
{{- range $i, $screenName := .ScreenNames }}
 {{- if eq $i 0 }}

	_{{ call $DOT.Funcs.LowerCase $screenName }}_ "{{ $DOT.ImportPrefix }}/frontend/screens/{{ $screenName }}"
 {{- else }}
	_{{ call $DOT.Funcs.LowerCase $screenName }}_ "{{ $DOT.ImportPrefix }}/frontend/screens/{{ $screenName }}"
 {{- end }}
	_{{ call $DOT.Funcs.LowerCase $screenName }}presetting_ "{{ $DOT.ImportPrefix }}/frontend/screens/{{ $screenName }}/presetting"
{{- end }}
)

var opener *_types_.WindowContentConsumer

func Start(ctx context.Context, ctxCancelFunc context.CancelFunc, application fyne.App, window fyne.Window) (err error) {
	buildScreenMap()
	if err = setOpeningScreen(ctx, ctxCancelFunc, application, window); err != nil {
		return
	}

	if usingMainMenu {
		if err = validateMainMenuItems(); err != nil {
			return
		}
		if err = _mainmenu_.Start(ctx, ctxCancelFunc, application, window, mainMenuItems); err != nil {
			return
		}
	}

	opener.Show(true)
	return
}

func buildScreenMap() {
	// Set the screen map.
	// Set the screen map.
{{ range $screenName := .ScreenNames }}
	_screenmap_.Map["{{ $screenName }}"] = &_screenmap_.API{
		NewBorderAreaContentConsumer:     _{{ call $DOT.Funcs.LowerCase $screenName }}_.NewBorderAreaContentConsumer,
		NewWindowContentConsumer:         _{{ call $DOT.Funcs.LowerCase $screenName }}_.NewWindowContentConsumer,
		NewAppTabsTabItemContentConsumer: _{{ call $DOT.Funcs.LowerCase $screenName }}_.NewAppTabsTabItemContentConsumer,
		NewDocTabsTabItemContentConsumer: _{{ call $DOT.Funcs.LowerCase $screenName }}_.NewDocTabsTabItemContentConsumer,
		NewAccordionItemContentConsumer:  _{{ call $DOT.Funcs.LowerCase $screenName }}_.NewAccordionItemContentConsumer,
	}
{{- end }}

	// Set the screen presets map.
{{ range $screenName := .ScreenNames }}
	_screenmap_.PresetsMap["{{ $screenName }}"] = make(map[string]any)
	for k, v := range _{{ call $DOT.Funcs.LowerCase $screenName }}presetting_.Presets {
		_screenmap_.PresetsMap["{{ $screenName }}"][k] = v
	}
{{- end }}
}

func validateMainMenuItems() (err error) {
	length := len(mainMenuItems)
	if length == 0 {
		err = fmt.Errorf("menuItems not found in frontend/settings.go var mainMenuItems")
		return
	}
	last := length - 1
	for i, item := range mainMenuItems {
		var isValid bool
		if _, isValid = _screenmap_.Map[item.ScreenName]; !isValid {
			err = fmt.Errorf(".ScreenName %q is not found in frontend/settings.go var mainMenuItems", item.ScreenName)
			return
		}
		presets := _screenmap_.PresetsMap[item.ScreenName]
		if _, isValid = presets[item.PresetName]; !isValid {
			err = fmt.Errorf(".PresetName %q not found for .ScreenName %q in frontend/settings.go var mainMenuItems", item.PresetName, item.ScreenName)
			return
		}
		if i < last {
			// Check for duplicate labels.
			checks := mainMenuItems[i+1:]
			for _, check := range checks {
				if check.Label == item.Label {
					err = fmt.Errorf("The .Label %q is used more than once in frontend/settings.go var mainMenuItems", item.Label)
					return
				}
			}
		}
	}
	return
}

func setOpeningScreen(ctx context.Context, ctxCancelFunc context.CancelFunc, app fyne.App, w fyne.Window) (err error) {
	var api *_screenmap_.API
	if api = _screenmap_.Map[openingScreenName]; api == nil {
		err = fmt.Errorf("%q is not a valid ScreenName in frontend/settings.go", openingScreenName)
		return
	}
	var presets map[string]any
	var preset any
	presets = _screenmap_.PresetsMap[openingScreenName]
	if preset = presets[openingScreenPresetName]; preset == nil {
		err = fmt.Errorf("%q is not a valid openingScreenPresetName in frontend/settings.go.", openingScreenPresetName)
		return
	}
	opener, _, err = api.NewWindowContentConsumer(ctx, ctxCancelFunc, app, w, true, preset)
	return
}
`
)
