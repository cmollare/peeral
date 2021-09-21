package presentation

import "fmt"

// ConsoleCallbacks ...
type ConsoleCallbacks struct {
}

// NewConsoleCallbacks ...
func NewConsoleCallbacks() *ConsoleCallbacks {
	return &ConsoleCallbacks{}
}

// OnReceive ...
func (c *ConsoleCallbacks) OnReceive(s string, err string) {
	fmt.Printf("callback called %s\n", s)
}

// CustomHostCallbacks ...
type CustomHostCallbacks struct {
}

// NewCustomHostCallbacks ...
func NewCustomHostCallbacks() *CustomHostCallbacks {
	return &CustomHostCallbacks{}
}

// OnListening Called when host starts to listen on adresses
func (c *CustomHostCallbacks) OnListening(id string, err string) {

}

// OnPeersDiscovered Called when a discover process ends
func (c *CustomHostCallbacks) OnPeersDiscovered(peersIds []string) {

}
