package manifest

import (
	"fmt"
	"slices"
	"strings"
)

type InfoKind string

const (
	NilInfoKind             InfoKind = ""
	FrameworkInfoKind       InfoKind = "Framework"
	SimpleScreenInfoKind    InfoKind = "Simple"
	AccordionScreenInfoKind InfoKind = "Accordion"
	AppTabsScreenInfoKind   InfoKind = "AppTabs"
	DocTabsScreenInfoKind   InfoKind = "DocTabs"
	BorderScreenInfoKind    InfoKind = "Border"
	SplitScreenInfoKind     InfoKind = "Split"
)

type Info struct {
	Kind  InfoKind
	Items []string
	Log   []*LogItem
}

func (info *Info) cleanItemNames() (cleanNames []string) {
	cleanNames = make([]string, len(info.Items))
	for i, item := range info.Items {
		parts := strings.Split(item, "=")
		cleanNames[i] = parts[0]
	}
	return
}

func newScreenInfo(screenName string, screenKind InfoKind, itemNames []string) (info *Info) {
	info = &Info{
		Kind:  screenKind,
		Items: make([]string, len(itemNames)),
		Log:   make([]*LogItem, 1, 10),
	}
	copy(info.Items, itemNames)
	info.Log[0] = newAddScreenLogItem(screenName, screenKind, itemNames)
	return
}

func newEmptyFrameworkInfo() (info *Info) {
	info = &Info{
		Kind:  FrameworkInfoKind,
		Items: nil,
		Log:   make([]*LogItem, 0, 5),
	}
	return
}

func (info *Info) Copy() (infoCopy *Info) {
	infoCopy = &Info{
		Kind:  info.Kind,
		Items: make([]string, len(info.Items)),
		Log:   make([]*LogItem, len(info.Log)),
	}
	copy(infoCopy.Items, info.Items)
	for i, logItem := range info.Log {
		infoCopy.Log[i] = logItem.Copy()
	}
	return
}

func (info *Info) String() (str string) {
	items := make([]string, len(info.Items))
	copy(items, info.Items)
	str = fmt.Sprintf("%s: %s", info.Kind, strings.Join(items, ", "))
	return
}

func (info *Info) AddLogItem(logItem *LogItem) {
	info.Log = slices.Insert(info.Log, 0, logItem)
}

func (info *Info) LastLogItem() (lastLogItem *LogItem) {
	if len(info.Log) == 0 {
		return
	}
	lastLogItem = info.Log[0].Copy()
	return
}

func (info *Info) CountItems() (count int) {
	count = len(info.Items)
	return
}

func (info *Info) HasItem(itemName string) (has bool) {
	has = info.itemIndex(itemName) >= 0
	return
}

func (info *Info) itemIndex(itemName string) (index int) {
	index = slices.Index(info.Items, itemName)
	return
}

func (info *Info) AddItems(itemNames ...string) {
	info.Items = append(info.Items, itemNames...)
}

func (info *Info) GetItems() (all, local, remote []string) {
	all = make([]string, len(info.Items))
	copy(all, info.Items)
	local = make([]string, 0, len(info.Items))
	remote = make([]string, 0, len(info.Items))
	for _, item := range info.Items {
		if info.Kind == BorderScreenInfoKind {
			parts := strings.Split(item, "=")
			if len(parts) == 2 {
				remote = append(remote, parts[1][1:])
			} else {
				local = append(local, item)
			}
		} else {
			if item[:1] == "*" {
				remote = append(remote, item[1:])
			} else {
				local = append(local, item)
			}
		}
	}
	return
}

func (info *Info) RemoveItems(itemNames ...string) {
	for _, itemName := range itemNames {
		at := info.itemIndex(itemName)
		if at >= 0 {
			info.Items = slices.Delete(info.Items, at, at+1)
		}
	}
}
