package lib

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"howett.net/plist"
)

const wechatAppID = "com.tencent.xin"

// BackupInformation struct
type BackupInformation struct {
	BackupDir string
	db        *sqlx.DB
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
	db, err := sqlx.Connect("sqlite3", filepath.Join(dir, "Manifest.db"))
	if err != nil {
		return nil, err
	}
	return &BackupInformation{
		BackupDir: dir,
		db:        db,
	}, nil
}

// BackupFile for itunes
type BackupFile struct {
	FileID       string `db:"fileID"`
	Domain       string `db:"domain"`
	RelativePath string `db:"relativePath"`
	Flag         int    `db:"flags"`
}

// FindAllFilesByName func to get the real path by fuzzy search
func (backup *BackupInformation) FindAllFilesByName(filename string) (rt []string) {
	files := []BackupFile{}
	sql := fmt.Sprintf("SELECT fileID,domain,relativePath,flags from Files where relativePath like \"%%%s\"", filename)
	err := backup.db.Select(&files, sql)
	if err == nil && len(files) > 0 {
		fileID := files[0].FileID
		rt = append(rt, filepath.Join(backup.BackupDir, fileID[:2], fileID))
	}
	return rt
}

// GetRealPath func to get the real path
func (backup *BackupInformation) GetRealPath(iosPath string) (string, error) {
	files := []BackupFile{}
	backup.db.Select(
		&files,
		"SELECT * from Files where relativePath = $1",
		iosPath,
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
