package lib

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

// BackupInformation struct
type BackupInformation struct {
	File           os.FileInfo
	BackupDateTime time.Time
}

// LookupBackupDirectories to choose backup directory
func LookupBackupDirectories() (rt []*BackupInformation, err error) {
	if runtime.GOOS == "windows" {

	}
	if runtime.GOOS == "darwin" {
		home, err := os.UserHomeDir()
		if err != nil {
			return nil, err
		}
		files, err := ioutil.ReadDir(filepath.Join(home, "Library/Application Support/MobileSync/Backup"))
		if err != nil {
			return nil, err
		}
		for _, file := range files {
			rt = append(rt, &BackupInformation{
				File: file,
			})
		}
		return rt, nil
	}
	return nil, fmt.Errorf("Not support os: %s", runtime.GOOS)
}
