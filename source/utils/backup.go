package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	folderBackupSuffix = "_backup"
	fileBackupExt      = ".bak"
)

func BackupFile(originalPath string) (err error) {
	var backupPath string
	if backupPath, err = BackupPath(originalPath); err != nil {
		return
	}
	err = os.Rename(originalPath, backupPath)
	return
}

func UnBackupFile(originalPath string) (err error) {
	var backupPath string
	if backupPath, err = BackupPath(originalPath); err != nil {
		return
	}
	err = os.Rename(backupPath, originalPath)
	return
}

func BackupPath(path string) (backupPath string, err error) {
	dirPath := filepath.Dir(path)
	if path == dirPath {
		// This is a folder path
		err = fmt.Errorf("BackupPath path is not a file path")
		return
	}
	// This is a file path.
	base := filepath.Base(path)
	newBase := strings.ReplaceAll(base, ".", "_") + ".bak"
	backupPath = filepath.Join(dirPath, newBase)
	return
}

func UnBackupPath(backupPath string) (originalPath string) {
	if strings.HasSuffix(backupPath, folderBackupSuffix) {
		// backupPath is a folder path.
		originalPath = strings.Replace(backupPath, folderBackupSuffix, "", 1)
		return
	}
	if strings.HasSuffix(backupPath, fileBackupExt) {
		// backupPath is a file path.
		originalPath = strings.Replace(backupPath, fileBackupExt, "", 1)
		originalPath = strings.Replace(originalPath, "_go", ".go", 1)
		return
	}
	// No a backup path.
	return
}
