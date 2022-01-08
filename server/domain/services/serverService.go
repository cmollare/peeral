package services

import (
	"errors"

	"peeral.com/proxy-libp2p/domain/ports"
)

//ServerService ...
type ServerService struct {
	serverRepo ports.IServerRepository
}

//NewServerService ...
func NewServerService(serverRepo ports.IServerRepository) *ServerService {
	return &ServerService{serverRepo: serverRepo}
}

//CreateServer server creation service
func (s *ServerService) CreateServer(name string) error {
	if name == "" {
		return errors.New("Empty name")
	}

	return s.serverRepo.Create(name)
}

//JoinServer server connection service
func (s *ServerService) JoinServer(name string) error {
	if name == "" {
		return errors.New("Empty name")
	}

	return s.serverRepo.Join(name)
}
