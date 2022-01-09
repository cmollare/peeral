package api_test

import (
	"peeral.com/proxy-libp2p/api"
	"peeral.com/proxy-libp2p/domain/services"
	"peeral.com/proxy-libp2p/mocks"
)

type BaseTest struct {
	serverHdl      *api.ServerHandler
	userHdl        *api.UserHandler
	serverRepoMock *mocks.IServerRepository
}

func setup() *BaseTest {
	serverRepoMock := &mocks.IServerRepository{}

	srvService := services.NewServerService(serverRepoMock)
	userService := services.NewUserService(serverRepoMock)

	return &BaseTest{api.NewServerHandler(srvService), api.NewUserHandler(userService), serverRepoMock}
}
