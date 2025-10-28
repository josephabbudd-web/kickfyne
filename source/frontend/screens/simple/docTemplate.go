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
1. The Simple screen displays only one of its panels at a time.
2. The package has 1 producer that all of the panels use.
3. That producer which gives the content's of its only visible panel to the package's consumer.
   The package provides 1 of 4 different consumers in api.go.
   1. A WindowContentConsumer consumers the container's content for the entire application window.
   2. An AppTabsTabItemContentConsumer consumes the container's content for a single TabItem in a separate AppTab screen.
   3. A SimpleTabItemContentConsumer consumes the container's content for a single TabItem in a separate DocTab screen.
   4. An AccordionConsumer consumes the container's content for a single AccordionItem in a separate Accordion screen.
*/
`
)
