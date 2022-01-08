package ports

//IServerRepository to manage servers persistances
type IServerRepository interface {
	Create(name string) error
	Join(name string) error
}
