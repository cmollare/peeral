package model

import mrand "math/rand"

// NetworkInterface network interface identifier
type NetworkInterface struct {
	Handler string
	Port    int
	Address string
}

// ConnectionOptions contains option for connection
type ConnectionOptions struct {
	Port       int
	RandomSeed *mrand.Rand
}

// NewLibp2pConnectionOptions factory for new connection parameters
func NewLibp2pConnectionOptions(port int) ConnectionOptions {
	r := mrand.New(mrand.NewSource(int64(port)))

	return ConnectionOptions{
		Port:       port,
		RandomSeed: r,
	}
}
