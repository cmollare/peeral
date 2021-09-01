package error

import "errors"

var (
	//ErrP2pConnectingHost Can't connect host
	ErrP2pConnectingHost = errors.New("Failed to make libp2p host")
)
