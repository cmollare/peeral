package serverrepository

import (
	"context"
	"log"

	"github.com/libp2p/go-libp2p-core/host"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	ma "github.com/multiformats/go-multiaddr"
	"peeral.com/proxy-libp2p/domain/models"
	"peeral.com/proxy-libp2p/domain/singleton"
	"peeral.com/proxy-libp2p/libp2p/ipfshost"
)

type peerHost struct {
	ctx          context.Context
	host         host.Host
	dht          *dht.IpfsDHT
	hostAdresses []ma.Multiaddr
}

func (ph *peerHost) connect() error {

	listenF := 0
	var seed int64 = 0

	// Make a host that listens on the given multiaddress
	bootstrapPeers := ipfshost.IPFS_PEERS

	err := ph.makeRoutedHost(listenF, seed, bootstrapPeers)

	if err != nil {
		log.Fatal(err)
		//TODO : make callbacks
		//p.onListenCallback("", err.Error())
	}

	//TODO : implement
	go ph.discoverPeers()

	singleton.LogEvent(models.NewLogData("Host listening with ID " + ph.host.ID().Pretty()))

	return nil
}
