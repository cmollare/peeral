package serverrepository

import (
	"context"
	"encoding/json"
	"log"

	"github.com/libp2p/go-libp2p-core/host"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	ma "github.com/multiformats/go-multiaddr"
	"peeral.com/proxy-libp2p/domain/models"
	"peeral.com/proxy-libp2p/domain/singleton"
	"peeral.com/proxy-libp2p/libp2p/ipfshost"
)

type peerHost struct {
	ctx          context.Context
	host         host.Host
	dht          *dht.IpfsDHT
	pubsub       *pubsub.PubSub
	hostAdresses []ma.Multiaddr
}

type chatRoom struct {
	// Messages is a channel of messages received from other peers in the chat room
	ExitEvent chan struct{}

	pubsub       *pubsub.PubSub
	topic        *pubsub.Topic
	subscription *pubsub.Subscription

	roomName string
}

type ChatMessage struct {
	ChatRoomName string
	Message      string
	SenderID     string
	SenderNick   string
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

	ps, err := pubsub.NewGossipSub(ph.ctx, ph.host)
	if err != nil {
		return err
	}
	ph.pubsub = ps

	return nil
}

// JoinNetwork Join a pubsub topic a.k.a. a network
func (ph *peerHost) joinChatRoom(name string) error {

	topic, err := ph.pubsub.Join(name)
	if err != nil {
		return err
	}

	// and subscribe to it
	sub, err := topic.Subscribe()
	if err != nil {
		return err
	}

	cr := &chatRoom{
		pubsub:       ph.pubsub,
		topic:        topic,
		subscription: sub,
		roomName:     name,
		ExitEvent:    make(chan struct{}, 1),
	}

	//p.chatRooms = append(p.chatRooms, cr)

	// start reading messages from the subscription in a loop
	go ph.readLoop(cr)

	return nil
}

// readLoop pulls messages from the pubsub topic and pushes them onto the Messages channel.
func (ph *peerHost) readLoop(cr *chatRoom) {
	for {
		select {
		case <-cr.ExitEvent:
			close(cr.ExitEvent)
			return
		default:
			msg, err := cr.subscription.Next(ph.ctx)
			if err != nil {
				close(cr.ExitEvent)
				return
			}
			// only forward messages delivered by others
			if msg.ReceivedFrom == ph.host.ID() {
				continue
			}
			cm := new(ChatMessage)
			err = json.Unmarshal(msg.Data, cm)
			if err != nil {
				//TODO implement error message callback
				continue
			}
			// TODO : send valid messages to the callback with more details
			singleton.LogEvent(models.NewMessageData(string(msg.Data)))
		}
	}
}
