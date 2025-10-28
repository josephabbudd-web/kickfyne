package panel

const (
	PresetTemplate = `package {{ .PanelName }}Panel

import (
	"fyne.io/fyne/v2"

	// "fyne.io/fyne/v2/theme"
)

type Preset struct {
	TabItemIconName fyne.ThemeIconName
	TabItemLabel    string
	Heading         string
	Description     string
}

func NewDefaultPreset() (preset *Preset) {
	preset = &Preset{
		// TabItemIconName: theme.IconName????,
		TabItemLabel: "{{ .PanelName }}",
		Heading:      "This is the {{ .PanelName }} panel heading.",
		Description:  "This is the {{ .PanelName }} panel description.",
	}
	return
}
`

	ConfigPresetTemplate = `package {{ .PanelName }}Panel

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type Preset struct {
	TabItemIconName fyne.ThemeIconName
	TabItemLabel    string
	Heading         string
	Description     string
}

func NewDefaultPreset() (preset *Preset) {
	preset = &Preset{
		TabItemIconName: theme.IconNameSettings,
		TabItemLabel:    "",
		Heading:         "Tab Settings.",
		Description:     "Select where you want the tabs located.",
	}
	return
}
`
)
