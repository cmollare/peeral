package api

import "peeral.com/proxy-libp2p/domain/services"

//UserHandler ...
type UserHandler struct {
	userService services.UserService
}

// NewUserHandler ...
func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{userService: *userService}
}

//Connect User
func (u *UserHandler) Connect(login string, pwd string) error {
	return u.userService.Connect(login, pwd)
}
