package panelers

type PanelersTemplateData struct {
	ImportPrefix    string
	LocalPanelNames []string
}

const (
	PanelersFileName = "panelers.go"

	PanelersTemplate = `package panelers

import (
	_types_ "{{ .ImportPrefix }}/frontend/types"
)

// Panelers is this screen's panels.
type Panelers struct {
{{- range $panelName := .LocalPanelNames }}
	{{ $panelName }}  _types_.Paneler
{{- end }}
	DefaultPanel _types_.Paneler
}

// Map returns each panel's name mapped to it's implementation.
func (panelers *Panelers) Map () (panelMap map[string]_types_.Paneler) {
	panelMap = make(map[string]_types_.Paneler)
{{- range $panelName := .LocalPanelNames}}
	panelMap["{{ $panelName }}"] = panelers.{{ $panelName }}
{{- end }}
	return
}
`
)
