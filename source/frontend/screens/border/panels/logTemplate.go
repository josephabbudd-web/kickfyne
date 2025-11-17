package panels

import (
	"fmt"
	"strings"

	_data_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/border/data"
	_utils_ "github.com/josephabbudd-web/kickfyne/source/utils"
)

const (
	LogFileName = "README.log"

	howPanelsWork = `%s Border screen package.:
How panels work.
A panel only produces content for it's border area.
1. content.go is the panel's content.
   * The content's members are defined.
   * The content is layed out.
   * The content handles user input events.
   KICKFYNE TODO: Modify content for your user interface.
2. state.go handles the complexities of setting the content.
   KICKFYNE TODO: Modify state.go for content.go.
3. preset.go defines each initial panel state.
   KICKFYNE TODO: Modify preset.go.
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
   4. An AccordionContentConsumer consumes the package's content for a single AccordionItem in a separate Accordion screen.
   5. A BorderContentConsumer consumes the package's content for a single area in a separate Border screen.
`
)

func LogContent(
	screenPackageName string,
	areas []*_data_.AreaData,
	folderPaths *_utils_.FolderPaths,
) (logMessage string) {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf(howPanelsWork, screenPackageName))
	for _, area := range areas {
		if !area.IsLocal {
			continue
		}
		contentPath := _utils_.PanelContentFullFilePath(screenPackageName, area.ItemName, folderPaths)
		statePath := _utils_.PanelStateFullFilePath(screenPackageName, area.ItemName, folderPaths)
		presetPath := _utils_.PanelPresetFullFilePath(screenPackageName, area.ItemName, folderPaths)
		builder.WriteString("  " + area.ItemName + "Panel\n")
		builder.WriteString(fmt.Sprintf("    Content:   %s.\n", contentPath))
		builder.WriteString(fmt.Sprintf("    State:     %s.\n", statePath))
		builder.WriteString(fmt.Sprintf("    Preset:    %s.\n", presetPath))
	}
	// The ending. The boring part.
	builder.WriteString(ending)
	logMessage = builder.String()
	return
}
