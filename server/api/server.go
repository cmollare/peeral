package api

import (
	"peeral.com/proxy-libp2p/domain/services"
)

// ServerHandler server object
type ServerHandler struct {
	serverService *services.ServerService
}

// NewServerHandler create new server service
func NewServerHandler(server *services.ServerService) *ServerHandler {
	return &ServerHandler{serverService: server}
}

// CreateServer create a server with following name
func (p *ServerHandler) CreateServer(name string) error {
	return p.serverService.CreateServer(name)
}

// ConnectToServer connect to an existing server
func (p *ServerHandler) ConnectToServer(name string) error {
	return p.serverService.JoinServer(name)
}
