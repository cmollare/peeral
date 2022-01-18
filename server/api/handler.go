package api

import (
	"peeral.com/proxy-libp2p/domain/services"
	"peeral.com/proxy-libp2p/domain/singleton"
)

// Handler server object
type Handler struct {
	serverService *services.ServerService
	userService   *services.UserService
}

// NewHandler create new server service
func NewHandler(conf *Conf) *Handler {
	singleton.NewConfiguration(conf.ctx, conf.eventCb)
	return &Handler{serverService: conf.serverService, userService: conf.userService}
}
