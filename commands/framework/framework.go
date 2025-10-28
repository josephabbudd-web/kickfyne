package framework

import (
	"fmt"

	_manifest_ "github.com/josephabbudd-web/kickfyne/manifest"
	_source_ "github.com/josephabbudd-web/kickfyne/source"
	_utils_ "github.com/josephabbudd-web/kickfyne/source/utils"
)

const (
	Cmd         = "framework"
	verbHelp    = "help"
	verbRestart = "restart"
)

// Handler passes control to the correct handler.
func Handler(args []string, importPrefix string, folderPaths *_utils_.FolderPaths) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("framework.Handler: %w", err)
		}
	}()

	switch folderPaths.IsBuilt() {
	case true:
		// Because this framework is already build there needs to be 1 verb.
		if len(args) == 0 {
			fmt.Println(Usage())
			return
		}
		// The framework is built in this folder.
		// Check the verb.
		switch args[0] {
		case verbRestart:
			if err = handleFrameworkRestart(folderPaths, args, importPrefix); err != nil {
				return
			}
			// fyneAppTOMLFilePath := _utils_.FyneAppTOMLFilePath(folderPaths)
			// fmt.Printf("KICKFYNE TODO: Check MenuItems in %s.\n", _utils_.Clickable(fyneAppTOMLFilePath))
		case verbHelp:
			fmt.Println(Usage())
			return
		default:
			fmt.Println(Usage())
			return
		}
	case false:
		// The framework is not built in this folder.
		// There may or may not be a verb.
		if len(args) > 0 {
			fmt.Println(Usage())
		}
		if err = handleFramework(folderPaths, args, importPrefix); err != nil {
			return
		}
	}
	return
}

// handleFramework creates the framework.
func handleFramework(folderPaths *_utils_.FolderPaths, args []string, importPrefix string) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("framework.handleFramework: %w", err)
		}
	}()

	_ = args // Add optional backend signal.

	// Create the framework code.
	if err = folderPaths.Build(); err != nil {
		return
	}
	if err = _source_.CreateFramework(importPrefix, folderPaths); err != nil {
		return
	}

	fmt.Println("Success. The framework is created.")
	return
}

// handleFrameworkRestart creates the framework.
func handleFrameworkRestart(folderPaths *_utils_.FolderPaths, args []string, importPrefix string) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("framework.handleFrameworkRestart: %w", err)
		}
	}()

	// Reset the manifest.
	var manifest _manifest_.Manifest
	if manifest, err = _manifest_.New(folderPaths); err != nil {
		return
	}
	manifest.Reset()
	err = manifest.Write(folderPaths)

	_ = args
	// Create the framework code.
	if err = folderPaths.Rebuild(); err != nil {
		return
	}
	if err = _source_.CreateFramework(importPrefix, folderPaths); err != nil {
		return
	}

	fmt.Println("Success. The framework is recreated.")
	return
}
