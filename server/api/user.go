package api

//UserHandler ...
type UserHandler struct {
}

// NewUserHandler ...
func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

//Connect User
func (u *UserHandler) Connect(login string, pwd string) error {
	return nil
}
