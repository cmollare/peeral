package libp2p

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"context"

	"io/ioutil"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/protocol"
)

const chatProtocol = protocol.ID("/libp2p/chat/1.0.0")

func (p *Peer) chatHandler(s network.Stream) {
	data, err := ioutil.ReadAll(s)

	fmt.Println("Received:", string(data))
	p.onMessageReceived(string(data), err.Error())
}

// Send ...
func (p *Peer) Send(msg string) {
	for _, peer := range p.host.Network().Peers() {
		if _, err := p.host.Peerstore().SupportsProtocols(peer, string(chatProtocol)); err == nil {
			s, err := p.host.NewStream(p.ctx, peer, chatProtocol)
			defer func() {
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
				}
			}()
			if err != nil {
				continue
			}
			err = chatSend(msg, s)
		}
	}
}

func chatSend(msg string, s network.Stream) error {
	fmt.Println("Sending:", msg)
	w := bufio.NewWriter(s)
	n, err := w.WriteString(msg)
	if n != len(msg) {
		return fmt.Errorf("expected to write %d bytes, wrote %d", len(msg), n)
	}
	if err != nil {
		return err
	}
	if err = w.Flush(); err != nil {
		return err
	}
	s.Close()
	data, err := ioutil.ReadAll(s)
	if err != nil {
		return err
	}
	if len(data) > 0 {
		fmt.Println("Received:", string(data))
	}
	return nil
}

func (p *Peer) chatInputLoop(ctx context.Context, h host.Host, donec chan struct{}) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		msg := scanner.Text()

		log.Println("SEND SOMETHING")

		toSend := &ChatMessage{
			Message:    msg,
			SenderID:   "test",
			SenderNick: "nickname",
		}
		jsonToSend, _ := json.Marshal(toSend)
		p.topic.Publish(p.ctx, jsonToSend)
		log.Printf("Message sent : %s\n", msg)

		//p.Send(msg)
	}
	donec <- struct{}{}
}
