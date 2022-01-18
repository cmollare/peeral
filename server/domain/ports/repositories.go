package ports

import "peeral.com/proxy-libp2p/domain/models"

//IServerRepository to manage servers persistances
type IServerRepository interface {
	Connect(login string, pwd string) error
	Create(name string) error
	Join(name string) error
	SendMessage(message *models.MessageDto) error
}
