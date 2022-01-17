package main

import (
	"context"
	"flag"
	"log"
	"os"

	"peeral.com/proxy-libp2p/api"
	"peeral.com/proxy-libp2p/domain/services"
	serverrepository "peeral.com/proxy-libp2p/infra/serverRepository"
	"peeral.com/proxy-libp2p/libp2p"
	"peeral.com/proxy-libp2p/libp2p/presentation"
)

func mainBis() {

	target := flag.String("d", "", "target peer to dial")
	flag.Parse()

	sh := presentation.NewConsoleCallbacks()
	HostCb := presentation.NewCustomHostCallbacks()
	peer := libp2p.NewPeer(HostCb, sh)

	peer.Connect()
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

var (
	apiHdl injection
)

type injection struct {
	hdl *api.Handler
}

func main() {
	//ctx := context.Background()

	inject()
	err := apiHdl.hdl.Connect("login", "pwd")

	if err != nil {
		log.Printf("Unable to connect")
		os.Exit(1)
	}

	log.Printf("Start loop")
	select {}

}

func inject() {
	serverRepo := serverrepository.NewServerRepository(context.Background())

	serverService := services.NewServerService(serverRepo)
	userService := services.NewUserService(serverRepo)

	apiHdl.hdl = api.NewHandler(serverService, userService)
}
