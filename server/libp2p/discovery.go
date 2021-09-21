package libp2p

import (
	disc "github.com/libp2p/go-libp2p-discovery"
)

func (p *Peer) discoverPeers() []string {
	routingDiscovery := disc.NewRoutingDiscovery(p.dht)
	disc.Advertise(p.ctx, routingDiscovery, string(chatProtocol))
	peers, err := disc.FindPeers(p.ctx, routingDiscovery, string(chatProtocol))
	if err != nil {
		panic(err)
	}
	var res []string
	for _, peer := range peers {
		res = append(res, peer.ID.Pretty())
	}
	p.onPeerDiscoveredCallback(res)

	return res
}
