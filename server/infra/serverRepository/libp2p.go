package serverrepository

import (
	"context"
	"log"

	"github.com/libp2p/go-libp2p-core/host"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	ma "github.com/multiformats/go-multiaddr"
	"peeral.com/proxy-libp2p/libp2p/ipfshost"
)

type peerHost struct {
	ctx          context.Context
	host         host.Host
	dht          *dht.IpfsDHT
	hostAdresses []ma.Multiaddr
}

func (ph *peerHost) connect(ctx context.Context) error {
	ph.ctx = ctx

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

	return nil

	//TODO : make callbacks
	//p.onListenCallback(p.host.ID().Pretty(), "")
}
