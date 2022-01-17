package api

// CreateServer create a server with following name
func (h *Handler) CreateServer(name string) error {
	return h.serverService.CreateServer(name)
}

// ConnectToServer connect to an existing server
func (h *Handler) ConnectToServer(name string) error {
	return h.serverService.JoinServer(name)
}
