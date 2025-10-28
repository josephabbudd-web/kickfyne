package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	FolderNameAccordionItems      = "accordionItems"
	FolderNameContent             = "content"
	FolderNameDeps                = "deps"
	folderNameFrontend            = "frontend"
	FolderNameLandingScreen       = "landingscreen"
	FolderNameLayout              = "layout"
	FolderNameLayoutTabItems      = "layoutTabItems"
	FolderNameLayoutAccorionItems = "layoutAccordionItems"
	folderNameMainMenu            = "mainmenu"
	folderNameMessage             = "message"
	FolderNameMisc                = "misc"
	FolderNamePanelers            = "panelers"
	FolderNamePanels              = "panels"
	FolderNamePresetting          = "presetting"
	FolderNameProducer            = "producer"
	FolderNameScreens             = "screens"
	folderNameScreenMap           = "screenmap"
	folderNameDeps                = "deps"
	// FolderNamePresets        = "presets"
	FolderNameTabItems = "tabItems"
	folderNameThread   = "thread"
	folderNameTypes    = "types"
	FolderNameDocTabs  = "doctabs"
)

var (
	frontendMainMenu  = filepath.Join(folderNameFrontend, folderNameMainMenu)
	frontendScreens   = filepath.Join(folderNameFrontend, FolderNameScreens)
	frontendScreenMap = filepath.Join(folderNameFrontend, folderNameScreenMap)
	frontendTypes     = filepath.Join(folderNameFrontend, folderNameTypes)

	depsMetaData = filepath.Join(folderNameDeps, "metadata")
	depsPaths    = filepath.Join(folderNameDeps, "paths")
	depsThread   = filepath.Join(folderNameDeps, folderNameThread)
)

type FolderPaths struct {
	App string

	Frontend                           string
	FrontendMainMenu                   string
	FrontendScreens, FrontendScreenMap string
	FrontendTypes                      string

	Deps         string
	DepsMetaData string
	DepsPaths    string
	DepsThread   string
}

// NewFolderPaths constructs paths and then makes them on the disk.
func NewFolderPaths(rootPath string) (folderPaths *FolderPaths, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("utils.NewFolderPaths: %w", err)
		}
	}()

	folderPaths = &FolderPaths{
		App: rootPath,

		Frontend:          filepath.Join(rootPath, folderNameFrontend),
		FrontendMainMenu:  filepath.Join(rootPath, frontendMainMenu),
		FrontendScreens:   filepath.Join(rootPath, frontendScreens),
		FrontendScreenMap: filepath.Join(rootPath, frontendScreenMap),
		FrontendTypes:     filepath.Join(rootPath, frontendTypes),

		Deps:         filepath.Join(rootPath, folderNameDeps),
		DepsMetaData: filepath.Join(rootPath, depsMetaData),
		DepsPaths:    filepath.Join(rootPath, depsPaths),
		DepsThread:   filepath.Join(rootPath, depsThread),
	}
	return
}

// IsBuilt returns the name of the app's folder.
func (folderPaths *FolderPaths) Base() (appFolderName string) {
	appFolderName = filepath.Base(folderPaths.App)
	return
}

// IsBuilt returns if the framework was built in this folder.
// It does so by checking for 2 folders.
func (folderPaths *FolderPaths) IsBuilt() (isBuilt bool) {
	isBuilt, _ = FolderHasFolders(
		folderPaths.App,
		folderNameFrontend, folderNameDeps,
	)
	return
}

// Build makes the folder paths on disk.
func (folderPaths *FolderPaths) Build() (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("FolderPaths.Build: %w", err)
		}
	}()

	err = folderPaths.buildFolderPaths()
	return
}

// Rebuild removes the frontend and deps folder and then builds.
// Useful for restarting the framework.
func (folderPaths *FolderPaths) Rebuild() (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("FolderPaths.Build: %w", err)
		}
	}()

	// Remove the folders.
	if err = os.RemoveAll(folderPaths.Frontend); err != nil {
		if os.IsNotExist(err) {
			err = nil
		} else {
			return
		}
	}
	if err = os.RemoveAll(folderPaths.Deps); err != nil {
		if os.IsNotExist(err) {
			err = nil
		} else {
			return
		}
	}

	err = folderPaths.buildFolderPaths()
	return
}

// buildFolderPaths constructs the paths onto the disk.
func (folderPaths *FolderPaths) buildFolderPaths() (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("utils.buildFolderPaths: %w", err)
		}
	}()

	if folderPaths.IsBuilt() {
		// The folders have already been created.
		return
	}

	// Create the folders.

	// Frontend.
	if err = os.Mkdir(folderPaths.Frontend, DMode); err != nil {
		return
	}
	if err = os.Mkdir(folderPaths.FrontendMainMenu, DMode); err != nil {
		return
	}
	if err = os.Mkdir(folderPaths.FrontendScreens, DMode); err != nil {
		return
	}
	if err = os.Mkdir(folderPaths.FrontendScreenMap, DMode); err != nil {
		return
	}
	if err = os.Mkdir(folderPaths.FrontendTypes, DMode); err != nil {
		return
	}

	// Deps
	if err = os.Mkdir(folderPaths.Deps, DMode); err != nil {
		return
	}
	if err = os.Mkdir(folderPaths.DepsMetaData, DMode); err != nil {
		return
	}
	if err = os.Mkdir(folderPaths.DepsPaths, DMode); err != nil {
		return
	}
	if err = os.Mkdir(folderPaths.DepsThread, DMode); err != nil {
		return
	}
	return

}
