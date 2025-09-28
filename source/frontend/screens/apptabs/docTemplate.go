package simple

import (
	"fmt"
	"strings"

	_utils_ "github.com/josephabbudd-web/kickfyne/source/utils"
)

type docTemplateData struct {
	PackageName string
	PackageDoc  string
	Files       string
	Funcs       _utils_.Funcs
}

func docTemplateSuccessMessage(
	screenPackageName string,
	localPanelNames []string,
	folderPaths *_utils_.FolderPaths,
) (successMessage string) {
	builder := strings.Builder{}
	builder.WriteString("Panels:\n")
	for _, panelName := range localPanelNames {
		contentPath := _utils_.PanelContentFilePath(screenPackageName, panelName, folderPaths)
		statePath := _utils_.PanelStateFilePath(screenPackageName, panelName, folderPaths)
		builder.WriteString("  " + panelName + " Panel\n")
		builder.WriteString(fmt.Sprintf("    Content:   %s.\n", contentPath))
		builder.WriteString(fmt.Sprintf("    State:     %s.\n", statePath))
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
1. Each panel produces content for it's own tabItem in the AppTabs tabbar.
   * The panel's producer gives the panel's content to the panel's consumer.
   * The panel's consumer gives the content to the panel's tabItem.
   * The panel's tabItem is part of the package's AppTabs tabbar.
2. The package has a producer which gives the AppTabs tabbar's content to the package's consumer.
   The package will provide 1 of 4 differenct consumers.
   1. A WindowContentConsumer consumers the content for the entire application window.
   2. An AppTabsTabItemContentConsumer consumes the content for a single TabItem in a separate AppTab screen.
   3. A AppTabsTabItemContentConsumer consumes the content for a single TabItem in a separate AppTabs screen.
   4. An AccordionConsumer consumes the content for a single AccordionItem in a separate Accordion screen.

Tabs:
1. Each tab's open func is in tab/tabs.go.
2. There are also 2 close funcs in tab/tabs.go.
*/
`
)
