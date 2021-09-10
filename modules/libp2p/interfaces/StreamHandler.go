package interfaces

// StreamHandler ...
type StreamHandler interface {
	OnReceive(s string)
}
