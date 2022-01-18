package api

import "peeral.com/proxy-libp2p/domain/models"

//Connect User
func (h *Handler) Connect(login string, pwd string) error {
	return h.userService.Connect(login, pwd)
}

func (h *Handler) Send(message string, topic string) error {
	return h.userService.SendMessage(&models.MessageDto{Data: message, Topic: topic})
}
