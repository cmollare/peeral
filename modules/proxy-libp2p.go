package main

import (
	"peeral.com/proxy-libp2p/libp2p"
	"peeral.com/proxy-libp2p/libp2p/presentation"
)

func main() {
	sh := presentation.NewConsoleCallbacks()
	HostCb := presentation.NewCustomHostCallbacks()
	peer := libp2p.NewPeer(HostCb, sh)

	peer.Connect()
	peer.StartListening()
	//peer.ConnectToPeer("")

	//peer.StartInputLoop()

	select {}
}
