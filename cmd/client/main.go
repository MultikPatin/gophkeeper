package main

import (
	"fmt"
	"main/internal/client/app/proto"
	"main/internal/client/cli"
	"main/internal/client/config"
	l "main/internal/logger"
)

func main() {
	fmt.Printf("Build version: %s\n", buildVersion)
	fmt.Printf("Build date: %s\n", buildDate)
	fmt.Printf("Build commit: %s\n", buildCommit)

	logger := l.GetLogger()
	defer l.SyncLogger()

	c := config.Parse(logger)

	client, err := proto.NewGothKeeperClient(c.GRPCAddr)
	if err != nil {
		logger.Fatalw(err.Error(), "event", "initialize client")
	}
	defer client.Close()

	cli.Execute(client)
}
