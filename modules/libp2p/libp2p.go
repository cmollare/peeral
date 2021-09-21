package libp2p

import (
	"bufio"
	"context"
	"fmt"
	"log"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/routing"
	kaddht "github.com/libp2p/go-libp2p-kad-dht"
	mplex "github.com/libp2p/go-libp2p-mplex"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	tls "github.com/libp2p/go-libp2p-tls"
	yamux "github.com/libp2p/go-libp2p-yamux"
	"github.com/libp2p/go-tcp-transport"
	ws "github.com/libp2p/go-ws-transport"
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
	ctx             context.Context
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

// Connect Announce peer on IPFS DHT
func (p *Peer) Connect() {

	ctx := context.Background()
	p.ctx = ctx

	global := true
	listenF := 0
	var seed int64 = 0

	// Make a host that listens on the given multiaddress
	var bootstrapPeers []peer.AddrInfo
	var globalFlag string
	if global {
		log.Println("using global bootstrap")
		bootstrapPeers = IPFS_PEERS
		globalFlag = " -global"
	} else {
		log.Println("using local bootstrap")
		bootstrapPeers = getLocalPeerInfo()
		globalFlag = ""
	}
	ha, err := p.makeRoutedHost(listenF, seed, bootstrapPeers, globalFlag)

	p.host = ha
	if err != nil {
		log.Fatal(err)
		p.onListenCallback("", err.Error())
	}
	p.onListenCallback(p.host.ID().Pretty(), "")

	go p.discoverPeers()

	/*ps, err := pubsub.NewGossipSub(p.ctx, p.host)
	if err != nil {
		panic(err)
	}

	topic, err := ps.Join("/peeral/test/1.0.0")
	if err != nil {
		panic(err)
	}

	// and subscribe to it
	sub, err := topic.Subscribe()
	if err != nil {
		panic(err)
	}

	log.Printf("%s", sub)

	peers := ps.ListPeers("/peeral/test/1.0.0")
	for _, p := range peers {
		log.Printf("PEERS FOUNDS %s", p)
	}*/
}

// StartListening register listener on given protocole
func (p *Peer) StartListening() {
	p.host.SetStreamHandler("/echo/2.0.0", func(s network.Stream) {
		log.Println("Got a new stream!")
		if err := doEcho(p.streamCallbacks, s); err != nil {
			log.Println(err)
			s.Reset()
		} else {
			s.Close()
		}
	})
}

func doEcho(hdl interfaces.StreamCallbacks, s network.Stream) error {
	buf := bufio.NewReader(s)
	str, err := buf.ReadString('\n')

	hdl.OnReceive(str, err.Error())
	return err
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

	/*if add != "" {
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
