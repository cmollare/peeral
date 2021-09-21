package libp2p

import (
	"encoding/json"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

type ChatRoom struct {
	// Messages is a channel of messages received from other peers in the chat room
	ExitEvent chan struct{}

	ps    *pubsub.PubSub
	topic *pubsub.Topic
	sub   *pubsub.Subscription

	roomName string
}

// JoinNetwork Join a pubsub topic a.k.a. a network
func (p *Peer) JoinNetwork(networkName string) error {
	ps, err := pubsub.NewGossipSub(p.ctx, p.host)
	if err != nil {
		return err
	}

	topic, err := ps.Join(networkName)
	if err != nil {
		return err
	}

	// and subscribe to it
	sub, err := topic.Subscribe()
	if err != nil {
		return err
	}

	cr := &ChatRoom{
		ps:        ps,
		topic:     topic,
		sub:       sub,
		roomName:  networkName,
		ExitEvent: make(chan struct{}, 1),
	}

	p.chatRooms = append(p.chatRooms, cr)

	// start reading messages from the subscription in a loop
	go p.readLoop(cr)

	return nil
}

// ChatMessage gets converted to/from JSON and sent in the body of pubsub messages.
type ChatMessage struct {
	ChatRoomName string
	Message      string
	SenderID     string
	SenderNick   string
}

// readLoop pulls messages from the pubsub topic and pushes them onto the Messages channel.
func (p *Peer) readLoop(cr *ChatRoom) {
	for {
		select {
		case <-cr.ExitEvent:
			close(cr.ExitEvent)
			return
		default:
			msg, err := cr.sub.Next(p.ctx)
			if err != nil {
				close(cr.ExitEvent)
				return
			}
			// only forward messages delivered by others
			if msg.ReceivedFrom == p.RoutedHost.PeerID {
				continue
			}
			cm := new(ChatMessage)
			err = json.Unmarshal(msg.Data, cm)
			if err != nil {
				p.onMessageReceived("", err.Error())
				continue
			}
			// send valid messages to the callback
			p.onMessageReceived(string(msg.Data), "")
		}
	}
}

// Send ...
func (p *Peer) Send(msg string) {
	for _, chatRoom := range p.chatRooms {
		toSend := &ChatMessage{
			Message:      msg,
			SenderID:     p.RoutedHost.PeerID.Pretty(),
			SenderNick:   p.RoutedHost.PeerID.Pretty(),
			ChatRoomName: chatRoom.roomName,
		}

		jsonToSend, _ := json.Marshal(toSend)
		chatRoom.topic.Publish(p.ctx, jsonToSend)
	}

}
