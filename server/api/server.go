package api

import (
	"peeral.com/proxy-libp2p/domain/services"
)

// ServerCmds server object
type ServerCmds struct {
	serverService *services.ServerService
}

// NewServerCmds create new server service
func NewServerCmds(server *services.ServerService) *ServerCmds {
	return &ServerCmds{serverService: server}
}

// CreateServer create a server with following name
func (p *ServerCmds) CreateServer(name string) error {
	return p.serverService.CreateServer(name)
}

// ConnectToServer connect to an existing server
func (p *ServerCmds) ConnectToServer(name string) error {
	return p.serverService.JoinServer(name)
}
