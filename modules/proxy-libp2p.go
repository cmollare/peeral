package main

import (
	"peeral.com/proxy-libp2p/libp2p"
	"peeral.com/proxy-libp2p/libp2p/presentation"
)

func main() {
	sh := presentation.NewConsoleStreamHandler()
	peer := libp2p.NewPeer(sh)

	peer.Start()
	peer.ConnectToPeer("")

	peer.StartInputLoop()

	select {}
}
