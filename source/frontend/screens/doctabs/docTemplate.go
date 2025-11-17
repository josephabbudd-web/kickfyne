package doctabs

import (
	"fmt"
	"strings"

	_utils_ "github.com/josephabbudd-web/kickfyne/source/utils"
)

type docTemplateData struct {
	PackageName  string
	Files        string
	UseConfigTab bool
	Funcs        _utils_.Funcs
}

func files(
	screenPackageName string,
	localPanelNames []string,
	folderPaths *_utils_.FolderPaths,
) (successMessage string) {
	builder := strings.Builder{}

	builder.WriteString("Panel Files:\n")
	for _, panelName := range localPanelNames {
		contentPath := _utils_.PanelContentFileRelativePath(panelName, folderPaths)
		statePath := _utils_.PanelStateFileRelativePath(panelName, folderPaths)
		presetPath := _utils_.PanelPresetFileRelativePath(panelName, folderPaths)
		builder.WriteString("\n  " + panelName + " Panel\n")
		builder.WriteString(fmt.Sprintf("    Content: %s.\n", contentPath))
		builder.WriteString(fmt.Sprintf("    State:   %s.\n", statePath))
		builder.WriteString(fmt.Sprintf("    Presets: %s.\n", presetPath))
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
Package {{ call .Funcs.LowerCase .PackageName }} is an DocTabs screen package.
A DocTabs screen lays out a tabbar above, below, left or right of where the TabItem's content is displayed.
A TabItem either displays the content from a panel or from another screen.
A TabItem has the exact same name as it's content panel or it's content screen.
*/
package {{ call .Funcs.LowerCase .PackageName }}
/*
Files:
{{ .Files }}

DocTabs layout of the TabItem:
1. TabItems are laid out in layoutitems/layout.go.
2. You may want to modify layout.go whenever a tab is appended or removed from this package.

Tabs Internals:
1. Each tab's open func is in tab/tabs.go.
2. There are also 2 close funcs in tab/tabs.go.

Content producers and consumers.
1. The DocTabs container displays a tabbar above, below, to the left or to the right of the selected TabItem's content.
2. Some TabItems may have their content provided by another screen. That TabItem and that screen will have the same name.
3. Some TabItems may have their content provided by a panel in this screen. That TabItem and that panel will have the same name.
4. Therefore, each panel produces content for it's own TabItem in the DocTabs container.
   * The panel's producer gives the panel's content to the panel's consumer.
   * The panel's consumer gives the content to the panel's TabItem.
   * The panel's TabItem is part of the package's DocTabs container.
5. The package has a producer which gives the DocTabs container's entire content to the package's consumer.
   The package provides 1 of 4 different consumers in api.go.
   1. A WindowContentConsumer consumers the container's content for the entire application window.
   2. An DocTabsTabItemContentConsumer consumes the container's content for a single TabItem in a separate DocTab screen.
   3. A DocTabsTabItemContentConsumer consumes the container's content for a single TabItem in a separate DocTabs screen.
   4. An AccordionConsumer consumes the container's content for a single AccordionItem in a separate Accordion screen.
*/
`
)
