package lib

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"howett.net/plist"
)

const macosBackupDir = "Library/Application Support/MobileSync/Backup"

// BackupInformation struct
type BackupInformation struct {
	BackupDir    string
	BackupDirRef os.FileInfo
}

// BackupMetadata struct
type BackupMetadata struct {
	Title          string
	BackupDateTime time.Time
}

// ManifestPlist struct
type ManifestPlist struct {
	Version     string
	Date        time.Time
	IsEncrypted bool
	Lockdown    struct {
		DeviceName     string
		SerialNumber   string
		UniqueDeviceID string
		ProductType    string
		ProductVersion string
	}
	Applications map[string]map[string]interface{}
}

// GetBackupMetadata func
func (backup *BackupInformation) GetBackupMetadata() (*ManifestPlist, error) {
	plistManifest := &ManifestPlist{}
	plistFile, err := os.Open(filepath.Join(backup.BackupDir, "Manifest.plist"))
	if err != nil {
		return nil, err
	}
	plistContent, err := ioutil.ReadAll(plistFile)
	if err != nil {
		return nil, err
	}
	_, err = plist.Unmarshal(plistContent, plistManifest)
	if err != nil {
		return nil, err
	}
	return plistManifest, nil
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
		backupsDir := filepath.Join(home, macosBackupDir)
		files, err := ioutil.ReadDir(backupsDir)
		if err != nil {
			return nil, err
		}
		for _, file := range files {
			if file.IsDir() {
				rt = append(rt, &BackupInformation{
					BackupDirRef: file,
					BackupDir:    filepath.Join(backupsDir, file.Name()),
				})
			}
		}
		return rt, nil
	}
	return nil, fmt.Errorf("Not support os: %s", runtime.GOOS)
}
