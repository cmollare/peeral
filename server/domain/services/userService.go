package services

import "peeral.com/proxy-libp2p/domain/ports"

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
	return u.serverRepo.Connect(login, pwd)
}
