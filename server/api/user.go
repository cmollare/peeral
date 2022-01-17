package api

//Connect User
func (h *Handler) Connect(login string, pwd string) error {
	return h.userService.Connect(login, pwd)
}
