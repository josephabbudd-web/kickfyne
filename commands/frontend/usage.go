package frontend

import (
	"fmt"
	"os"
	"strings"
)

const (
	screenPackageNameParam = "«screen-package-name»"
	panelNameParam         = "«panel-name» ..."
	itemNameParam          = "«panel-name»|*«screen-name» ..."
	borderAreaNameParam    = "[Top|Bottom|Left|Right|Center[=*«screen-name»]] ..."

	usage3F = "💲 %s %s %s"
	usage4F = "💲 %s %s %s %s"
	usage5F = "💲 %s %s %s %s %s"
)

var (
	UsageScreenAddSimple      = fmt.Sprintf(usage5F, os.Args[0], CmdScreen, verbAddSimple, screenPackageNameParam, panelNameParam)
	UsageScreenAddAccordion   = fmt.Sprintf(usage5F, os.Args[0], CmdScreen, verbAddAccordion, screenPackageNameParam, itemNameParam)
	UsageScreenAddAppTabs     = fmt.Sprintf(usage5F, os.Args[0], CmdScreen, verbAddAppTabs, screenPackageNameParam, itemNameParam)
	UsageScreenAddAppTabsPlus = fmt.Sprintf(usage5F, os.Args[0], CmdScreen, verbAddAppTabsPlus, screenPackageNameParam, itemNameParam)
	UsageScreenAddDocTabs     = fmt.Sprintf(usage5F, os.Args[0], CmdScreen, verbAddDocTabs, screenPackageNameParam, itemNameParam)
	UsageScreenAddDocTabsPlus = fmt.Sprintf(usage5F, os.Args[0], CmdScreen, verbAddDocTabsPlus, screenPackageNameParam, itemNameParam)
	UsageAddBorder            = fmt.Sprintf(usage5F, os.Args[0], CmdScreen, verbAddBorder, screenPackageNameParam, borderAreaNameParam)
	UsageScreenAddItem        = fmt.Sprintf(usage5F, os.Args[0], CmdScreen, verbAddItem, screenPackageNameParam, itemNameParam)
	UsageScreenRemoveItem     = fmt.Sprintf(usage5F, os.Args[0], CmdScreen, verbRemoveItem, screenPackageNameParam, itemNameParam)
	usageScreenHelp           = fmt.Sprintf(usage3F, os.Args[0], CmdScreen, subCmdHelp)
	usageScreenList           = fmt.Sprintf(usage3F, os.Args[0], CmdScreen, verbList)
	UsageScreenRemove         = fmt.Sprintf(usage4F, os.Args[0], CmdScreen, verbRemove, screenPackageNameParam)
)

func UsageScreen() (usage string) {
	screenCommands := []string{
		"Add a Simple screen:     " + UsageScreenAddSimple,
		"Add an Accordion screen: " + UsageScreenAddAccordion,
		"Add an AppTabs screen:   " + UsageScreenAddAppTabs,
		" Plus (+) configuration: " + UsageScreenAddAppTabsPlus,
		"Add a DocTabs screen:    " + UsageScreenAddDocTabs,
		" Plus (+) configuration: " + UsageScreenAddDocTabsPlus,
		"Add a Border screen:     " + UsageAddBorder,
		"Remove a screen:         " + UsageScreenRemove,
	}
	helpCommands := []string{
		usageScreenList,
		usageScreenHelp,
	}
	itemCommands := []string{
		UsageScreenAddItem,
		UsageScreenRemoveItem,
	}
	usage = `📺 MANAGING SCREENS:

ADD & REMOVE A SCREEN:
` + strings.Join(screenCommands, "\n") + `

MODIFY A SCREEN:
` + strings.Join(itemCommands, "\n") + `

HELP:
` + strings.Join(helpCommands, "\n") + `

TabItem names:
* A tab-item-name prefixed with '*':
  Will get its content from the screen package of the same name.
  That screen must already exist.
* A tab-item-name not prefixed with '*':
  Will get its content from a panel of the same name.
  That panel will be created in the same tabbar screen package.

AccordionItem names:
* An accordion-item-name prefixed with '*':
  Will get its content from the screen package of the same name.
  That screen must already exist.
* An accordion-item-name not prefixed with '*':
  Will get its content from a panel of the same name.
  That panel will be created in the same accordion screen package.

BorderArea names:
* An area-name prefixed with '*':
  Will get its content from the screen package of the same name.
  That screen must already exist.
* An area-name not prefixed with '*':
  Will get its content from a panel of the same name.
  That panel will be created in the same border screen package.
* There must be 2 or more areas defined for a border screen.

After a screen is added:
1. A link to it's screen.zig file is displayed.
2. A search for KICKFYNE TODO in the screen package files will reveal instructions for proper developement and management of the screen operation.
`
	return
}

func Usage() (usage string) {
	usage = `👀 THE front-end:
Screen names must be in PascalCase.
Panel names must be in PascalCase.
TabItem names must be in PascalCase.
AccordionItem names must be in PascalCase.
BorderArea names must be in PascalCase.

` + UsageScreen()
	return
}
