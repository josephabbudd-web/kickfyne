package frontend

import (
	"fmt"

	_utils_ "github.com/josephabbudd-web/kickfyne/source/utils"
)

const (
	CmdScreen          = "screen"
	subCmdHelp         = "help"
	verbRemove         = "remove"
	verbList           = "list"
	verbAddSimple      = "add-simple"
	verbAddDocTabs     = "add-doctabs"
	verbAddAppTabs     = "add-apptabs"
	verbAddDocTabsPlus = "add-doctabs+"
	verbAddAppTabsPlus = "add-apptabs+"
	verbAddAccordion   = "add-accordion"
	verbAddBorder      = "add-border"
	verbAddSplit       = "add-split"
	verbAddItem        = "add-item"
	verbRemoveItem     = "remove-item"
)

// Handler passes control to the correct handlers.
func Handler(args []string, importPrefix string, folderPaths *_utils_.FolderPaths) (err error) {

	if len(args) == 0 {
		fmt.Println(Usage())
		return
	}

	defer func() {
		if err != nil {
			err = fmt.Errorf("frontend.Handler: %w", err)
		}
	}()

	switch args[0] {
	case CmdScreen:
		err = handleScreen(args, importPrefix, folderPaths)
	case subCmdHelp:
		fmt.Println(Usage())
	default:
		fmt.Println(Usage())
	}
	return
}
