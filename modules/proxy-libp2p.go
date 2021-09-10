package main

import "peeral.com/proxy-libp2p/libp2p"

func main() {
	peer := libp2p.Peer{}

	peer.Start()
	peer.ConnectToPeer("")

	select {}
}
