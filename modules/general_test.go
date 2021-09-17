package main_test

import (
	"testing"

	"peeral.com/proxy-libp2p/libp2p"
)

func TestNilCallbacksShouldNotCrash(t *testing.T) {

	peer := libp2p.NewPeer(nil, nil)
	peer.Connect()
	peer.StartListening()
}
