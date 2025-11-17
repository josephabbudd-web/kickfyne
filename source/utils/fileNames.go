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
	MainMenuFileName      = "mainmenu.go"

	ralativeFilePathSuffix = ":1:1"
)

func Clickable(path string) (clickable string) {
	clickable = path + ralativeFilePathSuffix
	return
}

// FyneAppTOMLFullFilePath
func FyneAppTOMLFullFilePath(folderPaths *FolderPaths) (metaDataTOMLFilePath string) {
	metaDataTOMLFilePath = filepath.Join(folderPaths.App, FyneAppTOMLFileName)
	return
}

// MainMenuFilePath
func MainMenuFilePath(folderPaths *FolderPaths) (mainMenuFilePath string) {
	mainMenuFilePath = filepath.Join(folderPaths.FrontendMainMenu, MainMenuFileName)
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

// PanelContentFullFilePath returns the relative path for a panel content file.
func PanelContentFullFilePath(screenPackageName, panelName string, folderPaths *FolderPaths) (filePath string) {
	contentFolderName := PanelContentFolderName(panelName)
	filePath = path.Join(folderPaths.FrontendScreens, screenPackageName, FolderNamePanels, contentFolderName, ContentFileName)
	return
}

// PanelStateFullFilePath returns the full path for a panel's content state file.
func PanelStateFullFilePath(screenPackageName, panelName string, folderPaths *FolderPaths) (filePath string) {
	contentFolderName := PanelContentFolderName(panelName)
	filePath = path.Join(folderPaths.FrontendScreens, screenPackageName, FolderNamePanels, contentFolderName, StateFileName)
	return
}

// PanelPresetFullFilePath returns the relative path for a panel's content state file.
func PanelPresetFullFilePath(screenPackageName, panelName string, folderPaths *FolderPaths) (filePath string) {
	contentFolderName := PanelContentFolderName(panelName)
	filePath = path.Join(folderPaths.FrontendScreens, screenPackageName, FolderNamePanels, contentFolderName, PresetFileName)
	return
}

// PanelContentFileRelativePath returns the relative path for a panel content file.
func PanelContentFileRelativePath(panelName string, folderPaths *FolderPaths) (filePath string) {
	contentFolderName := PanelContentFolderName(panelName)
	filePath = path.Join(FolderNamePanels, contentFolderName, ContentFileName)
	return
}

// PanelStateFileRelativePath returns the relative path for a panel's content state file.
func PanelStateFileRelativePath(panelName string, folderPaths *FolderPaths) (filePath string) {
	contentFolderName := PanelContentFolderName(panelName)
	filePath = path.Join(FolderNamePanels, contentFolderName, StateFileName)
	return
}

// PanelPresetRelativeFilePath returns the relative path for a panel's content state file.
func PanelPresetFileRelativePath(panelName string, folderPaths *FolderPaths) (filePath string) {
	contentFolderName := PanelContentFolderName(panelName)
	filePath = path.Join(FolderNamePanels, contentFolderName, PresetFileName)
	return
}

// ScreenFileRelativePath returns the relative path for a screen's screen.go file.
func ScreenFileRelativePath(screenPackageName string) (relativeFilePath string) {
	relativeFilePath = path.Join(folderNameFrontend, FolderNameScreens, screenPackageName, ScreenFileName+ralativeFilePathSuffix)
	return
}

// ScreenDocFileFullPath returns the relative path for a screen's doc.go file.
func ScreenDocFileFullPath(screenPackageName string, folderPaths *FolderPaths) (relativeFilePath string) {
	relativeFilePath = path.Join(folderPaths.FrontendScreens, screenPackageName, DocFileName)
	return
}

// ScreenPresettingAPIFileRelativePath returns the relative path for a screen's doc.go file.
func ScreenPresettingAPIFileRelativePath() (relativeFilePath string) {
	relativeFilePath = path.Join(FolderNamePresetting, APIFileName)
	return
}

// ScreenPresettingAPIFilePath returns the relative path for a screen's doc.go file.
func ScreenPresettingAPIFilePath(screenPackageName string, folderPaths *FolderPaths) (relativeFilePath string) {
	relativeFilePath = path.Join(folderPaths.FrontendScreens, screenPackageName, FolderNamePresetting, APIFileName)
	return
}

// ScreenPresettingDefaultPresetFileFullPath returns the full path for a screen's doc.go file.
func ScreenPresettingDefaultPresetFileFullPath(screenPackageName string, folderPaths *FolderPaths) (relativeFilePath string) {
	relativeFilePath = path.Join(folderPaths.FrontendScreens, screenPackageName, FolderNamePresetting, DefaultPresetFileName)
	return
}

// ScreenPresettingDefaultPresetFileRelativePath returns the relative path for a screen's doc.go file.
func ScreenPresettingDefaultPresetFileRelativePath() (relativeFilePath string) {
	relativeFilePath = path.Join(FolderNamePresetting, DefaultPresetFileName)
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
