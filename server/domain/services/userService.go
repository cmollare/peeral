package services

import (
	"peeral.com/proxy-libp2p/domain/models"
	"peeral.com/proxy-libp2p/domain/ports"
)

//UserService ...
type UserService struct {
	serverRepo ports.IServerRepository
}

//NewUserService ...
func NewUserService(serverRepo ports.IServerRepository) *UserService {
	return &UserService{serverRepo: serverRepo}
}

//Connect ...
func (u UserService) Connect(login string, pwd string) error {
	err := u.serverRepo.Connect(login, pwd)
	if err != nil {
		return err
	}

	err = u.serverRepo.Join("CustomRoom")
	if err != nil {
		return err
	}

	return nil
}

//SendMessage ...
func (u UserService) SendMessage(message *models.MessageDto) error {
	return u.serverRepo.SendMessage(message)
}
