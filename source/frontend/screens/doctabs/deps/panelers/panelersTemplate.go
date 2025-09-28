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
// This screen has 2 panels.
// The default panel is Home.
type Panelers struct {
{{- range $panelName := .LocalPanelNames }}
	{{ $panelName }}  _types_.Paneler
{{- end }}
	DefaultPanel _types_.Paneler
}
`
)
