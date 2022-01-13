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
	peerHost *peerHost
}

// NewServerRepository ...
func NewServerRepository() *ServerRepository {
	once.Do(func() {
		instance = &ServerRepository{peerHost: &peerHost{}}
	})

	return instance
}

// Connect implementation of interface with libp2p
func (s *ServerRepository) Connect(login string, pwd string) error {
	return s.peerHost.connect(context.Background())
}

// Create implementation of interface with libp2p
func (s *ServerRepository) Create(name string) error {
	return nil
}

// Join implementation of interface with libp2p
func (s *ServerRepository) Join(name string) error {
	return nil
}
