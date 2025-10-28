package panel

const (
	PresetTemplate = `package {{ .PanelName }}Panel

type Preset struct {
	AccordionItemTitle string
	Heading            string
	Description        string
}

func NewDefaultPreset() (preset *Preset) {
	preset = &Preset{
		AccordionItemTitle: "{{ .PanelName }}",
		Heading:            "This is the {{ .PanelName }} panel heading.",
		Description:        "This is the {{ .PanelName }} panel description.",
	}
	return
}
`
)
