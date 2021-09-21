package main

import (
	"flag"
	"log"

	"peeral.com/proxy-libp2p/libp2p"
	"peeral.com/proxy-libp2p/libp2p/presentation"
)

func main() {

	target := flag.String("d", "", "target peer to dial")
	flag.Parse()

	sh := presentation.NewConsoleCallbacks()
	HostCb := presentation.NewCustomHostCallbacks()
	peer := libp2p.NewPeer(HostCb, sh)

	peer.Connect()
	//peer.StartListening()
	peer.JoinNetwork("test")
	//peer.ConnectToPeer("")

	if *target != "" {
		log.Printf("CONNECT TO PEER %s\n", *target)
		e := peer.ConnectToPeer(*target)
		if e != nil {
			log.Printf("ERROR CONNECTING ? %s\n", e.Error())
		}
	}

	peer.StartInputLoop()

	select {}
}
