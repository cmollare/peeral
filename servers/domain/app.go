package domain

import (
	"context"

	i "example.com/server/domain/adapter/infra"
	"example.com/server/domain/model"
)

// Application application structure
type Application struct {
	NetworkInfra i.Network
}

// New ...
func New(networkInfra i.Network) *Application {
	return &Application{NetworkInfra: networkInfra}
}

// Connect connect the appto the p2p network
func (a *Application) Connect(context context.Context, options model.ConnectionOptions) {
	a.NetworkInfra.Connect(context, options)
}
