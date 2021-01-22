package lib

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"os"

	"github.com/square/squalor"
	"howett.net/plist"
)

const userProfileName = "mmsetting.archive" // plist file for user settings

// WechatTraverser struct
type WechatTraverser struct {
	backup  *BackupInformation
	friends *squalor.Model
}

// WechatFriend struct
type WechatFriend struct {
	UserName string `db:"userName"`
	UserType int    `db:"type"`
	remark   []byte `db:"dbContactRemark"`
}

// WechatUser struct
type WechatUser struct {
	Version  int           `plist:"$version"`
	Archiver string        `plist:"$archiver"`
	Items    []interface{} `plist:"$objects"`
}

// GetUserWeChatID string
func (s *WechatUser) GetUserWeChatID() string {
	return s.Items[2].(string)
}

// GetUserWeChatIDMD5 string
func (s *WechatUser) GetUserWeChatIDMD5() string {
	hash := md5.Sum([]byte(s.GetUserWeChatID()))
	return hex.EncodeToString(hash[:])
}

// GetUserWeChatNickName string
func (s *WechatUser) GetUserWeChatNickName() string {
	return s.Items[3].(string)
}

// GetFriendsList of current user
func (s *WechatUser) GetFriendsList() {
}

// NewWechatTraverser instance
func NewWechatTraverser(backup *BackupInformation) *WechatTraverser {
	rt := &WechatTraverser{backup: backup}
	return rt
}

// GetUserList func
func (t *WechatTraverser) GetUserList() (rt []*WechatUser) {
	userProfileFiles := t.backup.FindAllFilesByName(userProfileName)

	for _, userProfileFile := range userProfileFiles {
		user := &WechatUser{}
		plistFile, err := os.Open(userProfileFile)
		if err == nil {
			content, err := ioutil.ReadAll(plistFile)
			if err == nil {
				plist.Unmarshal(content, user)
				rt = append(rt, user)
			}
		}

	}

	return rt
}
