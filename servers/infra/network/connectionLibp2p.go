package network

import (
	"context"
	"fmt"
	"io"
	"log"

	model "example.com/server/domain/model"
	e "example.com/server/domain/model/error"
	libp2p "github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/multiformats/go-multiaddr"
)

// P2p ...
type P2p struct {
	host *host.Host
}

// Connect connect with p2p network
func (c P2p) Connect(context context.Context, options model.ConnectionOptions) (*model.NetworkInterface, error) {

	host, err := makeHost(context, options.Port, options.RandomSeed)

	c.host = &host

	if err != nil {
		return nil, e.ErrP2pConnectingHost
	}

	startPeer(context, *c.host, handleStream)

	return &model.NetworkInterface{}, nil
}

func makeHost(ctx context.Context, port int, randomness io.Reader) (host.Host, error) {
	// Creates a new RSA key pair for this host.
	prvKey, _, err := crypto.GenerateKeyPairWithReader(crypto.RSA, 2048, randomness)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// 0.0.0.0 will listen on any interface device.
	println(fmt.Sprintf("/ip4/0.0.0.0/tcp/%d", port))
	sourceMultiAddr, _ := multiaddr.NewMultiaddr(fmt.Sprintf("/ip4/0.0.0.0/tcp/%d", port))

	// libp2p.New constructs a new libp2p Host.
	// Other options can be added here.
	return libp2p.New(
		ctx,
		libp2p.ListenAddrs(sourceMultiAddr),
		libp2p.Identity(prvKey),
	)
}

func startPeer(ctx context.Context, h host.Host, streamHandler network.StreamHandler) {
	// Set a function as stream handler.
	// This function is called when a peer connects, and starts a stream with this protocol.
	// Only applies on the receiving side.
	h.SetStreamHandler("/chat/1.0.0", streamHandler)

	// Let's get the actual TCP port from our listen multiaddr, in case we're using 0 (default; random available port).
	var port string
	for _, la := range h.Network().ListenAddresses() {
		if p, err := la.ValueForProtocol(multiaddr.P_TCP); err == nil {
			port = p
			break
		}
	}

	if port == "" {
		log.Println("was not able to find actual local port")
		return
	}

	log.Printf("Run './chat -d /ip4/127.0.0.1/tcp/%v/p2p/%s' on another console.\n", port, h.ID().Pretty())
	log.Println("You can replace 127.0.0.1 with public IP as well.")
	log.Println("Waiting for incoming connection")
	log.Println()
}
