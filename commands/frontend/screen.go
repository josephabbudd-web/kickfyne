package frontend

import (
	"fmt"

	_manifest_ "github.com/josephabbudd-web/kickfyne/manifest"
	_frontend_ "github.com/josephabbudd-web/kickfyne/source/frontend"
	_utils_ "github.com/josephabbudd-web/kickfyne/source/utils"
)

// handleScreen passes control to the correct handlers.
func handleScreen(args []string, importPrefix string, folderPaths *_utils_.FolderPaths) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("frontend.handleScreen: %w", err)
		}
	}()

	var manifest _manifest_.Manifest
	if manifest, err = _manifest_.New(folderPaths); err != nil {
		return
	}

	if len(args) == 1 {
		fmt.Println(UsageScreen())
		return
	}
	// args[0] is "screen"
	// args[1] is the verb
	switch args[1] {
	case verbList:
		// args[0] is "screen"
		// args[1] is "list"
		if len(args) != 2 {
			fmt.Println(UsageScreen())
			return
		}
		handleScreenList(manifest)
	case verbRemoveItem:
		// args[0] is "screen"
		// args[1] is "remove-item"
		// args[2] is the «screen-package-name»
		// args[3..] is «item-name, ...».
		if len(args) < 4 {
			fmt.Println(UsageScreen())
			return
		}
		// Validate the screen package name.
		if isValid, failureMessage := _utils_.ValidatePascalCase(args[2], "screen"); !isValid || err != nil {
			fmt.Println(failureMessage)
			return
		}
		screenKind := manifest.ScreenKind(args[2])
		if screenKind == _manifest_.NilScreen {
			fmt.Printf("Failure: The screen package %q does not exists.\n", args[2])
			return
		}
		// Validate item names.
		var haveInvalidItemName bool
		for _, itemName := range args[3:] {
			if !manifest.HasScreenItem(args[2], itemName) {
				haveInvalidItemName = true
				switch screenKind {
				case _manifest_.SimpleScreen:
					fmt.Printf("Failure: The %q screen package does not have a panel named %q.\n", args[2], itemName)
				default:
					fmt.Printf("Failure: The %q screen package does not have an item named \"%s\".\n", args[2], itemName)
				}
				continue
			}
			isRemote, isScreenName := manifest.IsRemoteNameIsScreenName(itemName)
			switch screenKind {
			case _manifest_.SimpleScreen:
				if isRemote && isScreenName {
					haveInvalidItemName = true
					fmt.Printf("Failure: The panel name \"%s\" can not be a reference to a screen package.\n    Simple screens only have local panels.\n", itemName)
				}
			case _manifest_.AccordionScreen, _manifest_.AppTabsScreen, _manifest_.DocTabsScreen:
				if isRemote && !isScreenName {
					haveInvalidItemName = true
					fmt.Printf("Failure: The item name \"%s\" is not a reference to an existing screen package.\n", itemName)
				}
			}
		}
		if haveInvalidItemName {
			return
		}
		// Names must be unique.
		if isValid, failureMessage := _utils_.ValidatePanelNames(args[3:]); !isValid {
			fmt.Println(failureMessage)
			return
		}
		// Count
		lengthArgsItems := len(args[3:])
		lengthScreenItems := manifest.CountScreenItems(args[2])
		if lengthArgsItems == lengthScreenItems {
			switch screenKind {
			case _manifest_.AccordionScreen:
				fmt.Printf("Failure: The Accordion screen %q needs at least 1 AccordionItem.\n", args[2])
			case _manifest_.AppTabsScreen:
				fmt.Printf("Failure: The AppTabs screen %q needs at least 1 TabItem.\n", args[2])
			case _manifest_.DocTabsScreen:
				fmt.Printf("Failure: The DocTabs screen %q needs at least 1 TabItem.\n", args[2])
			case _manifest_.SimpleScreen:
				fmt.Printf("Failure: The Simple screen %q needs at least 1 panel.\n", args[2])
			}
			return
		}
		// Remove items from the screen.
		switch screenKind {
		case _manifest_.SimpleScreen:
			screenDocComment := fmt.Sprintf("Package %s is a Simple screen package.\nA simple screen displays only one panel at a time.\nKICKFYNE TODO: Complete this package doc commment.", args[2])
			if err = handleScreenRemoveSimplePanels(args[2], args[3:], screenDocComment, importPrefix, folderPaths); err != nil {
				return
			}
		case _manifest_.AccordionScreen:
			screenDocComment := fmt.Sprintf("Package %s is an AppTabs screen package.\nA tabbar screen where a tab displays it's panel or another screen.\nKICKFYNE TODO: Complete this package doc commment.", args[2])
			if err = handleScreenRemoveAccordionItems(args[2], args[3:], screenDocComment, importPrefix, folderPaths); err != nil {
				return
			}
		case _manifest_.AppTabsScreen:
			screenDocComment := fmt.Sprintf("Package %s is an AppTabs screen package.\nA tabbar screen where a tab displays it's panel or another screen.\nKICKFYNE TODO: Complete this package doc commment.", args[2])
			if err = handleScreenRemoveAppTabsItems(args[2], args[3:], screenDocComment, importPrefix, folderPaths); err != nil {
				return
			}
		case _manifest_.DocTabsScreen:
			screenDocComment := fmt.Sprintf("Package %s is a DocTabs screen package.\nA tabbar screen where a tab displays it's panel or another screen.\nKICKFYNE TODO: Complete this package doc commment.", args[2])
			if err = handleScreenRemoveDocTabsItems(args[2], args[3:], screenDocComment, importPrefix, folderPaths); err != nil {
				return
			}
		}
		if err = _frontend_.RebuildFrontendGo(importPrefix, folderPaths); err != nil {
			return
		}
		// Update the manifest.
		manifest.RemoveItems(args[2], args[3:]...)
		err = manifest.Write(folderPaths)
		return
	case verbAddItem:
		// args[0] is "screen"
		// args[1] is "add-item"
		// args[2] is the «screen-package-name»
		// args[3..] is «item-name, ...».
		if len(args) < 4 {
			fmt.Println(UsageScreen())
			return
		}
		// Validate the screen package name.
		if isValid, failureMessage := _utils_.ValidatePascalCase(args[2], "screen"); !isValid || err != nil {
			fmt.Println(failureMessage)
			return
		}
		screenKind := manifest.ScreenKind(args[2])
		if screenKind == _manifest_.NilScreen {
			fmt.Printf("Failure: The screen package %q does not exists.\n", args[2])
			return
		}
		// Validate item names.
		var haveInvalidItemName bool
		for _, itemName := range args[3:] {
			if manifest.HasScreenItem(args[2], itemName) {
				haveInvalidItemName = true
				switch screenKind {
				case _manifest_.SimpleScreen:
					fmt.Printf("Failure: The %q screen package already has a panel named %q.\n", args[2], itemName)
				default:
					fmt.Printf("Failure: The %q screen package already has an item named \"*%s\".\n", args[2], itemName)
				}
				continue
			}
			isRemote, isScreenName := manifest.IsRemoteNameIsScreenName(itemName)
			switch screenKind {
			case _manifest_.SimpleScreen:
				if isRemote && isScreenName {
					haveInvalidItemName = true
					fmt.Printf("Failure: The panel name \"*%s\" can not be a reference to a screen package.\n    Simple screens only have local panels.\n", itemName)
				}
			case _manifest_.AccordionScreen, _manifest_.AppTabsScreen, _manifest_.DocTabsScreen:
				if isRemote && !isScreenName {
					haveInvalidItemName = true
					fmt.Printf("Failure: The item name \"*%s\" is not a reference to an existing screen package.\n", itemName)
				}
			}
		}
		if haveInvalidItemName {
			return
		}
		if isValid, failureMessage := _utils_.ValidatePanelNames(args[3:]); !isValid {
			fmt.Println(failureMessage)
			return
		}
		// Add items to the screen.
		switch screenKind {
		case _manifest_.SimpleScreen:
			screenDocComment := fmt.Sprintf("Package %s is a Simple screen package.\nA simple screen displays only one panel at a time.\nKICKFYNE TODO: Complete this package doc commment.", args[2])
			if err = handleScreenAddSimplePanels(args[2], args[3:], screenDocComment, importPrefix, folderPaths); err != nil {
				return
			}
		case _manifest_.AccordionScreen:
			screenDocComment := fmt.Sprintf("Package %s is an AppTabs screen package.\nA tabbar screen where a tab displays it's panel or another screen.\nKICKFYNE TODO: Complete this package doc commment.", args[2])
			if err = handleScreenAddAccordionItems(args[2], args[3:], screenDocComment, importPrefix, folderPaths); err != nil {
				return
			}
		case _manifest_.AppTabsScreen:
			screenDocComment := fmt.Sprintf("Package %s is an AppTabs screen package.\nA tabbar screen where a tab displays it's panel or another screen.\nKICKFYNE TODO: Complete this package doc commment.", args[2])
			if err = handleScreenAddAppTabsItems(args[2], args[3:], screenDocComment, importPrefix, folderPaths); err != nil {
				return
			}
		case _manifest_.DocTabsScreen:
			screenDocComment := fmt.Sprintf("Package %s is a DocTabs screen package.\nA tabbar screen where a tab displays it's panel or another screen.\nKICKFYNE TODO: Complete this package doc commment.", args[2])
			if err = handleScreenAddDocTabsItems(args[2], args[3:], screenDocComment, importPrefix, folderPaths); err != nil {
				return
			}
		}
		if err = _frontend_.RebuildFrontendGo(importPrefix, folderPaths); err != nil {
			return
		}
		// Update the manifest.
		manifest.AddItems(args[2], args[3:]...)
		err = manifest.Write(folderPaths)
		return
	case verbAddSimple:
		// args[0] is "screen"
		// args[1] is "add-simple"
		// args[2] is the «screen-package-name»
		// args[3..] is «panel-name, ...».
		if len(args) < 4 {
			fmt.Println(UsageScreen())
			return
		}
		// Validate the screen package name.
		if isValid, failureMessage := _utils_.ValidatePascalCase(args[2], "screen"); !isValid || err != nil {
			fmt.Println(failureMessage)
			return
		}
		if manifest.HasScreen(args[2]) {
			fmt.Printf("Failure: The screen package %q already exists.\n", args[2])
			return
		}
		// Validate item names.
		var haveInvalidItemName bool
		for _, itemName := range args[3:] {
			isRemote, _ := manifest.IsRemoteNameIsScreenName(itemName)
			if isRemote {
				haveInvalidItemName = true
				fmt.Printf("Failure: The panel name %q can not be a reference to a screen package.\n", itemName)
			}
		}
		if haveInvalidItemName {
			return
		}
		if isValid, failureMessage := _utils_.ValidatePanelNames(args[3:]); !isValid {
			fmt.Println(failureMessage)
			return
		}
		// Build the new screen.
		screenDocComment := fmt.Sprintf("Package %s is a Simple screen package.\nA simple screen displays only one panel at a time.\nKICKFYNE TODO: Complete this package doc commment.", args[2])
		if err = handleScreenAddSimple(args[2], args[3:], screenDocComment, importPrefix, folderPaths); err != nil {
			return
		}
		if err = _frontend_.RebuildFrontendGo(importPrefix, folderPaths); err != nil {
			return
		}
		// Update the manifest.
		manifest.AddScreen(args[2], _manifest_.SimpleScreen, args[3:]...)
		err = manifest.Write(folderPaths)
		return
	case verbAddAccordion:
		// args[0] is "screen"
		// args[1] is "add-accordion"
		// args[2] is the «screen-package-name»
		// args[3..] is «[*]accordion-item-name, ...».
		if len(args) < 4 {
			fmt.Println(UsageScreen())
			return
		}
		// Validate the screen package name.
		if isValid, failureMessage := _utils_.ValidatePascalCase(args[2], "screen"); !isValid || err != nil {
			fmt.Println(failureMessage)
			return
		}
		if manifest.HasScreen(args[2]) {
			fmt.Printf("Failure: The screen package %q already exists.\n", args[2])
			return
		}
		// Validate item names.
		var haveInvalidItemName bool
		for _, itemName := range args[3:] {
			isRemote, isScreen := manifest.IsRemoteNameIsScreenName(itemName)
			if isRemote && !isScreen {
				haveInvalidItemName = true
				fmt.Printf("Failure: The AccordionItem name %q does not reference an existing screen package.\n", itemName)
			}
		}
		if haveInvalidItemName {
			return
		}
		if isValid, failureMessage := _utils_.ValidateAccordionItemNames(args[3:]); !isValid {
			fmt.Println(failureMessage)
			return
		}
		// Build the new screen.
		screenDocComment := fmt.Sprintf("Package %s is an Accordion screen package.\nAn accordion screen displays each panel as a titled item that can be opened.\nKICKFYNE TODO: Complete this package doc commment.", args[2])
		if err = handleScreenAddAccordion(args[2], args[3:], screenDocComment, importPrefix, folderPaths); err != nil {
			return
		}
		if err = _frontend_.RebuildFrontendGo(importPrefix, folderPaths); err != nil {
			return
		}
		// Update the manifest.
		manifest.AddScreen(args[2], _manifest_.AccordionScreen, args[3:]...)
		err = manifest.Write(folderPaths)
		return
	case verbAddAppTabs:
		// args[0] is "screen"
		// args[1] is "add-apptabs"
		// args[2] is the «screen-package-name»
		// args[3..] is «[*]tab-item-name, ...».
		if len(args) < 4 {
			fmt.Println(UsageScreen())
			return
		}
		// Validate the screen package name.
		if isValid, failureMessage := _utils_.ValidatePascalCase(args[2], "screen"); !isValid || err != nil {
			fmt.Println(failureMessage)
			return
		}
		if manifest.HasScreen(args[2]) {
			fmt.Printf("Failure: The screen package %q already exists.\n", args[2])
			return
		}
		// Validate item names.
		var haveInvalidItemName bool
		for _, itemName := range args[3:] {
			isRemote, isScreen := manifest.IsRemoteNameIsScreenName(itemName)
			if isRemote && !isScreen {
				haveInvalidItemName = true
				fmt.Printf("Failure: The TabItem name %q does not reference an existing screen package.\n", itemName)
			}
		}
		if haveInvalidItemName {
			return
		}
		if isValid, failureMessage := _utils_.ValidateTabNames(args[3:]); !isValid {
			fmt.Println(failureMessage)
			return
		}
		// Build the new screen.
		screenDocComment := fmt.Sprintf("Package %s is an AppTabs screen package.\nA tabbar screen where a tab displays it's panel or another screen.\nThe user is not able to close a tab but your app can.\nKICKFYNE TODO: Complete this package doc commment.", args[2])
		if err = handleScreenAddAppTabs(args[2], args[3:], screenDocComment, importPrefix, folderPaths); err != nil {
			return
		}
		if err = _frontend_.RebuildFrontendGo(importPrefix, folderPaths); err != nil {
			return
		}
		// Update the manifest.
		manifest.AddScreen(args[2], _manifest_.AppTabsScreen, args[3:]...)
		err = manifest.Write(folderPaths)
		return
	case verbAddAppTabsPlus:
		// args[0] is "screen"
		// args[1] is "add-apptabs+"
		// args[2] is the «screen-package-name»
		// args[3..] is «[*]tab-item-name, ...».
		if len(args) < 4 {
			fmt.Println(UsageScreen())
			return
		}
		// Validate the screen package name.
		if isValid, failureMessage := _utils_.ValidatePascalCase(args[2], "screen"); !isValid || err != nil {
			fmt.Println(failureMessage)
			return
		}
		if manifest.HasScreen(args[2]) {
			fmt.Printf("Failure: The screen package %q already exists.\n", args[2])
			return
		}
		// Validate item names.
		var haveInvalidItemName bool
		for _, itemName := range args[3:] {
			isRemote, isScreen := manifest.IsRemoteNameIsScreenName(itemName)
			if isRemote && !isScreen {
				haveInvalidItemName = true
				fmt.Printf("Failure: The TabItem name %q does not reference an existing screen package.\n", itemName)
			}
		}
		if haveInvalidItemName {
			return
		}
		if isValid, failureMessage := _utils_.ValidateTabNames(args[3:]); !isValid {
			fmt.Println(failureMessage)
			return
		}
		// Build the new screen.
		screenDocComment := fmt.Sprintf("Package %s is an AppTabs screen package.\nA tabbar screen where a tab displays it's panel or another screen.\nThe user is not able to close a tab but your app can.\nThe first tab is a configuration tab and allows the user to configure where the tabbar is located.\nKICKFYNE TODO: Complete this package doc commment.", args[2])
		itemNames := make([]string, len(args[3:])+1)
		itemNames[0] = _utils_.ConfigTabName
		for i, itemName := range args[3:] {
			itemNames[i+1] = itemName
		}
		if err = handleScreenAddAppTabs(args[2], itemNames, screenDocComment, importPrefix, folderPaths); err != nil {
			return
		}
		if err = _frontend_.RebuildFrontendGo(importPrefix, folderPaths); err != nil {
			return
		}
		// Update the manifest.
		manifest.AddScreen(args[2], _manifest_.AppTabsScreen, itemNames...)
		err = manifest.Write(folderPaths)
		return
	case verbAddDocTabs:
		// args[0] is "screen"
		// args[1] is "add-doctabs"
		// args[2] is the «screen-package-name»
		// args[3..] is «[*]tab-item-name, ...».
		if len(args) < 4 {
			fmt.Println(UsageScreen())
			return
		}
		// Validate the screen package name.
		if isValid, failureMessage := _utils_.ValidatePascalCase(args[2], "screen"); !isValid || err != nil {
			fmt.Println(failureMessage)
			return
		}
		if manifest.HasScreen(args[2]) {
			fmt.Printf("Failure: The screen package %q already exists.\n", args[2])
			return
		}
		// Validate item names.
		var haveInvalidItemName bool
		for _, itemName := range args[3:] {
			isRemote, isScreen := manifest.IsRemoteNameIsScreenName(itemName)
			if isRemote && !isScreen {
				haveInvalidItemName = true
				fmt.Printf("Failure: The TabItem name %q does not reference an existing screen package.\n", itemName)
			}
		}
		if haveInvalidItemName {
			return
		}
		if isValid, failureMessage := _utils_.ValidateTabNames(args[3:]); !isValid {
			fmt.Println(failureMessage)
			return
		}
		// Build the new screen.
		screenDocComment := fmt.Sprintf("Package %s is a DocTabs screen package.\nA tabbar screen where a tab displays it's panel or another screen.\nKICKFYNE TODO: Complete this package doc commment.", args[2])
		if err = handleScreenAddDocTabs(args[2], args[3:], screenDocComment, importPrefix, folderPaths); err != nil {
			return
		}
		if err = _frontend_.RebuildFrontendGo(importPrefix, folderPaths); err != nil {
			return
		}
		// Update the manifest.
		manifest.AddScreen(args[2], _manifest_.DocTabsScreen, args[3:]...)
		err = manifest.Write(folderPaths)
		return
	case verbAddDocTabsPlus:
		// args[0] is "screen"
		// args[1] is "add-doctabs"
		// args[2] is the «screen-package-name»
		// args[3..] is «[*]tab-item-name, ...».
		if len(args) < 4 {
			fmt.Println(UsageScreen())
			return
		}
		// Validate the screen package name.
		if isValid, failureMessage := _utils_.ValidatePascalCase(args[2], "screen"); !isValid || err != nil {
			fmt.Println(failureMessage)
			return
		}
		if manifest.HasScreen(args[2]) {
			fmt.Printf("Failure: The screen package %q already exists.\n", args[2])
			return
		}
		// Validate item names.
		var haveInvalidItemName bool
		for _, itemName := range args[3:] {
			isRemote, isScreen := manifest.IsRemoteNameIsScreenName(itemName)
			if isRemote && !isScreen {
				haveInvalidItemName = true
				fmt.Printf("Failure: The TabItem name %q does not reference an existing screen package.\n", itemName)
			}
		}
		if haveInvalidItemName {
			return
		}
		if isValid, failureMessage := _utils_.ValidateTabNames(args[3:]); !isValid {
			fmt.Println(failureMessage)
			return
		}
		// Build the new screen.
		screenDocComment := fmt.Sprintf("Package %s is a DocTabs screen package.\nA tabbar screen where a tab displays it's panel or another screen.\nThe first tab is a configuration tab and allows the user to configure where the tabbar is located.\nKICKFYNE TODO: Complete this package doc commment.", args[2])
		itemNames := make([]string, len(args[3:])+1)
		itemNames[0] = _utils_.ConfigTabName
		for i, itemName := range args[3:] {
			itemNames[i+1] = itemName
		}
		if err = handleScreenAddDocTabs(args[2], itemNames, screenDocComment, importPrefix, folderPaths); err != nil {
			return
		}
		if err = _frontend_.RebuildFrontendGo(importPrefix, folderPaths); err != nil {
			return
		}
		// Update the manifest.
		manifest.AddScreen(args[2], _manifest_.DocTabsScreen, itemNames...)
		err = manifest.Write(folderPaths)
		return
	case verbRemove:
		// args[0] is "screen"
		// args[1] is "remove"
		// args[2] is the «screen-package-name»
		if len(args) != 3 {
			fmt.Println(UsageScreen())
			return
		}
		if !manifest.HasScreen(args[2]) {
			fmt.Printf("Failure: Screen package %q was not found.\n", args[2])
			return
		}
		if count, screenNames := manifest.CountScreenReferences(args[2]); count > 0 {
			fmt.Printf("Failure: Screen package %q is referenced by %s.\n", args[2], screenNames)
			return
		}
		if err = handleScreenRemove(args[2], folderPaths); err != nil {
			return
		}
		err = _frontend_.RebuildFrontendGo(importPrefix, folderPaths)
		// Update the manifest.
		manifest.RemoveScreen(args[2])
		err = manifest.Write(folderPaths)
		return
	case subCmdHelp:
		// args[0] is "screen"
		// args[1] is "help"
		fmt.Println(UsageScreen())
	default:
		// args[0] is "screen"
		fmt.Println(UsageScreen())
	}
	return
}

// handleScreenList handles the listing of the screen packages.
func handleScreenList(manifest _manifest_.Manifest) {

	// Display the list.
	fmt.Printf("List of %d screen packages.\n", manifest.CountScreens())
	fmt.Println(manifest.String())
}

func separatePanelNames(panelNames []string) (all []string, local []string, remote []string) {
	all = make([]string, 0, len(panelNames))
	local = make([]string, 0, len(panelNames))
	remote = make([]string, 0, len(panelNames))
	for _, panelName := range panelNames {
		if panelName[0] == '*' {
			fixed := panelName[1:]
			remote = append(remote, fixed)
			all = append(all, fixed)
		} else {
			local = append(local, panelName)
			all = append(all, panelName)
		}
	}
	return
}
