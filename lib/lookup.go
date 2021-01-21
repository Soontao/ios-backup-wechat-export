package lib

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
)

const macosBackupDir = "Library/Application Support/MobileSync/Backup"

// LookupBackupDirectories to choose backup directory
func LookupBackupDirectories() (rt []*BackupInformation, err error) {
	if runtime.GOOS == "windows" {

	}
	if runtime.GOOS == "darwin" {
		home, err := os.UserHomeDir()
		if err != nil {
			return nil, err
		}
		backupsDir := filepath.Join(home, macosBackupDir)
		files, err := ioutil.ReadDir(backupsDir)
		if err != nil {
			return nil, err
		}
		for _, file := range files {
			if file.IsDir() {
				rt = append(rt, &BackupInformation{
					BackupDir: filepath.Join(backupsDir, file.Name()),
				})
			}
		}
		return rt, nil
	}
	return nil, fmt.Errorf("Not support os: %s", runtime.GOOS)
}
