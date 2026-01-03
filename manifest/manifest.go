package manifest

import (
	"fmt"
	"log"
	"maps"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/goccy/go-yaml"

	_utils_ "github.com/josephabbudd-web/kickfyne/source/utils"
)

// Manifest[screen-name]Info
type Manifest map[string]*Info

var _manifest Manifest = nil

const (
	Framework = "_Framework_"
)

// NewAgain returns the *Manifest.
func NewAgain() (manifest Manifest) {
	manifest = _manifest
	return
}

func (manifest Manifest) Copy() (manifestCopy Manifest) {
	manifestCopy = make(map[string]*Info, len(manifest))
	maps.Copy(manifestCopy, manifest)
	return
}

// New constructs a new manifest from the screen packages manifest.yaml file.
func New(folderPaths *_utils_.FolderPaths) (manifest Manifest, err error) {

	if _manifest != nil {
		manifest = _manifest
		return
	}

	defer func() {
		if err != nil {
			err = fmt.Errorf("Manifest.Read: %w", err)
		}
	}()

	path := filepath.Join(folderPaths.App, _utils_.ManifestFileName)
	var contents []byte
	if contents, err = os.ReadFile(path); err != nil {
		if os.IsNotExist(err) {
			err = nil
			_manifest = make(map[string]*Info)
			_manifest[Framework] = newEmptyFrameworkInfo()
			manifest = _manifest
		}
		return
	}
	if err = yaml.Unmarshal(contents, &_manifest); err != nil {
		return
	}
	manifest = _manifest
	return
}

func (manifest Manifest) ScreenNames() (screenNames []string) {
	screenNames = make([]string, 0, len(manifest))
	for screenName := range manifest {
		if screenName != Framework {
			screenNames = append(screenNames, screenName)
		}
	}
	return
}

func (manifest Manifest) InfoCopy(screenName string) (infoCopy *Info) {
	var info *Info
	if info = manifest[screenName]; info == nil {
		return
	}
	infoCopy = info.Copy()
	return
}

func (manifest Manifest) String() (str string) {
	ss := make([]string, 0, len(manifest))
	for screenName, info := range manifest {
		if screenName != Framework {
			ss = append(ss, fmt.Sprintf("%s: %s", screenName, info.String()))
		}
	}
	str = strings.Join(ss, "\n")
	return
}

// Reset resets the manifest so that there are no more screens.
func (manifest Manifest) Reset() {
	for k := range manifest {
		delete(manifest, k)
	}
	manifest[Framework] = newEmptyFrameworkInfo()

}

// CountScreenReferences returns if the any screen is using this item.
func (manifest Manifest) CountScreenReferences(referenceScreenName string) (count int, screenNames string) {
	var itemName string
	builder := make([]string, 0, len(manifest))

	if referenceScreenName[:1] == "*" {
		itemName = referenceScreenName
	} else {
		itemName = "*" + referenceScreenName
	}
	for screenName, info := range manifest {
		if info.HasItem(itemName) {
			builder = append(builder, screenName)
		}
	}

	count = len(builder)
	switch {
	case count == 1:
		screenNames = "the screen " + builder[0]
	case count == 2:
		screenNames = "the screens " + strings.Join(builder, " and ")
	case count > 2:
		last := count - 1
		screenNames = "the screens " + strings.Join(builder[:last], ", ") + " and " + builder[last]
	}
	return
}

// HasScreen returns if the manifest has the screen and it has not been removed.
func (manifest Manifest) HasScreen(screenName string) (hasScreenName bool) {
	info := manifest[screenName]
	hasScreenName = (info != nil)
	return
}

// InfoKind returns the screen's kind.
func (manifest Manifest) InfoKind(screenName string) (screenKind InfoKind) {
	if info := manifest[screenName]; info != nil {
		screenKind = info.Kind
	} else {
		screenKind = NilInfoKind
	}
	return
}

// IsRemoteNameIsScreenName returns if the manifest has the screen.
func (manifest Manifest) IsRemoteNameIsScreenName(itemName string) (isRemoteName, isScreenName bool) {
	if !strings.Contains(itemName, "=") {
		// *Edit
		if len(itemName) <= 1 {
			// "*" is not valid syntax.
			return
		}
		if itemName[:1] == "*" {
			isRemoteName = true
			screenName := itemName[1:]
			isScreenName = manifest[screenName] != nil
			return
		}
		// Simple local item name.
		return
	}
	isValidSyntax, _, screenName := manifest.ParseBorderAreaParam(itemName)
	if !isValidSyntax {
		return
	}
	isRemoteName = len(screenName) > 0
	isScreenName = manifest[screenName] != nil
	return
}

// IsRemoteNameIsScreenName returns if the manifest has the screen.
func (manifest Manifest) ParseBorderAreaParam(paramName string) (isValidSyntax bool, areaName, screenName string) {
	parts := strings.Split(paramName, "=")
	// No "=" is ok.
	if len(parts) == 1 {
		// The syntax is valid.
		isValidSyntax = true
		areaName = parts[0]
		return
	}
	// Invalid syntaxes.
	if len(parts) != 2 {
		return
	}
	if len(parts[0]) == 0 {
		return
	}
	if len(parts[1]) <= 1 {
		return
	}
	if parts[1][:1] != "*" {
		return
	}
	// The syntax is valid.
	isValidSyntax = true
	areaName = parts[0]
	screenName = parts[1][1:]
	return
}

