package main_test

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/cucumber/godog"
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

type ctxPeerId struct{}
type ctxMessageChannel struct{}
type ctxMyPeer struct{}
type ctxDiscoveredPeerIDs struct{}
type ctxListeningError struct{}
type ctxPeerFoundError struct{}

func aPeerIsConnected(ctx context.Context) (context.Context, error) {
	e := make(chan error, 1)

	var c context.Context

	mockHostCallbacks := &testutils.MockHostCallbacks{
		OnListeningStub: func(peerId string, err string) {
			if err != "" {
				e <- errors.New(err)
				return
			}
			c = context.WithValue(ctx, ctxPeerId{}, peerId)
			e <- nil
		},
		OnPeersDiscoveredStub: func(peersId []string) {
			//e <- nil
		},
	}

	mockStreamCallbacks := &testutils.MockStreamCallbacks{
		OnReceiveStub: func(msg string, err string) {

			msgChan := ctx.Value(ctxMessageChannel{}).(chan string)
			msgStruct := &libp2p.ChatMessage{}
			json.Unmarshal([]byte(msg), msgStruct)
			msgChan <- msgStruct.Message
		},
	}

	peer1 := libp2p.NewPeer(mockHostCallbacks, mockStreamCallbacks)
	peer1.Connect()
	peer1.JoinNetwork("MyChatRoom")

	failure := <-e
	if failure != nil {
		return nil, failure
	}

	//<-e

	return c, nil
}

func hasBeenForSeconds(sec int) error {
	time.Sleep(time.Duration(sec) * time.Second)
	return nil
}

func itShouldFindPeers(ctx context.Context) error {
	err, _ := ctx.Value(ctxPeerFoundError{}).(error)
	return err
}

func itShouldStartListening(ctx context.Context) error {
	err, _ := ctx.Value(ctxListeningError{}).(error)
	return err
}

func myPeerIsStarted(ctx context.Context) (context.Context, error) {
	e := make(chan error, 1)
	var peerFound []string

	mockHostCallbacks := &testutils.MockHostCallbacks{
		OnListeningStub: func(peerId string, err string) {
			if err != "" {
				e <- errors.New(err)
				return
			}
			e <- nil
		},
		OnPeersDiscoveredStub: func(peersId []string) {
			peerId := ctx.Value(ctxPeerId{}).(string)
			ok := findInArray(peersId, peerId)
			if !ok {
				e <- errors.New("Cannot find peer with ID " + peerId)
				return
			}
			e <- nil
			peerFound = peersId
		},
	}

	peer2 := libp2p.NewPeer(mockHostCallbacks, nil)
	peer2.Connect()
	if err := peer2.JoinNetwork("MyChatRoom"); err != nil {
		return ctx, errors.New("Can't join the room named : MyChatRoom")
	}

	c := context.WithValue(ctx, ctxListeningError{}, <-e)
	c = context.WithValue(c, ctxPeerFoundError{}, <-e)
	c = context.WithValue(c, ctxDiscoveredPeerIDs{}, peerFound)
	c = context.WithValue(c, ctxMyPeer{}, peer2)

	return c, nil
}

func isConnectedToFirstPeer(ctx context.Context) error {
	peer := ctx.Value(ctxMyPeer{}).(*libp2p.Peer)
	if peer == nil {
		return errors.New("Peer not found")
	}

	peerList := ctx.Value(ctxDiscoveredPeerIDs{}).([]string)
	/*if len(peerList) > 15 {
		peerList = peerList[0:15]
	}*/
	connectedIds, err := peer.ConnectToPeers(peerList) //Try first 15 peers

	if err != nil {
		return errors.New(err.Error())
	}

	peerID := ctx.Value(ctxPeerId{}).(string)
	if peerID == "" {
		return errors.New("peer ID not found")
	}

	ok := findInArray(connectedIds, peerID)
	if !ok {
		return errors.New("Not connected to peer with ID " + peerID)
	}

	return nil
}

func itShouldJoinRoomNamedCustomChatRoomAndSendAMessageHelloWorld(ctx context.Context) error {
	peer := ctx.Value(ctxMyPeer{}).(*libp2p.Peer)
	if peer == nil {
		return errors.New("Peer not found")
	}

	peer.Send("Hello world")

	time.Sleep(2 * time.Second)

	return nil
}

func otherPeerShouldReceiveAMessageHelloWorld(ctx context.Context) error {
	msgChan := ctx.Value(ctxMessageChannel{}).(chan string)
	if msg := <-msgChan; msg == "Hello world" {
		return nil
	} else {
		return errors.New("Message should be 'Hello world' and is : " + msg)
	}
}

func iCanGetItsID(ctx context.Context) error {
	peerID := ctx.Value(ctxPeerId{}).(string)
	if peerID == "" {
		return errors.New("peer ID not found")
	}
	return nil
}

func myPeerCanConnectToIt(ctx context.Context) error {
	peer := ctx.Value(ctxMyPeer{}).(*libp2p.Peer)
	if peer == nil {
		return errors.New("Peer not found")
	}

	peerID := ctx.Value(ctxPeerId{}).(string)
	if peerID == "" {
		return errors.New("peer ID not found")
	}

	return peer.ConnectToPeer(peerID)
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		//log.SetOutput(ioutil.Discard)
		c := context.WithValue(ctx, ctxMessageChannel{}, make(chan string, 1))

		return c, nil
	})

	ctx.Step(`^a peer is connected$`, aPeerIsConnected)
	ctx.Step(`^has been for (\d+) seconds$`, hasBeenForSeconds)
	ctx.Step(`^it should find peers$`, itShouldFindPeers)
	ctx.Step(`^it should start listening$`, itShouldStartListening)
	ctx.Step(`^my peer is started$`, myPeerIsStarted)
	ctx.Step(`^is connected to first peer$`, isConnectedToFirstPeer)
	ctx.Step(`^it should join room named customChatRoom and send a message hello world$`, itShouldJoinRoomNamedCustomChatRoomAndSendAMessageHelloWorld)
	ctx.Step(`^other peer should receive a message hello world$`, otherPeerShouldReceiveAMessageHelloWorld)
	ctx.Step(`^I can get it\'s ID$`, iCanGetItsID)
	ctx.Step(`^my peer can connect to it$`, myPeerCanConnectToIt)
}
