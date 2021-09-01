package infra

import (
	"context"

	model "example.com/server/domain/model"
)

// Network management interface
type Network interface {
	Connect(context context.Context, options model.ConnectionOptions) (*model.NetworkInterface, error)
}
