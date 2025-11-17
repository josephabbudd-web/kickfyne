package frontend

const (
	settingsFileName = "settings.go"
	settingsTemplate = `package frontend

import (
	_mainmenu_ "{{ .ImportPrefix }}/frontend/deps/mainmenu"
)

const (
	// usingMainMenu. Is the application using a main menu.
	usingMainMenu = true

	// This is the screen that the application opens with.
	// It does not have to be referenced in var mainMenuItems.
	openingScreenName       = "HelloWorld"
	openingScreenPresetName = "Default"
)

// mainMenuItems is the list of items for the main menu.
// The following issues will generate an error.
//   - Repeated Label.
//   - Unknown ScreenName.
//   - Unknown PresetName.
var mainMenuItems = []_mainmenu_.MainMenuItem{
	{
		Label:      "Hello World!",
		ScreenName: "HelloWorld",
		PresetName: "Default",
	},
}
`
)
