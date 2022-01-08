package api_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldFailCreateServer(t *testing.T) {
	assert := assert.New(t)
	base := setup()

	err := base.serverHdl.CreateServer("")

	assert.NotEqual(nil, err, "Should fail create a server")
}

func TestShouldCreateAServer(t *testing.T) {
	assert := assert.New(t)
	base := setup()

	base.serverRepoMock.On("Create", "newServer").Return(nil).Once()

	err := base.serverHdl.CreateServer("newServer")

	assert.Equal(nil, err, "Should create a server")
}

func TestShouldFailToConnectToServer(t *testing.T) {
	assert := assert.New(t)
	base := setup()
	base.serverRepoMock.On("Join", "nonExistingServer").Return(errors.New("Non Existing Server")).Once()

	err := base.serverHdl.ConnectToServer("nonExistingServer")

	assert.NotEqual(nil, err, "Should fail connecting to a server")
}

func TestShouldConnectToServer(t *testing.T) {
	assert := assert.New(t)
	base := setup()

	base.serverRepoMock.On("Join", "ExistingServer").Return(nil).Once()

	err := base.serverHdl.ConnectToServer("ExistingServer")

	assert.Equal(nil, err, "Should connect to a server")
}
