package api_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldConnectUser(t *testing.T) {
	assert := assert.New(t)
	base := setup()

	base.serverRepoMock.On("Connect", "ExistingUser", "GoodPwd").Return(nil).Once()

	err := base.handler.Connect("ExistingUser", "GoodPwd")

	assert.Equal(nil, err, "Should connect user")
}

func TestShouldFailToConnectUser(t *testing.T) {
	assert := assert.New(t)
	base := setup()

	base.serverRepoMock.On("Connect", "NonExistingUser", "GoodPwd").Return(errors.New("User Not known")).Once()

	err := base.handler.Connect("NonExistingUser", "GoodPwd")

	assert.NotEqual(nil, err, "Should fail connecting user")
}