// HasScreenItem returns if the screen has the item.
// Param itemName is "[*]ItemName"
func (manifest Manifest) HasScreenItem(screenName string, itemName string) (cleanItemName string, hasItemName bool) {
	info := manifest[screenName]
	if info == nil {
		log.Printf("HasScreenItem: screenName %q is not a valid screen name.", screenName)
		return
	}
	cleanNames := info.cleanItemNames()
	log.Printf("HasScreenItem: cleanNames is %#v", cleanNames)

	if hasItemName = slices.Contains(cleanNames, itemName); hasItemName {
		// Ex: "Edit", "Center", "Center=*Edit"
		cleanItemName = itemName
		log.Printf("HasScreenItem: 1 screenName %q has item %q.", screenName, itemName)
		return
	}
	log.Printf("HasScreenItem: 1 screenName %q does not have item %q.", screenName, itemName)

	if info.Kind == BorderScreenInfoKind {
		// Border area syntax.
		// area-name=*screen-name
		parts := strings.Split(itemName, "=")
		if len(parts) == 2 {
			if hasItemName = slices.Contains(cleanNames, parts[0]); hasItemName {
				cleanItemName = parts[0]
				log.Printf("HasScreenItem: 2 screenName %q has item %q.", screenName, cleanItemName)
			}
		}
	}
	log.Printf("HasScreenItem: 2 screenName %q does not have item %q.", screenName, itemName)

	return
}

// AddFramework adds a framework and it's info.
func (manifest Manifest) AddFramework() {
	logItem := newFrameworkLogItem()
	manifest[Framework].AddLogItem(logItem)
}

// AddScreen adds a screen and it's info.
// Param itemNames is [] "[*]ItemName"
// See Manifest.HasScreen(screenName string)
func (manifest Manifest) AddScreen(screenName string, screenKind InfoKind, itemNames ...string) {
	manifest[screenName] = newScreenInfo(screenName, screenKind, itemNames)
	// Update the framework log.
	logItem := newAddScreenLogItem(screenName, screenKind, itemNames)
	manifest[Framework].AddLogItem(logItem)
}

// AddScreen adds a screen and it's info.
// Param itemNames is [] "[*]ItemName"
func (manifest Manifest) RemoveScreen(screenName string) {
	var info *Info
	var found bool
	if info, found = manifest[screenName]; !found {
		return
	}
	delete(manifest, screenName)
	logItem := newFrameworkRemoveScreenLogItem(screenName, info.Kind, info.Items)
	manifest[Framework].AddLogItem(logItem)
}

func (manifest Manifest) CountScreens() (count int) {
	for screenName := range manifest {
		if screenName != Framework {
			count++
		}
	}
	return
}

func (manifest Manifest) CountScreenItems(screenName string) (count int) {
	var info *Info
	var found bool
	if info, found = manifest[screenName]; !found {
		return
	}
	count = info.CountItems()
	return
}

func (manifest Manifest) Kind(screenName string) (kind InfoKind) {
	var info *Info
	var found bool
	if info, found = manifest[screenName]; !found {
		kind = NilInfoKind
		return
	}
	kind = info.Kind
	return
}

// Items returns the local and remote item names.
func (manifest Manifest) Items(screenName string) (all, local, remote []string) {
	var info *Info
	if info = manifest[screenName]; info == nil {
		return
	}
	all, local, remote = info.GetItems()
	return
}

// AddItems adds info to a screen.
// Param itemNames is [] "[*]ItemName"
func (manifest Manifest) AddItems(screenName string, itemNames ...string) {
	var info *Info
	if info = manifest[screenName]; info == nil {
		return
	}
	// Add the item names.
	info.AddItems(itemNames...)
}

// AddItemsLogAction adds info to a screen.
// Param itemNames is [] "[*]ItemName"
func (manifest Manifest) AddItemsLogAction(screenName string, itemNames ...string) {
	var info *Info
	if info = manifest[screenName]; info == nil {
		return
	}
	// Add the item names.
	info.AddItems(itemNames...)
	// Log this action.
	logItem := newAddItemsLogItem(screenName, itemNames)
	info.AddLogItem(logItem)
}

// RemoveItems removes a local or remote item name.
func (manifest Manifest) RemoveItems(screenName string, itemNames ...string) {
	var info *Info
	if info = manifest[screenName]; info == nil {
		return
	}
	// Remove the item names.
	info.RemoveItems(itemNames...)
}

// RemoveItemsLogAction removes a local or remote item name.
func (manifest Manifest) RemoveItemsLogAction(screenName string, itemNames ...string) {
	var info *Info
	if info = manifest[screenName]; info == nil {
		return
	}
	// Remove the item names.
	info.RemoveItems(itemNames...)
	// Log this action.
	logItem := newRemoveItemsLogItem(screenName, itemNames)
	info.AddLogItem(logItem)
}

