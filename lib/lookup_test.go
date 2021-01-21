package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLookupBackupDirectories(t *testing.T) {
	assert := assert.New(t)
	backups, err := LookupBackupDirectories()
	assert.Nil(err)
	assert.NotEmpty(backups)
}
