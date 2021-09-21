package libp2p

import (
	"context"
	"fmt"
	"log"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	kaddht "github.com/libp2p/go-libp2p-kad-dht"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"peeral.com/proxy-libp2p/libp2p/interfaces"
	"peeral.com/proxy-libp2p/libp2p/ipfshost"
)

type discoveryNotifee struct {
	h   host.Host
	ctx context.Context
}

func (m *discoveryNotifee) HandlePeerFound(pi peer.AddrInfo) {
	if m.h.Network().Connectedness(pi.ID) != network.Connected {
		fmt.Printf("Found %s!\n", pi.ID.ShortString())
		err := m.h.Connect(m.ctx, pi)
		if err != nil {
			fmt.Printf("Can't connect %s: !\n", err)
		} else {
			fmt.Printf("Connected to %s !\n", pi.ID.ShortString())
		}
	}
}

// Peer ...
type Peer struct {
	ctx             context.Context
	RoutedHost      *ipfshost.RoutedHost
	host            host.Host
	dht             *kaddht.IpfsDHT
	hostCallbacks   interfaces.HostCallbacks
	streamCallbacks interfaces.StreamCallbacks
	topic           *pubsub.Topic
	sub             *pubsub.Subscription
}

// NewPeer ...
func NewPeer(hostCallbacks interfaces.HostCallbacks, streamCallbacks interfaces.StreamCallbacks) *Peer {
	return &Peer{hostCallbacks: hostCallbacks, streamCallbacks: streamCallbacks}
}

// Connect Announce peer on IPFS network
func (p *Peer) Connect() {

	ctx := context.Background()
	p.ctx = ctx

	listenF := 0
	var seed int64 = 0

	// Make a host that listens on the given multiaddress
	bootstrapPeers := ipfshost.IPFS_PEERS

	routedHost, err := ipfshost.MakeRoutedHost(listenF, seed, bootstrapPeers)

	p.RoutedHost = routedHost

	p.host = routedHost.Host
	p.dht = routedHost.Dht
	if err != nil {
		log.Fatal(err)
		p.onListenCallback("", err.Error())
	}

	go p.discoverPeers()

	p.onListenCallback(p.host.ID().Pretty(), "")
}

// Close peer
func (p *Peer) Close() {
	p.host.Close()
}

// ConnectToPeer Connect to peer with given addr
func (p *Peer) ConnectToPeer(peerID string) error {

	peerid, err := peer.Decode(peerID)
	if err != nil {
		log.Fatalln(err)
	}

	// peerinfo := peer.AddrInfo{ID: peerid}
	// make a new stream from host B to host A
	// it should be handled on host A by the handler we set above because
	// we use the same /echo/1.0.0 protocol
	adr, err := p.dht.FindPeer(p.ctx, peerid)
	return p.host.Connect(context.Background(), adr)

}

// StartInputLoop ...
func (p *Peer) StartInputLoop() {
	donec := make(chan struct{}, 1)
	go p.chatInputLoop(p.ctx, p.host, donec)
}
