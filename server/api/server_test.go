package api_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"peeral.com/proxy-libp2p/api"
	"peeral.com/proxy-libp2p/domain/services"
	"peeral.com/proxy-libp2p/mocks"
)

type BaseServerTest struct {
	serverCmd      *api.ServerCmds
	serverRepoMock *mocks.IServerRepository
}

func setup() *BaseServerTest {
	serverRepoMock := &mocks.IServerRepository{}
	srvService := services.NewServerService(serverRepoMock)
	return &BaseServerTest{api.NewServerCmds(srvService), serverRepoMock}
}

func TestShouldFailCreateServer(t *testing.T) {
	assert := assert.New(t)
	base := setup()

	err := base.serverCmd.CreateServer("")

	assert.NotEqual(nil, err, "Should fail create a server")
}

func TestShouldCreateAServer(t *testing.T) {
	assert := assert.New(t)
	base := setup()

	base.serverRepoMock.On("Create", "newServer").Return(nil).Once()

	err := base.serverCmd.CreateServer("newServer")

	assert.Equal(nil, err, "Should create a server")
}

func TestShouldFailToConnectToServer(t *testing.T) {
	assert := assert.New(t)
	base := setup()
	base.serverRepoMock.On("Join", "nonExistingServer").Return(errors.New("Non Existing Server")).Once()

	err := base.serverCmd.ConnectToServer("nonExistingServer")

	assert.NotEqual(nil, err, "Should fail connecting to a server")
}

func TestShouldConnectToServer(t *testing.T) {
	assert := assert.New(t)
	base := setup()

	base.serverRepoMock.On("Join", "ExistingServer").Return(nil).Once()

	err := base.serverCmd.ConnectToServer("ExistingServer")

	assert.Equal(nil, err, "Should connect to a server")
}
