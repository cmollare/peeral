package main

import (
	"context"

	app "example.com/server/domain"
	i "example.com/server/domain/adapter/infra"
	"example.com/server/domain/model"
	net "example.com/server/infra/network"
	golog "github.com/ipfs/go-log/v2"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	golog.SetAllLoggers(golog.LevelInfo)

	var connection i.Network = net.P2p{}
	app := app.New(connection)
	app.Connect(ctx, model.NewLibp2pConnectionOptions(8000))

	select {}
}
