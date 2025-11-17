package layoutitems

type TemplateData struct {
	PackageName  string
	ImportPrefix string
	AllItemNames []string
	UseConfigTab bool
}

const (
	DocsTemplate = `// This file is updated each time kickfyne is used to add or remove a tab in this package.
// Do not modify this file.
//
// Below is an example tab layout func.
// It is provided as a reference for func TabLayout in tabLayout.go
//
// func ExampleTabLayout(screen *_misc_.Miscellaneous, preset *_presetting_.Preset) (err error) {
{{- range $itemName := .AllItemNames}}
 {{- if eq (slice $itemName 0 1) "*" }}
// 	if err = _tabs_.Open{{ slice $itemName 1 }}TabItem(screen, preset.{{ slice $itemName 1 }}Screen); err != nil {
// 		return
// 	}
 {{- else }}
// 	if err = _tabs_.Open{{ $itemName }}TabItem(screen, preset.{{ $itemName }}Panel); err != nil {
// 		return
// 	}
 {{- end }}
{{- end }}
{{- if .UseConfigTab }}
// Tabbar configuration.
// The first tab is the configuration tab. 
//	tabbar := screen.Layout.Tabbar()
//	tabbar.SelectIndex(1)
//	tabbar.SetTabLocation(defaultTabbarLocation)
{{- else }}
//  Tabbar configuration.
//	tabbar := screen.Layout.Tabbar()
//	tabbar.SelectIndex(0)
//	tabbar.SetTabLocation(defaultTabbarLocation)
{{- end }}
// 	return
// }
package layouttabitems
`
	LayoutTemplate = `package layouttabitems

import (
	"fyne.io/fyne/v2/container"

	_tabs_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/deps/tabItems"
	_misc_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/misc"
	_presetting_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/presetting"
)

/* KICKFYNE TODO:
- You may need to customize func LayoutTabItems.
  Kickfyne always keeps updated example code for you to review in example.go.
  The example in docs.go is updated whenever you add or remove a tab from this package.
*/

// KICKFYNE TODO:
// You may want to set a custom tabbar location.
// defaultTabbarLocation is where the tabbar is located at startup.
const defaultTabbarLocation = container.TabLocationTop

func LayoutTabItems(screen *_misc_.Miscellaneous, preset *_presetting_.Preset) (err error) {
{{- range $itemName := .AllItemNames}}
 {{- if eq (slice $itemName 0 1) "*" }}
 	if err = _tabs_.Open{{ slice $itemName 1 }}TabItem(screen, preset.{{ slice $itemName 1 }}Screen); err != nil {
 		return
 	}
 {{- else }}
 	if err = _tabs_.Open{{ $itemName }}TabItem(screen, preset.{{ $itemName }}Panel); err != nil {
 		return
 	}
 {{- end }}
{{- end }}
{{- if .UseConfigTab }}
	// Tabbar configuration.
	// The first tab is the configuration tab. 
	tabbar := screen.Layout.Tabbar()
	tabbar.SelectIndex(1)
	tabbar.SetTabLocation(defaultTabbarLocation)
{{- else }}
	// Tabbar configuration.
	tabbar := screen.Layout.Tabbar()
	tabbar.SelectIndex(0)
	tabbar.SetTabLocation(defaultTabbarLocation)
{{- end }}
	return
}
`
)
