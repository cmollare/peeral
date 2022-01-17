package api

import "peeral.com/proxy-libp2p/domain/services"

// Handler server object
type Handler struct {
	serverService *services.ServerService
	userService   *services.UserService
}

// NewHandler create new server service
func NewHandler(server *services.ServerService, userService *services.UserService) *Handler {
	return &Handler{serverService: server, userService: userService}
}
