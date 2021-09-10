package presentation

import "fmt"

// ConsoleStreamHandler ...
type ConsoleStreamHandler struct {
}

// NewConsoleStreamHandler ...
func NewConsoleStreamHandler() *ConsoleStreamHandler {
	return &ConsoleStreamHandler{}
}

// OnReceive ...
func (c *ConsoleStreamHandler) OnReceive(s string) {
	fmt.Printf("interface called %s\n", s)
}
