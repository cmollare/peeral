package serverrepository

import "sync"

var (
	once     sync.Once
	instance *ServerRepository
)

//ServerRepository ...
type ServerRepository struct {
}

// NewServerRepository ...
func NewServerRepository() *ServerRepository {
	once.Do(func() {
		instance = &ServerRepository{}
	})

	return instance
}

// Connect implementation of interface with libp2p
func (s *ServerRepository) Connect(login string, pwd string) error {
	return nil
}

// Create implementation of interface with libp2p
func (s *ServerRepository) Create(name string) error {
	return nil
}

// Join implementation of interface with libp2p
func (s *ServerRepository) Join(name string) error {
	return nil
}
