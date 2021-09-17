package libp2p

func (p *Peer) onListenCallback(id string, err string) {
	if p.hostCallbacks == nil {
		return
	}
	p.hostCallbacks.OnListening(id, err)
}

func (p *Peer) onPeerDiscoveredCallback(peerIds []string) {
	if p.hostCallbacks == nil {
		return
	}
	p.hostCallbacks.OnPeersDiscovered(peerIds)
}

func (p *Peer) onMessageReceived(s string, err string) {
	if p.streamCallbacks == nil {
		return
	}
	p.streamCallbacks.OnReceive(s, err)
}
