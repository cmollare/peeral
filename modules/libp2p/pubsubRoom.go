package libp2p

import (
	"context"
	"encoding/json"
	"log"

	"github.com/libp2p/go-libp2p-core/peer"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

type ChatRoom struct {
	// Messages is a channel of messages received from other peers in the chat room
	Messages chan *ChatMessage

	ctx   context.Context
	ps    *pubsub.PubSub
	topic *pubsub.Topic
	sub   *pubsub.Subscription

	roomName string
	self     peer.ID
	nick     string
}

// JoinNetwork Join a pubsub topic a.k.a. a network
func (p *Peer) JoinNetwork(networkName string) {
	ps, err := pubsub.NewGossipSub(p.ctx, p.host)
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

	p.topic = topic
	p.sub = sub

	log.Printf("Connected to /peeral/test/1.0.0\n")

	cr := &ChatRoom{
		ctx:      p.ctx,
		ps:       ps,
		topic:    topic,
		sub:      sub,
		self:     p.host.ID(),
		nick:     "sdfsfsdf",
		roomName: "/peeral/test/1.0.0",
		Messages: make(chan *ChatMessage, 128),
	}

	// start reading messages from the subscription in a loop
	go cr.readLoop()
}

// ChatMessage gets converted to/from JSON and sent in the body of pubsub messages.
type ChatMessage struct {
	Message    string
	SenderID   string
	SenderNick string
}

// readLoop pulls messages from the pubsub topic and pushes them onto the Messages channel.
func (cr *ChatRoom) readLoop() {
	for {
		msg, err := cr.sub.Next(cr.ctx)
		if err != nil {
			close(cr.Messages)
			return
		}
		// only forward messages delivered by others
		if msg.ReceivedFrom == cr.self {
			continue
		}
		cm := new(ChatMessage)
		err = json.Unmarshal(msg.Data, cm)
		if err != nil {
			continue
		}
		// send valid messages onto the Messages channel
		log.Printf("Message received %s\n", cm.Message)
		//cr.Messages <- cm
	}
}
