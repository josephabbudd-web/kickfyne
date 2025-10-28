package layoutitems

type TemplateData struct {
	PackageName  string
	ImportPrefix string
	AllItemNames []string
}

const (
	DocsTemplate = `// This file is updated each time kickfyne is used to add or remove a accordionitem in this package.
// Do not modify this file.
//
// Below is an example accordionitem layout func.
// It is provided as a reference for func AccordionLayout in accordionitemLayout.go
//
// func ExampleAccordionLayout(screen *_misc_.Miscellaneous, preset *_presetting_.Preset) (err error) {
{{- range $itemName := .AllItemNames}}
 {{- if eq (slice $itemName 0 1) "*" }}
// 	if err = _accordionitems_.Open{{ slice $itemName 1 }}AccordionItem(screen, preset.{{ slice $itemName 1 }}Screen); err != nil {
// 		return
// 	}
 {{- else }}
// 	if err = _accordionitems_.Open{{ $itemName }}AccordionItem(screen, preset.{{ $itemName }}Panel); err != nil {
// 		return
// 	}
 {{- end }}
{{- end }}
// 	return
// }
package layoutaccordionitems
`
	LayoutTemplate = `package layoutaccordionitems

import (
	_accordionitems_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/deps/accordionItems"
	_misc_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/misc"
	_presetting_ "{{ .ImportPrefix }}/frontend/screens/{{ .PackageName }}/presetting"
)

/* KICKFYNE TODO:
- You may need to customize func LayoutAccordionItems.
  Kickfyne always keeps updated example code for you to review in example.go.
  Example.go is updated whenever you add or remove a accordionitem from this package.
*/

func LayoutAccordionItems(screen *_misc_.Miscellaneous, preset *_presetting_.Preset) (err error) {
{{- range $itemName := .AllItemNames}}
 {{- if eq (slice $itemName 0 1) "*" }}
 	if err = _accordionitems_.Open{{ slice $itemName 1 }}AccordionItem(screen, preset.{{ slice $itemName 1 }}Screen); err != nil {
 		return
 	}
 {{- else }}
 	if err = _accordionitems_.Open{{ $itemName }}AccordionItem(screen, preset.{{ $itemName }}Panel); err != nil {
 		return
 	}
 {{- end }}
{{- end }}
	return
}
`
)
