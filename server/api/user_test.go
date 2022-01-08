package api_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldConnectUser(t *testing.T) {
	assert := assert.New(t)
	base := setup()

	err := base.userHdl.Connect("ExistingUser", "GoodPwd")

	assert.Equal(nil, err, "Should connect user")
}

func TestShouldFailToConnectUser(t *testing.T) {
	assert := assert.New(t)
	base := setup()

	err := base.userHdl.Connect("NonExistingUser", "GoodPwd")

	assert.NotEqual(nil, err, "Should fail connecting user")
}
