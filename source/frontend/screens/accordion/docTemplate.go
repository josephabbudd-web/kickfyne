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
1. The Accordion widget is a vertical list of AccordionItems.
2. Some AccordionItems may have their content provided by another screen.
3. Some AccordionItems may have their content provided by a panel by the same name.
4. Therefore, each panel produces content for it's own AccordionItem in the Accordion tabbar.
   * The panel's producer gives the panel's content to the panel's consumer.
   * The panel's consumer gives the content to the panel's AccordionItem.
   * The panel's AccordionItem is part of the package's Accordion tabbar.
4. The package has a producer which gives the Accordion widget's entire content to the package's consumer.
   The package will provide 1 of 4 differenct consumers in api.go.
   1. A WindowContentConsumer consumers the content for the entire application window.
   2. An AppTabsTabItemContentConsumer consumes the content for a single TabItem in a separate AppTab screen.
   3. A AppTabsTabItemContentConsumer consumes the content for a single TabItem in a separate Accordion screen.
   4. An AccordionConsumer consumes the content for a single AccordionItem in a separate Accordion screen.


Accordion layout of the AccordionItems:
1. AccordionItems are laid out in layoutitems/layout.go.
2. You may want to modify layout.go whenever an AccordionItem is added or removed from this package.

Tabs Internals:
1. Each AccordionItem's open func is in deps/accordionItems/accordionItems.go.
2. There are also 2 close funcs in deps/accordionItems/accordionItems.go.
*/
`
)