// Write writes the manifest to the screen package's manifest.yaml file.
func (manifest Manifest) Write(folderPaths *_utils_.FolderPaths) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("Manifest.Write: %w", err)
		}
	}()

	path := filepath.Join(folderPaths.App, _utils_.ManifestFileName)
	var contents []byte
	if contents, err = yaml.Marshal(manifest); err != nil {
		return
	}
	err = _utils_.WriteFile(path, contents)
	return
}

func (manifest Manifest) LogAddItems(screenPackageName string, items []string) {

}

func (manifest Manifest) LogRemoveItems(screenPackageName string, items []string) {

}

func (manifest Manifest) LastLogMesssage(screenPackageName string, folderPaths *_utils_.FolderPaths) (successMessage string) {
	var screenInfo *Info
	if screenInfo = manifest.InfoCopy(screenPackageName); screenInfo == nil {
		return
	}
	lastLog := screenInfo.LastLogItem()
	builder := strings.Builder{}

	switch lastLog.Kind() {
	case LogItemKindCreateScreen:
		builder.WriteString(string(lastLog.Action))
		builder.WriteString(
			reviewDocs(screenPackageName, folderPaths),
		)
		builder.WriteString("If you want this package referenced in the main menu then you will want to add it.\n")
		builder.WriteString("mainmenu.go: ")
		mainMenuFilePath := _utils_.MainMenuFilePath(folderPaths)
		builder.WriteString(
			_utils_.Clickable(mainMenuFilePath),
		)
		builder.WriteString("\n")
	case LogItemKindAddItem:
		switch screenInfo.Kind {
		case SimpleScreenInfoKind:
			builder.WriteString(fmt.Sprintf("Added %d panels to the %s screen package named %q.\n", len(lastLog.Items), screenInfo.Kind, screenPackageName))
		case AppTabsScreenInfoKind, DocTabsScreenInfoKind:
			builder.WriteString(fmt.Sprintf("Added %d TabItems to the %s screen package named %q.\n", len(lastLog.Items), screenInfo.Kind, screenPackageName))
		case AccordionScreenInfoKind:
			builder.WriteString(fmt.Sprintf("Added %d AccordionItems to the %s screen package named %q.\n", len(lastLog.Items), screenInfo.Kind, screenPackageName))
		}
		builder.WriteString(
			reviewDocs(screenPackageName, folderPaths),
		)
		builder.WriteString(
			addRemoveItemSuccessMesssage(screenPackageName, folderPaths),
		)
	case LogItemKindRemoveItem:
		switch screenInfo.Kind {
		case SimpleScreenInfoKind:
			builder.WriteString(fmt.Sprintf("Removed %d panels from the %s screen package named %q.\n", len(lastLog.Items), screenInfo.Kind, screenPackageName))
		case AppTabsScreenInfoKind, DocTabsScreenInfoKind:
			builder.WriteString(fmt.Sprintf("Removed %d TabItems from the %s screen package named %q.\n", len(lastLog.Items), screenInfo.Kind, screenPackageName))
		case AccordionScreenInfoKind:
			builder.WriteString(fmt.Sprintf("Removed %d AccordionItems from the %s screen package named %q.\n", len(lastLog.Items), screenInfo.Kind, screenPackageName))
		}
		builder.WriteString(
			reviewDocs(screenPackageName, folderPaths),
		)
		builder.WriteString(
			addRemoveItemSuccessMesssage(screenPackageName, folderPaths),
		)
	}

	successMessage = builder.String()
	return
}

func (manifest Manifest) LastFrameworkLogMesssage() (successMessage string) {
	var frameworkInfo *Info
	if frameworkInfo = manifest.InfoCopy(Framework); frameworkInfo == nil {
		return
	}
	lastLog := frameworkInfo.LastLogItem()
	if lastLog.Action != ActionCreateFramework {
		return
	}
	successMessage = "Created the framework.\n"
	return
}

func reviewDocs(screenPackageName string, folderPaths *_utils_.FolderPaths) (message string) {
	docFilePath := _utils_.ScreenDocFileFullPath(screenPackageName, folderPaths)
	builder := strings.Builder{}
	builder.WriteString("Review the package's docs.go file.\n")
	builder.WriteString("Package docs: ")
	builder.WriteString(_utils_.Clickable(docFilePath))
	builder.WriteString("\n")
	message = builder.String()
	return
}

func addRemoveItemSuccessMesssage(screenPackageName string, folderPaths *_utils_.FolderPaths) (successMessage string) {
	builder := strings.Builder{}
	presettingAPIFilePath := _utils_.ScreenPresettingAPIFilePath(screenPackageName, folderPaths)
	docFilePath := _utils_.ScreenDocFileFullPath(screenPackageName, folderPaths)

	builder.WriteString("ATTENTION: These changes may have made breaking changes to the package's presettings.\n")
	builder.WriteString("Review: var Presets in " + _utils_.Clickable(presettingAPIFilePath) + "\n")
	builder.WriteString("The package's docs.go file has been updated.\n")
	builder.WriteString(_utils_.Clickable(docFilePath) + "\n")

	successMessage = builder.String()
	return
}
