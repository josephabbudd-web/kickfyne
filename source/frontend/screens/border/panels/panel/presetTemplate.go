package panel

const (
	PresetTemplate = `package {{ .PanelName }}Panel

type Preset struct {
	Heading         string
	Description     string
}

func NewDefaultPreset() (preset *Preset) {
	preset = &Preset{
		Heading:      "{{ .PanelName }} panel.",
		Description:  "{{ .Area }} border area.",
	}
	return
}
`
)
