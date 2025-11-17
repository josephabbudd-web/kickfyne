package manifest

import (
	"fmt"
	"strings"
	"time"
)

type Action string

const (
	ActionCreateFramework Action = "Created the framework."
	ActionAddScreen       Action = "Created the screen with items."
	ActionAddItem         Action = "Added items to the screen."
	ActionRemoveItem      Action = "Removed items from the screen."
)

type LogItemKind int

const (
	LogItemKindCreateFramework LogItemKind = iota
	LogItemKindCreateScreen
	LogItemKindRemoveScreen
	LogItemKindAddItem
	LogItemKindRemoveItem
)

func (logItem *LogItem) Kind() (kind LogItemKind) {
	switch {
	case logItem.Action == ActionCreateFramework:
		kind = LogItemKindCreateFramework
	case strings.HasPrefix(string(logItem.Action), "Created the "):
		kind = LogItemKindCreateScreen
	case strings.HasPrefix(string(logItem.Action), "Deleted the "):
		kind = LogItemKindRemoveScreen
	case strings.HasPrefix(string(logItem.Action), "Added the "):
		kind = LogItemKindAddItem
	case strings.HasPrefix(string(logItem.Action), "Removed the "):
		kind = LogItemKindRemoveItem
	}
	return
}

type LogItem struct {
	Date   string
	Action Action
	Items  []string
}

func newAddScreenLogItem(screenName string, screenKind InfoKind, itemNames []string) (logItem *LogItem) {
	nowbb, _ := time.Now().MarshalText()
	if len(itemNames) == 1 {
		logItem = &LogItem{
			Date:   string(nowbb),
			Action: Action(fmt.Sprintf("Created the %s screen named %s with the item named %s.", string(screenKind), screenName, itemNames[0])),
			Items:  itemNames,
		}
	} else {
		itemsLine := commaAnd(itemNames)
		logItem = &LogItem{
			Date:   string(nowbb),
			Action: Action(fmt.Sprintf("Created the %s screen named %s with the %d items named %s.", string(screenKind), screenName, len(itemNames), itemsLine)),
			Items:  itemNames,
		}
	}
	return
}

func newFrameworkRemoveScreenLogItem(screenName string, screenKind InfoKind, itemNames []string) (logItem *LogItem) {
	nowbb, _ := time.Now().MarshalText()
	var actionString string
	if len(itemNames) == 1 {
		actionString = fmt.Sprintf("Deleted the %s screen named %s with item named %s.", string(screenKind), screenName, itemNames[0])
	} else {
		itemsLine := commaAnd(itemNames)
		actionString = fmt.Sprintf("Deleted the %s screen named %s with %d items named %s.", string(screenKind), screenName, len(itemNames), itemsLine)
	}
	logItem = &LogItem{
		Date:   string(nowbb),
		Action: Action(actionString),
		Items:  itemNames,
	}
	return
}

func newFrameworkLogItem() (logItem *LogItem) {
	nowbb, _ := time.Now().MarshalText()
	logItem = &LogItem{
		Date:   string(nowbb),
		Action: ActionCreateFramework,
		Items:  nil,
	}
	return
}

func newAddItemsLogItem(screenName string, itemNames []string) (logItem *LogItem) {
	nowbb, _ := time.Now().MarshalText()
	var actionString string
	if len(itemNames) == 1 {
		actionString = fmt.Sprintf("Added the item named %s to the screen named %s.", itemNames[0], screenName)
	} else {
		itemsLine := commaAnd(itemNames)
		actionString = fmt.Sprintf("Added the %d items named %s to the screen named %s.", len(itemNames), itemsLine, screenName)
	}
	logItem = &LogItem{
		Date:   string(nowbb),
		Action: Action(actionString),
		Items:  itemNames,
	}
	return
}

func newRemoveItemsLogItem(screenName string, itemNames []string) (logItem *LogItem) {
	nowbb, _ := time.Now().MarshalText()
	var actionString string
	if len(itemNames) == 1 {
		actionString = fmt.Sprintf("Removed the item named %s from the screen named %s.", itemNames[0], screenName)
	} else {
		itemsLine := commaAnd(itemNames)
		actionString = fmt.Sprintf("Removed the %d items named %s from the screen named %s.", len(itemNames), itemsLine, screenName)
	}
	logItem = &LogItem{
		Date:   string(nowbb),
		Action: Action(actionString),
		Items:  itemNames,
	}
	return
}

func (logItem *LogItem) Copy() (logItemCopy *LogItem) {
	logItemCopy = &LogItem{
		Date:   logItem.Date,
		Action: logItem.Action,
		Items:  make([]string, len(logItem.Items)),
	}
	copy(logItemCopy.Items, logItem.Items)
	return
}

func commaAnd(list []string) (line string) {
	last := len(list) - 1
	if last < 0 {
		return
	}
	var builder strings.Builder
	builder.WriteString(strings.Join(list[:last], ", "))
	builder.WriteString(" and ")
	builder.WriteString(list[last])
	line = builder.String()
	return
}
