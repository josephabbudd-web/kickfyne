package container

import (
	"fmt"
	"os"
	"path/filepath"

	_utils_ "github.com/josephabbudd-web/kickfyne/source/utils"
)

func Remove(
	packageName string,
	folderPaths *_utils_.FolderPaths,
) {
	path := filepath.Join(folderPaths.Deps, packageName)
	_ = os.RemoveAll(path)
}

func CreateFramework(
	folderPaths *_utils_.FolderPaths,
) (err error) {
	err = _utils_.WriteGitKeepFile(folderPaths.DepsContainer)
	return
}

func AddAccordion(
	packageName string,
	allPanelNames []string,
	folderPaths *_utils_.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("deps.AddAccordion: %w", err)
		}
	}()

	var fileName string
	var path string
	var data any

	// Accordion IDs
	fileName = _utils_.ContainerFileName(packageName)
	path = filepath.Join(folderPaths.DepsContainer, packageName)
	if err = os.Mkdir(path, _utils_.DMode); err != nil {
		return
	}
	path = filepath.Join(path, fileName)
	data = &accordionTemplateData{
		PackageName:   packageName,
		AllPanelNames: cleanPanelNames(allPanelNames),
		Funcs:         _utils_.GetFuncs(),
	}
	err = _utils_.ProcessTemplate(fileName, path, accordionTemplate, data)
	return
}

func AddTabbar(
	packageName string,
	allPanelNames []string,
	folderPaths *_utils_.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("deps.AddTabbar: %w", err)
		}
	}()

	var fileName string
	var path string
	var data any

	// Tabbar IDs
	fileName = _utils_.ContainerFileName(packageName)
	path = filepath.Join(folderPaths.DepsContainer, packageName)
	if err = os.Mkdir(path, _utils_.DMode); err != nil {
		return
	}
	path = filepath.Join(path, fileName)
	data = &tabbarTemplateData{
		PackageName:   packageName,
		AllPanelNames: cleanPanelNames(allPanelNames),
		Funcs:         _utils_.GetFuncs(),
	}
	err = _utils_.ProcessTemplate(fileName, path, tabbarTemplate, data)
	return
}

func cleanPanelNames(panelNames []string) (cleanPanelNames []string) {
	l := len(panelNames)
	cleanPanelNames = make([]string, l, l)
	for i, panelName := range panelNames {
		if panelName[0] == '*' {
			cleanPanelNames[i] = panelName[1:]
		} else {
			cleanPanelNames[i] = panelName
		}
	}
	return
}
