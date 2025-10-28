package panel

const (
	PresetTemplate = `package {{ .PanelName }}Panel

type Preset struct {
	Heading     string
	Description string
}

func NewDefaultPreset() (preset *Preset) {
	preset = &Preset{
		Heading:     "This is the {{ .PanelName }} panel heading.",
		Description: "This is the {{ .PanelName }} panel description.",
	}
	return
}
`
)
