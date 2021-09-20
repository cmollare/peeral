package main_test

import (
	"context"
	"errors"
	"io/ioutil"
	"log"
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
			e <- nil
		},
	}

	peer1 := libp2p.NewPeer(mockHostCallbacks, nil)
	peer1.Connect()

	failure := <-e
	if failure != nil {
		return nil, failure
	}

	<-e

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
		},
	}

	peer2 := libp2p.NewPeer(mockHostCallbacks, nil)
	peer2.Connect()

	c := context.WithValue(ctx, ctxListeningError{}, <-e)
	c = context.WithValue(c, ctxPeerFoundError{}, <-e)

	return c, nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		log.SetOutput(ioutil.Discard)

		return ctx, nil
	})

	ctx.Step(`^a peer is connected$`, aPeerIsConnected)
	ctx.Step(`^has been for (\d+) seconds$`, hasBeenForSeconds)
	ctx.Step(`^it should find peers$`, itShouldFindPeers)
	ctx.Step(`^it should start listening$`, itShouldStartListening)
	ctx.Step(`^my peer is started$`, myPeerIsStarted)
}
