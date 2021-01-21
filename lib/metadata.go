package lib

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/square/squalor"
	"howett.net/plist"
)

const wechatAppID = "com.tencent.xin"

// BackupInformation struct
type BackupInformation struct {
	BackupDir string
	db        *squalor.DB
	files     *squalor.Model
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

// NewBackupInformation func
func NewBackupInformation(dir string) (*BackupInformation, error) {
	db, err := sql.Open("sqlite3", filepath.Join(dir, "Manifest.db"))
	if err != nil {
		return nil, err
	}
	_db, err := squalor.NewDB(db)
	if err != nil {
		return nil, err
	}
	files, err := _db.BindModel("Files", &BackupFile{})
	if err != nil {
		return nil, err
	}

	return &BackupInformation{
		BackupDir: dir,
		db:        _db,
		files:     files,
	}, nil
}

// BackupFile for itunes
type BackupFile struct {
	FileID       string `db:"fileID"`
	Domain       string `db:"domain"`
	RelativePath string `db:"relativePath"`
	Flag         int    `db:"flags"`
}

// GetRealPath func to get the real path
func (backup *BackupInformation) GetRealPath(iosPath string) (string, error) {
	files := []BackupFile{}
	backup.db.Select(
		&files,
		backup.files.Select(backup.files.All()).Where(backup.files.C("relativePath").Eq(iosPath)),
	)
	if len(files) > 0 {
		fileID := files[0].FileID
		return filepath.Join(backup.BackupDir, fileID[:2], fileID), nil
	}
	return "", fmt.Errorf("not found file %s", iosPath)
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
