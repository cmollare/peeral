package testutils

import "log"

// MockHostCallbacks ...
type MockHostCallbacks struct {
	OnListeningStub       func(string, string)
	OnPeersDiscoveredStub func([]string)
}

func OnReceive(s string, err string) {

}

func (m *MockHostCallbacks) OnListening(peerId string, err string) {
	if m.OnListeningStub == nil {
		m.OnListeningStub = func(peerId string, err string) {
			if err != "" {
				log.Printf("Error : %s\n", err)
				return
			}
			log.Printf("Peer %s is listening\n", peerId)
		}
	}

	m.OnListeningStub(peerId, err)
}

func (m *MockHostCallbacks) OnPeersDiscovered(peersIds []string) {
	if m.OnPeersDiscoveredStub == nil {
		m.OnPeersDiscoveredStub = func(s []string) {
			for _, p := range s {
				log.Printf("Peer found with id %s", p)
			}
		}
	}

	m.OnPeersDiscoveredStub(peersIds)
}
