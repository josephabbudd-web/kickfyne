package mainmenu

type mainMenuTemplateData struct {
	ImportPrefix string
}

const (
	mainMenuFileName = "mainmenu.go"

	mainMenuTemplate = `package mainmenu

import (
	"log"
)

// label is the text displayed in the application's main menu.
// screen is the name of the screen package folder in frontend/screens/.
// preset is the name of the screen package's preset setting in the package's startup/ folder.
type mainMenuItem struct {
	label  string
	screen string
	preset string
}

// openingScreen is the screen that will open the application.
// If there is an error identifying the opening screen then the first working screen in mainMenuItems will be used.
var openingScreen = mainMenuItem{
		screen: "HelloWorld",
		preset: "Default",
}

// The first mainMenuItem is also the opening screen.
// The following items in mainMenuItems are ignored and logged without an error.
//  - Repeated labels.
//  - Invalid screen package names.
//  - Invalid presets names.
var mainMenuItems = []mainMenuItem{
	{
		label:  "Hello World!",
		screen: "HelloWorld",
		preset: "Default",
	},
}

// Show the screen labeled screenLabel.
func Show(screenLabel string, isMainThread bool) {
	if consumer := screenLabelConsumer[screenLabel]; consumer != nil {
		consumer.Show(isMainThread)
		return
	}
	log.Printf("Unable to show a screen labeled %q.", screenLabel)
}
`
)
