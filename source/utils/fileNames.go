package utils

import (
	"path"
	"path/filepath"
)

const (
	APIFileName           = "api.go"
	ContentFileName       = "content.go"
	StateFileName         = "state.go"
	panelFileSuffix       = "Panel.go"
	GoFileExt             = ".go"
	FyneAppTOMLFileName   = "FyneApp.toml"
	ScreenFileName        = "screen.go"
	DocFileName           = "doc.go"
	ExampleFileName       = "example.go"
	LayoutFileName        = "layout.go"
	PanelsFileName        = "panels.go"
	PanelsStateFileName   = "panelsState.go"
	ComponentsFileName    = "components.go"
	KeepFileName          = ".gitkeep"
	IDsFileName           = "ids.go"
	PresetFileName        = "preset.go"
	RemotePresetFileName  = "remotePreset.go"
	DefaultPresetFileName = "defaultPreset.go"
	ManifestFileName      = "manifest.yaml"

	ralativeFilePathSuffix = ":1:1"
)

func Clickable(path string) (clickable string) {
	clickable = path + ralativeFilePathSuffix
	return
}

// FyneAppTOMLFilePath
func FyneAppTOMLFilePath(folderPaths *FolderPaths) (metaDataTOMLFilePath string) {
	metaDataTOMLFilePath = filepath.Join(folderPaths.App, FyneAppTOMLFileName)
	return
}

// PanelFileName returns the file name for a panel file.
func PanelFileName(panelName string) (fileName string) {
	fileName = panelName + GoFileExt
	return
}

// PanelContentFolderName returns the content folder name for a panel.
func PanelContentFolderName(panelName string) (fileName string) {
	fileName = panelName + "Panel"
	return
}

// PanelContentFilePath returns the relative path for a panel content file.
func PanelContentFilePath(screenPackageName, panelName string, folderPaths *FolderPaths) (filePath string) {
	contentFolderName := PanelContentFolderName(panelName)
	filePath = path.Join(folderPaths.FrontendScreens, screenPackageName, FolderNamePanels, contentFolderName, ContentFileName)
	return
}

// PanelStateFilePath returns the relative path for a panel's content state file.
func PanelStateFilePath(screenPackageName, panelName string, folderPaths *FolderPaths) (filePath string) {
	contentFolderName := PanelContentFolderName(panelName)
	filePath = path.Join(folderPaths.FrontendScreens, screenPackageName, FolderNamePanels, contentFolderName, StateFileName)
	return
}

// PanelPresetFilePath returns the relative path for a panel's content state file.
func PanelPresetFilePath(screenPackageName, panelName string, folderPaths *FolderPaths) (filePath string) {
	contentFolderName := PanelContentFolderName(panelName)
	filePath = path.Join(folderPaths.FrontendScreens, screenPackageName, FolderNamePanels, contentFolderName, PresetFileName)
	return
}

// ScreenFileRelativeFilePath returns the relative path for a screen's screen.go file.
func ScreenFileRelativeFilePath(screenPackageName string) (relativeFilePath string) {
	relativeFilePath = path.Join(folderNameFrontend, FolderNameScreens, screenPackageName, ScreenFileName+ralativeFilePathSuffix)
	return
}

// DocFileRelativeFilePath returns the relative path for a screen's screen.go file.
func DocFileRelativeFilePath(screenPackageName string, folderPaths *FolderPaths) (relativeFilePath string) {
	relativeFilePath = path.Join(folderPaths.FrontendScreens, screenPackageName, DocFileName)
	return
}

// ScreenDocFilePath returns the relative path for a screen's doc.go file.
func ScreenDocFilePath(screenPackageName string, folderPaths *FolderPaths) (relativeFilePath string) {
	relativeFilePath = path.Join(folderPaths.FrontendScreens, screenPackageName, DocFileName)
	return
}

// ScreenPresettingAPIFilePath returns the relative path for a screen's doc.go file.
func ScreenPresettingAPIFilePath(screenPackageName string, folderPaths *FolderPaths) (relativeFilePath string) {
	relativeFilePath = path.Join(folderPaths.FrontendScreens, screenPackageName, FolderNamePresetting, APIFileName)
	return
}

// LayoutFileRelativeFilePath returns the relative path for a screen's layout.go file.
func LayoutFileRelativeFilePath(screenPackageName string) (relativeFilePath string) {
	relativeFilePath = path.Join(folderNameFrontend, FolderNameScreens, screenPackageName, LayoutFileName+ralativeFilePathSuffix)
	return
}

func ContainerFileName(screenPakcageName string) (fileName string) {
	fileName = screenPakcageName + GoFileExt
	return
}

// MessageFileName returns the file name for a messsage.
func MessageFileName(messageName string) (fileName string) {
	fileName = DeCap(messageName) + GoFileExt
	return
}

// MessageFileRelativeFilePath returns the relative path for a message file.
func MessageFileRelativeFilePath(messageName string) (relativeFilePath string) {
	fName := MessageFileName(messageName) + ralativeFilePathSuffix
	relativeFilePath = path.Join(folderNameDeps, folderNameMessage, fName)
	return
}

// RecordFileName returns the file name for a record.
func RecordFileName(recordName string) (fileName string) {
	fileName = DeCap(recordName) + GoFileExt
	return
}
