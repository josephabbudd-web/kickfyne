package panels

import (
	"fmt"
	"strings"

	_utils_ "github.com/josephabbudd-web/kickfyne/source/utils"
)

const (
	LogFileName = "README.log"

	howPanelsWork = `%s AppTabs screen package.:
How panels work.
A panel only produces content for it's tab.
1. content.go is the panel's content.
   * The tab's icon, label and content are defined.
   * The content is layed out for display.
   * The content handles user input events.
   KICKFYNE TODO: Modify content for your user interface.
2. state.go sets and gets the panel's content.go.
   KICKFYNE TODO: Modify state.go for content.go.
`

	ending = `
Content producers and consumers. (Not required learning.)
1. Each panel produces content for it's own tabItem in the AppTabs tabbar.
   * The panel's producer gives the panel's content to the tabItem's consumer.
   * The tabItem's consumer gives the content to the panel's tabItem.
   * The panel's tabItem is part of the package's AppTabs tabbar.
2. The package has a producer which gives the AppTabs tabbar's content to the package's consumer.
   The package's consumer will be 1 of these 4 consumer types.
   1. A WindowContentConsumer consumes the package's content for the entire application window.
   2. An AppTabsTabItemContentConsumer consumes the package's content for a single TabItem in a separate AppTab screen.
   3. A DocTabsTabItemContentConsumer consumes the package's content for a single TabItem in a separate DocTabs screen.
   4. An AccordionConsumer consumes the package's content for a single AccordionItem in a separate Accordion screen.
`
)

func LogContent(
	screenPackageName string,
	localPanelNames []string,
	folderPaths *_utils_.FolderPaths,
) (logMessage string) {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf(howPanelsWork, screenPackageName))
	for _, panelName := range localPanelNames {
		contentPath := _utils_.PanelContentFilePath(screenPackageName, panelName, folderPaths)
		statePath := _utils_.PanelStateFilePath(screenPackageName, panelName, folderPaths)
		builder.WriteString("  " + panelName + "Panel\n")
		builder.WriteString(fmt.Sprintf("    Content:   %s.\n", contentPath))
		builder.WriteString(fmt.Sprintf("    State:     %s.\n", statePath))
	}
	// The ending. The boring part.
	builder.WriteString(ending)
	logMessage = builder.String()
	return
}
