package testutils

import "log"

// MockHostCallbacks is a mock for host event callbacks
type MockHostCallbacks struct {
	OnListeningStub       func(string, string)
	OnPeersDiscoveredStub func([]string)
}

// MockStreamCallbacks is a mock for message received event callbacks
type MockStreamCallbacks struct {
	OnReceiveStub func(s string, err string)
}

// OnReceive implements StreamCallbacks interface
func (m *MockStreamCallbacks) OnReceive(s string, err string) {
	if m.OnReceiveStub == nil {
		m.OnReceiveStub = func(s string, err string) {
			log.Printf("Message received %s\n", s)
		}
	}

	m.OnReceiveStub(s, err)
}

// OnListening implements HostCallbacks interface
func (m *MockHostCallbacks) OnListening(peerID string, err string) {
	if m.OnListeningStub == nil {
		m.OnListeningStub = func(peerID string, err string) {
			if err != "" {
				log.Printf("Error : %s\n", err)
				return
			}
			log.Printf("Peer %s is listening\n", peerID)
		}
	}

	m.OnListeningStub(peerID, err)
}

// OnPeersDiscovered implements HostCallbacks interface
func (m *MockHostCallbacks) OnPeersDiscovered(peersIDs []string) {
	if m.OnPeersDiscoveredStub == nil {
		m.OnPeersDiscoveredStub = func(s []string) {
			for _, p := range s {
				log.Printf("Peer found with id %s", p)
			}
		}
	}

	m.OnPeersDiscoveredStub(peersIDs)
}
