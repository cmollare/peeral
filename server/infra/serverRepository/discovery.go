package serverrepository

import (
	"fmt"
	"log"

	"github.com/libp2p/go-libp2p-core/protocol"
	disc "github.com/libp2p/go-libp2p-discovery"
	"peeral.com/proxy-libp2p/domain/models"
	"peeral.com/proxy-libp2p/domain/singleton"
)

const chatProtocol = protocol.ID("/libp2p/chat/1.0.0")

func (ph *peerHost) discoverPeers() []string {
	routingDiscovery := disc.NewRoutingDiscovery(ph.dht)
	disc.Advertise(ph.ctx, routingDiscovery, string(chatProtocol))
	peers, err := disc.FindPeers(ph.ctx, routingDiscovery, string(chatProtocol))
	if err != nil {
		panic(err)
	}

	peerList := ph.dht.RoutingTable().ListPeers()
	singleton.LogEvent(models.NewLogData(fmt.Sprintf("nb total peer %d", len(peerList))))

	var res []string
	for _, peer := range peers {
		res = append(res, peer.ID.Pretty())
		//TODO : replace by logEvent
		log.Println(ph.ctx, "peer Discovered ", peer.ID.Pretty())
	}

	log.Println(ph.ctx, "End discovery ")

	//TODO : add callback
	//p.onPeerDiscoveredCallback(res)

	return res
}
