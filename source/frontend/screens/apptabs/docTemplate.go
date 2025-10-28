package apptabs

import (
	"fmt"
	"strings"

	_utils_ "github.com/josephabbudd-web/kickfyne/source/utils"
)

type docTemplateData struct {
	PackageName  string
	PackageDoc   string
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
	builder.WriteString("Panels:\n")
	for _, panelName := range localPanelNames {
		contentPath := _utils_.PanelContentFilePath(screenPackageName, panelName, folderPaths)
		statePath := _utils_.PanelStateFilePath(screenPackageName, panelName, folderPaths)
		presetPath := _utils_.PanelPresetFilePath(screenPackageName, panelName, folderPaths)
		builder.WriteString("  " + panelName + " Panel\n")
		builder.WriteString(fmt.Sprintf("    Content: %s.\n", contentPath))
		builder.WriteString(fmt.Sprintf("    State:   %s.\n", statePath))
		builder.WriteString(fmt.Sprintf("    Presets: %s.\n", presetPath))
	}
	successMessage = builder.String()
	return
}

const (
	docFileName = _utils_.DocFileName

	docTemplate = `{{ call .Funcs.Comment .PackageDoc }}
package {{ call .Funcs.LowerCase .PackageName }}
/*
Files:
{{ .Files }}

Content producers and consumers.
1. The AppTabs container is a horizontal list of TabItems displayed above the selected TabItem's content.
2. Some TabItems may have their content provided by another screen.
3. Some TabItems may have their content provided by a panel by the same name.
4. Therefore, each panel produces content for it's own TabItem in the AppTabs container.
   * The panel's producer gives the panel's content to the panel's consumer.
   * The panel's consumer gives the content to the panel's TabItem.
   * The panel's TabItem is part of the package's AppTabs container.
5. The package has a producer which gives the DocTabs container's entire content to the package's consumer.
   The package provides 1 of 4 different consumers in api.go.
   1. A WindowContentConsumer consumers the container's content for the entire application window.
   2. An AppTabsTabItemContentConsumer consumes the container's content for a single TabItem in a separate AppTab screen.
   3. A DocTabsTabItemContentConsumer consumes the container's content for a single TabItem in a separate AppTabs screen.
   4. An AccordionConsumer consumes the container's content for a single AccordionItem in a separate Accordion screen.

AppTabs layout of the TabItem:
1. TabItems are laid out in layoutitems/layout.go.
2. You may want to modify layout.go whenever a tab is appended or removed from this package.

Tabs Internals:
1. Each tab's open func is in tab/tabs.go.
2. There are also 2 close funcs in tab/tabs.go.
*/
`
)
