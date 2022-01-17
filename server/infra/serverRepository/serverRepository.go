package serverrepository

import (
	"context"
	"sync"
)

var (
	once     sync.Once
	instance *ServerRepository
)

//ServerRepository ...
type ServerRepository struct {
	ctx      context.Context
	peerHost *peerHost
}

// NewServerRepository ...
func NewServerRepository(ctx context.Context) *ServerRepository {
	once.Do(func() {
		instance = &ServerRepository{ctx: ctx, peerHost: &peerHost{ctx: ctx}}
	})

	return instance
}

// Connect implementation of interface with libp2p
func (s *ServerRepository) Connect(login string, pwd string) error {
	return s.peerHost.connect()
}

// Create implementation of interface with libp2p
func (s *ServerRepository) Create(name string) error {
	return nil
}

// Join implementation of interface with libp2p
func (s *ServerRepository) Join(name string) error {
	return nil
}
