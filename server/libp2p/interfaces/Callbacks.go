package interfaces

// StreamCallbacks ...
type StreamCallbacks interface {
	OnReceive(s string, err string)
}

//HostCallbacks to Handle host events
type HostCallbacks interface {
	OnListening(id string, err string)
	OnPeersDiscovered(peersIds []string)
}
