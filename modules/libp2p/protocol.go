package libp2p

import (
	"bufio"
	"fmt"
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
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		p.streamHandler.OnReceive(err.Error())
	}
	fmt.Println("Received:", string(data))
	p.streamHandler.OnReceive(string(data))
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

		p.Send(msg)
	}
	donec <- struct{}{}
}
