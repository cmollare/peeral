package api

import (
	"context"

	"peeral.com/proxy-libp2p/domain/services"
	serverrepository "peeral.com/proxy-libp2p/infra/serverRepository"
)

// Conf Configuration of api handler
type Conf struct {
	ctx           context.Context
	eventCb       func(result string)
	serverService *services.ServerService
	userService   *services.UserService
}

//EventCallback inherite from this class to configure callback
type EventCallback interface {
	Event(result string)
}

// NewConf ...
func NewConf() *Conf {
	conf := &Conf{ctx: context.Background()}
	conf.eventCb = func(results string) {
		println("Unimplemented event callback called with :\n%s", results)
	}
	conf.defaultInjection()
	return conf
}

// SetContext context setter, will not be accessible from jar
func (c *Conf) SetContext(ctx context.Context) {
	c.ctx = ctx
}

// SetEventCallback event callback setter
func (c *Conf) SetEventCallback(eventCb EventCallback) {
	c.eventCb = eventCb.Event
}

//Inject injection function, will not be accessible from jar
func (c *Conf) Inject(serverService *services.ServerService, userService *services.UserService) {
	c.serverService = serverService
	c.userService = userService
}

func (c *Conf) defaultInjection() {
	serverRepo := serverrepository.NewServerRepository(context.Background())

	serverService := services.NewServerService(serverRepo)
	userService := services.NewUserService(serverRepo)

	c.serverService = serverService
	c.userService = userService
}
