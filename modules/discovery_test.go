package main_test

import (
	"testing"
	"time"

	"peeral.com/proxy-libp2p/libp2p"
	"peeral.com/proxy-libp2p/libp2p/testutils"
)

func findInArray(array []string, txt string) bool {
	for _, str := range array {
		if str == txt {
			return true
		}
	}

	return false
}

func TestDiscovery(t *testing.T) {
	var peers []string
	wait2Calls := make(chan bool, 2)
	nbCalls := 0

	mockHostCallbacks := &testutils.MockHostCallbacks{
		OnListeningStub: func(peerId string, err string) {
			if err != "" {
				t.Errorf("Error in listening setup : %s\n", err)
			}
			peers = append(peers, peerId)
		},
		OnPeersDiscoveredStub: func(peersId []string) {
			wait2Calls <- true
			nbCalls += 1
			if nbCalls >= 2 {
				ok := findInArray(peersId, peers[0])
				if !ok {
					t.Errorf("Peer not discovered %s", peers[0])
				}
			}
		},
	}

	peer1 := libp2p.NewPeer(mockHostCallbacks, nil)
	peer1.Connect()

	// wait 4 sec before second launch to wait that first peer records on the network
	time.Sleep(4 * time.Second)

	peer2 := libp2p.NewPeer(mockHostCallbacks, nil)
	peer2.Connect()

	deadline, ok := t.Deadline()
	if !ok {
		t.Errorf("Test Timed out : %s", deadline)
	}

	<-wait2Calls
	<-wait2Calls
}
