package libp2p

import (
	"bufio"
	"os"

	"context"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/protocol"
)

const chatProtocol = protocol.ID("/libp2p/chat/1.0.0")

func (p *Peer) chatInputLoop(ctx context.Context, h host.Host, donec chan struct{}) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		msg := scanner.Text()

		p.Send(msg)
	}
	donec <- struct{}{}
}
