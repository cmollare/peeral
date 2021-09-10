package libp2p

import (
	"context"
	"fmt"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/routing"
	disc "github.com/libp2p/go-libp2p-discovery"
	kaddht "github.com/libp2p/go-libp2p-kad-dht"
	mplex "github.com/libp2p/go-libp2p-mplex"
	tls "github.com/libp2p/go-libp2p-tls"
	yamux "github.com/libp2p/go-libp2p-yamux"
	"github.com/libp2p/go-libp2p/p2p/discovery/mdns"
	"github.com/libp2p/go-tcp-transport"
	ws "github.com/libp2p/go-ws-transport"
	"github.com/multiformats/go-multiaddr"
	"peeral.com/proxy-libp2p/libp2p/interfaces"
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
	ctx           context.Context
	host          host.Host
	dht           *kaddht.IpfsDHT
	streamHandler interfaces.StreamHandler
}

// NewPeer ...
func NewPeer(streamHandler interfaces.StreamHandler) *Peer {
	return &Peer{streamHandler: streamHandler}
}

// Start ...
func (p *Peer) Start() {
	ctx := context.Background()
	p.ctx = ctx

	transports := libp2p.ChainOptions(
		libp2p.Transport(tcp.NewTCPTransport),
		libp2p.Transport(ws.New),
	)

	listenAddrs := libp2p.ListenAddrStrings(
		"/ip4/0.0.0.0/tcp/0",
		"/ip4/0.0.0.0/tcp/0/ws",
	)

	muxers := libp2p.ChainOptions(
		libp2p.Muxer("/yamux/1.0.0", yamux.DefaultTransport),
		libp2p.Muxer("/mplex/6.7.0", mplex.DefaultTransport),
	)

	security := libp2p.Security(tls.ID, tls.New)

	var dht *kaddht.IpfsDHT
	newDHT := func(h host.Host) (routing.PeerRouting, error) {
		var err error
		dht, err = kaddht.New(ctx, h)
		p.dht = dht
		return dht, err
	}

	routing := libp2p.Routing(newDHT)

	host, err := libp2p.New(ctx, transports, listenAddrs, muxers, security, routing)
	if err != nil {
		panic(err)
	}
	p.host = host

	host.SetStreamHandler(chatProtocol, p.chatHandler)

	for _, addr := range host.Addrs() {
		fmt.Println("Listening on", addr)
	}
}

// Close peer
func (p *Peer) Close() {
	p.host.Close()
}

// ConnectToPeer Connect to peer with given addr
func (p *Peer) ConnectToPeer(add string) {

	if add != "" {
		targetAddr, err := multiaddr.NewMultiaddr(add)
		if err != nil {
			panic(err)
		}

		targetInfo, err := peer.AddrInfoFromP2pAddr(targetAddr)
		if err != nil {
			panic(err)
		}

		err = p.host.Connect(p.ctx, *targetInfo)
		if err != nil {
			panic(err)
		}

		fmt.Println("Connected to", targetInfo.ID)
	}

	mdns := mdns.NewMdnsService(p.host, "")
	notifee := &discoveryNotifee{h: p.host, ctx: p.ctx}
	mdns.RegisterNotifee(notifee)

	err := p.dht.Bootstrap(p.ctx)
	if err != nil {
		panic(err)
	}

	routingDiscovery := disc.NewRoutingDiscovery(p.dht)
	disc.Advertise(p.ctx, routingDiscovery, string(chatProtocol))
	peers, err := disc.FindPeers(p.ctx, routingDiscovery, string(chatProtocol))
	if err != nil {
		panic(err)
	}
	for _, peer := range peers {
		notifee.HandlePeerFound(peer)
	}

	/*stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT)

	select {
	case <-stop:
		(p.host).Close()
		os.Exit(0)
	case <-donec:
		(p.host).Close()
	}*/
}

// StartInputLoop ...
func (p *Peer) StartInputLoop() {
	donec := make(chan struct{}, 1)
	go p.chatInputLoop(p.ctx, p.host, donec)
}
