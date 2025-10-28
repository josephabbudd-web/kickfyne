package frontend

import (
	"fmt"
	"os"
	"strings"
)

const (
	screenPackageNameParam = "«screen-package-name»"
	panelNameParam         = "«panel-name, ...»"
	tabItemNameParam       = "«[*]tab-item-name, ...»"
	accordionItemNameParam = "«[*]accordion-item-name, ...»"
	screenItemNameParam    = "«[*]item-name or panel-name, ...»"

	usage3F = "💲 %s %s %s"
	usage4F = "💲 %s %s %s %s"
	usage5F = "💲 %s %s %s %s %s"
)

var (
	UsageScreenAddSimple      = fmt.Sprintf(usage5F, os.Args[0], CmdScreen, verbAddSimple, screenPackageNameParam, panelNameParam)
	UsageScreenAddAccordion   = fmt.Sprintf(usage5F, os.Args[0], CmdScreen, verbAddAccordion, screenPackageNameParam, accordionItemNameParam)
	UsageScreenAddAppTabs     = fmt.Sprintf(usage5F, os.Args[0], CmdScreen, verbAddAppTabs, screenPackageNameParam, tabItemNameParam)
	UsageScreenAddAppTabsPlus = fmt.Sprintf(usage5F, os.Args[0], CmdScreen, verbAddAppTabsPlus, screenPackageNameParam, tabItemNameParam)
	UsageScreenAddDocTabs     = fmt.Sprintf(usage5F, os.Args[0], CmdScreen, verbAddDocTabs, screenPackageNameParam, tabItemNameParam)
	UsageScreenAddDocTabsPlus = fmt.Sprintf(usage5F, os.Args[0], CmdScreen, verbAddDocTabsPlus, screenPackageNameParam, tabItemNameParam)
	UsageScreenAddItem        = fmt.Sprintf(usage5F, os.Args[0], CmdScreen, verbAddItem, screenPackageNameParam, screenItemNameParam)
	UsageScreenRemoveItem     = fmt.Sprintf(usage5F, os.Args[0], CmdScreen, verbRemoveItem, screenPackageNameParam, screenItemNameParam)
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
		"Add a DocTabs screen:    " + UsageScreenAddAppTabs,
		" Plus (+) configuration: " + UsageScreenAddAppTabsPlus,
		"Add an DocTabs screen: " + UsageScreenAddDocTabs,
		UsageScreenRemove,
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
  
After a screen is added:
1. A link to it's screen.zig file is displayed.
2. A search for KICKFYNE TODO in the screen package files will reveal instructions for proper developement and management of the screen operation.
`
	return
}

func Usage() (usage string) {
	usage = `👀 THE front-end:
Screen names must be in TitleCase.
Panel names must be in TitleCase.
TabItem names must be in TitleCase.
AccordionItem names must be in TitleCase.

` + UsageScreen()
	return
}
