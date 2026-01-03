package split

import (
	"fmt"
	"strings"

	_data_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/split/data"
	_utils_ "github.com/josephabbudd-web/kickfyne/source/utils"
)

type docTemplateData struct {
	Data  *_data_.TemplateData
	Files string
}

func files(
	general *_data_.TemplateData,
	folderPaths *_utils_.FolderPaths,
) (successMessage string) {
	builder := strings.Builder{}

	builder.WriteString("The split areas and their panel files:\n")

	if general.Leading.IsLocal {
		builder.WriteString("The Leading area displays the content from the Leading panel.\n")
		contentPath := _utils_.PanelContentFileRelativePath("Leading", folderPaths)
		statePath := _utils_.PanelStateFileRelativePath("Leading", folderPaths)
		presetPath := _utils_.PanelPresetFileRelativePath("Leading", folderPaths)
		builder.WriteString("Leading Panel\n")
		builder.WriteString(fmt.Sprintf("    Content: %s.\n", contentPath))
		builder.WriteString(fmt.Sprintf("    State:   %s.\n", statePath))
		builder.WriteString(fmt.Sprintf("    Presets: %s.\n", presetPath))
	} else {
		builder.WriteString(fmt.Sprintf("The Leading area displays the content from the %s screen.\n", general.Leading.ScreenName))
	}

	if general.Trailing.IsLocal {
		builder.WriteString("The Trailing area displays the content from the Trailing panel.\n")
		contentPath := _utils_.PanelContentFileRelativePath("Trailing", folderPaths)
		statePath := _utils_.PanelStateFileRelativePath("Trailing", folderPaths)
		presetPath := _utils_.PanelPresetFileRelativePath("Trailing", folderPaths)
		builder.WriteString("Trailing Panel\n")
		builder.WriteString(fmt.Sprintf("    Content: %s.\n", contentPath))
		builder.WriteString(fmt.Sprintf("    State:   %s.\n", statePath))
		builder.WriteString(fmt.Sprintf("    Presets: %s.\n", presetPath))
	} else {
		builder.WriteString(fmt.Sprintf("The Trailing area displays the content from the %s screen.\n", general.Trailing.ScreenName))
	}

	// Screen preset files.
	builder.WriteString("\nScreen Preset Files:\n")
	apiPath := _utils_.ScreenPresettingAPIFileRelativePath()
	defaultPresetPath := _utils_.ScreenPresettingDefaultPresetFileRelativePath()
	builder.WriteString(fmt.Sprintf("  The API: %s.\n", apiPath))
	builder.WriteString(fmt.Sprintf("  The Default Preset: %s.\n", defaultPresetPath))

	successMessage = builder.String()
	return
}

const (
	docFileName = _utils_.DocFileName

	docTemplate = `/*
Package {{ call .Data.Funcs.LowerCase .Data.PackageName }} is an Split screen package.
An Split screen lays out a 2 areas; Leading and Trailing.
*/
package {{ call .Data.Funcs.LowerCase .Data.PackageName }}
/*
Files:
{{ .Files }}

Split layout of the Leading and Trailing areas:
1. Leading and Trailing areas are laid out in layoutAreas/layoutAreas.go.
2. modify layoutAreas.go func Layout to layout the Split vertically or horizontally.

Content producers and consumers.
1. The Split container displays leading and trailing content vertically or horizontally.
2. The Leading and Trailing areas may have their content provided by another screen.
3. The Leading and Trailing areas may have their content provided by a panel of the same name (LeadingPanel, TrailingPanel).
4. Therefore, each panel produces content for it's own area in the Split container.
   * The panel's producer gives the panel's content to the panel's consumer.
   * The panel's consumer gives the content to it's Split area (Leading or Trailing).
5. The package has a producer which gives the Split container's entire content to the package's consumer.
   The package provides 1 of 6 different consumers in api.go.
   1. A WindowContentConsumer consumers the container's content for the entire application window.
   2. An AccordionConsumer consumes the container's content for a single AccordionItem in a separate Accordion screen.
   3. An AppTabsTabItemContentConsumer consumes the container's content for a single TabItem in a separate AppTab screen.
   4. A BorderCenterAreaContentConsumer consumes the container's content for a Center Border area in a separate Border screen.
   5. An DocTabsTabItemContentConsumer consumes the container's content for a single TabItem in a separate DocTab screen.
   6. A SplitAreaContentConsumer consumes the container's content for a Leading or Trailing Split area in a separate Split screen.
*/
`
)
