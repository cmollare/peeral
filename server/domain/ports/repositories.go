package ports

//IServerRepository to manage servers persistances
type IServerRepository interface {
	Connect(login string, pwd string) error
	Create(name string) error
	Join(name string) error
}
