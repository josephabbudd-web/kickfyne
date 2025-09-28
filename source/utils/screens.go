package utils

import (
	"fmt"
	"slices"
)

// ScreenPackageNames returns the names of the screen packages.
func ScreenPackageNames(folderPaths *FolderPaths) (screenNames []string, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("utils.ScreenNames: %w", err)
		}
	}()

	screenNames, err = FolderNames(folderPaths.FrontendScreens)
	return
}

// IsCurrentScreenName returns if the screenPackageName exists.
func IsCurrentScreenName(screenPackageName string, folderPaths *FolderPaths) (is bool, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("utils.IsCurrentScreenName: %w", err)
		}
	}()

	var packageNames []string
	if packageNames, err = ScreenPackageNames(folderPaths); err != nil {
		return
	}
	for _, packageName := range packageNames {
		if is = packageName == screenPackageName; is {
			break
		}
	}
	return
}

// ValidateNewScreenPackageName validates a new panel name for a screenName.
func ValidateNewScreenPackageName(
	screenPackageName string,
	folderPaths *FolderPaths,
) (isValid bool, failureMessage string, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("utils.ValidateNewScreenPackageName: %w", err)
		}
	}()

	// Valid. Not title case.
	if isValid, failureMessage = validateScreenName(screenPackageName); !isValid {
		return
	}
	// Unique. Not an existing screen package name.
	var currentPackageNames []string
	if currentPackageNames, err = ScreenPackageNames(folderPaths); err != nil {
		return
	}
	if slices.Contains(currentPackageNames, screenPackageName) {
		// Not a new screen package name.
		isValid = false
		failureMessage = fmt.Sprintf("The screen package name %q is not a new screen package name.", screenPackageName)
		return
	}
	// The screen package name is valid and unique.
	isValid = true
	return
}

// ValidateCurrentScreenPackageName validates a new panel name for a screenName.
func ValidateCurrentScreenPackageName(
	screenPackageName string,
	folderPaths *FolderPaths,
) (isValid bool, failureMessage string, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("utils.ValidateCurrentScreenPackageName: %w", err)
		}
	}()

	// Valid. Not title case.
	if isValid, failureMessage = validateScreenName(screenPackageName); !isValid {
		return
	}
	// Not Unique. Is an existing screen package name.
	var currentPackageNames []string
	if currentPackageNames, err = ScreenPackageNames(folderPaths); err != nil {
		return
	}
	if !slices.Contains(currentPackageNames, screenPackageName) {
		// Not an existing screen package name.
		failureMessage = fmt.Sprintf("The screen package name %q is not an existing screen package name.", screenPackageName)
		return
	}
	// The screen package name is valid and is an existing screen package name.
	isValid = true
	return
}
