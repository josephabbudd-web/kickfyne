package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/goccy/go-yaml"

	_utils_ "github.com/josephabbudd-web/kickfyne/source/utils"
)

type ScreenKind string

const (
	NilScreen       ScreenKind = ""
	SimpleScreen    ScreenKind = "Simple"
	AccordionScreen ScreenKind = "Accordion"
	AppTabsScreen   ScreenKind = "AppTabs"
	DocTabsScreen   ScreenKind = "DocTabs"
)

type Info struct {
	Kind  ScreenKind
	Items []string
}

func (info *Info) Copy() (infoCopy *Info) {
	infoCopy = &Info{
		Kind:  info.Kind,
		Items: make([]string, len(info.Items)),
	}
	copy(infoCopy.Items, info.Items)
	return
}

func (info *Info) String() (str string) {
	items := make([]string, len(info.Items))
	copy(items, info.Items)
	str = fmt.Sprintf("%s: %s", info.Kind, strings.Join(items, ", "))
	return
}

func (info *Info) Count() (count int) {
	count = len(info.Items)
	return
}

func (info *Info) Has(itemName string) (has bool) {
	has = slices.Contains(info.Items, itemName)
	return
}

func (info *Info) Add(itemNames ...string) {
	fmt.Printf("> info.Items is %#v\n", info.Items)
	info.Items = append(info.Items, itemNames...)
	fmt.Printf("< info.Items is %#v\n", info.Items)
}

func (info *Info) GetItems() (all, local, remote []string) {
	all = make([]string, len(info.Items))
	copy(all, info.Items)
	local = make([]string, 0, len(info.Items))
	remote = make([]string, 0, len(info.Items))
	for _, item := range info.Items {
		if item[:1] == "*" {
			remote = append(remote, item[1:])
		} else {
			local = append(local, item)
		}
	}
	return
}

func (info *Info) Remove(itemNames ...string) {
	for _, itemName := range itemNames {
		at := slices.Index(info.Items, itemName)
		if at >= 0 {
			info.Items = slices.Delete(info.Items, at, at+1)
		}
	}
}

func newItems() (info *Info) {
	info = &Info{
		Kind:  "",
		Items: make([]string, 0, 10),
	}
	return
}

// Manifest[screen-name]Info
type Manifest map[string]*Info

var _manifest Manifest = nil

// NewAgain returns the *Manifest.
func NewAgain() (manifest Manifest) {
	manifest = _manifest
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
		ss = append(ss, fmt.Sprintf("%s: %s", screenName, info.String()))
	}
	str = strings.Join(ss, "\n")
	return
}

// SetNewInfo sets the manifest's info.
func (manifest Manifest) SetNewInfo(screenName string, newInfo *Info) {
	manifest[screenName] = newInfo
}

// Reset resets the manifest so that there are no more screens.
func (manifest Manifest) Reset() {
	for k := range manifest {
		delete(manifest, k)
	}
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
		if info.Has(itemName) {
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

// HasScreen returns if the manifest has the screen.
func (manifest Manifest) HasScreen(screenName string) (hasScreenName bool) {
	hasScreenName = manifest[screenName] != nil
	return
}

// ScreenKind returns the screen's kind.
func (manifest Manifest) ScreenKind(screenName string) (screenKind ScreenKind) {
	if info := manifest[screenName]; info != nil {
		screenKind = info.Kind
	} else {
		screenKind = NilScreen
	}
	return
}

// IsHasScreenName returns if the manifest has the screen.
func (manifest Manifest) IsRemoteNameIsScreenName(itemName string) (isRemoteName, isScreenName bool) {
	if itemName[:1] == "*" {
		isRemoteName = true
		isScreenName = manifest[itemName[1:]] != nil
	}
	return
}

// HasScreenItem returns if the screen has the item.
// Param itemName is "[*]ItemName"
func (manifest Manifest) HasScreenItem(screenName string, itemName string) (hasItemName bool) {
	info := manifest[screenName]
	if info == nil {
		return
	}
	hasItemName = info.Has(itemName)
	return
}

// AddScreen adds a screen and it's info.
// Param itemNames is [] "[*]ItemName"
func (manifest Manifest) AddScreen(screenName string, screenKind ScreenKind, itemNames ...string) {
	info := newItems()
	info.Kind = screenKind
	info.Add(itemNames...)
	manifest[screenName] = info
}

// AddScreen adds a screen and it's info.
// Param itemNames is [] "[*]ItemName"
func (manifest Manifest) RemoveScreen(screenName string) {
	delete(manifest, screenName)
}

func (manifest Manifest) CountScreens() (count int) {
	count = len(manifest)
	return
}

func (manifest Manifest) CountScreenItems(screenName string) (count int) {
	var info *Info
	var found bool
	if info, found = manifest[screenName]; !found {
		return
	}
	count = info.Count()
	return
}

func (manifest Manifest) Kind(screenName string) (kind ScreenKind) {
	var info *Info
	var found bool
	if info, found = manifest[screenName]; !found {
		kind = NilScreen
		return
	}
	kind = info.Kind
	return
}

// Items returns the local and remote item names.
func (manifest Manifest) Items(screenName string) (local, remote []string) {
	var info *Info
	if info = manifest[screenName]; info == nil {
		return
	}
	local = make([]string, len(info.Items))
	remote = make([]string, len(info.Items))
	for _, itemName := range info.Items {
		if itemName[:1] == "*" {
			remote = append(remote, itemName[1:])
		} else {
			local = append(local, itemName)
		}
	}
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
	info.Add(itemNames...)
}

// Remove removes a local or remote item name.
func (manifest Manifest) RemoveItems(screenName string, itemNames ...string) {
	var info *Info
	if info = manifest[screenName]; info == nil {
		return
	}
	// Remove the item names.
	info.Remove(itemNames...)
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
