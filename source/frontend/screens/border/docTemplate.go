package border

import (
	"fmt"
	"strings"

	_data_ "github.com/josephabbudd-web/kickfyne/source/frontend/screens/border/data"
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

	builder.WriteString("The border areas and their panel files:\n")
	if len(general.Center.Area) == 0 {
		builder.WriteString("The Center area is not being used.\n")
	} else {
		if general.Center.IsLocal {
			builder.WriteString(fmt.Sprintf("The Center area displays the content from the %sPanel.\n", general.Center.ItemName))
			contentPath := _utils_.PanelContentFileRelativePath(general.Center.ItemName, folderPaths)
			statePath := _utils_.PanelStateFileRelativePath(general.Center.ItemName, folderPaths)
			presetPath := _utils_.PanelPresetFileRelativePath(general.Center.ItemName, folderPaths)
			builder.WriteString("\n  " + general.Center.ItemName + " Panel\n")
			builder.WriteString(fmt.Sprintf("    Content: %s.\n", contentPath))
			builder.WriteString(fmt.Sprintf("    State:   %s.\n", statePath))
			builder.WriteString(fmt.Sprintf("    Presets: %s.\n", presetPath))
		} else {
			builder.WriteString(fmt.Sprintf("The Center area displays the content from the %s screen.\n", general.Center.ItemName))
		}
	}

	if len(general.Bottom.Area) == 0 {
		builder.WriteString("The Bottom area is not being used.\n")
	} else {
		if general.Bottom.IsLocal {
			builder.WriteString(fmt.Sprintf("The Bottom area displays the content from the %sPanel.\n", general.Bottom.ItemName))
			contentPath := _utils_.PanelContentFileRelativePath(general.Bottom.ItemName, folderPaths)
			statePath := _utils_.PanelStateFileRelativePath(general.Bottom.ItemName, folderPaths)
			presetPath := _utils_.PanelPresetFileRelativePath(general.Bottom.ItemName, folderPaths)
			builder.WriteString("\n  " + general.Bottom.ItemName + " Panel\n")
			builder.WriteString(fmt.Sprintf("    Content: %s.\n", contentPath))
			builder.WriteString(fmt.Sprintf("    State:   %s.\n", statePath))
			builder.WriteString(fmt.Sprintf("    Presets: %s.\n", presetPath))
		} else {
			builder.WriteString(fmt.Sprintf("The Bottom area displays the content from the %s screen.\n", general.Bottom.ItemName))
		}
	}

	if len(general.Left.Area) == 0 {
		builder.WriteString("The Left area is not being used.\n")
	} else {
		if general.Left.IsLocal {
			builder.WriteString(fmt.Sprintf("The Left area displays the content from the %sPanel.\n", general.Left.ItemName))
			contentPath := _utils_.PanelContentFileRelativePath(general.Left.ItemName, folderPaths)
			statePath := _utils_.PanelStateFileRelativePath(general.Left.ItemName, folderPaths)
			presetPath := _utils_.PanelPresetFileRelativePath(general.Left.ItemName, folderPaths)
			builder.WriteString("\n  " + general.Left.ItemName + " Panel\n")
			builder.WriteString(fmt.Sprintf("    Content: %s.\n", contentPath))
			builder.WriteString(fmt.Sprintf("    State:   %s.\n", statePath))
			builder.WriteString(fmt.Sprintf("    Presets: %s.\n", presetPath))
		} else {
			builder.WriteString(fmt.Sprintf("The Left area displays the content from the %s screen.\n", general.Left.ItemName))
		}
	}

	if len(general.Right.Area) == 0 {
		builder.WriteString("The Right area is not being used.\n")
	} else {
		if general.Right.IsLocal {
			builder.WriteString(fmt.Sprintf("The Right area displays the content from the %sPanel.\n", general.Right.ItemName))
			contentPath := _utils_.PanelContentFileRelativePath(general.Right.ItemName, folderPaths)
			statePath := _utils_.PanelStateFileRelativePath(general.Right.ItemName, folderPaths)
			presetPath := _utils_.PanelPresetFileRelativePath(general.Right.ItemName, folderPaths)
			builder.WriteString("\n  " + general.Right.ItemName + " Panel\n")
			builder.WriteString(fmt.Sprintf("    Content: %s.\n", contentPath))
			builder.WriteString(fmt.Sprintf("    State:   %s.\n", statePath))
			builder.WriteString(fmt.Sprintf("    Presets: %s.\n", presetPath))
		} else {
			builder.WriteString(fmt.Sprintf("The Right area displays the content from the %s screen.\n", general.Right.ItemName))
		}
	}

	if len(general.Center.Area) == 0 {
		builder.WriteString("The Center area is not being used.\n")
	} else {
		if general.Center.IsLocal {
			builder.WriteString(fmt.Sprintf("The Center area displays the content from the %sPanel.\n", general.Center.ItemName))
			contentPath := _utils_.PanelContentFileRelativePath(general.Center.ItemName, folderPaths)
			statePath := _utils_.PanelStateFileRelativePath(general.Center.ItemName, folderPaths)
			presetPath := _utils_.PanelPresetFileRelativePath(general.Center.ItemName, folderPaths)
			builder.WriteString("\n  " + general.Center.ItemName + " Panel\n")
			builder.WriteString(fmt.Sprintf("    Content: %s.\n", contentPath))
			builder.WriteString(fmt.Sprintf("    State:   %s.\n", statePath))
			builder.WriteString(fmt.Sprintf("    Presets: %s.\n", presetPath))
		} else {
			builder.WriteString(fmt.Sprintf("The Center area displays the content from the %s screen.\n", general.Center.ItemName))
		}
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
Package {{ call .Data.Funcs.LowerCase .Data.PackageName }} is an Border screen package.
An Border screen lays out 5 areas; Center, Bottom, Left, Right and Center.
A border area either displays the content from a panel or from another screen.
*/
package {{ call .Data.Funcs.LowerCase .Data.PackageName }}
/*
Files:
{{ .Files }}

AppTabs layout of the TabItems:
1. TabItems are laid out in layoutitems/layout.go.
2. You may want to modify layout.go whenever a tab is appended or removed from this package.

AppTabs callbacks for TabItem OnSelected and OnUnselected are in tabitemHandlers.go

Tabs Internals:
1. Each tab's open func is in tab/tabs.go.
2. There are also 2 close funcs in tab/tabs.go.

Content producers and consumers.
1. The AppTabs container displays a tabbar above, below, to the left or to the right of the selected TabItem's content.
2. Some TabItems may have their content provided by another screen. That TabItem and that screen will have the same name.
3. Some TabItems may have their content provided by a panel in this screen. That TabItem and that panel will have the same name.
4. Therefore, each panel produces content for it's own TabItem in the AppTabs container.
   * The panel's producer gives the panel's content to the panel's consumer.
   * The panel's consumer gives the content to the panel's TabItem.
   * The panel's TabItem is part of the package's AppTabs container.
5. The package has a producer which gives the AppTabs container's entire content to the package's consumer.
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
